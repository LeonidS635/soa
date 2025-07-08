package usecase

import (
	"context"
	"time"

	"github.com/LeonidS635/soa/internal/kafka/events"
)

func (uc *PostsUseCase) CommentPost(postID, userID int32) error {
	event := &events.PostCommentedEvent{
		Moment:    time.Now(),
		PostID:    postID,
		CommentID: 0,
		UsedID:    userID,
	}
	if err := uc.eventsProducer.Publish(context.TODO(), "posts-comments", event); err != nil {
		return err
	}
	return nil
}
