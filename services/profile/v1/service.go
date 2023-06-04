package v1

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
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

		store *store.Store
	}
)

func New(jwt *jwt.JWT, repo *repository2.Queries, store *store.Store) *Service {
	return &Service{
		repo:  repo,
		jwt:   jwt,
		store: store,
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

func (s *Service) UpdateProfile(ctx context.Context, request *profilepb.UpdateProfileRequest) (*emptypb.Empty, error) {
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
		Username: request.GetUsername(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not update user")
	}

	_ = s.DeleteAvatar(ctx, id)

	img, err := base64.StdEncoding.DecodeString(request.Image)
	if err != nil {
		return nil, err
	}

	_ = s.CreateAvatar(ctx, id, bytes.NewBuffer(img))

	return &emptypb.Empty{}, nil
}

func (s *Service) GetProfileByDID(ctx context.Context, request *emptypb.Empty) (*profilepb.Person, error) {
	userID, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "could not parse user id")
	}

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve user")
	}

	return &profilepb.Person{
		Id:       id.String(),
		Name:     user.Name,
		Username: user.Username,
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

	return &profilepb.Person{
		Id:       id.String(),
		Name:     user.Name,
		Username: user.Username,
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
		pbProfile := &profilepb.Person{
			Id:       user.ID.String(),
			Name:     user.Name,
			Username: user.Username,
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
