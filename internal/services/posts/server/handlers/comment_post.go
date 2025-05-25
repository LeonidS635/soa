package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
)

func (h *PostsHandlers) CommentPost(ctx context.Context, req *postspb.CommentPostRequest) (
	*postspb.CommentPostResponse, error,
) {
	err := h.useCase.CommentPost(req.GetPostId(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &postspb.CommentPostResponse{}, nil
}
