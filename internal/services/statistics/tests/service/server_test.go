package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/server/handlers"
	"github.com/LeonidS635/soa/internal/services/statistics/storage"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase"
	"github.com/ory/dockertest/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var db driver.Conn
var client statspb.StatsServiceClient

func initDB(ctx context.Context) (func(), error) {
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("could not construct pool: %w", err)
	}

	if err = dockerPool.Client.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to docker: %w", err)
	}

	resource, err := dockerPool.Run(
		"clickhouse/clickhouse-server", "latest",
		[]string{"CLICKHOUSE_USER=test", "CLICKHOUSE_PASSWORD=test", "CLICKHOUSE_DB=test_db"},
	)
	if err != nil {
		return nil, fmt.Errorf("could not start resource: %w", err)
	}

	dockerPool.MaxWait = time.Minute
	if err = dockerPool.Retry(
		func() error {
			var err error
			db, err = clickhouse.Open(
				&clickhouse.Options{
					Addr: []string{resource.GetHostPort("9000/tcp")},
					Auth: clickhouse.Auth{
						Database: "test_db",
						Username: "test",
						Password: "test",
					},
					Protocol: clickhouse.Native,
				},
			)
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
	schemaStats := `
	CREATE TABLE daily_stats
	(
		post_id   Int32,
		author_id Int32,
		post_date DATETIME,
		views     UInt8 DEFAULT 0,
		likes     UInt8 DEFAULT 0,
		comments  UInt8 DEFAULT 0
	)
	ENGINE = MergeTree
	PRIMARY KEY post_id`

	return db.Exec(ctx, schemaStats)
}

func createTestStatisticsService(context.Context) (func(), error) {
	statsStorage, err := storage.NewStatisticsStorage(db)
	if err != nil {
		log.Fatalf("statistics service: failed to initialize storage: %v", err)
	}

	statsUseCase := usecase.NewStatsUseCase(statsStorage)
	statsHandlers := handlers.NewStatsHandlers(statsUseCase)

	listener := bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer()
	statspb.RegisterStatsServiceServer(grpcServer, statsHandlers)

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

	client = statspb.NewStatsServiceClient(conn)
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

	log.Println("Started clickhouse database")

	if err = fillDB(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Println("Created tables")

	cleanup, err := createTestStatisticsService(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer cleanup()

	log.Println("Started statistics service")

	m.Run()
}
