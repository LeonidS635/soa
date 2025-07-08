# ----------
# Tests
# ----------

# Services

test-user-service:
	go test -v ./internal/services/user/tests -count=1

test-post-service:
	go test -v ./internal/services/posts/tests -count=1

test-stats-service:
	go test -v ./internal/services/statistics/tests/service -count=1

test-services: test-user-service test-post-service test-stats-service

# Kafka

test-stats-kafka:
	go test -v ./internal/services/statistics/tests/kafka -count=1

# e2e

test-e2e:
	go test -v ./internal/tests/e2e

# All

test-all: test-services test-stats-kafka test-e2e

# --------------
# Docker compose
# --------------

compose-up:
	docker-compose up --build

compose-down:
	docker-compose down

# ---------------------
# Proto code generation
# ---------------------

protogen-posts:
	protoc --go_out=./internal/pkg/services/postspb --go_opt=paths=source_relative \
		--go-grpc_out=./internal/pkg/services/postspb --go-grpc_opt=paths=source_relative \
		-I ./internal/services/posts/api posts_service.proto

protogen-stats:
	protoc --go_out=./internal/pkg/services/statspb --go_opt=paths=source_relative \
		--go-grpc_out=./internal/pkg/services/statspb --go-grpc_opt=paths=source_relative \
		-I ./internal/services/statistics/api stats_service.proto
