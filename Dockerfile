FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

COPY .env /app/.env

RUN chmod +x /app/main

EXPOSE 8080

CMD ["/app/main"]
