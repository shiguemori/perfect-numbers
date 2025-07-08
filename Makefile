# Perfect Numbers API Makefile

# Variables
APP_NAME=perfect-numbers-api
DOCKER_IMAGE=$(APP_NAME):latest
DOCKER_CONTAINER=$(APP_NAME)-container
PORT=8080

# Go variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=main
BINARY_PATH=./cmd/api

# Colors for output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

.PHONY: help build run test clean deps docker-build docker-run docker-stop docker-clean compose-up compose-down lint fmt vet coverage benchmark

# Default target
all: clean deps test build

# Help
help: ## Show this help message
	@echo "$(BLUE)Perfect Numbers API - Available commands:$(NC)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
deps: ## Download dependencies
	@echo "$(YELLOW)Downloading dependencies...$(NC)"
	$(GOMOD) download
	$(GOMOD) tidy

build: ## Build the application
	@echo "$(YELLOW)Building application...$(NC)"
	$(GOBUILD) -o $(BINARY_NAME) $(BINARY_PATH)
	@echo "$(GREEN)Build completed!$(NC)"

run: ## Run the application locally
	@echo "$(YELLOW)Starting application...$(NC)"
	$(GOCMD) run $(BINARY_PATH)

clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning...$(NC)"
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	@echo "$(GREEN)Clean completed!$(NC)"

# Testing
test: ## Run tests
	@echo "$(YELLOW)Running tests...$(NC)"
	$(GOTEST) -v ./tests/...
	@echo "$(GREEN)Tests completed!$(NC)"

test-coverage: ## Run tests with coverage
	@echo "$(YELLOW)Running tests with coverage...$(NC)"
	$(GOTEST) -v -coverprofile=coverage.out ./tests/...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

benchmark: ## Run benchmarks
	@echo "$(YELLOW)Running benchmarks...$(NC)"
	$(GOTEST) -bench=. -benchmem ./tests/...

# Code quality
fmt: ## Format code
	@echo "$(YELLOW)Formatting code...$(NC)"
	$(GOCMD) fmt ./...
	@echo "$(GREEN)Code formatted!$(NC)"

vet: ## Run go vet
	@echo "$(YELLOW)Running go vet...$(NC)"
	$(GOCMD) vet ./...
	@echo "$(GREEN)Vet completed!$(NC)"

lint: fmt vet ## Run linting tools

# Docker
docker-build: ## Build Docker image
	@echo "$(YELLOW)Building Docker image...$(NC)"
	docker build -t $(DOCKER_IMAGE) .
	@echo "$(GREEN)Docker image built: $(DOCKER_IMAGE)$(NC)"

docker-run: ## Run Docker container
	@echo "$(YELLOW)Running Docker container...$(NC)"
	docker run -d --name $(DOCKER_CONTAINER) -p $(PORT):$(PORT) $(DOCKER_IMAGE)
	@echo "$(GREEN)Container running on port $(PORT)$(NC)"

docker-stop: ## Stop Docker container
	@echo "$(YELLOW)Stopping Docker container...$(NC)"
	-docker stop $(DOCKER_CONTAINER)
	-docker rm $(DOCKER_CONTAINER)
	@echo "$(GREEN)Container stopped!$(NC)"

docker-clean: docker-stop ## Clean Docker images and containers
	@echo "$(YELLOW)Cleaning Docker images...$(NC)"
	-docker rmi $(DOCKER_IMAGE)
	-docker system prune -f
	@echo "$(GREEN)Docker cleanup completed!$(NC)"

docker-logs: ## Show Docker container logs
	docker logs -f $(DOCKER_CONTAINER)

# Docker Compose
compose-up: ## Start services with docker-compose
	@echo "$(YELLOW)Starting services with docker-compose...$(NC)"
	docker-compose up -d
	@echo "$(GREEN)Services started!$(NC)"

compose-down: ## Stop services with docker-compose
	@echo "$(YELLOW)Stopping services with docker-compose...$(NC)"
	docker-compose down
	@echo "$(GREEN)Services stopped!$(NC)"

compose-logs: ## Show docker-compose logs
	docker-compose logs -f

compose-build: ## Build and start services with docker-compose
	@echo "$(YELLOW)Building and starting services...$(NC)"
	docker-compose up --build -d
	@echo "$(GREEN)Services built and started!$(NC)"

# Development workflow
dev: clean deps lint test build ## Complete development workflow

# Production deployment
deploy: clean deps test docker-build compose-up ## Deploy to production

# API testing
test-api: ## Test API endpoints (requires running server)
	@echo "$(YELLOW)Testing API endpoints...$(NC)"
	@echo "Testing health endpoint..."
	curl -s http://localhost:$(PORT)/health | jq .
	@echo "\nTesting perfect numbers endpoint..."
	curl -s -X POST http://localhost:$(PORT)/perfect-numbers \
		-H "Content-Type: application/json" \
		-d '{"start": 1, "end": 10000}' | jq .
	@echo "$(GREEN)API tests completed!$(NC)"

# Monitoring
status: ## Show application status
	@echo "$(BLUE)Application Status:$(NC)"
	@echo "Docker containers:"
	@docker ps --filter name=$(APP_NAME)
	@echo "\nDocker images:"
	@docker images --filter reference=$(APP_NAME)

# Installation
install: ## Install the application binary
	@echo "$(YELLOW)Installing application...$(NC)"
	$(GOBUILD) -o /usr/local/bin/$(APP_NAME) $(BINARY_PATH)
	@echo "$(GREEN)Application installed to /usr/local/bin/$(APP_NAME)$(NC)"

# Quick start
quick-start: deps build run ## Quick start for development

# Show project info
info: ## Show project information
	@echo "$(BLUE)Perfect Numbers API$(NC)"
	@echo "Version: 2.0.0"
	@echo "Go version: $(shell go version)"
	@echo "Docker version: $(shell docker --version 2>/dev/null || echo 'Docker not installed')"
	@echo "Project structure:"
	@find . -type f -name "*.go" | head -10

