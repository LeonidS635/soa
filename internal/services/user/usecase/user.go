package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/LeonidS635/soa/internal/services/user/dto"
)

type (
	userStorage interface {
		AddUser(user *dto.RegistrationData) (int, error)
		CheckUser(user *dto.LoginData) (int, error)
		GetUserProfile(userId int) (*dto.Profile, error)
		UpdateUserProfile(userId int, newProfile *dto.Profile) error
	}

	eventsProducer interface {
		Publish(ctx context.Context, event events.Event) error
	}
)

type UserUseCase struct {
	storage        userStorage
	eventsProducer eventsProducer
}

func NewUserUseCase(storage userStorage, producer eventsProducer) *UserUseCase {
	return &UserUseCase{storage: storage, eventsProducer: producer}
}
