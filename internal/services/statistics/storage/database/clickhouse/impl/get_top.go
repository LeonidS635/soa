package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

const (
	TopPosts = iota
	TopAuthors
)

func (c *ClickHouseStatisticsStorageImpl) GetTop(ctx context.Context, postOrAuthor, by int32) (dto.Top, error) {
	var selected string
	switch postOrAuthor {
	case TopPosts:
		selected = "post_id"
	case TopAuthors:
		selected = "author_id"
	default:
		return nil, errors.New("invalid argument \"postOrAuthor\"")
	}

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
		`SELECT %s, sum(%s) AS total FROM daily_stats GROUP BY %s ORDER BY total DESC`,
		selected, metric, selected,
	)

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var top dto.Top
	for rows.Next() {
		var id int32
		var metricValue uint64
		if err = rows.Scan(&id, &metricValue); err != nil {
			return nil, err
		}

		top = append(
			top, struct {
				ID          int32
				MetricValue uint64
			}{ID: id, MetricValue: metricValue},
		)
	}

	return top, nil
}
