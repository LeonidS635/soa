package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto/converters"
)

func (h *PostsHandlers) UpdatePost(ctx context.Context, req *postspb.UpdatePostRequest) (
	*postspb.UpdatePostResponse, error,
) {
	err := h.useCase.UpdatePost(req.GetPostId(), req.GetUserId(), converters.PostToDTO(req.GetPost()))
	if err != nil {
		return nil, err
	}

	return &postspb.UpdatePostResponse{}, nil
}
