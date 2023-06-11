package v1

import (
	"bytes"
	"context"
	"encoding/base64"
	pb_comments "fivi/gen/go/comments/v1"
	pb_followers "fivi/gen/go/followers/v1"
	pb_likes "fivi/gen/go/likes/v1"
	pb_posts "fivi/gen/go/posts/v1"
	profilepb "fivi/gen/go/profile/v1"
	"fivi/lib/jwt"
	"fivi/lib/store"
	repository2 "fivi/services/posts/v1/repository"
	"io"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ pb_posts.ServiceServer = (*Server)(nil)

type (
	Server struct {
		repo *repository2.Queries

		store     *store.Store
		profiles  profilepb.ProfileServiceClient
		comments  pb_comments.ServiceClient
		likes     pb_likes.ServiceClient
		followers pb_followers.FollowersServiceClient
	}
)

func New(repo *repository2.Queries, store *store.Store, profiles profilepb.ProfileServiceClient, comments pb_comments.ServiceClient, likes pb_likes.ServiceClient, followers pb_followers.FollowersServiceClient) *Server {
	return &Server{
		repo:      repo,
		store:     store,
		profiles:  profiles,
		comments:  comments,
		likes:     likes,
		followers: followers,
	}
}

func (s *Server) CreatePost(ctx context.Context, request *pb_posts.CreatePostRequest) (*pb_posts.CreatePostResponse, error) {
	userIDStr, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "could not parse user id")
	}

	postID := uuid.New()

	err = s.repo.CreatePosts(ctx, repository2.CreatePostsParams{
		ID:        postID,
		Payload:   request.GetText(),
		CreatorID: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not create post")
	}

	for i, image := range request.Images {
		img, err := base64.StdEncoding.DecodeString(image)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "could not decode image")
		}

		name := strconv.Itoa(i) + ".png"

		err = s.CreateImage(ctx, postID, name, bytes.NewBuffer(img))
		if err != nil {
			return nil, errors.Wrap(err, "could not create image")
		}
	}

	post, err := s.repo.Get(ctx, postID)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve post")
	}

	pbPost, err := s.toPbPost(userIDStr, post)
	if err != nil {
		return nil, errors.Wrap(err, "could not cast post")
	}

	return &pb_posts.CreatePostResponse{
		Post: pbPost,
	}, nil
}

func (s *Server) GetPost(ctx context.Context, request *pb_posts.GetPostRequest) (*pb_posts.GetPostResponse, error) {
	id, err := uuid.Parse(request.Identifier)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
	}

	userIDStr, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	post, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve post")
	}

	pbPost, err := s.toPbPost(userIDStr, post)
	if err != nil {
		return nil, errors.Wrap(err, "could not cast post")
	}

	return &pb_posts.GetPostResponse{
		Post: pbPost,
	}, nil
}

