package handlers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func (h *UserHandlers) setJWT(w http.ResponseWriter, userId int) error {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"user_id": userId})
	signedToken, err := token.SignedString(h.jwtPrivate)
	if err != nil {
		return fmt.Errorf("%w: %w", SigningError, err)
	}

	http.SetCookie(
		w, &http.Cookie{
			Name:  "jwt",
			Value: signedToken,
		},
	)

	return nil
}
