package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/kafka/events"
)

func (uc *PostsUseCase) LikePost(postId, userId int32) error {
	event := events.PostLiked(userId, postId)
	if err := uc.eventsProducer.Publish(context.TODO(), "posts-likes", event); err != nil {
		return err
	}
	return nil
}
