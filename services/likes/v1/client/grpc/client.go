package grpc

import (
	pb_likes "fivi/gen/go/likes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewLikesServiceClient returns a grpc client to the comments service.
func NewLikesServiceClient(addr string) (pb_likes.ServiceClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return pb_likes.NewServiceClient(conn), nil
}
