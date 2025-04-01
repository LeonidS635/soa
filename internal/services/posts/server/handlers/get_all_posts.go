package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/server/handlers/converters"
)

func (h *PostsHandlers) GetAllPosts(ctx context.Context, req *postspb.GetAllPostsRequest) (
	*postspb.GetAllPostsResponse, error,
) {
	posts, err := h.useCase.GetAllPosts(req.GetPostsPerPageN())
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*postspb.PostsList, len(posts))
	for page := 0; page < len(posts); page++ {
		pbPosts[page].PostsList = make([]*postspb.PostFullInfo, len(posts[page]))
		for i, post := range posts[page] {
			pbPosts[page].PostsList[i] = converters.PostFullInfoToPB(post)
		}
	}

	return &postspb.GetAllPostsResponse{Posts: pbPosts}, nil
}
