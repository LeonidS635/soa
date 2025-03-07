package usecase

import (
	"errors"
	"testing"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/LeonidS635/soa/internal/services/user/usecase/mocks"
)

func TestUserUseCase_SignIn(t *testing.T) {
	mockStorage := mocks.NewMockStorage()
	uc := NewUserUseCase(mockStorage)

	tests := []struct {
		name           string
		credentials    dto.LoginData
		mockStorage    mocks.MockStorage
		expectedUserId int
		wantErr        bool
	}{
		{"empty", dto.LoginData{}, mocks.MockStorage{}, 0, true},
		{
			"incorrect login data",
			dto.LoginData{
				Username: "123",
				Password: "123456",
			},
			mocks.MockStorage{}, 0, true,
		},
		{
			"storage error",
			dto.LoginData{
				Username: "correct",
				Password: "123456",
			},
			mocks.MockStorage{
				UserId: 0, UserProfile: nil, Err: errors.New("mock error"),
			}, 0, true,
		},
		{
			"correct",
			dto.LoginData{
				Username: "correct",
				Password: "123456",
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

				userId, err := uc.SignIn(&tt.credentials)
				if (err != nil) != tt.wantErr {
					t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr && userId != tt.expectedUserId {
					t.Errorf("SignIn() got = %v, want %v", userId, tt.expectedUserId)
				}
			},
		)
	}
}
