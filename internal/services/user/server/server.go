package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/LeonidS635/soa/internal/services/user/server/handlers"
	"github.com/LeonidS635/soa/internal/services/user/storage"
	"github.com/LeonidS635/soa/internal/services/user/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	jwtPrivateFile = "internal/services/user/server/signatures/signature.pem"
	jwtPublicFile  = "internal/services/user/server/signatures/signature.pub"

	connString = "postgres://postgres:postgres@users_postgres_db:5432/postgres?sslmode=disable"

	port = 8082
)

func main() {
	absoluteJWTPrivateFile, err := filepath.Abs(jwtPrivateFile)
	if err != nil {
		log.Fatalln(err)
	}

	absoluteJWTPublicFile, err := filepath.Abs(jwtPublicFile)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalln(err)
	}
	defer pool.Close()

	userStorage, err := storage.NewUserStorage(ctx, pool)
	if err != nil {
		log.Fatalln(err)
	}
	userUseCase := usecase.NewUserUseCase(userStorage)
	authHandlers, err := handlers.NewUserHandlers(userUseCase, absoluteJWTPrivateFile, absoluteJWTPublicFile)
	if err != nil {
		log.Fatalln(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("POST /signup", authHandlers.SignUp)
	router.HandleFunc("POST /signin", authHandlers.SignIn)
	router.HandleFunc("GET /profile", authHandlers.GetProfile)
	router.HandleFunc("POST /profile", authHandlers.UpdateProfile)
	router.HandleFunc("GET /user_id", authHandlers.GetUserId)

	log.Println(
		"Starting server on port", port, "with jwt private key file", absoluteJWTPrivateFile, "and jwt public key file",
		absoluteJWTPublicFile,
	)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatalln(err)
	}
}
