APP_NAME=main
MAIN_FILE=./cmd/server/main.go

# Detect OS (Linux, Darwin=Mac, Windows_NT=Windows)
ifeq ($(OS),Windows_NT)
    EXE_EXT=.exe
    AIR_BIN=$(shell go env GOPATH)/bin/air.exe
else
    EXE_EXT=
    AIR_BIN=$(shell go env GOPATH)/bin/air
endif

BINARY=./tmp/$(APP_NAME)$(EXE_EXT)

all: build

build:
	@echo "⚒️ Building..."
	@mkdir -p tmp
	@go build -o $(BINARY) $(MAIN_FILE)

# Run binary
run: build
	@echo "🚀 Running $(APP_NAME) binary..."
	@$(BINARY)

# Create DB container
docker-up:
	@if docker compose up -d --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up -d --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Clean the binary
clean:
	@echo "🧹 Cleaning..."
	@rm -rf tmp

# Live Reload
watch:
	@echo "👀 Starting Air..."
	@$(AIR_BIN)

# Create a new migration
migrate-diff:
	@if [ -z "$(name)" ]; then \
		echo "❌ Please provide a migration name. Usage: make migrate-diff name=add_users"; \
	else \
		atlas migrate diff $${name} --env gorm; \
	fi

# Apply migrations
migrate-apply:
	@atlas migrate apply --env gorm

# Show migration status
migrate-status:
	@atlas migrate status --env gorm

# Rollback the last migration
migrate-down:
	@atlas migrate down --env gorm; 


.PHONY: all build run clean watch docker-up docker-down migrate-diff migrate-apply migrate-status migrate-down migrate-reset
