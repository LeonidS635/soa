package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/LeonidS635/soa/internal/kafka"
	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/server/handlers"
	"github.com/LeonidS635/soa/internal/services/statistics/storage"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase"
	"google.golang.org/grpc"
)

const (
	port = 8084

	kafkaConnString         = "kafka:9092"
	kafkaPostsViewsTopic    = "posts-views"
	kafkaPostsLikesTopic    = "posts-likes"
	kafkaPostsCommentsTopic = "posts-comments"
)

func main() {
	conn, err := clickhouse.Open(
		&clickhouse.Options{
			Addr: []string{"clickhouse_stats_db:9000"},
			Auth: clickhouse.Auth{
				Database: "clickhouse",
				Username: "clickhouse",
				Password: "12345",
			},
			Protocol: clickhouse.Native,
		},
	)
	if err != nil {
		log.Fatalf("statistics service: failed to open clickhouse: %v", err)
	}
	defer conn.Close()

	statsStorage, err := storage.NewStatisticsStorage(conn)
	if err != nil {
		log.Fatalf("statistics service: failed to initialize storage: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventsConsumer := kafka.NewConsumer(kafkaConnString)
	eventsConsumer.RegisterTopic(kafkaPostsViewsTopic)
	eventsConsumer.RegisterTopic(kafkaPostsLikesTopic)
	eventsConsumer.RegisterTopic(kafkaPostsCommentsTopic)

	statsUseCase := usecase.NewStatsUseCase(statsStorage)
	statsHandlers := handlers.NewStatsHandlers(statsUseCase)

	go func() {
		if err := eventsConsumer.Consume(ctx, kafkaPostsViewsTopic, statsUseCase.PostViewed); err != nil {
			log.Fatalf("statistics service: failed to consume 'post-viewes' events: %v", err)
		}
	}()
	go func() {
		if err := eventsConsumer.Consume(ctx, kafkaPostsLikesTopic, statsUseCase.PostLiked); err != nil {
			log.Fatalf("statistics service: failed to consume 'post-likes' events: %v", err)
		}
	}()
	go func() {
		if err := eventsConsumer.Consume(ctx, kafkaPostsCommentsTopic, statsUseCase.PostCommented); err != nil {
			log.Fatalf("statistics service: failed to consume 'post-comments' events: %v", err)
		}
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("statistics service: failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	statspb.RegisterStatsServiceServer(grpcServer, statsHandlers)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("statistics service: failed to serve: %v", err)
	}
}
