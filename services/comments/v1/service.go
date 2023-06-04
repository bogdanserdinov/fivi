package v1

import (
	"context"
	pb_comments "fivi/gen/go/comments/v1"
	profilepb "fivi/gen/go/profile/v1"
	"fivi/lib/jwt"
	repository2 "fivi/services/comments/v1/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ pb_comments.ServiceServer = (*Service)(nil)

type (
	Service struct {
		repo *repository2.Queries

		profiles profilepb.ProfileServiceClient
	}
)

func New(repo *repository2.Queries, profiles profilepb.ProfileServiceClient) *Service {
	return &Service{
		repo:     repo,
		profiles: profiles,
	}
}

func (s *Service) Create(ctx context.Context, request *pb_comments.CreateRequest) (*emptypb.Empty, error) {
	did, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	postID, err := uuid.Parse(request.GetPostId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid postID")
	}

	err = s.repo.Create(ctx, repository2.CreateParams{
		ID:        uuid.New(),
		Text:      request.Text,
		PostID:    postID,
		CreatorID: did,
	})

	return &emptypb.Empty{}, err
}

func (s *Service) GetByID(ctx context.Context, request *pb_comments.GetByIDRequest) (*pb_comments.GetByIDResponse, error) {
	postID, err := uuid.Parse(request.GetPostId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid postID")
	}

	comments, err := s.repo.ListPostComments(ctx, postID)
	if err != nil {
		return nil, err
	}

	pbComments := make([]*pb_comments.Comment, 0, len(comments))

	for _, comment := range comments {
		usernameResp, err := s.profiles.GetProfileByDIDNoAuth(ctx, &profilepb.GetProfileByDIDRequest{
			UserDid: comment.CreatorID,
		})
		if err != nil {
			return nil, err
		}

		pbComment := &pb_comments.Comment{
			Identifier: comment.ID.String(),
			Text:       comment.Text,
			PostId:     postID.String(),
			Username:   usernameResp.GetUsername(),
		}

		pbComments = append(pbComments, pbComment)
	}

	return &pb_comments.GetByIDResponse{
		Comments: pbComments,
	}, nil
}
