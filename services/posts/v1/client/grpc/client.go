package grpc

import (
	pb_posts "fivi/gen/go/posts/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewPostsServiceClient returns a grpc client to the comments service.
func NewPostsServiceClient(addr string) (pb_posts.ServiceClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return pb_posts.NewServiceClient(conn), nil
}
