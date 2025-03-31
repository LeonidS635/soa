FROM golang:1.24-alpine

WORKDIR /app
COPY go.* .

RUN go mod download

COPY /api ./api

RUN go build -o api-gateway ./api/server/server.go

CMD ["./api-gateway"]