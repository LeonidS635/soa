package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto/converters"
)

func (h *StatsHandlers) GetTopPosts(
	ctx context.Context, req *statspb.GetTopPostsRequest,
) (*statspb.GetTopPostsResponse, error) {
	top, err := h.useCase.GetTopPosts(ctx, req.Metric)
	if err != nil {
		return nil, err
	}

	return &statspb.GetTopPostsResponse{
		Posts: converters.TopPostsToPB(top),
	}, nil
}
