package usecase

import (
	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/LeonidS635/soa/internal/services/user/usecase/validators"
)

func (uc *UserUseCase) SignIn(credentials *dto.LoginData) (int, error) {
	err := validators.LoginData(credentials)
	if err != nil {
		return -1, err
	}

	credentials.ChangePasswordWithHash()

	userId, err := uc.storage.CheckUser(credentials)
	if err != nil {
		return -1, err
	}
	return userId, nil
}
