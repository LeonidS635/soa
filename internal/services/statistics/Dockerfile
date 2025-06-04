FROM golang:1.24-alpine

WORKDIR /app
COPY go.* .

RUN go mod download

COPY internal/services/statistics ./internal/services/statistics/
COPY internal/pkg/services/statspb ./internal/pkg/services/statspb
COPY internal/kafka ./internal/kafka

RUN go build -o stats-service ./internal/services/statistics/server/server.go

CMD ["./stats-service"]