package kafka

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/LeonidS635/soa/internal/kafka"
	"github.com/LeonidS635/soa/internal/services/statistics/storage"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const (
	postsViewsTopic    = "posts-views"
	postsLikesTopic    = "posts-likes"
	postsCommentsTopic = "posts-comments"
)

var db driver.Conn
var kafkaAddr string

func setupKafka(ctx context.Context) (func(), error) {
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("could not construct pool: %w", err)
	}

	if err = dockerPool.Client.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to docker: %w", err)
	}

	zookeeperOpts := &dockertest.RunOptions{
		Repository: "confluentinc/cp-zookeeper",
		Tag:        "latest",
		Env: []string{
			"ZOOKEEPER_CLIENT_PORT=2181",
			"ZOOKEEPER_TICK_TIME=2000",
		},
		ExposedPorts: []string{"2181/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"2181/tcp": {{HostIP: "0.0.0.0", HostPort: "2181"}},
		},
		Hostname: "zookeeper",
		Name:     "zookeeper",
	}
	zookeeper, err := dockerPool.RunWithOptions(zookeeperOpts)
	if err != nil {
		return nil, fmt.Errorf("could not start zookeeper: %w", err)
	}

	dockerPool.MaxWait = time.Minute
	if err = dockerPool.Retry(
		func() error {
			zookeeperAddr := zookeeper.GetHostPort("2181/tcp")
			c, err := net.DialTimeout("tcp", zookeeperAddr, 10*time.Second)
			if err != nil {
				return err
			}
			return c.Close()
		},
	); err != nil {
		return nil, fmt.Errorf("could not connect to zookeeper: %w", err)
	}

	purgeZookeeper := func() {
		if err := dockerPool.Purge(zookeeper); err != nil {
			log.Printf("could not purge zookeeper: %v", err)
		}
	}

	kafkaOpts := &dockertest.RunOptions{
		Repository: "confluentinc/cp-kafka",
		Tag:        "7.9.0",
		Env: []string{
			"KAFKA_BROKER_ID=1",
			"KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181",
			"KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092",
			"KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
		},
		ExposedPorts: []string{"9092/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9092/tcp": {{HostIP: "", HostPort: "9092"}},
		},
		Hostname: "kafka",
		Links:    []string{"zookeeper:zookeeper"},
		Name:     "kafka",
	}
	kafkaContainer, err := dockerPool.RunWithOptions(kafkaOpts)
	if err != nil {
		return purgeZookeeper, fmt.Errorf("could not start kafka: %w", err)
	}

	dockerPool.MaxWait = time.Minute
	if err = dockerPool.Retry(
		func() error {
			kafkaAddr = kafkaContainer.GetHostPort("9092/tcp")
			c, err := net.DialTimeout("tcp", kafkaAddr, 10*time.Second)
			if err != nil {
				return err
			}
			return c.Close()
		},
	); err != nil {
		return purgeZookeeper, fmt.Errorf("could not connect to kafka: %w", err)
	}

	purge := func() {
		if err := dockerPool.Purge(kafkaContainer); err != nil {
			log.Printf("Could not purge resource: %s", err)
		}
		if err := dockerPool.Purge(zookeeper); err != nil {
			log.Printf("Could not purge resource: %s", err)
		}
	}

	topics := []string{"users-registrations", "posts-views", "posts-likes", "posts-comments"}
	for _, topic := range topics {
		cmd := []string{
			"kafka-topics",
			"--bootstrap-server", kafkaAddr,
			"--create",
			"--if-not-exists",
			"--topic", topic,
			"--partitions", "1",
			"--replication-factor", "1",
		}
		exitCode, execErr := kafkaContainer.Exec(cmd, dockertest.ExecOptions{})
		if execErr != nil {
			return purge, fmt.Errorf("failed to create topic %s: %s", topic, execErr)
		}
		if exitCode != 0 {
			return purge, fmt.Errorf("failed to create topic %s: exited with code %d", topic, exitCode)
		}
	}

	return purge, nil
}

func consumeKafkaTopics(ctx context.Context) error {
	statsStorage, err := storage.NewStatisticsStorage(db)
	if err != nil {
		return fmt.Errorf("statistics service: failed to initialize storage: %w", err)
	}

	eventsConsumer := kafka.NewConsumer(kafkaAddr)
	eventsConsumer.RegisterTopic(postsViewsTopic)
	eventsConsumer.RegisterTopic(postsLikesTopic)
	eventsConsumer.RegisterTopic(postsCommentsTopic)

	statsUseCase := usecase.NewStatsUseCase(statsStorage)

	go func() {
		if err := eventsConsumer.Consume(ctx, postsViewsTopic, statsUseCase.PostViewed); err != nil {
			log.Fatalf("statistics service: failed to consume 'post-viewes' events: %v", err)
		}
	}()
	go func() {
		if err := eventsConsumer.Consume(ctx, postsLikesTopic, statsUseCase.PostLiked); err != nil {
			log.Fatalf("statistics service: failed to consume 'post-likes' events: %v", err)
		}
	}()
	go func() {
		if err := eventsConsumer.Consume(ctx, postsCommentsTopic, statsUseCase.PostCommented); err != nil {
			log.Fatalf("statistics service: failed to consume 'post-comments' events: %v", err)
		}
	}()

	return nil
}

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
	schemaPosts := `
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

	return db.Exec(ctx, schemaPosts)
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	purgeKafka, err := setupKafka(ctx)
	if purgeKafka != nil {
		defer purgeKafka()
	}
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Started kafka")

	purgeClick, err := initDB(ctx)
	if purgeClick != nil {
		defer purgeClick()
	}
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Started clickhouse database")

	if err = fillDB(ctx); err != nil {
		log.Println(err)
		return
	}

	log.Println("Created tables")

	if err = consumeKafkaTopics(ctx); err != nil {
		log.Println(err)
		return
	}

	log.Println("Started consuming topics")

	m.Run()
}
