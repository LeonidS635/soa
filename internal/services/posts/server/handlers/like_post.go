package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
)

func (h *PostsHandlers) LikePost(ctx context.Context, req *postspb.LikePostRequest) (*postspb.LikePostResponse, error) {
	err := h.useCase.LikePost(req.GetPostId(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &postspb.LikePostResponse{}, nil
}
