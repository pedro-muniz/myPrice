# Load environment variables from .env
include .env
export

.PHONY: run test build db-up db-down

run:
	go run main.go

test:
	go test -v ./...

build:
	go build ./...

# Helper to start a quick postgres container matching your .env file
db-up:
	docker run --name auth-postgres -e POSTGRES_USER=$(postgres_user) -e POSTGRES_PASSWORD=$(postgres_pass) -e POSTGRES_DB=$(postgres_db) -p $(postgres_port):5432 -d postgres:14

# Helper to stop and remove the postgres container
db-down:
	docker stop auth-postgres || true
	docker rm auth-postgres || true
