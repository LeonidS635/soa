package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase/mocks"
	"github.com/stretchr/testify/require"
)

func TestStatsUseCase_GetTopUsers(t *testing.T) {
	mockStorage := mocks.NewStatsStorage()
	uc := NewStatsUseCase(mockStorage)

	correctTop := dto.Top{
		{
			ID:          3,
			MetricValue: 123,
		},
		{
			ID:          2,
			MetricValue: 456,
		},
		{
			ID:          10123,
			MetricValue: 789,
		},
	}

	testCases := []struct {
		name   string
		metric statspb.Metric

		mock    mocks.StatsStorage
		wantErr bool
	}{
		{
			name:   "storage error",
			metric: statspb.Metric_VIEWS,
			mock: mocks.StatsStorage{
				Err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:   "incorrect metric",
			metric: 123,
			mock: mocks.StatsStorage{
				Err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:   "correct",
			metric: statspb.Metric_COMMENTS,
			mock: mocks.StatsStorage{
				TopAuthors: correctTop,
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				mockStorage.TopAuthors = tc.mock.TopAuthors
				mockStorage.Err = tc.mock.Err

				top, err := uc.GetTopUsers(ctx, tc.metric)
				if (err != nil) != tc.wantErr {
					t.Errorf("GetTopUsers() error = %v, wantErr %v", err, tc.wantErr)
					return
				}

				if err == nil {
					require.Equal(t, correctTop, top, "GetTopUsers() returns wrong stats")
				}
			},
		)
	}
}
