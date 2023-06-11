package v1

import (
	"context"
	"fivi/gen/go/followers/v1"
	profilepb "fivi/gen/go/profile/v1"
	"fivi/lib/jwt"
	repository2 "fivi/services/followers/v1/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ followers.FollowersServiceServer = (*Service)(nil)

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

func (s *Service) DeleteFollower(ctx context.Context, request *followers.DeleteFollowerRequest) (*emptypb.Empty, error) {
	userID, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	followeeID, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	id, err := s.repo.GetFollower(ctx, repository2.GetFollowerParams{
		FollowerID: followeeID.String(),
		FolloweeID: userID,
	})

	err = s.repo.DeleteFollow(ctx, id)

	return &emptypb.Empty{}, err
}

func (s *Service) Follow(ctx context.Context, request *followers.FollowRequest) (*followers.Follower, error) {
	id := uuid.New()
	err := s.repo.CreateFollow(ctx, repository2.CreateFollowParams{
		ID:         id,
		FollowerID: request.UserId,
		FolloweeID: request.UserToFollowId,
	})

	usernameResp, err := s.profiles.GetProfileByDIDNoAuth(ctx, &profilepb.GetProfileByDIDRequest{
		UserDid: request.UserId,
	})
	if err != nil {
		return nil, err
	}

	pbFollower := &followers.Follower{
		Id:             id.String(),
		UserId:         request.GetUserId(),
		Username:       usernameResp.GetUsername(),
		IsAvatarExists: usernameResp.GetIsAvatarExists(),
		IsSubscribed:   true,
	}

	return pbFollower, err
}

func (s *Service) Unfollow(ctx context.Context, request *followers.UnFollowRequest) (*emptypb.Empty, error) {
	userID, err := jwt.DIDFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "could not retrieve user id")
	}

	followeeID, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	id, err := s.repo.GetFollower(ctx, repository2.GetFollowerParams{
		FollowerID: userID,
		FolloweeID: followeeID.String(),
	})

	err = s.repo.DeleteFollow(ctx, id)

	return &emptypb.Empty{}, err
}

func (s *Service) ListFollowers(ctx context.Context, request *followers.ListFollowersRequest) (*followers.ListFollowersResponse, error) {
	userIDs, err := s.repo.ListFollowers(ctx, request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbFollowers := make([]*followers.Follower, 0, len(userIDs))

	for _, userID := range userIDs {
		usernameResp, err := s.profiles.GetProfileByDIDNoAuth(ctx, &profilepb.GetProfileByDIDRequest{
			UserDid: userID.FollowerID,
		})
		if err != nil {
			return nil, err
		}

		isFollowed, err := s.IsFollowing(ctx, &followers.IsFollowingRequest{
			UserId:         userID.FollowerID,
			UserToFollowId: request.UserId,
		})
		if err != nil {
			return nil, err
		}

		pbFollower := &followers.Follower{
			Id:             userID.ID.String(),
			UserId:         userID.FollowerID,
			Username:       usernameResp.GetUsername(),
			IsAvatarExists: usernameResp.GetIsAvatarExists(),
			IsSubscribed:   isFollowed.GetIsFollow(),
		}

		pbFollowers = append(pbFollowers, pbFollower)
	}

	return &followers.ListFollowersResponse{
		Followers: pbFollowers,
	}, nil
}

func (s *Service) ListFollowings(ctx context.Context, request *followers.ListFollowingsRequest) (*followers.ListFollowingsResponse, error) {
	userIDs, err := s.repo.ListFollowings(ctx, request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbFollowers := make([]*followers.Follower, 0, len(userIDs))

	for _, userID := range userIDs {
		usernameResp, err := s.profiles.GetProfileByDIDNoAuth(ctx, &profilepb.GetProfileByDIDRequest{
			UserDid: userID.FolloweeID,
		})
		if err != nil {
			return nil, err
		}

		isFollowed, err := s.IsFollowing(ctx, &followers.IsFollowingRequest{
			UserId:         request.UserId,
			UserToFollowId: userID.FolloweeID,
		})
		if err != nil {
			return nil, err
		}

		pbFollower := &followers.Follower{
			Id:             userID.ID.String(),
			UserId:         userID.FolloweeID,
			Username:       usernameResp.GetUsername(),
			IsAvatarExists: usernameResp.GetIsAvatarExists(),
			IsSubscribed:   isFollowed.GetIsFollow(),
		}

		pbFollowers = append(pbFollowers, pbFollower)
	}

	return &followers.ListFollowingsResponse{
		Followings: pbFollowers,
	}, nil
}

func (s *Service) CountFollowers(ctx context.Context, request *followers.CountFollowersRequest) (*followers.CountFollowersResponse, error) {
	countFollowers, err := s.repo.CountFollowers(ctx, request.UserId)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		return nil, err
	}

	return &followers.CountFollowersResponse{
		Count: int32(countFollowers),
	}, nil
}

func (s *Service) CountFollowings(ctx context.Context, request *followers.CountFollowingsRequest) (*followers.CountFollowingsResponse, error) {
	countFollowings, err := s.repo.CountFollowings(ctx, request.UserId)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		return nil, err
	}

	return &followers.CountFollowingsResponse{
		Count: int32(countFollowings),
	}, nil
}

func (s *Service) IsFollowing(ctx context.Context, request *followers.IsFollowingRequest) (*followers.IsFollowingResponse, error) {
	in := repository2.IsFollowUserParams{
		FollowerID: request.UserId,
		FolloweeID: request.UserToFollowId,
	}

	isFollowUser, err := s.repo.IsFollowUser(ctx, in)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		return nil, err
	}

	resp := followers.IsFollowingResponse{
		IsFollow: isFollowUser,
	}

	return &resp, nil
}
