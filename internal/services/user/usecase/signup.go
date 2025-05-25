package usecase

import (
	"context"

	"github.com/LeonidS635/soa/internal/kafka/events"
	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/LeonidS635/soa/internal/services/user/usecase/validators"
)

func (uc *UserUseCase) SignUp(credentials *dto.RegistrationData) (int, error) {
	err := validators.RegistrationData(credentials)
	if err != nil {
		return -1, err
	}

	credentials.ChangePasswordWithHash()

	userId, err := uc.storage.AddUser(credentials)
	if err != nil {
		return -1, err
	}

	event, err := events.UserRegistered(int32(userId))
	if err != nil {
		return -1, err
	}
	if err = uc.eventsProducer.Publish(context.TODO(), event); err != nil {
		return -1, err
	}
	return userId, nil
}
