package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase/validators"
)

func (uc *StatsUseCase) GetPostDynamics(ctx context.Context, postID int32, by statspb.Metric) (
	dto.PostDynamics, error,
) {
	metric, err := validators.Metric(by)
	if err != nil {
		return nil, err
	}

	postDynamics, err := uc.storage.GetPostDynamics(ctx, postID, metric)
	if err != nil {
		return nil, err
	}

	return postDynamics, nil
}
