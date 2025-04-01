package impl

import (
	"context"
	"errors"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/jackc/pgx/v5"
)

func (pg *PgPostsStorageImpl) getPostsServiceInfo(ctx context.Context, postId int32) (
	*dto.PostServiceInfo, error,
) {
	postServiceInfo := &dto.PostServiceInfo{}
	err := pg.db.QueryRow(
		ctx,
		`SELECT post_id, author_id, created_at, updated_at FROM posts_info WHERE post_id = $1`,
		postId,
	).Scan(postServiceInfo.PostId, postServiceInfo.AuthorId, postServiceInfo.CreatedAt, postServiceInfo.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = PostNotFoundError
		}
		return nil, err
	}

	return postServiceInfo, nil
}
