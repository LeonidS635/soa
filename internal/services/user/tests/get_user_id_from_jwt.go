package tests

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func getUserIDFromJWT(cookies []*http.Cookie) (int, error) {
	var jwtToken string
	for _, cookie := range cookies {
		if cookie.Name == "jwt" {
			jwtToken = cookie.Value
			break
		}
	}

	if jwtToken == "" {
		return 0, fmt.Errorf("JWT coockie not found in response")
	}

	token, err := jwt.Parse(
		jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtPublic, nil
		},
	)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid JWT claims")
	}

	userID, exists := claims["user_id"]
	if !exists {
		return 0, fmt.Errorf("user_id not found in JWT claims")
	}

	return int(userID.(float64)), nil
}
