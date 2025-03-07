package usecase

import "github.com/LeonidS635/soa/internal/services/user/dto"

func (uc *UserUseCase) GetProfile(userId int) (*dto.Profile, error) {
	profile, err := uc.storage.GetUserProfile(userId)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
