.PHONY: proto, run-collector, run-gateway, docker-up, docker-down, swagger

proto:
	protoc --proto_path=api/proto \
    	--go_out=api/gen --go_opt=paths=source_relative \
    	--go-grpc_out=api/gen --go-grpc_opt=paths=source_relative \
    	api/proto/collector.proto
run-collector:
	go run services/collector/cmd/main.go

run-gateway:
	go run services/gateway/cmd/main.go

docker-up:
	docker-compose up --build\

swagger:
	cd services/gateway && swag init -g cmd/main.go -o docs

docker-down:
	docker-compose down