func (s *Server) toPbPost(did string, post repository2.Post) (*pb_posts.Post, error) {
	numOfImages, _ := s.CountImages(context.Background(), post.ID)

	creatorProfile, err := s.profiles.GetProfileByDIDNoAuth(context.Background(), &profilepb.GetProfileByDIDRequest{
		UserDid: post.CreatorID.String(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not attach profile to post")
	}

	comments, err := s.comments.GetByID(context.Background(), &pb_comments.GetByIDRequest{
		PostId: post.ID.String(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not attach comments to post")
	}

	numOfLikes, err := s.likes.GetNumberOfLikes(context.Background(), &pb_likes.CountPostLikesRequest{
		PostId: post.ID.String(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not attach num of likes to post")
	}

	isLiked, err := s.likes.IsLiked(context.Background(), &pb_likes.IsLikedRequest{
		Like: &pb_likes.Like{
			Id:     did,
			PostId: post.ID.String(),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not attach is liked to post")
	}

	return &pb_posts.Post{
		Identifier:      post.ID.String(),
		Text:            post.Payload,
		CreatorId:       post.CreatorID.String(),
		NumOfImages:     int32(numOfImages),
		Comments:        comments.GetComments(),
		NumOfComments:   int32(len(comments.GetComments())),
		CreatorUsername: creatorProfile.Username,
		NumOfLikes:      numOfLikes.GetCount(),
		CreatorProfile:  creatorProfile,
		IsLiked:         isLiked.GetIsLiked(),
	}, nil
}

func (s *Server) toPbPosts(did string, posts []repository2.Post) ([]*pb_posts.Post, error) {
	pbPosts := make([]*pb_posts.Post, 0, len(posts))

	for _, post := range posts {
		pbPost, err := s.toPbPost(did, post)
		if err != nil {
			return nil, errors.Wrap(err, "could not cast post")
		}

		pbPosts = append(pbPosts, pbPost)
	}

	return pbPosts, nil
}

func (s *Server) GetPostsByCreator(ctx context.Context, request *pb_posts.GetPostsByCreatorRequest) (*pb_posts.GetPostsByCreatorResponse, error) {
	userIDStr := request.GetUserId()
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "could not parse user id")
	}

	posts, err := s.repo.ListUserPosts(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "could not list posts")
	}

	pbPosts, err := s.toPbPosts(userIDStr, posts)
	if err != nil {
		return nil, errors.Wrap(err, "could not cast posts")
	}

	return &pb_posts.GetPostsByCreatorResponse{
		Posts: pbPosts,
	}, nil
}

func (s *Server) UpdatePost(ctx context.Context, request *pb_posts.UpdatePostRequest) (*pb_posts.UpdatePostResponse, error) {
	id, err := uuid.Parse(request.Identifier)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
	}

	userIDStr, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	relatedPath := filepath.Join("posts", id.String())
	err = s.store.DeleteFolder(ctx, relatedPath)
	if err != nil {
		return nil, errors.Wrap(err, "could not delete posts folder")
	}

	err = s.repo.UpdatePost(ctx, repository2.UpdatePostParams{
		ID:      id,
		Payload: request.Text,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not delete post")
	}

	if len(request.GetImages()) != 0 {
		for i, image := range request.Images {
			img, err := base64.StdEncoding.DecodeString(image)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, "could not decode image")
			}

			name := strconv.Itoa(i) + ".png"

			err = s.CreateImage(ctx, id, name, bytes.NewBuffer(img))
			if err != nil {
				return nil, errors.Wrap(err, "could not create image")
			}
		}
	}

	post, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve post")
	}

	pbPost, err := s.toPbPost(userIDStr, post)
	if err != nil {
		return nil, errors.Wrap(err, "could not cast post")
	}

	return &pb_posts.UpdatePostResponse{
		Post: pbPost,
	}, nil
}

func (s *Server) DeletePost(ctx context.Context, request *pb_posts.DeletePostRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(request.Identifier)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
	}

	relatedPath := filepath.Join("posts", id.String())
	err = s.store.DeleteFolder(ctx, relatedPath)
	if err != nil {
		return nil, errors.Wrap(err, "could not delete posts folder")
	}

	err = s.repo.DeletePost(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not delete post")
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListPosts(ctx context.Context, empty *emptypb.Empty) (*pb_posts.ListPostsResponse, error) {
	userIDStr, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	followings, err := s.followers.ListFollowings(ctx, &pb_followers.ListFollowingsRequest{
		UserId: userIDStr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not list followings")
	}

	users := make([]uuid.UUID, 0, len(followings.GetFollowings()))

	for _, following := range followings.GetFollowings() {
		id, err := uuid.Parse(following.Id)
		if err != nil {
			return nil, errors.Wrap(err, "invalid following id")
		}
		users = append(users, id)
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	users = append(users, userID)

	posts, err := s.repo.ListUsersPosts(ctx, users)
	if err != nil {
		return nil, errors.Wrap(err, "could not list posts by followings")
	}

	pbPosts, err := s.toPbPosts(userIDStr, posts)
	if err != nil {
		return nil, errors.Wrap(err, "could not cast posts")
	}

	return &pb_posts.ListPostsResponse{
		Posts: pbPosts,
	}, nil
}

func (s *Server) CreateImage(ctx context.Context, postID uuid.UUID, imageName string, reader io.Reader) error {
	pathFromRoot := filepath.Join("posts", postID.String(), imageName)

	return s.store.Create(ctx, pathFromRoot, reader)
}

func (s *Server) CountImages(ctx context.Context, productID uuid.UUID) (int, error) {
	pathFromRoot := filepath.Join("posts", productID.String())
	return s.store.Count(ctx, pathFromRoot)
}
