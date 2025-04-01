package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/posts/server/handlers"
	"github.com/LeonidS635/soa/internal/services/posts/storage"
	"github.com/LeonidS635/soa/internal/services/posts/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

const (
	connString = "postgres://postgres:postgres@posts_postgres_db:5432/postgres?sslmode=disable"

	port = 8083
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalln(err)
	}
	defer pool.Close()

	postsStorage, err := storage.NewPostsStorage(ctx, pool)
	if err != nil {
		log.Fatalln(err)
	}
	userUseCase := usecase.NewPostsUseCase(postsStorage)
	postsHandlers := handlers.NewPostsHandlers(userUseCase)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("posts service: failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	postspb.RegisterPostsServiceServer(grpcServer, postsHandlers)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("posts service: failed to serve: %v", err)
	}
}
