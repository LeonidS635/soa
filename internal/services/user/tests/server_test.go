package tests

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/LeonidS635/soa/internal/services/posts/tests/mocks"
	"github.com/LeonidS635/soa/internal/services/user/server/handlers"
	"github.com/LeonidS635/soa/internal/services/user/storage"
	"github.com/LeonidS635/soa/internal/services/user/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
)

var db *pgxpool.Pool
var serverRouter *http.ServeMux

var jwtPrivate *rsa.PrivateKey
var jwtPublic *rsa.PublicKey

func initDB(ctx context.Context) (func(), error) {
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("could not construct pool: %w", err)
	}

	if err = dockerPool.Client.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to docker: %w", err)
	}

	resource, err := dockerPool.Run(
		"postgres", "17-alpine", []string{"POSTGRES_USER=test", "POSTGRES_PASSWORD=test", "POSTGRES_DB=test_db"},
	)
	if err != nil {
		return nil, fmt.Errorf("could not start resource: %w", err)
	}

	dbUrl := fmt.Sprintf(
		"postgres://test:test@localhost:%s/test_db?sslmode=disable", resource.GetPort("5432/tcp"),
	)

	if err = dockerPool.Retry(
		func() error {
			var err error
			db, err = pgxpool.New(ctx, dbUrl)
			if err != nil {
				return err
			}
			return db.Ping(ctx)
		},
	); err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	purge := func() {
		if err := dockerPool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	return purge, nil
}

func fillDB(ctx context.Context) error {
	schemaProfiles := `
	CREATE TABLE profiles
	(
		user_id     int  not null,
		name        text default '',
		surname     text default '',
		age         int default 0,
		email       text not null,
		phone       text default '',
		city        text default '',
		description text default '',
		birthdate   timestamptz default now(),
		updated_at  timestamptz default now(),
	
		primary key (user_id)
	)`

	schemaUsers := `
	CREATE TABLE users
	(
		user_id       serial not null,
		username      text   not null,
		password_hash bytea  not null,
		email         text   not null,
		created_at    timestamptz default now(),
	
		primary key (user_id)
	)`

	_, err := db.Exec(ctx, schemaProfiles)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, schemaUsers)
	if err != nil {
		return err
	}

	return nil
}

const (
	jwtPrivateFile = "../server/signatures/signature.pem"
	jwtPublicFile  = "../server/signatures/signature.pub"
)

func createTestPostService(ctx context.Context) (func(), error) {
	absoluteJWTPrivateFile, err := filepath.Abs(jwtPrivateFile)
	if err != nil {
		log.Fatalln(err)
	}
	private, err := os.ReadFile(absoluteJWTPrivateFile)
	if err != nil {
		return nil, fmt.Errorf("colud not read private signature file: %w", err)
	}
	jwtPrivate, err = jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		return nil, fmt.Errorf("colud not parse private signature file: %w", err)
	}

	absoluteJWTPublicFile, err := filepath.Abs(jwtPublicFile)
	if err != nil {
		log.Fatalln(err)
	}
	public, err := os.ReadFile(absoluteJWTPublicFile)
	if err != nil {
		return nil, fmt.Errorf("colud not read public signature file: %w", err)
	}
	jwtPublic, err = jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		return nil, fmt.Errorf("colud not parse public signature file: %w", err)
	}

	userStorage, err := storage.NewUserStorage(ctx, db)
	if err != nil {
		log.Fatalln(err)
	}

	eventsProducer := &mocks.KafkaProducer{}

	userUseCase := usecase.NewUserUseCase(userStorage, eventsProducer)
	authHandlers, err := handlers.NewUserHandlers(userUseCase, absoluteJWTPrivateFile, absoluteJWTPublicFile)
	if err != nil {
		log.Fatalln(err)
	}

	serverRouter = http.NewServeMux()
	serverRouter.HandleFunc("POST /signup", authHandlers.SignUp)
	serverRouter.HandleFunc("POST /signin", authHandlers.SignIn)
	serverRouter.HandleFunc("GET /profile", authHandlers.GetProfile)
	serverRouter.HandleFunc("POST /profile", authHandlers.UpdateProfile)
	serverRouter.HandleFunc("GET /user_id", authHandlers.GetUserId)

	cleanup := func() {
		db.Close()
	}

	return cleanup, nil
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	purge, err := initDB(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer purge()

	log.Println("Started postgres database")

	if err = fillDB(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Println("Created tables")

	cleanup, err := createTestPostService(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer cleanup()

	log.Println("Started user service")

	m.Run()
}
