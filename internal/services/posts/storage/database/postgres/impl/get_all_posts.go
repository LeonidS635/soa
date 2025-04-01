package impl

import (
	"context"
	"errors"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/jackc/pgx/v5"
)

func (pg *PgPostsStorageImpl) GetAllPosts(
	ctx context.Context, postsPerPageN, userId int32, onlyAuthor bool,
) ([][]*dto.PostFullInfo, error) {
	rows, err := pg.db.Query(
		ctx,
		`SELECT post_id, created_at, updated_at FROM posts_info WHERE author_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts [][]*dto.PostFullInfo
	for pagesCounter, postsCounter := int32(0), int32(0); rows.Next(); postsCounter++ {
		if postsCounter == postsPerPageN-1 {
			postsCounter = 0
			pagesCounter++
		}

		post := &dto.PostFullInfo{}
		if err = rows.Scan(post.PostId, post.CreatedAt, post.UpdatedAt); err != nil {
			return nil, err
		}

		err = pg.db.QueryRow(
			ctx,
			`SELECT is_private, title, tags, text FROM posts WHERE post_id = $1`,
			post.PostId,
		).Scan(post.IsPrivate, post.Title, post.Tags, post.Text)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				err = PostNotFoundError
			}
			return nil, err
		}

		if !post.IsPrivate || !onlyAuthor || post.AuthorId != userId {
			posts[pagesCounter] = append(posts[pagesCounter], post)
		}
	}
	if rows.Err() != nil {
		return nil, err
	}

	return posts, nil
}
