package impl

import (
	"context"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

func (pg *PgPostsStorageImpl) AddPost(ctx context.Context, userId int32, post *dto.Post) (int32, error) {
	var postId int32
	err := pg.db.QueryRow(
		ctx,
		`INSERT INTO posts (is_private, title, tags, text) VALUES ($1, $2, $3, $4) RETURNING post_id`,
		post.IsPrivate, post.Title, post.Tags, post.Text,
	).Scan(&postId)
	if err != nil {
		return -1, err
	}

	_, err = pg.db.Exec(
		ctx,
		`INSERT INTO posts_info (post_id, author_id) VALUES ($1, $2)`,
		postId, userId,
	)
	if err != nil {
		return -1, err
	}

	return postId, err
}
