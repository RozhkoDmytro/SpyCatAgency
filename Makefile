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
	migrate -path migrations -database "$(DB_URL)" up

# Rollback the last migration
.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

# Create a new migration file
.PHONY: create-migration
create-migration:
	migrate create -ext sql -dir migrations -seq $(NAME)

# Stop and remove Docker containers
.PHONY: docker-stop
docker-stop:
	docker-compose down
