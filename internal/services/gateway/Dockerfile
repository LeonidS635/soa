FROM golang:1.24-alpine

WORKDIR /app
COPY go.* .

RUN go mod download

COPY internal/services/gateway ./internal/services/gateway/
COPY internal/pkg/services/postspb ./internal/pkg/services/postspb

RUN go build -o gateway ./internal/services/gateway/server/server.go

CMD ["./gateway"]