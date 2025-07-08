package service

import (
	"context"
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/stretchr/testify/require"
)

func TestStatisticsService_GetPostStats(t *testing.T) {
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

	for i := 1; i < postsN+1; i++ {
		req := &statspb.GetPostStatsRequest{
			PostId: int32(i),
		}
		resp, err := client.GetPostStats(t.Context(), req)
		require.NoError(t, err)

		require.Equal(t, uint64(viewsN*i), resp.Views)
		require.Equal(t, uint64(likesN*i), resp.Likes)
		require.Equal(t, uint64(commentsN*i), resp.Comments)
	}
}
