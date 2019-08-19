.PHONY: help build clean test all
.DEFAULT_GOAL := help

PACKAGE = ./pkg/markdown
GO_SOURCES = $(shell find . -type f -name '*.go')

all: build test ## do everything

build: $(PACKAGE) ## compile project

test: ## run test
	GO111MODULE=on go test -v ./...

$(OUTPUT): $(GO_SOURCES)
	GO111MODULE=on go build ./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

