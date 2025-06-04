package validators

import (
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

func TestMetric(t *testing.T) {
	testCases := []struct {
		name   string
		metric statspb.Metric

		metricValue int32
		wantErr     bool
	}{
		{
			name:        "views",
			metric:      statspb.Metric_VIEWS,
			metricValue: dto.MetricViews,
			wantErr:     false,
		},
		{
			name:        "likes",
			metric:      statspb.Metric_LIKES,
			metricValue: dto.MetricLikes,
			wantErr:     false,
		},
		{
			name:        "comments",
			metric:      statspb.Metric_COMMENTS,
			metricValue: dto.MetricComments,
			wantErr:     false,
		},
		{
			name:    "incorrect metric",
			metric:  -1,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				metric, err := Metric(tc.metric)
				if (err != nil) != tc.wantErr {
					t.Errorf("Metric() error = %v, wantErr %v", err, tc.wantErr)
				}

				if err == nil && metric != tc.metricValue {
					t.Errorf("Metric() got = %v, want %v", metric, tc.metricValue)
				}
			},
		)
	}
}
