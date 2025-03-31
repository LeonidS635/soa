package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/LeonidS635/soa/internal/services/user/dto"
	"github.com/LeonidS635/soa/internal/services/user/usecase/mocks"
)

func TestUserUseCase_GetProfile(t *testing.T) {
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
		name                string
		userId              int
		mockStorage         mocks.MockStorage
		expectedUserProfile *dto.Profile
		wantErr             bool
	}{
		{"empty", 0, mocks.MockStorage{}, nil, false},
		{
			"storage error",
			0,
			mocks.MockStorage{
				UserId: 0, UserProfile: nil, Err: errors.New("mock error"),
			},
			nil,
			true,
		},
		{
			"correct",
			0,
			mocks.MockStorage{
				UserId: 0, UserProfile: &testUserProfile, Err: nil,
			},
			&testUserProfile,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockStorage.UserId = tt.mockStorage.UserId
				mockStorage.UserProfile = tt.mockStorage.UserProfile
				mockStorage.Err = tt.mockStorage.Err

				userProfile, err := uc.GetProfile(tt.userId)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetProfile() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if !reflect.DeepEqual(userProfile, tt.expectedUserProfile) {
					t.Errorf("GetProfile() got = %v, want %v", userProfile, tt.expectedUserProfile)
				}
			},
		)
	}
}
