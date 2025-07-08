package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

func (uc *StatsUseCase) GetPostStats(ctx context.Context, postID int32) (dto.PostStats, error) {
	postStats, err := uc.storage.GetPostStats(ctx, postID)
	if err != nil {
		return dto.PostStats{}, err
	}

	return postStats, nil
}
