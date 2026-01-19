# Variables
BUF_BIN := $(shell which buf)
PROTO_SRC := proto

# Database Config
DB_CONTAINER := totle-db
DB_USER := postgres
DB_PASS := postgres
DB_NAME := totle_tasks
DB_TEST_NAME := totle_tasks_test
DB_PORT := 5433
SCHEMA_FILE := internal/sql/schema.sql

# DSNs for Go tools
DB_URL := postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable
DB_TEST_URL := postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_TEST_NAME)?sslmode=disable

# === Targets ===

.PHONY: all lint format proto test db/up db/init-test db/shell clean help

all: proto test

# --- Proto Management ---

## lint: Run buf lint to check for AIP/Buf style violations
lint:
	@echo "Running buf lint..."
	@$(BUF_BIN) lint

## format: Automatically format proto files to standard style
format:
	@echo "Formatting proto files..."
	@$(BUF_BIN) format -w

## proto: Generate Go code from proto files
proto: lint
	@echo "🔨 Generating code..."
	@$(BUF_BIN) generate

# --- Database Management ---

## db/up: Start the Postgres container
db/up:
	@echo "Starting Postgres container..."
	@docker run --name $(DB_CONTAINER) \
		-e POSTGRES_PASSWORD=$(DB_PASS) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p $(DB_PORT):5432 \
		-d postgres:15-alpine || docker start $(DB_CONTAINER)

## db/load-schema: Push local schema.sql into the dev database
db/load-schema:
	@echo "Loading schema from $(SCHEMA_FILE)..."
	@docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) < $(SCHEMA_FILE)

## db/init-test: Create test DB if missing and load schema
db/init-test:
	@echo "Initializing test database from $(SCHEMA_FILE)..."
	@docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -tc "SELECT 1 FROM pg_database WHERE datname = '$(DB_TEST_NAME)'" | grep -q 1 || \
		docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -c "CREATE DATABASE $(DB_TEST_NAME);"
	@docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_TEST_NAME) < $(SCHEMA_FILE)

## db/shell: Jump into a psql session
db/shell:
	@docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME)

# --- Testing ---

## test: Run unit tests
test:
	@echo "Running unit tests..."
	go clean -testcache
	go test ./... -v

## test/integration: Setup test DB and run integration tests
test/integration: db/init-test
	@echo "Running integration tests..."
	@TEST_DATABASE_URL="$(DB_TEST_URL)" go test -v -tags=integration ./internal/repository/postgres/...

# --- Cleanup ---

## clean: Stop and remove the database container
clean:
	@echo "Cleaning up..."
	@docker stop $(DB_CONTAINER) || true
	@docker rm $(DB_CONTAINER) || true

help:
	@grep -E '^[a-zA-Z_/]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
