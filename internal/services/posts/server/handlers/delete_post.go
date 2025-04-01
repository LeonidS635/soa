package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
)

func (h *PostsHandlers) DeletePost(ctx context.Context, req *postspb.DeletePostRequest) (
	*postspb.DeletePostResponse, error,
) {
	err := h.useCase.DeletePost(req.GetPostId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &postspb.DeletePostResponse{}, nil
}
