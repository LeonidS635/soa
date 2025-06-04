package validators

import (
	"errors"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

func Metric(metric statspb.Metric) (int32, error) {
	switch metric {
	case statspb.Metric_VIEWS:
		return dto.MetricViews, nil
	case statspb.Metric_LIKES:
		return dto.MetricLikes, nil
	case statspb.Metric_COMMENTS:
		return dto.MetricComments, nil
	default:
		return -1, errors.New("unknown metric type")
	}
}
