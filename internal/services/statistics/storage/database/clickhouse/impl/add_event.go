package impl

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

func (c *ClickHouseStatisticsStorageImpl) AddEvent(
	ctx context.Context, postID, userID int32, date time.Time, eventType int32,
) error {
	var metric string
	switch eventType {
	case dto.MetricViews:
		metric = "views"
	case dto.MetricLikes:
		metric = "likes"
	case dto.MetricComments:
		metric = "comments"
	default:
		return errors.New("invalid argument \"by\"")
	}

	query := fmt.Sprintf(
		`INSERT INTO daily_stats (post_id, author_id, post_date, %s) VALUES (?, ?, ?, 1)`,
		metric,
	)

	return c.db.Exec(ctx, query, postID, userID, date)
}
