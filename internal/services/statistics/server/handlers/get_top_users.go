package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto/converters"
)

func (h *StatsHandlers) GetTopUsers(
	ctx context.Context, req *statspb.GetTopUsersRequest,
) (*statspb.GetTopUsersResponse, error) {
	top, err := h.useCase.GetTopUsers(ctx, req.Metric)
	if err != nil {
		return nil, err
	}

	return &statspb.GetTopUsersResponse{
		Users: converters.TopUsersToPB(top),
	}, nil
}
