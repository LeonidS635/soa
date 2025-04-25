package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto/converters"
)

func (h *PostsHandlers) GetAllPosts(ctx context.Context, req *postspb.GetAllPostsRequest) (
	*postspb.GetAllPostsResponse, error,
) {
	posts, err := h.useCase.GetAllPosts(req.GetPage(), postsPerPageN)
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*postspb.PostFullInfo, len(posts))
	for i := 0; i < len(posts); i++ {
		pbPosts[i] = converters.PostFullInfoToPB(posts[i])
	}

	return &postspb.GetAllPostsResponse{Posts: pbPosts}, nil
}
