package usecase

import (
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
	return userId, nil
}
