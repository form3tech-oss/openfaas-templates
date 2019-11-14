SHELL := /bin/bash

ROOT := $(shell git rev-parse --show-toplevel)

.PHONY: docker.build
docker.build:
	@find $(ROOT)/template -maxdepth 1 -mindepth 1 -type d -exec bash -c 'cd {} && docker build -t form3tech/openfaas-template-$${PWD##*/}:latest .' \;

.PHONY: install-golangci-lint
install-golangci-lint:
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $$(go env GOPATH)/bin latest

.PHONY: install-deps
install-deps: install-golangci-lint

.PHONY: lint
lint:
	@find $(ROOT)/template -maxdepth 1 -mindepth 1 -type d -exec bash -c 'cd {} && golangci-lint run ./... --enable-all --disable lll,wsl' \;
