package usecase

import (
	"github.com/LeonidS635/soa/internal/services/user/dto"
)

type userStorage interface {
	AddUser(user *dto.RegistrationData) (int, error)
	CheckUser(user *dto.LoginData) (int, error)
	GetUserProfile(userId int) (*dto.Profile, error)
	UpdateUserProfile(userId int, newProfile *dto.Profile) error
}

type UserUseCase struct {
	storage userStorage
}

func NewUserUseCase(storage userStorage) *UserUseCase {
	return &UserUseCase{storage: storage}
}
