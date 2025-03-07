package usecase

import (
	"errors"
	"testing"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/LeonidS635/soa/internal/services/user/usecase/mocks"
)

func TestUserUseCase_SignUp(t *testing.T) {
	mockStorage := mocks.NewMockStorage()
	uc := NewUserUseCase(mockStorage)

	tests := []struct {
		name           string
		credentials    dto.RegistrationData
		mockStorage    mocks.MockStorage
		expectedUserId int
		wantErr        bool
	}{
		{"empty", dto.RegistrationData{}, mocks.MockStorage{}, 0, true},
		{
			"incorrect registration data",
			dto.RegistrationData{
				LoginData: dto.LoginData{
					Username: "123",
					Password: "123456",
				}, Email: "some@email",
			},
			mocks.MockStorage{}, 0, true,
		},
		{
			"storage error",
			dto.RegistrationData{
				LoginData: dto.LoginData{
					Username: "correct",
					Password: "123456",
				}, Email: "some@email",
			},
			mocks.MockStorage{
				UserId: 0, UserProfile: nil, Err: errors.New("mock error"),
			}, 0, true,
		},
		{
			"correct",
			dto.RegistrationData{
				LoginData: dto.LoginData{
					Username: "correct",
					Password: "123456",
				}, Email: "some@email",
			},
			mocks.MockStorage{
				UserId: 0, UserProfile: nil, Err: nil,
			}, 0, false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockStorage.UserId = tt.mockStorage.UserId
				mockStorage.UserProfile = tt.mockStorage.UserProfile
				mockStorage.Err = tt.mockStorage.Err

				userId, err := uc.SignUp(&tt.credentials)
				if (err != nil) != tt.wantErr {
					t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr && userId != tt.expectedUserId {
					t.Errorf("SignUp() got = %v, want %v", userId, tt.expectedUserId)
				}
			},
		)
	}
}
