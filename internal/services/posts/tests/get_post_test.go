package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/stretchr/testify/require"
)

func TestPostsService_GetPost(t *testing.T) {
	const N = 10

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

	for i := 0; i < N; i++ {
		req := &postspb.GetPostRequest{
			UserId: posts[i].authorID,
			PostId: posts[i].postID,
		}
		resp, err := client.GetPost(t.Context(), req)
		require.NoError(t, err)

		require.Equal(t, posts[i].authorID, resp.Post.Details.AuthorId)
		require.Equal(t, posts[i].postID, resp.Post.Details.PostId)

		require.Equal(t, posts[i].post.IsPrivate, resp.Post.Post.IsPrivate)
		require.Equal(t, posts[i].post.Title, resp.Post.Post.Title)
		require.Equal(t, posts[i].post.Tags, resp.Post.Post.Tags)
		require.Equal(t, posts[i].post.Text, resp.Post.Post.Text)
	}
}
