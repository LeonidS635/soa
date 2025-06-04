package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase/mocks"
	"github.com/stretchr/testify/require"
)

func TestStatsUseCase_GetPostDynamics(t *testing.T) {
	mockStorage := mocks.NewStatsStorage()
	uc := NewStatsUseCase(mockStorage)

	correctDynamics := dto.PostDynamics{
		{
			Date:        time.Now(),
			MetricValue: 123,
		},
		{
			Date:        time.Now(),
			MetricValue: 456,
		},
		{
			Date:        time.Now(),
			MetricValue: 789,
		},
	}

	testCases := []struct {
		name   string
		postID int32
		metric statspb.Metric

		mock    mocks.StatsStorage
		wantErr bool
	}{
		{
			name:   "storage error",
			postID: 1,
			metric: statspb.Metric_VIEWS,
			mock: mocks.StatsStorage{
				Err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:   "incorrect metric",
			metric: 123,
			postID: 1,
			mock: mocks.StatsStorage{
				Err: errors.New("something went wrong"),
			},
			wantErr: true,
		},
		{
			name:   "correct",
			postID: 1,
			metric: statspb.Metric_COMMENTS,
			mock: mocks.StatsStorage{
				Dynamics: correctDynamics,
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				mockStorage.Dynamics = tc.mock.Dynamics
				mockStorage.Err = tc.mock.Err

				dynamics, err := uc.GetPostDynamics(ctx, tc.postID, tc.metric)
				if (err != nil) != tc.wantErr {
					t.Errorf("GetPostDynamics() error = %v, wantErr %v", err, tc.wantErr)
					return
				}

				if err == nil {
					require.Equal(t, correctDynamics, dynamics, "GetPostDynamics() returns wrong stats")
				}
			},
		)
	}
}
