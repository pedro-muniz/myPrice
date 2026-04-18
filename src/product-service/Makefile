.PHONY: build test clean coverage migrate migrate-info migrate-repair migrate-baseline

-include .env
export

# Database configuration

# Flyway configuration
FLYWAY_IMAGE = flyway/flyway:latest
FLYWAY_COMMAND = docker run --rm --network="host" -v $(PWD)/flyway/sql:/flyway/sql $(FLYWAY_IMAGE) -url=$(postgres_host)$(postgres_db) -user=$(postgres_user) -password=$(postgres_pass) -locations=filesystem:/flyway/sql -schemas=service -connectRetries=60

build:
	go build ./...

test:
	go test -v ./...

coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	go clean
	rm -f coverage.out

migrate:
	$(FLYWAY_COMMAND) migrate

migrate-info:
	$(FLYWAY_COMMAND) info

migrate-repair:
	$(FLYWAY_COMMAND) repair

migrate-baseline:
	$(FLYWAY_COMMAND) baseline
