package storage

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
	"github.com/LeonidS635/soa/internal/services/statistics/storage/database/clickhouse/impl"
)

type storageImpl interface {
	AddEvent(ctx context.Context, postID, userID int32, date time.Time, eventType int32) error

	GetPostStats(ctx context.Context, postID int32) (dto.PostStats, error)
	GetPostDynamics(ctx context.Context, postID, by int32) (dto.PostDynamics, error)
	GetTop(ctx context.Context, postOrAuthor, by int32) (dto.Top, error)

	Close()
}

type StatisticsStorage struct {
	impl_ storageImpl
}

func NewStatisticsStorage(conn driver.Conn) (*StatisticsStorage, error) {
	storage, err := impl.NewClickHouseStatisticsStorageImpl(conn)
	if err != nil {
		return nil, err
	}

	return &StatisticsStorage{storage}, nil
}

func (s StatisticsStorage) AddEvent(ctx context.Context, postID, userID int32, date time.Time, eventType int32) error {
	return s.impl_.AddEvent(ctx, postID, userID, date, eventType)
}

func (s StatisticsStorage) GetPostStats(ctx context.Context, postID int32) (dto.PostStats, error) {
	return s.impl_.GetPostStats(ctx, postID)
}

func (s StatisticsStorage) GetPostDynamics(ctx context.Context, postID, by int32) (
	dto.PostDynamics, error,
) {
	return s.impl_.GetPostDynamics(ctx, postID, by)
}

func (s StatisticsStorage) GetTopPosts(ctx context.Context, by int32) (dto.Top, error) {
	return s.impl_.GetTop(ctx, impl.TopPosts, by)
}

func (s StatisticsStorage) GetTopUsers(ctx context.Context, by int32) (dto.Top, error) {
	return s.impl_.GetTop(ctx, impl.TopAuthors, by)
}
