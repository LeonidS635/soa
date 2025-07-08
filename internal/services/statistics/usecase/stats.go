package usecase

import (
	"context"
	"time"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

type (
	statisticsStorage interface {
		AddEvent(ctx context.Context, postID, userID int32, date time.Time, eventType int32) error

		GetPostStats(ctx context.Context, postID int32) (dto.PostStats, error)
		GetPostDynamics(ctx context.Context, postID, by int32) (dto.PostDynamics, error)
		GetTopPosts(ctx context.Context, by int32) (dto.Top, error)
		GetTopUsers(ctx context.Context, by int32) (dto.Top, error)
	}

	eventsConsumer interface {
		Publish(ctx context.Context, topic string, event events.Event) error
	}
)

type StatsUseCase struct {
	storage statisticsStorage
}

func NewStatsUseCase(storage statisticsStorage) *StatsUseCase {
	return &StatsUseCase{storage: storage}
}
