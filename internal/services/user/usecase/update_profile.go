package usecase

import (
	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/LeonidS635/soa/internal/services/user/usecase/validators"
)

func (uc *UserUseCase) UpdateProfile(userId int, newProfile *dto.Profile) error {
	err := validators.ProfileData(newProfile)
	if err != nil {
		return err
	}

	err = uc.storage.UpdateUserProfile(userId, newProfile)
	if err != nil {
		return err
	}
	return nil
}
