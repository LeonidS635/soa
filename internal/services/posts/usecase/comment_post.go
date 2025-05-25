package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/kafka/events"
)

func (uc *PostsUseCase) CommentPost(postId, userId int32) error {
	event := events.PostCommented(userId, postId)
	if err := uc.eventsProducer.Publish(context.TODO(), "posts-comments", event); err != nil {
		return err
	}
	return nil
}
