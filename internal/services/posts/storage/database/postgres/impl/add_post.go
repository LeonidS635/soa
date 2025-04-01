package impl

import (
	"context"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

func (pg *PgPostsStorageImpl) AddPost(ctx context.Context, userId int32, post *dto.Post) (int32, error) {
	var postId int32
	err := pg.db.QueryRow(
		ctx,
		`INSERT INTO posts_info (author_id, is_private) VALUES ($1, $2,) RETURNING post_id`,
		userId, post.IsPrivate,
	).Scan(&postId)
	if err != nil {
		return -1, err
	}

	_, err = pg.db.Exec(
		ctx,
		`INSERT INTO posts (post_id, title, tags, text) VALUES ($1, $2, $3, $4)`,
		postId, post.Title, post.Tags, post.Text,
	)
	return postId, err
}
