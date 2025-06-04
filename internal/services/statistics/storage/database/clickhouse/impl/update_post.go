package impl

import (
	"context"
	"errors"
	"time"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/jackc/pgx/v5"
)

func (pg *PgPostsStorageImpl) UpdatePost(ctx context.Context, postId, userId int32, newPost *dto.Post) error {
	postServiceInfo, err := pg.getPostsServiceInfo(ctx, postId)
	if err != nil {
		return err
	}

	if postServiceInfo.AuthorId != userId {
		return PostAccessError
	}

	_, err = pg.db.Exec(
		ctx,
		`UPDATE posts SET is_private=$1, title=$2, tags=$3, text=$4 WHERE post_id=$5`,
		newPost.IsPrivate, newPost.Title, newPost.Tags, newPost.Text, postId,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = PostNotFoundError
		}
		return err
	}

	_, err = pg.db.Exec(
		ctx,
		`UPDATE posts_info SET updated_at=$1 WHERE post_id=$2`,
		time.Now(), postId,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		err = PostNotFoundError
	}
	return err
}
