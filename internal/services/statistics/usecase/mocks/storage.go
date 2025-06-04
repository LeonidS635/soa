package mocks

import (
	"context"
	"time"

	"github.com/LeonidS635/soa/internal/services/statistics/dto"
)

type StatsStorage struct {
	Stats      dto.PostStats
	Dynamics   dto.PostDynamics
	TopPosts   dto.Top
	TopAuthors dto.Top

	Err error
}

func NewStatsStorage() *StatsStorage {
	return &StatsStorage{}
}

func (s *StatsStorage) AddEvent(ctx context.Context, postID, userID int32, date time.Time, eventType int32) error {
	return s.Err
}

func (s *StatsStorage) GetPostStats(ctx context.Context, postID int32) (dto.PostStats, error) {
	return s.Stats, s.Err
}

func (s *StatsStorage) GetPostDynamics(ctx context.Context, postID, by int32) (dto.PostDynamics, error) {
	return s.Dynamics, s.Err
}

func (s *StatsStorage) GetTopPosts(ctx context.Context, by int32) (dto.Top, error) {
	return s.TopPosts, s.Err
}

func (s *StatsStorage) GetTopUsers(ctx context.Context, by int32) (dto.Top, error) {
	return s.TopAuthors, s.Err
}
