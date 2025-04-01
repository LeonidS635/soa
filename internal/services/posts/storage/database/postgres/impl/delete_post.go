package impl

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (pg *PgPostsStorageImpl) DeletePost(ctx context.Context, postId, userId int32) error {
	postServiceInfo, err := pg.getPostsServiceInfo(ctx, postId)
	if err != nil {
		return err
	}

	if postServiceInfo.AuthorId != userId {
		return PostAccessError
	}

	_, err = pg.db.Exec(
		ctx,
		`DELETE FROM posts WHERE post_id=$1`,
		postId,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = PostNotFoundError
		}
		return err
	}

	_, err = pg.db.Exec(
		ctx,
		`DELETE FROM posts_info WHERE post_id=$1`,
		postId,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		err = PostNotFoundError
	}
	return err
}
