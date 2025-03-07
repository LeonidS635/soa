package validators

import (
	"errors"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func LoginData(credentials *dto.LoginData) error {
	if len(credentials.Username) < 4 {
		return errors.New("username is too short")
	}
	if len(credentials.Username) > 64 {
		return errors.New("username is too long")
	}

	if len(credentials.Password) < 4 {
		return errors.New("password is too short")
	}
	if len(credentials.Password) > 64 {
		return errors.New("password is too long")
	}

	return nil
}
