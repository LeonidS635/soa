package handlers

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto/converters"
)

func (h *StatsHandlers) GetPostDynamics(
	ctx context.Context, req *statspb.GetPostDynamicsRequest,
) (*statspb.GetPostDynamicsResponse, error) {
	dynamics, err := h.useCase.GetPostDynamics(ctx, req.PostId, req.Metric)
	if err != nil {
		return nil, err
	}

	return &statspb.GetPostDynamicsResponse{
		Dynamics: converters.DynamicsToPB(dynamics),
	}, nil
}
