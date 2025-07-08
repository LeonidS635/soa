package service

import (
	"context"
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/stretchr/testify/require"
)

func TestStatisticsService_GetTopPosts(t *testing.T) {
	t.Cleanup(
		func() {
			if err := clearTables(context.Background()); err != nil {
				t.Fatal(err)
			}
		},
	)

	const viewsN, likesN, commentsN, postsN = 3, 2, 1, 10

	for i := 1; i < postsN+1; i++ {
		if err := addMultiplePostStats(
			t.Context(), int32(i), int32(i)+123, viewsN*i, likesN*i, commentsN*i,
		); err != nil {
			t.Fatal(err)
		}
	}

	req := &statspb.GetTopPostsRequest{
		Metric: statspb.Metric_VIEWS,
	}
	resp, err := client.GetTopPosts(t.Context(), req)
	require.NoError(t, err)

	require.Len(t, resp.Posts, 10)
	for i, post := range resp.Posts {
		require.Equal(t, postsN-int32(i), post.PostId)
		require.Equal(t, viewsN*(postsN-uint64(i)), post.MetricValue)
	}
}
