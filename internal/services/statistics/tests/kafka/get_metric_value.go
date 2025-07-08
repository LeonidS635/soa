package kafka

import (
	"context"
	"fmt"
)

func getMetricValue(ctx context.Context, postID int32, metric string) (uint64, error) {
	var metricValue uint64
	err := db.QueryRow(
		ctx,
		fmt.Sprintf(`SELECT sum(%s) FROM daily_stats WHERE post_id = $1`, metric),
		postID,
	).Scan(&metricValue)
	if err != nil {
		return 0, err
	}
	return metricValue, nil
}
