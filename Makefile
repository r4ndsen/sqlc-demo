include .env
export

build:
	@go build -o bin/sqlc-demo ./cmd/sqlc-demo/main.go

server: build
	@./bin/sqlc-demo --port 3000

types:
	docker-compose run --rm sqlc

migrate-up:
	docker-compose run --rm migrate -path . -database "$(DATABASE_URI)" up

migrate-down:
	docker-compose run --rm migrate -path . -database "$(DATABASE_URI)" down -all

migrate-new:
	docker-compose run --rm migrate create -ext sql -dir . -seq $(name)

update:
	go get -u ./... && \
	go mod tidy
