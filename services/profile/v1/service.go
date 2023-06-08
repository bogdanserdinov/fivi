package v1

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	pb_followers "fivi/gen/go/followers/v1"
	pb_posts "fivi/gen/go/posts/v1"
	profilepb "fivi/gen/go/profile/v1"
	"fivi/lib/jwt"
	"fivi/lib/store"
	repository2 "fivi/services/profile/v1/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tyler-smith/go-bip39"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"path/filepath"
	"strings"
)

var _ profilepb.ProfileServiceServer = (*Service)(nil)

type (
	Service struct {
		jwt *jwt.JWT

		repo *repository2.Queries

		store     *store.Store
		posts     pb_posts.ServiceClient
		followers pb_followers.FollowersServiceClient
	}
)

func New(jwt *jwt.JWT, repo *repository2.Queries, store *store.Store, posts pb_posts.ServiceClient, followers pb_followers.FollowersServiceClient) *Service {
	return &Service{
		repo:      repo,
		jwt:       jwt,
		store:     store,
		posts:     posts,
		followers: followers,
	}
}

func (s *Service) GenerateMnemonic(ctx context.Context, empty *emptypb.Empty) (*profilepb.GenerateMnemonicResponse, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, errors.Wrap(err, "could not generate entropy")
	}

	mnemonicPhrases, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, errors.Wrap(err, "could not generate mnemonic")
	}

	words := strings.Split(mnemonicPhrases, " ")
	if len(words) > 12 {
		words = words[:12]
	}

	return &profilepb.GenerateMnemonicResponse{
		Mnemonic: words,
	}, nil
}

func (s *Service) Register(ctx context.Context, request *profilepb.RegisterRequest) (*profilepb.RegisterResponse, error) {
	userID := uuid.New()
	err := s.repo.CreateUser(ctx, repository2.CreateUserParams{
		ID:       userID,
		Name:     request.GetFullName(),
		Email:    request.GetEmail(),
		Username: request.GetName(),
		Mnemonic: strings.Join(request.GetMnemonic(), " "),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not create user")
	}

	_, token, err := s.jwt.NewWithUserData(userID.String())
	if err != nil {
		log.Errorf("couldn't create jwt token, err: %v", err)
		return nil, errors.Wrap(err, "could not create users jwt token")
	}

	return &profilepb.RegisterResponse{
		Jwt: token,
	}, nil
}

func (s *Service) Login(ctx context.Context, request *profilepb.LoginRequest) (*profilepb.LoginResponse, error) {
	user, err := s.repo.GetByUsername(ctx, request.GetUsername())
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve user by username")
	}

	if user.Mnemonic != strings.Join(request.GetMnemonic(), " ") {
		return nil, errors.Wrap(err, "invalid mnemonic phrase")
	}

	_, token, err := s.jwt.NewWithUserData(user.ID.String())
	if err != nil {
		log.Errorf("couldn't create jwt token, err: %v", err)
		return nil, errors.Wrap(err, "could not create users jwt token")
	}

	return &profilepb.LoginResponse{
		Jwt: token,
	}, nil
}

