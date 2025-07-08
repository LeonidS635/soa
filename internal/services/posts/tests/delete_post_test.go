package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestPostsService_DeletePost(t *testing.T) {
	const N = 2

	posts := make(
		[]struct {
			authorID int32
			postID   int32
			post     dto.Post
		}, N,
	)
	for i := 0; i < N; i++ {
		posts[i].post = dto.Post{
			IsPrivate: i&1 == 0,
			Title:     fmt.Sprintf("Title (%d)", i),
			Tags:      []string{fmt.Sprintf("Tag (%d)", i), fmt.Sprintf("Tag (%d)", i+1)},
			Text:      fmt.Sprintf("Text (%d)", i),
		}

		var err error
		posts[i].postID, posts[i].authorID, err = addPost(context.Background(), posts[i].post)
		if err != nil {
			t.Fatal(err)
		}
	}

	req := &postspb.DeletePostRequest{
		UserId: posts[0].authorID,
		PostId: posts[0].postID,
	}
	_, err := client.DeletePost(t.Context(), req)
	require.NoError(t, err)

	err = db.QueryRow(t.Context(), `SELECT FROM posts_info WHERE post_id = $1`, posts[0].postID).Scan()
	require.Error(t, err)
	require.Equal(t, pgx.ErrNoRows, err)

	err = db.QueryRow(t.Context(), `SELECT FROM posts WHERE post_id = $1`, posts[0].postID).Scan()
	require.Error(t, err)
	require.Equal(t, pgx.ErrNoRows, err)

	err = db.QueryRow(t.Context(), `SELECT FROM posts_info WHERE post_id = $1`, posts[1].postID).Scan()
	require.NoError(t, err)

	err = db.QueryRow(t.Context(), `SELECT FROM posts WHERE post_id = $1`, posts[1].postID).Scan()
	require.NoError(t, err)
}
