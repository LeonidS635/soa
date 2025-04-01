package impl

import (
	"context"
	"errors"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/jackc/pgx/v5"
)

func (pg *PgPostsStorageImpl) GetPost(ctx context.Context, postId, userId int32) (*dto.PostFullInfo, error) {
	postServiceInfo, err := pg.getPostsServiceInfo(ctx, postId)
	if err != nil {
		return nil, err
	}

	post := &dto.Post{}
	err = pg.db.QueryRow(
		ctx,
		`SELECT is_private, title, tags, text FROM posts WHERE post_id = $1`,
		postId,
	).Scan(post.IsPrivate, post.Title, post.Tags, post.Text)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = PostNotFoundError
		}
		return nil, err
	}

	if post.IsPrivate && postServiceInfo.AuthorId != userId {
		return nil, PostAccessError
	}

	return &dto.PostFullInfo{
		Post:            post,
		PostServiceInfo: postServiceInfo,
	}, err
}
