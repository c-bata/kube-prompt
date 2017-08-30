NAME := kube-prompt
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'

.DEFAULT_GOAL := help

.PHONY: setup
setup:  ## Setup for required tools.
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/dep/cmd/dep

.PHONY: fmt
fmt: ## Formatting source codes.
	@goimports -w ./kube

.PHONY: lint
lint: ## Run golint and go vet.
	@golint ./kube/...
	@go vet ./kube/...

.PHONY: test
test:  ## Run the tests.
	@go test ./kube/...

.PHONY: build
build: main.go  ## Build a binary.
	go build -ldflags "$(LDFLAGS)"

.PHONY: cross
cross: main.go  ## Build binaries for cross platform.
	mkdir -p pkg
	@for os in "darwin" "linux"; do \
		for arch in "amd64" "386"; do \
			GOOS=$${os} GOARCH=$${arch} make build; \
			zip pkg/kube-prompt_$(VERSION)_$${os}_$${arch}.zip kube-prompt; \
		done; \
	done

.PHONY: help
help: ## Show help text
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
