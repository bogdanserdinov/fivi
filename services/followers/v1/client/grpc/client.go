package grpc

import (
	pb_followers "fivi/gen/go/followers/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewFollowersServiceClient returns a grpc client to the comments service.
func NewFollowersServiceClient(addr string) (pb_followers.FollowersServiceClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return pb_followers.NewFollowersServiceClient(conn), nil
}
