package handlers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func (h *UserHandlers) getJWT(req *http.Request) (*jwt.Token, error) {
	tokenCookie, err := req.Cookie("jwt")
	if err != nil {
		return nil, fmt.Errorf("%w: token not provided", JWTError)
	}
	token, err := jwt.Parse(
		tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("%w: unexpected signing method (expected RSA)", SigningError)
			}
			return h.jwtPublic, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", JWTError, err)
	}

	return token, nil
}
