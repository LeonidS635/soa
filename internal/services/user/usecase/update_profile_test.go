package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/LeonidS635/soa/internal/services/user/usecase/mocks"
)

func TestUserUseCase_UpdateProfile(t *testing.T) {
	mockStorage := mocks.NewMockStorage()
	uc := NewUserUseCase(mockStorage)

	testUserProfile := dto.Profile{
		Name:        "John",
		Surname:     "Doe",
		Age:         19,
		Email:       "some@email",
		Phone:       "1 (234) 567-89-00",
		City:        "Berlin",
		Description: "some description",
		Birthdate:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name        string
		userId      int
		mockStorage mocks.MockStorage
		wantErr     bool
	}{
		{"empty", 0, mocks.MockStorage{}, false},
		{
			"storage error",
			0,
			mocks.MockStorage{
				UserId: 0, UserProfile: nil, Err: errors.New("mock error"),
			},
			true,
		},
		{
			"correct",
			0,
			mocks.MockStorage{
				UserId: 0, UserProfile: &testUserProfile, Err: nil,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockStorage.UserId = tt.mockStorage.UserId
				mockStorage.UserProfile = tt.mockStorage.UserProfile
				mockStorage.Err = tt.mockStorage.Err

				err := uc.UpdateProfile(tt.userId, tt.mockStorage.UserProfile)
				if (err != nil) != tt.wantErr {
					t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
