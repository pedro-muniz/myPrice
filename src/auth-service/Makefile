# Load environment variables from .env
include .env
export

# Set a default auth_secret if not defined in .env
export auth_secret ?= $(shell openssl rand -base64 12)

.PHONY: run test build db-up db-down wire

wire:
	~/go/bin/wire ./infra/config/wire/

run:
	go run main.go

auth-test:
	curl -X POST -H "Content-Type: application/json" -d '{"client_id": "726a28ba-b07b-4848-808d-6b345448357e", "client_secret": "teste"}' http://localhost:8080/authorize

validate-test:
	@echo "Usage: make validate-test TOKEN=<token>"
	curl -X GET "http://localhost:8080/validate?token=$(TOKEN)"

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
