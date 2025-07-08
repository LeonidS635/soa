package e2e

import (
	"context"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	baseURL = "http://localhost:8081"

	dockerComposeFile = "docker-compose.test.yml"
)

var dbm dbsManager

func TestMain(m *testing.M) {
	ctx := context.Background()

	services, err := compose.NewDockerCompose(dockerComposeFile)
	if err != nil {
		log.Fatal(err)
	}

	services = services.WaitForService(
		"postgres_users", wait.ForListeningPort("5432/tcp"),
	).WaitForService(
		"postgres_posts", wait.ForListeningPort("5432/tcp"),
	).WaitForService(
		"clickhouse_stats", wait.ForListeningPort("9000/tcp"),
	).WaitForService(
		"zookeeper", wait.ForListeningPort("2181/tcp"),
	).WaitForService(
		"kafka", wait.ForListeningPort("9092/tcp"),
	).WaitForService(
		"kafka-init-topics", wait.ForExit(),
	).WaitForService(
		"gateway", wait.ForListeningPort("8081/tcp"),
	).WaitForService(
		"user_service", wait.ForListeningPort("8082/tcp"),
	).WaitForService(
		"posts_service", wait.ForListeningPort("8083/tcp"),
	).WaitForService(
		"stats_service", wait.ForListeningPort("8084/tcp"),
	).(*compose.DockerCompose)

	if err = services.Up(ctx, compose.Wait(true)); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := services.Down(
			context.Background(),
			compose.RemoveOrphans(true),
			compose.RemoveVolumes(true),
			compose.RemoveImagesLocal,
		); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("services are up")

	dbm, err = initDBsManager()
	if err != nil {
		log.Fatal(err)
	}

	if err = dbm.fillDB(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("filled databases")

	m.Run()
}
