BUILD?= $(CURDIR)/bin
$(shell mkdir -p $(BUILD))
VERSION?= $(shell git rev-list HEAD --count)
export GO111MODULE=on
export GOPATH=$(go env GOPATH)

.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u golang.org/x/tools
	go get -u golang.org/x/lint/golint

.PHONY: mod
mod: ## Runs mod
	go mod verify
	go mod vendor
	go mod tidy

.PHONY: fmt
fmt: ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: test
test: setup ## Runs all the tests
	echo 'mode: atomic' > coverage.txt && go test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: lint
lint: ## Runs all the linters
	golint ./internal/... ./cmd/... ./log/...

.PHONY: vet
vet: ## Runs go vet
	go vet ./...

.PHONY: checks
checks: setup fmt lint vet ## Runs all checks for the project
	echo 'Checks done!'

.PHONY: build
build: ## Builds the project
	go build -o $(BUILD)/abf-guard $(CURDIR)

.PHONY: gen
gen: ## Triggers code generation of
	protoc --go_out=plugins=grpc:$(CURDIR)/internal/grpc api/*.proto

.PHONY: dockerbuild
dockerbuild: mod ## Builds a docker image with a project
	docker build -t omer513/abf-guard:0.${VERSION} .

.PHONY: dockerpush
dockerpush: dockerbuild ## Publishes the docker image to the registry
	docker push omer513/abf-guard:0.${VERSION}

.PHONY: docker-compose-up
docker-compose-up: ## Runs docker-compose command to kick-start the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml up -d

.PHONY: docker-compose-down
docker-compose-down: ## Runs docker-compose command to remove the turn down the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml down -v

.PHONY: integration
integration: ##
	docker-compose -f ./deployments/docker-compose.test.yaml up --build -d;\
	test_status_code=0 ;\
# 	docker-compose -f ./deployments/docker-compose.test.yaml run integration_tests ./bin/integration-test || test_status_code=$$? ;
	docker-compose -f ./deployments/docker-compose.test.yaml down --volumes;\
	printf "Return code is $$test_status_code\n" ;\
	exit $$test_status_code ;\

.PHONY: clean
clean: ## Remove temporary files
	go clean $(CURDIR)
	rm -rf $(BUILD)
	rm -rf coverage.txt

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
