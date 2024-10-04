SHELL := /bin/bash

.PHONY: default clean build

SOURCES:=$(shell find . -name '*.go')
BINARY:=svitava-go

default: build

clean: ## Run go clean
	@go clean

build: ${BINARY} ## Build binary containing service executable

${BINARY}:
	@go build
