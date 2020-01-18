BUILD?= $(CURDIR)/bin
$(shell mkdir -p $(BUILD))
VERSION?= v$(shell git rev-list HEAD --count)
export GO111MODULE=on
export GOPATH=$(go env GOPATH)

.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u golang.org/x/tools
	go get -u golang.org/x/lint/golint

.PHONY: mod
mod: ## Runs go mod on a project
	go mod verify
	go mod vendor
	go mod tidy

.PHONY: fmt
fmt: ## Runs goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: test
test: ## Runs all unit tests
	echo 'mode: atomic' > coverage.txt && go test -covermode=atomic -coverprofile=coverage.txt -v -race \
	-timeout=30s ./log... ./internal/...

.PHONY: coverage
coverage: test ## Runs all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: lint
lint: ## Runs all the linters
	golint ./internal/... ./cmd/... ./log/...
	golint ./test/integration-test/main_test.go

.PHONY: vet
vet: ## Runs go vet
	go vet -atomic -bools -assign -copylocks -cgocall -asmdecl  ./...

.PHONY: checks
checks: fmt lint vet ## Runs all checks for the project (go fmt, go lint, go vet)

.PHONY: build
build: ## Builds the project
	go build -o $(BUILD)/abf-guard $(CURDIR)

.PHONY: run
run: build ## Runs the project in production mode
	$(BUILD)/abf-guard grpc-server -c ./configs/config.json

.PHONY: run-test
run-test: ## Runs the project for the local usage
	go run main.go grpc-server -c ./configs/config-test.json

.PHONY: gen
gen: ## Triggers code generation for the GRPC Server and Client API
	protoc --go_out=plugins=grpc:$(CURDIR)/internal/grpc ./api/*.proto

.PHONY: gen-test
gen-test: ## Triggers code generation for the GRPC Server and Client API for ITs
	protoc --go_out=plugins=grpc:$(CURDIR)/test/integration-test/ ./api/*.proto

.PHONY: dockerbuild
dockerbuild: ## Builds a docker image with the project
	docker build -t omer513/abf-guard:${VERSION} ./deployments/abfg-service/.

.PHONY: dockerpush
dockerpush: dockerbuild ## Publishes the docker image to the registry
	docker push omer513/abf-guard:${VERSION}

.PHONY: docker-compose-up
docker-compose-up: ## Runs docker-compose command to kick-start the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml up -d

.PHONY: docker-compose-down
docker-compose-down: ## Runs docker-compose command to turn down the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml down -v

.PHONY: integration
integration: ## Runs the integration tests for the project
	docker-compose -f ./deployments/docker-compose.test.yaml up --build -d;\
	exit_code=0 ;\
	docker-compose -f ./deployments/docker-compose.test.yaml run integration_tests \
	./abfg-integration-test || $exit_code=$$? ;
	docker-compose -f ./deployments/docker-compose.test.yaml down --volumes;\
	printf "Return code is $$exit_code\n" ;\
	exit $$exit_code ;\

.PHONY: clean
clean: ## Remove temporary files
	go clean $(CURDIR)
	rm -rf $(BUILD)
	rm -rf coverage.txt

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
