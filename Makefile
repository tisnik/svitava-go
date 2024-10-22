SHELL := /bin/bash

.PHONY: default clean build

SOURCES:=$(shell find . -name '*.go')
BINARY:=svitava-go
OUTDIR:=.
#OUTDIR:=/tmp/ramdisk

default: build

clean: ## Run go clean
	@go clean

build:	${OUTDIR}/${BINARY} ## Build binary containing service executable

run:	${OUTDIR}/${BINARY}
	${OUTDIR}/${BINARY} -s

${OUTDIR}/${BINARY}:	${SOURCES}
	@go build -o ${OUTDIR}

benchmark:
	@go test -bench . ./...