func (s *Service) UpdateProfile(ctx context.Context, request *profilepb.UpdateProfileRequest) (*profilepb.Person, error) {
	userID, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "could not parse user id")
	}

	err = s.repo.UpdateUser(ctx, repository2.UpdateUserParams{
		ID:       id,
		Name:     request.GetName(),
		Email:    request.GetEmail(),
		Username: request.GetUsername(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not update user")
	}

	if request.Image != "" {
		_ = s.DeleteAvatar(ctx, id)

		img, err := base64.StdEncoding.DecodeString(request.Image)
		if err != nil {
			return nil, err
		}

		err = s.CreateAvatar(ctx, id, bytes.NewBuffer(img))
		if err != nil {
			return nil, err
		}
	}

	return s.GetProfileByDIDNoAuth(ctx, &profilepb.GetProfileByDIDRequest{
		UserDid: userID,
	})
}

func (s *Service) GetProfileByDID(ctx context.Context, request *profilepb.GetProfileByDIDRequest) (*profilepb.Person, error) {
	var userID string

	if request.GetUserDid() == "" {
		var err error
		userID, err = jwt.DIDFromCtx(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
		}
	} else {
		userID = request.GetUserDid()
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "could not parse user id")
	}

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve user")
	}

	followers, err := s.followers.ListFollowers(ctx, &pb_followers.ListFollowersRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not list followers")
	}

	followings, err := s.followers.ListFollowings(ctx, &pb_followers.ListFollowingsRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not list followings")
	}

	var isFollowed bool
	if request.GetUserDid() != "" {
		userID, err = jwt.DIDFromCtx(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
		}

		isFollow, err := s.followers.IsFollowing(ctx, &pb_followers.IsFollowingRequest{
			UserId:         request.GetUserDid(),
			UserToFollowId: userID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "could not retrieve is user followed, %v", err)
		}
		isFollowed = isFollow.GetIsFollow()
	} else {
		isFollowed = true
	}

	posts, err := s.posts.GetPostsByCreator(ctx, &pb_posts.GetPostsByCreatorRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not list posts")
	}

	relatedPath := filepath.Join("users", user.ID.String()+".png")

	return &profilepb.Person{
		Id:             id.String(),
		Email:          user.Email,
		Username:       user.Username,
		NumOfPosts:     int64(len(posts.GetPosts())),
		Subscribers:    followers.GetFollowers(),
		Subscriptions:  followings.GetFollowings(),
		IsAvatarExists: s.store.Stat(ctx, relatedPath),
		IsFollowed:     isFollowed,
	}, nil
}

func (s *Service) GetProfileByDIDNoAuth(ctx context.Context, request *profilepb.GetProfileByDIDRequest) (*profilepb.Person, error) {
	id, err := uuid.Parse(request.GetUserDid())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "could not parse user id")
	}

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve user")
	}

	relatedPath := filepath.Join("users", user.ID.String()+".png")
	return &profilepb.Person{
		Id:             id.String(),
		Email:          user.Email,
		Username:       user.Username,
		IsAvatarExists: s.store.Stat(ctx, relatedPath),
	}, nil
}

func (s *Service) SearchDIDsByUsername(ctx context.Context, request *profilepb.SearchDIDsByUsernameRequest) (*profilepb.SearchDIDsByUsernameResponse, error) {
	users, err := s.repo.ListUserIDsWithName(ctx, sql.NullString{
		String: request.GetUsername(),
		Valid:  true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not list users by username")
	}

	pbProfiles := make([]*profilepb.Person, 0, len(users))
	for _, user := range users {
		followers, err := s.followers.ListFollowers(ctx, &pb_followers.ListFollowersRequest{
			UserId: user.ID.String(),
		})
		if err != nil {
			return nil, errors.Wrap(err, "could not list followers")
		}

		followings, err := s.followers.ListFollowings(ctx, &pb_followers.ListFollowingsRequest{
			UserId: user.ID.String(),
		})
		if err != nil {
			return nil, errors.Wrap(err, "could not list followings")
		}

		posts, err := s.posts.GetPostsByCreator(ctx, &pb_posts.GetPostsByCreatorRequest{
			UserId: user.ID.String(),
		})
		if err != nil {
			return nil, errors.Wrap(err, "could not list posts")
		}

		pbProfile := &profilepb.Person{
			Id:            user.ID.String(),
			Email:         user.Email,
			Username:      user.Username,
			NumOfPosts:    int64(len(posts.GetPosts())),
			Subscribers:   followers.GetFollowers(),
			Subscriptions: followings.GetFollowings(),
		}

		pbProfiles = append(pbProfiles, pbProfile)
	}

	return &profilepb.SearchDIDsByUsernameResponse{
		Profiles: pbProfiles,
	}, nil
}

func (s *Service) CreateAvatar(ctx context.Context, userID uuid.UUID, reader io.Reader) error {
	pathFromRoot := filepath.Join("users", userID.String()+".png")

	return s.store.Create(ctx, pathFromRoot, reader)
}

func (s *Service) DeleteAvatar(ctx context.Context, userID uuid.UUID) error {
	pathFromRoot := filepath.Join("users", userID.String()+".png")

	return s.store.Delete(ctx, pathFromRoot)
}
