package v1

import (
	"context"
	"database/sql"
	pb_likes "fivi/gen/go/likes/v1"
	"fivi/lib/jwt"
	repository2 "fivi/services/likes/v1/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ pb_likes.ServiceServer = (*Service)(nil)

type (
	Service struct {
		repo *repository2.Queries
	}
)

func New(repo *repository2.Queries) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Like(ctx context.Context, request *pb_likes.LikeRequest) (*emptypb.Empty, error) {
	did, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	postID, err := uuid.Parse(request.GetPostId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid postID")
	}

	log.Println("LIKEEEEDD", err)

	isLiked, err := s.IsLiked(ctx, &pb_likes.IsLikedRequest{
		Like: &pb_likes.Like{
			Id:     did,
			PostId: postID.String(),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not check is liked")
	}

	if isLiked.GetIsLiked() {
		err = s.repo.DeleteLike(ctx, repository2.DeleteLikeParams{
			PostID: postID,
			UserID: did,
		})
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.Internal, "could not delete like")
		}
	} else {
		err = s.repo.Like(ctx, repository2.LikeParams{
			PostID: postID,
			UserID: did,
		})

		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, err
}

func (s *Service) GetNumberOfLikes(ctx context.Context, request *pb_likes.CountPostLikesRequest) (*pb_likes.CountPostResponse, error) {
	postID, err := uuid.Parse(request.GetPostId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid postID")
	}

	count, err := s.repo.CountPostLikes(ctx, postID)

	return &pb_likes.CountPostResponse{
		Count: int32(count),
	}, err
}

func (s *Service) GetLikes(ctx context.Context, request *pb_likes.GetListByPostRequest) (*pb_likes.GetLikesByPostResponse, error) {
	postID, err := uuid.Parse(request.GetPostId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid postID")
	}

	likes, err := s.repo.ListLikes(ctx, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not list likes")
	}

	pbLikes := make([]*pb_likes.Like, 0, len(likes))

	for _, like := range likes {
		pbLike := &pb_likes.Like{
			Id:     like.UserID,
			PostId: like.PostID.String(),
		}

		pbLikes = append(pbLikes, pbLike)
	}

	return &pb_likes.GetLikesByPostResponse{
		Likes: pbLikes,
	}, nil
}

func (s *Service) IsLiked(ctx context.Context, request *pb_likes.IsLikedRequest) (*pb_likes.IsLikedResponse, error) {
	postID, err := uuid.Parse(request.GetLike().GetPostId())
	if err != nil {
		return nil, errors.Wrap(err, "invalid post id")
	}

	isLiked, err := s.repo.IsLiked(ctx, repository2.IsLikedParams{
		PostID: postID,
		UserID: request.GetLike().GetId(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not get is_liked value")
	}
	log.Println("isLiked", isLiked)

	return &pb_likes.IsLikedResponse{
		IsLiked: isLiked,
	}, nil
}
