package posts

import (
	"fmt"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GateWayPostsHandlers struct {
	userServiceURL     string
	postsServiceClient postspb.PostsServiceClient
}

func NewGateWayPostsHandlers(userServiceURL, postsServiceHost string) (*GateWayPostsHandlers, error) {
	conn, err := grpc.NewClient(postsServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create posts grpc client: %v", err)
	}

	return &GateWayPostsHandlers{
		userServiceURL:     userServiceURL,
		postsServiceClient: postspb.NewPostsServiceClient(conn),
	}, nil
}
