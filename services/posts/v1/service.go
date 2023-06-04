package v1

import (
	"bytes"
	"context"
	"encoding/base64"
	pb_posts "fivi/gen/go/posts/v1"
	"fivi/lib/jwt"
	"fivi/lib/store"
	repository2 "fivi/services/posts/v1/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"path/filepath"
	"strconv"
)

var _ pb_posts.ServiceServer = (*Server)(nil)

type (
	Server struct {
		repo *repository2.Queries

		store *store.Store
	}
)

func New(repo *repository2.Queries, store *store.Store) *Server {
	return &Server{
		repo:  repo,
		store: store,
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

	return &pb_posts.CreatePostResponse{}, nil
}

func (s *Server) GetPost(ctx context.Context, request *pb_posts.GetPostRequest) (*pb_posts.GetPostResponse, error) {
	id, err := uuid.Parse(request.Identifier)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
	}

	post, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve post")
	}

	return &pb_posts.GetPostResponse{
		Post: s.toPbPost(post),
	}, nil
}

func (s *Server) toPbPost(post repository2.Post) *pb_posts.Post {
	numOfImages, _ := s.CountImages(context.Background(), post.ID)

	return &pb_posts.Post{
		Identifier:  post.ID.String(),
		Text:        post.Payload,
		CreatorId:   post.CreatorID.String(),
		NumOfImages: int32(numOfImages),
	}
}

func (s *Server) toPbPosts(posts []repository2.Post) []*pb_posts.Post {
	pbPosts := make([]*pb_posts.Post, 0, len(posts))

	for _, post := range posts {
		pbPosts = append(pbPosts, s.toPbPost(post))
	}

	return pbPosts
}

func (s *Server) GetPostsByCreator(ctx context.Context, request *pb_posts.GetPostsByCreatorRequest) (*pb_posts.GetPostsByCreatorResponse, error) {
	userIDStr, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "could not parse user id")
	}

	posts, err := s.repo.ListUserPosts(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "could not list posts")
	}

	return &pb_posts.GetPostsByCreatorResponse{
		Posts: s.toPbPosts(posts),
	}, nil
}

func (s *Server) UpdatePost(ctx context.Context, request *pb_posts.UpdatePostRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(request.Identifier)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
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

	return &emptypb.Empty{}, nil
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
	//TODO implement me
	panic("implement me")
}

func (s *Server) CreateImage(ctx context.Context, postID uuid.UUID, imageName string, reader io.Reader) error {
	pathFromRoot := filepath.Join("posts", postID.String(), imageName)

	return s.store.Create(ctx, pathFromRoot, reader)
}

func (s *Server) CountImages(ctx context.Context, productID uuid.UUID) (int, error) {
	pathFromRoot := filepath.Join("posts", productID.String())
	return s.store.Count(ctx, pathFromRoot)
}
