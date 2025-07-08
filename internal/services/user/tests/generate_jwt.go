package tests

import "github.com/golang-jwt/jwt/v5"

func generateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(jwtPrivate)
}
