package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
)

func (h *StatsHandlers) GetPostStats(
	ctx context.Context, req *statspb.GetPostStatsRequest,
) (*statspb.GetPostStatsResponse, error) {
	stats, err := h.useCase.GetPostStats(ctx, req.PostId)
	if err != nil {
		return nil, err
	}

	return &statspb.GetPostStatsResponse{
		Views:    stats.Views,
		Likes:    stats.Likes,
		Comments: stats.Comments,
	}, nil
}
