# Project variables
APP_NAME=spycat-api
DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

# Default target
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  run             - Run the application locally"
	@echo "  build           - Build the Go binary"
	@echo "  clean           - Remove build artifacts"
	@echo "  test            - Run all tests"
	@echo "  docker-build    - Build the Docker image"
	@echo "  docker-run      - Run the application with Docker Compose"
	@echo "  migrate-up      - Apply database migrations"
	@echo "  migrate-down    - Rollback the last migration"
	@echo "  create-migration NAME=<name> - Create a new migration"

# Run the application locally
.PHONY: run
run:
	go run cmd/api/main.go

# Build the application
.PHONY: build
build:
	go build -o $(APP_NAME) cmd/api/main.go

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(APP_NAME)

# Run tests
.PHONY: test
test:
	go test ./...

# Build Docker image
.PHONY: docker-build
docker-build:
	docker build -t $(APP_NAME) .

# Run the application using Docker Compose
.PHONY: docker-run
docker-run:
	docker-compose up --build

# Apply database migrations
.PHONY: migrate-up
migrate-up:
	@command -v migrate >/dev/null 2>&1 || { echo >&2 "migrate is not installed. Run: go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; exit 1; }
	migrate -path migrations -database "$(DB_URL)" up

# Rollback the last migration
.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

.PHONY: docker-clean docker-rebuild

DOCKER := $(shell command -v docker 2>/dev/null)

docker-clean:
	@if [ -z "$(DOCKER)" ]; then \
		echo "❌ Error: Docker is not installed or not in PATH"; exit 1; \
	fi
	docker compose down --volumes --remove-orphans || true
	docker volume prune -f || true
	docker rmi $(shell docker images -q spycatagency-api) || true

docker-rebuild: docker-clean
	@if [ -z "$(DOCKER)" ]; then \
		echo "❌ Error: Docker is not installed or not in PATH"; exit 1; \
	fi
	docker compose up --build -d

