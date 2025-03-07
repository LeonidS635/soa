package validators

import (
	"errors"
	"strings"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func RegistrationData(credentials *dto.RegistrationData) error {
	err := LoginData(&credentials.LoginData)
	if err != nil {
		return err
	}

	if !strings.ContainsRune(credentials.Email, '@') {
		return errors.New("email address is not valid")
	}

	return nil
}
