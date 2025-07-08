package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase/validators"
)

func (uc *StatsUseCase) GetTopPosts(ctx context.Context, by statspb.Metric) (dto.Top, error) {
	metric, err := validators.Metric(by)
	if err != nil {
		return nil, err
	}

	return uc.storage.GetTopPosts(ctx, metric)
}
