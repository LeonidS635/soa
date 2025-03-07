package validators

import (
	"testing"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func TestRegistrationData(t *testing.T) {
	tests := []struct {
		name        string
		credentials dto.RegistrationData
		wantErr     bool
	}{
		{"empty", dto.RegistrationData{}, true},
		{
			"incorrect registration data",
			dto.RegistrationData{
				LoginData: dto.LoginData{
					Username: "123",
					Password: "123456",
				}, Email: "some@email",
			}, true,
		},
		{
			"incorrect email",
			dto.RegistrationData{
				LoginData: dto.LoginData{
					Username: "correct",
					Password: "123456",
				}, Email: "some_email",
			}, true,
		},
		{
			"correct",
			dto.RegistrationData{
				LoginData: dto.LoginData{
					Username: "correct",
					Password: "123456",
				}, Email: "some@email",
			}, false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := RegistrationData(&tt.credentials)
				if (err != nil) != tt.wantErr {
					t.Errorf("RegistrationData() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
