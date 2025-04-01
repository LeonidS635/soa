package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/server/handlers/converters"
)

func (h *PostsHandlers) GetAllPostsOfOneAuthor(
	ctx context.Context, req *postspb.GetAllPostsOfOneAuthorRequest,
) (*postspb.GetAllPostsOfOneAuthorResponse, error) {
	posts, err := h.useCase.GetAllPostsOfOneAuthor(req.GetUserId(), req.GetPostsPerPageN())
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

	return &postspb.GetAllPostsOfOneAuthorResponse{Posts: pbPosts}, nil
}
