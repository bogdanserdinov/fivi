package grpc

import (
	profilepb "fivi/gen/go/profile/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewProfileServiceClient returns a grpc client to the comments service.
func NewProfileServiceClient(addr string) (profilepb.ProfileServiceClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return profilepb.NewProfileServiceClient(conn), nil
}
