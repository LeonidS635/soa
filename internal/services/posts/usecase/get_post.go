package usecase

import (
	"context"
	"time"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/LeonidS635/soa/internal/services/posts/dto"
)

func (uc *PostsUseCase) GetPost(postID, userID int32) (*dto.PostFullInfo, error) {
	post, err := uc.storage.GetPost(postID, userID)
	if err != nil {
		return nil, err
	}

	event := &events.PostViewedEvent{
		Moment: time.Now(),
		PostID: postID,
		UsedID: userID,
	}
	if err = uc.eventsProducer.Publish(context.TODO(), "posts-views", event); err != nil {
		return nil, err
	}
	return post, nil
}
