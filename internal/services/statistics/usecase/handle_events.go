package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

func (uc *StatsUseCase) PostViewed(ctx context.Context, eventData []byte) error {
	var event events.PostViewedEvent
	if err := event.Unmarshal(eventData); err != nil {
		return err
	}

	return uc.storage.AddEvent(ctx, event.PostID, event.UsedID, event.Moment, dto.MetricViews)
}

func (uc *StatsUseCase) PostLiked(ctx context.Context, eventData []byte) error {
	var event events.PostLikedEvent
	if err := event.Unmarshal(eventData); err != nil {
		return err
	}

	return uc.storage.AddEvent(ctx, event.PostID, event.UsedID, event.Moment, dto.MetricLikes)
}

func (uc *StatsUseCase) PostCommented(ctx context.Context, eventData []byte) error {
	var event events.PostCommentedEvent
	if err := event.Unmarshal(eventData); err != nil {
		return err
	}

	return uc.storage.AddEvent(ctx, event.PostID, event.UsedID, event.Moment, dto.MetricComments)
}
