SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	GOARCH=amd64 $(GOCMD) build -v -o $(BINARY_DIR)/api ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/api

test: ## Run tests
	$(GOCMD) test ./... -cover

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

deps: ## Install dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod 

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

wire: ## Generate wire_gen.go
	cd pkg/di && wire

swag: ## Generate swagger docs
	swag init -g pkg/api/server.go -o ./cmd/api/docs

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

mockgen:
	mockgen -source=pkg/repository/interface/user.go -destination=pkg/repository/mockrepo/mock_user.go
	mockgen -source=pkg/repository/interface/refresh_token.go -destination=pkg/repository/mockrepo/refresh_token.go
	mockgen -source=pkg/usecase/interface/user.go -destination=pkg/usecase/mockusecase/mock_user.go
	mockgen -source=pkg/usecase/interface/cart.go -destination=pkg/usecase/mockusecase/mock_cart.go
	mockgen -source=pkg/usecase/interface/wallet.go -destination=pkg/usecase/mockusecase/mock_wallet.go

docker-up: ## To up the docker compose file
	docker-compose up 

docker-down: ## To down the docker compose file
	docker-compose down

docker-build: ## To build newdocker file for this project
	docker build -t kannan112/any-fashion-store . 