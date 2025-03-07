package storage

import (
	"context"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/LeonidS635/soa/internal/services/user/storage/database/postgres/impl"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storageImpl interface {
	AddUser(ctx context.Context, user *dto.RegistrationData) (int, error)
	CheckUser(ctx context.Context, user *dto.LoginData) (int, error)
	GetUserProfile(ctx context.Context, userId int) (*dto.Profile, error)
	UpdateUserProfile(ctx context.Context, userId int, newProfile *dto.Profile) error
	Close()
}

type UserStorage struct {
	ctx   context.Context
	impl_ storageImpl
}

func NewUserStorage(ctx context.Context, pool *pgxpool.Pool) (*UserStorage, error) {
	storage, err := impl.NewPgUserStorageImpl(ctx, pool)
	if err != nil {
		return nil, err
	}

	return &UserStorage{ctx, storage}, nil
}

func (s *UserStorage) AddUser(credentials *dto.RegistrationData) (int, error) {
	return s.impl_.AddUser(s.ctx, credentials)
}

func (s *UserStorage) CheckUser(user *dto.LoginData) (int, error) {
	return s.impl_.CheckUser(s.ctx, user)
}

func (s *UserStorage) GetUserProfile(userId int) (*dto.Profile, error) {
	return s.impl_.GetUserProfile(s.ctx, userId)
}

func (s *UserStorage) UpdateUserProfile(userId int, newProfile *dto.Profile) error {
	return s.impl_.UpdateUserProfile(s.ctx, userId, newProfile)
}
