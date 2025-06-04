package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/LeonidS635/soa/internal/services/statistics/dto"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase/mocks"
	"github.com/stretchr/testify/require"
)

func TestStatsUseCase_GetPostStats(t *testing.T) {
	mockStorage := mocks.NewStatsStorage()
	uc := NewStatsUseCase(mockStorage)

	correctStats := dto.PostStats{
		Views:    1000,
		Likes:    123,
		Comments: 82,
	}

	testCases := []struct {
		name   string
		postID int32

		mock    mocks.StatsStorage
		wantErr bool
	}{
		{
			name:   "storage error",
			postID: 1,
			mock: mocks.StatsStorage{
				Err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:   "correct",
			postID: 1,
			mock: mocks.StatsStorage{
				Stats: correctStats,
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				mockStorage.Stats = tc.mock.Stats
				mockStorage.Err = tc.mock.Err

				stats, err := uc.GetPostStats(ctx, tc.postID)
				if (err != nil) != tc.wantErr {
					t.Errorf("GetPostStats() error = %v, wantErr %v", err, tc.wantErr)
					return
				}

				if err == nil {
					require.Equal(t, correctStats, stats, "GetPostStats() returns wrong stats")
				}
			},
		)
	}
}
