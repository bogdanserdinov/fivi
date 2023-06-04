package grpc

import (
	pb_comments "fivi/gen/go/comments/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewCommentsServiceClient returns a grpc client to the comments service.
func NewCommentsServiceClient(addr string) (pb_comments.ServiceClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return pb_comments.NewServiceClient(conn), nil
}
