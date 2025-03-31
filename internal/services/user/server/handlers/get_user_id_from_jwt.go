package handlers

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

func getUserIdFromJWT(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1, fmt.Errorf("%w: claims not provided", ClaimsError)
	}

	userIdStr, ok := claims["user_id"]
	if !ok {
		return -1, fmt.Errorf("%w: user id not in claims", ClaimsError)
	}

	userId, err := strconv.Atoi(fmt.Sprintf("%v", userIdStr))
	if err != nil {
		return -1, fmt.Errorf("%w: user id is not a number", ClaimsError)
	}

	return userId, nil
}
