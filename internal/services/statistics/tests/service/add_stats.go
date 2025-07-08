package service

import (
	"context"
	"fmt"
	"time"

	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

func addPostStats(ctx context.Context, postID, userID int32, date time.Time, eventType int32) error {
	var metric string
	switch eventType {
	case dto.MetricViews:
		metric = "views"
	case dto.MetricLikes:
		metric = "likes"
	case dto.MetricComments:
		metric = "comments"
	default:
		return fmt.Errorf("invalid metric type")
	}

	query := fmt.Sprintf(
		`INSERT INTO daily_stats (post_id, author_id, post_date, %s) VALUES (?, ?, ?, 1)`,
		metric,
	)

	return db.Exec(ctx, query, postID, userID, date)
}

func addMultiplePostStats(ctx context.Context, postID, userID int32, viewsN, likesN, commentsN int) error {
	for i := 0; i < viewsN; i++ {
		if err := addPostStats(ctx, postID, userID, time.Now(), dto.MetricViews); err != nil {
			return err
		}
	}

	for i := 0; i < likesN; i++ {
		if err := addPostStats(ctx, postID, userID, time.Now(), dto.MetricLikes); err != nil {
			return err
		}
	}

	for i := 0; i < commentsN; i++ {
		if err := addPostStats(ctx, postID, userID, time.Now(), dto.MetricComments); err != nil {
			return err
		}
	}

	return nil
}
