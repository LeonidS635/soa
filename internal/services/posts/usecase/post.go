package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

type (
	postsStorage interface {
		AddPost(userId int32, post *dto.Post) (int32, error)
		GetPost(postId, userId int32) (*dto.PostFullInfo, error)
		GetAllPosts(page, postsPerPageN, userId int32, onlyAuthor bool) (
			[]*dto.PostFullInfo, error,
		)
		UpdatePost(postId, userId int32, newPost *dto.Post) error
		DeletePost(postId, userId int32) error
	}

	eventsProducer interface {
		Publish(ctx context.Context, event events.Event) error
	}
)

type PostsUseCase struct {
	storage        postsStorage
	eventsProducer eventsProducer
}

func NewPostsUseCase(storage postsStorage, producer eventsProducer) *PostsUseCase {
	return &PostsUseCase{storage: storage, eventsProducer: producer}
}
