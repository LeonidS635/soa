package validators

import (
	"strings"
	"testing"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func TestLoginData(t *testing.T) {
	tests := []struct {
		name        string
		credentials dto.LoginData
		wantErr     bool
	}{
		{"empty", dto.LoginData{}, true},
		{
			"short username",
			dto.LoginData{
				Username: "123",
				Password: "123456",
			}, true,
		},
		{
			"short password",
			dto.LoginData{
				Username: "correct",
				Password: "123",
			}, true,
		},
		{
			"long username",
			dto.LoginData{
				Username: strings.Repeat("too_long", 10),
				Password: "123",
			}, true,
		},
		{
			"long password",
			dto.LoginData{
				Username: "correct",
				Password: strings.Repeat("too_long", 10),
			}, true,
		},
		{
			"correct",
			dto.LoginData{
				Username: "correct",
				Password: "correct",
			}, false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := LoginData(&tt.credentials)
				if (err != nil) != tt.wantErr {
					t.Errorf("LoginData() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
