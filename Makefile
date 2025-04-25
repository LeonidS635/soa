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
