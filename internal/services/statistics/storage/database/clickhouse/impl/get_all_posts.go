package impl

import (
	"context"

	"github.com/LeonidS635/soa/internal/services/statistics/storage/dto"
)

const (
	views = iota
	likes
	comments
)

func (c *ClickHouseStatisticsStorageImpl) GetQuantity(ctx context.Context, postID int32) (dto.PostStats, error) {
	postStats := dto.PostStats{}
	err := c.db.QueryRow(
		ctx,
		`SELECT count(views), count(likes), count(comments) FROM daily_stats
        WHERE post_id = $1`,
		postID,
	).ScanStruct(&postStats)

	return postStats, err
}

func (c *ClickHouseStatisticsStorageImpl) GetDynamics(ctx context.Context, postID int32) (dto.PostDynamics, error) {
	postDynamics := dto.PostDynamics{}
	err := c.db.QueryRow(
		ctx,
		`SELECT post_date, count(views), count(likes), count(comments) FROM daily_stats
        WHERE post_id = $1
        GROUP BY post_date
        ORDER BY post_date DESC`,
		postID,
	).ScanStruct(&postDynamics)

	return postDynamics, err
}

func (c *ClickHouseStatisticsStorageImpl) GetTopPosts(ctx context.Context, postID, by int32) ([]int32, error) {
	var postsIDs []int32

	switch by {
	case views:
		err := c.db.QueryRow(
			ctx,
			`SELECT post_id FROM daily_stats
        GROUP BY post_date
        ORDER BY count(views) DESC`,
			postID,
		).Scan(&postsIDs)
		return postsIDs, err
	}
}
