package tests

import (
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/stretchr/testify/require"
)

func TestPostsService_CreatePost(t *testing.T) {
	req := &postspb.AddPostRequest{
		UserId: 123,
		Post: &postspb.Post{
			IsPrivate: false,
			Title:     "My post",
			Tags:      []string{"postTag"},
			Text:      "Text of my post",
		},
	}
	resp, err := client.AddPost(t.Context(), req)

	require.NoError(t, err)
	require.Positive(t, resp.PostId)

	post := dto.PostFullInfo{
		Post:            &dto.Post{},
		PostServiceInfo: &dto.PostServiceInfo{},
	}
	err = db.QueryRow(
		t.Context(),
		`SELECT is_private, title, tags, text FROM posts WHERE post_id = $1`,
		resp.PostId,
	).Scan(&post.IsPrivate, &post.Title, &post.Tags, &post.Text)
	if err != nil {
		t.Fatal(err)
	}

	err = db.QueryRow(
		t.Context(),
		`SELECT author_id FROM posts_info WHERE post_id = $1`,
		resp.PostId,
	).Scan(&post.AuthorId)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, req.UserId, post.AuthorId)
	require.Equal(t, req.Post.IsPrivate, post.IsPrivate)
	require.Equal(t, req.Post.Title, post.Title)
	require.Equal(t, req.Post.Tags, post.Tags)
	require.Equal(t, req.Post.Text, post.Text)
}
