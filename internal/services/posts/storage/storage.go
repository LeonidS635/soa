package storage

import (
	"context"

	"github.com/LeonidS635/soa/internal/services/posts/dto"
	"github.com/LeonidS635/soa/internal/services/posts/storage/database/postgres/impl"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storageImpl interface {
	AddPost(ctx context.Context, userId int32, post *dto.Post) (int32, error)
	GetPost(ctx context.Context, postId, userId int32) (*dto.PostFullInfo, error)
	GetAllPosts(ctx context.Context, page, postsPerPageN, userId int32, onlyAuthor bool) (
		[]*dto.PostFullInfo, error,
	)
	UpdatePost(ctx context.Context, postId, userId int32, newPost *dto.Post) error
	DeletePost(ctx context.Context, postId, userId int32) error

	Close()
}

type PostsStorage struct {
	ctx   context.Context
	impl_ storageImpl
}

func NewPostsStorage(ctx context.Context, pool *pgxpool.Pool) (*PostsStorage, error) {
	storage, err := impl.NewPgPostsStorageImpl(pool)
	if err != nil {
		return nil, err
	}

	return &PostsStorage{ctx, storage}, nil
}

func (s *PostsStorage) AddPost(userId int32, post *dto.Post) (int32, error) {
	return s.impl_.AddPost(s.ctx, userId, post)
}

func (s *PostsStorage) GetPost(postId, userId int32) (*dto.PostFullInfo, error) {
	return s.impl_.GetPost(s.ctx, postId, userId)
}

func (s *PostsStorage) GetAllPosts(page, postsPerPageN, userId int32, onlyAuthor bool) ([]*dto.PostFullInfo, error) {
	return s.impl_.GetAllPosts(s.ctx, page, postsPerPageN, userId, onlyAuthor)
}

func (s *PostsStorage) UpdatePost(postId, userId int32, newPost *dto.Post) error {
	return s.impl_.UpdatePost(s.ctx, postId, userId, newPost)
}

func (s *PostsStorage) DeletePost(postId, userId int32) error {
	return s.impl_.DeletePost(s.ctx, postId, userId)
}
