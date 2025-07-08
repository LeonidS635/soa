package service

import (
	"context"
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/stretchr/testify/require"
)

func TestStatisticsService_IncorrectRequests(t *testing.T) {
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

	testCases := []struct {
		name   string
		method func() error
	}{
		{
			name: "get dynamics with an unknown metric",
			method: func() error {
				_, err := client.GetPostDynamics(
					t.Context(), &statspb.GetPostDynamicsRequest{PostId: 1, Metric: -1},
				)
				return err
			},
		},
		{
			name: "get top with an unknown metric",
			method: func() error {
				_, err := client.GetTopPosts(t.Context(), &statspb.GetTopPostsRequest{Metric: 10})
				return err
			},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				require.Error(t, tc.method())
			},
		)
	}
}
