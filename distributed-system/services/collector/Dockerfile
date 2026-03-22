FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o collector services/collector/cmd/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/collector .
EXPOSE 50051
CMD ["./collector"]