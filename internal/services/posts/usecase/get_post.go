package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

func (uc *PostsUseCase) GetPost(postId, userId int32) (*dto.PostFullInfo, error) {
	post, err := uc.storage.GetPost(postId, userId)
	if err != nil {
		return nil, err
	}

	event, err := events.PostViewed(userId, postId)
	if err != nil {
		return nil, err
	}
	if err = uc.eventsProducer.Publish(context.TODO(), event); err != nil {
		return nil, err
	}
	return post, nil
}
