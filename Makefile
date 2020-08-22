BUILD?= $(CURDIR)/bin
$(shell mkdir -p $(BUILD))
VERSION?= v$(shell git rev-list HEAD --count)
ARCH?= $(shell uname -m)
BRANCH_NAME ?= $(shell git rev-parse --abbrev-ref HEAD)
ifneq ($(BRANCH_NAME), master)
	BRANCH_NAME = dev
endif
export GO111MODULE=on
export GOPATH=$(go env GOPATH)

.PHONY: setup mod fmt test coverage lint vet checks build run run-test gen gen-test clean help
.PHONY: dockerbuild dockerpush docker-compose-up docker-compose-down integration

setup: ## Install all the build and lint dependencies
	go get golang.org/x/tools
	go get golang.org/x/tools/cmd/goimports
	go get golang.org/x/lint/golint

mod: ## Runs go mod on a project
	go mod verify
	go mod vendor
	go mod tidy

fmt: ## Runs goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

test: ## Runs all unit tests
	echo 'mode: atomic' > coverage.txt && go test --count=1 -covermode=atomic -coverprofile=coverage.txt -v -race \
	-timeout=30s ./log... ./internal/...

coverage: test ## Runs all the tests and opens the coverage report
	go tool cover -html=coverage.txt

lint: ## Runs all the linters
	golint ./internal/... ./cmd/...
	golint ./test/integration/main_test.go

vet: ## Runs go vet
	go vet -atomic -bools -assign -copylocks -cgocall -asmdecl  ./...

checks: setup fmt lint vet ## Runs all checks for the project (go fmt, go lint, go vet)

build: ## Builds the project
	go build -o $(BUILD)/abf-guard $(CURDIR)

run: build ## Runs the project in production mode
	$(BUILD)/abf-guard grpc-server -c ./configs/config.json

run-test: ## Runs the project for the local usage
	go run main.go grpc-server -c ./configs/config-test.json

.PHONY: gen
gen: ## Triggers code generation for the GRPC Server and Client API
	protoc --go_out=plugins=grpc:$(CURDIR)/internal/grpc ./api/*.proto

.PHONY: gen-test
gen-test: ## Triggers code generation for the GRPC Server and Client API for ITs
	protoc --go_out=plugins=grpc:$(CURDIR)/test/integration-test/ ./api/*.proto

dockerbuild: ## Builds a docker image with the project
	docker build -t omer513/abf-guard-${BRANCH_NAME}-${ARCH}:${VERSION} -f ./deployments/abfg-service.Dockerfile .

dockerpush: dockerbuild ## Publishes the docker image to the registry
	docker push omer513/abf-guard-${BRANCH_NAME}-${ARCH}:${VERSION}

docker-compose-up: ## Runs docker-compose command to kick-start the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml up -d

docker-compose-down: ## Runs docker-compose command to turn down the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml down -v

integration: ## Runs the integration tests for the project
	docker-compose -f ./deployments/docker-compose.test.yaml up --build -d;\
	exit_code=0 ;\
	docker-compose -f ./deployments/docker-compose.test.yaml run integration_tests \
	./abfg-integration-test || $exit_code=$$? ;
	docker-compose -f ./deployments/docker-compose.test.yaml down --volumes;\
	printf "Return code is $$exit_code\n" ;\
	exit $$exit_code ;\

clean: ## Remove temporary files
	go clean $(CURDIR)
	rm -rf $(BUILD)
	rm -rf coverage.txt

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
