package tests

import (
	"context"
	"math/rand/v2"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

func addPost(ctx context.Context, post dto.Post) (int32, int32, error) {
	var postID int32
	authorID := rand.Int32()

	err := db.QueryRow(
		ctx,
		`INSERT INTO posts (is_private, title, tags, text) VALUES ($1, $2, $3, $4) RETURNING post_id`,
		post.IsPrivate, post.Title, post.Tags, post.Text,
	).Scan(&postID)
	if err != nil {
		return -1, -1, err
	}

	_, err = db.Exec(
		ctx,
		`INSERT INTO posts_info (post_id, author_id) VALUES ($1, $2)`,
		postID, authorID,
	)
	if err != nil {
		return -1, -1, err
	}

	return postID, authorID, nil
}
