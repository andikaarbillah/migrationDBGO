# Makefile

BINARY_NAME=main
SOURCE_FILES=cmd/main.go
BIN_DIR=bin
BIN_PATH=$(BIN_DIR)/$(BINARY_NAME)
MIGRATION_DIR=cmd/migrate/migrations

# Path to the migrate binary
MIGRATE_BIN=/home/ruter/go/bin/migrate

all: build

build:
	@mkdir -p $(BIN_DIR)  # Create bin directory if it doesn't exist
	@go build -o $(BIN_PATH) $(SOURCE_FILES)

run: build
	@$(BIN_PATH)

# Create a new migration
migration:
	@echo "Creating new migration..."
	@$(MIGRATE_BIN) create -ext sql -dir $(MIGRATION_DIR) $(name)

# Run migrations up
migrate-up:
	@echo "Running migrations up..."
	@go run cmd/migrate/main.go up

# Run migrations down
migrate-down:
	@echo "Running migrations down..."
	@go run cmd/migrate/main.go down

# Define the 'name' variable for migration
name = $(word 2, $(MAKECMDGOALS))

# Override make's default behavior for handling arguments
%:
	@:
.PHONY: all build run migration migrate-up migrate-down
