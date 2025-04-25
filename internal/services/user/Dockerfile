FROM golang:1.24-alpine

WORKDIR /app
COPY go.* .

RUN go mod download

COPY internal/services/user ./internal/services/user/
COPY internal/kafka ./internal/kafka

RUN go build -o user-service ./internal/services/user/server/server.go

CMD ["./user-service"]