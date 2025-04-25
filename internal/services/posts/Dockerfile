FROM golang:1.24-alpine

WORKDIR /app
COPY go.* .

RUN go mod download

COPY internal/services/posts ./internal/services/posts/
COPY internal/pkg/services/postspb ./internal/pkg/services/postspb
COPY internal/kafka ./internal/kafka

RUN go build -o posts-service ./internal/services/posts/server/server.go

CMD ["./posts-service"]