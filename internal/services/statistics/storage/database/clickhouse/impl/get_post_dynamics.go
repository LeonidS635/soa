package impl

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

func (c *ClickHouseStatisticsStorageImpl) GetPostDynamics(ctx context.Context, postID, by int32) (
	dto.PostDynamics, error,
) {
	var metric string
	switch by {
	case dto.MetricViews:
		metric = "views"
	case dto.MetricLikes:
		metric = "likes"
	case dto.MetricComments:
		metric = "comments"
	default:
		return nil, errors.New("invalid argument \"by\"")
	}

	query := fmt.Sprintf(
		`SELECT post_date, sum(%s) FROM daily_stats
		WHERE post_id = ? AND %s != 0
		GROUP BY post_date
		ORDER BY post_date DESC`,
		metric, metric,
	)

	rows, err := c.db.Query(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postDynamics dto.PostDynamics
	for rows.Next() {
		var date time.Time
		var value uint64
		if err = rows.Scan(&date, &value); err != nil {
			return nil, err
		}

		postDynamics = append(
			postDynamics, struct {
				Date        time.Time
				MetricValue uint64
			}{Date: date, MetricValue: value},
		)
	}

	return postDynamics, err
}
