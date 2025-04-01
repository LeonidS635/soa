package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/server/handlers/converters"
)

func (h *PostsHandlers) GetPost(ctx context.Context, req *postspb.GetPostRequest) (*postspb.GetPostResponse, error) {
	post, err := h.useCase.GetPost(req.GetPostId(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &postspb.GetPostResponse{Post: converters.PostFullInfoToPB(post)}, nil
}
