package mocks

import (
	"github.com/LeonidS635/soa/internal/services/user/dto"
)

type MockStorage struct {
	UserId      int
	UserProfile *dto.Profile
	Err         error
}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func (m *MockStorage) AddUser(user *dto.RegistrationData) (int, error) {
	return m.UserId, m.Err
}

func (m *MockStorage) CheckUser(user *dto.LoginData) (int, error) {
	return m.UserId, m.Err
}

func (m *MockStorage) GetUserProfile(userId int) (*dto.Profile, error) {
	return m.UserProfile, m.Err
}

func (m *MockStorage) UpdateUserProfile(userId int, newProfile *dto.Profile) error {
	return m.Err
}
