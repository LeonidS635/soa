package impl

import (
	"context"

	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

func (c *ClickHouseStatisticsStorageImpl) GetPostStats(ctx context.Context, postID int32) (dto.PostStats, error) {
	var postStats dto.PostStats
	err := c.db.QueryRow(
		ctx,
		`SELECT sum(views), sum(likes), sum(comments) FROM daily_stats
        WHERE post_id = $1`,
		postID,
	).Scan(&postStats.Views, &postStats.Likes, &postStats.Comments)

	return postStats, err
}
