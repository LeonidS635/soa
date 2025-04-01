package usecase

import (
	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

type postsStorage interface {
	AddPost(userId int32, post *dto.Post) (int32, error)
	GetPost(postId, userId int32) (*dto.PostFullInfo, error)
	GetAllPosts(postsPerPageN, userId int32, onlyAuthor bool) (
		[][]*dto.PostFullInfo, error,
	)
	UpdatePost(postId, userId int32, newPost *dto.Post) error
	DeletePost(postId, userId int32) error
}

type PostsUseCase struct {
	storage postsStorage
}

func NewPostsUseCase(storage postsStorage) *PostsUseCase {
	return &PostsUseCase{storage: storage}
}
