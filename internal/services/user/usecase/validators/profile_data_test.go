package validators

import (
	"testing"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func TestProfileData(t *testing.T) {
	tests := []struct {
		name    string
		profile dto.Profile
		wantErr bool
	}{
		{"empty", dto.Profile{}, false},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := ProfileData(&tt.profile)
				if (err != nil) != tt.wantErr {
					t.Errorf("ProfileData() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
