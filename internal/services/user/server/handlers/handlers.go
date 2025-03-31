package handlers

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"

	"github.com/LeonidS635/soa/internal/services/user/usecase"
)

type UserHandlers struct {
	useCase    *usecase.UserUseCase
	jwtPrivate *rsa.PrivateKey
	jwtPublic  *rsa.PublicKey
}

func NewUserHandlers(useCase *usecase.UserUseCase, jwtPrivateFile string, jwtPublicFile string) (*UserHandlers, error) {
	private, err := os.ReadFile(jwtPrivateFile)
	if err != nil {
		log.Fatalf("error while reading jwt private file: %v", err)
	}
	public, err := os.ReadFile(jwtPublicFile)
	if err != nil {
		log.Fatalf("error while reading jwt public file: %v", err)
	}
	jwtPrivate, err := jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		log.Fatalf("error while parsing jwt private file: %v", err)
	}
	jwtPublic, err := jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		log.Fatalf("error while parsing jwt private file: %v", err)
	}

	return &UserHandlers{
		useCase:    useCase,
		jwtPrivate: jwtPrivate,
		jwtPublic:  jwtPublic,
	}, nil
}
