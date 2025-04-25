package impl

import (
	"context"
	"errors"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/jackc/pgx/v5"
)

func (pg *PgPostsStorageImpl) GetAllPosts(
	ctx context.Context, page, postsPerPageN, userId int32, onlyAuthor bool,
) ([]*dto.PostFullInfo, error) {
	rows, err := pg.db.Query(
		ctx,
		`SELECT post_id, author_id, created_at, updated_at FROM posts_info
        WHERE NOT $1 OR author_id = $2
        ORDER BY updated_at DESC
        LIMIT $3 OFFSET $4`,
		onlyAuthor, userId, postsPerPageN, (page-1)*postsPerPageN,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]*dto.PostFullInfo, 0, postsPerPageN)
	for postsCounter := 0; rows.Next(); postsCounter++ {
		post := dto.PostFullInfo{
			Post:            &dto.Post{},
			PostServiceInfo: &dto.PostServiceInfo{},
		}
		if err = rows.Scan(&post.PostId, &post.AuthorId, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}

		err = pg.db.QueryRow(
			ctx,
			`SELECT is_private, title, tags, text FROM posts WHERE post_id = $1`,
			post.PostId,
		).Scan(&post.IsPrivate, &post.Title, &post.Tags, &post.Text)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				err = PostNotFoundError
			}
			return nil, err
		}

		if !post.IsPrivate || !onlyAuthor || post.AuthorId != userId {
			posts = append(posts, &post)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
