# Binary name
BINARY_NAME=kube-a2s

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard ./cmd/$(BINARY_NAME)/*.go)

# Go build flags
LDFLAGS=-ldflags "-s -w"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: all build clean test coverage vet lint fmt help

## Build:
all: clean build ## Build the binary

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(GOBIN)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

clean: ## Remove build related files
	@echo "Cleaning..."
	rm -rf $(GOBIN)
	rm -f $(BINARY_NAME)

## Test:
test: ## Run unit tests
	@echo "Running tests..."
	go test -v ./...

coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

## Code Quality:
vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

lint: ## Run golangci-lint
	@echo "Running golangci-lint..."
	golangci-lint run

fmt: ## Run go fmt
	@echo "Running go fmt..."
	go fmt ./...

## Help:
help: ## Show this help
	@echo "Usage:"
	@echo "  make \033[36m<target>\033[0m"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# Default target
.DEFAULT_GOAL := help 