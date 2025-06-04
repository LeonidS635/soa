package usecase

import (
	"context"
	"time"

	"github.com/LeonidS635/soa/internal/kafka/events"
)

func (uc *PostsUseCase) LikePost(postID, userID int32) error {
	event := &events.PostLikedEvent{
		Moment: time.Now(),
		PostID: postID,
		UsedID: userID,
	}
	if err := uc.eventsProducer.Publish(context.TODO(), "posts-likes", event); err != nil {
		return err
	}
	return nil
}
