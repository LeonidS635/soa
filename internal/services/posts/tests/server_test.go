package tests

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/server/handlers"
	"github.com/LeonidS635/soa/internal/services/posts/storage"
	"github.com/LeonidS635/soa/internal/services/posts/tests/mocks"
	"github.com/LeonidS635/soa/internal/services/posts/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var db *pgxpool.Pool
var client postspb.PostsServiceClient

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
	schemaPosts := `
	CREATE TABLE IF NOT EXISTS posts (
	    post_id    serial not null,

    	is_private boolean default false,
    	title      text    default '',
    	tags       text[]  default array []::text[],
    	text       text    default '',

    	primary key (post_id)
	)`

	schemaPostsInfo := `
	CREATE TABLE IF NOT EXISTS posts_info (
	    post_id    int not null,
    	author_id  int not null,

    	created_at timestamptz default now(),
    	updated_at timestamptz default now(),

    	primary key (post_id)
	)`

	_, err := db.Exec(ctx, schemaPosts)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, schemaPostsInfo)
	if err != nil {
		return err
	}

	return nil
}

func createTestPostService(ctx context.Context) (func(), error) {
	postsStorage, err := storage.NewPostsStorage(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("could not create post storage: %w", err)
	}

	eventsProducer := &mocks.KafkaProducer{}

	postsUseCase := usecase.NewPostsUseCase(postsStorage, eventsProducer)
	postsHandlers := handlers.NewPostsHandlers(postsUseCase)

	listener := bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer()
	postspb.RegisterPostsServiceServer(grpcServer, postsHandlers)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

	conn, err := grpc.NewClient(
		"passthrough://test", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	client = postspb.NewPostsServiceClient(conn)
	cleanup := func() {
		grpcServer.Stop()
		conn.Close()
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

	log.Println("Started posts service")

	m.Run()
}
