package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto/converters"
)

func (h *PostsHandlers) GetAllMyPosts(
	ctx context.Context, req *postspb.GetAllMyPostsRequest,
) (*postspb.GetAllMyPostsResponse, error) {
	posts, err := h.useCase.GetAllMyPosts(req.GetPage(), postsPerPageN, req.GetUserId())
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*postspb.PostFullInfo, len(posts))
	for i := 0; i < len(posts); i++ {
		pbPosts[i] = converters.PostFullInfoToPB(posts[i])
	}

	return &postspb.GetAllMyPostsResponse{Posts: pbPosts}, nil
}
