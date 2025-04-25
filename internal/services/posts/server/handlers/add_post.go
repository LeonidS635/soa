package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto/converters"
)

func (h *PostsHandlers) AddPost(ctx context.Context, req *postspb.AddPostRequest) (*postspb.AddPostResponse, error) {
	postId, err := h.useCase.AddPost(req.GetUserId(), converters.PostToDTO(req.GetPost()))
	if err != nil {
		return nil, err
	}
	return &postspb.AddPostResponse{PostId: postId}, nil
}
