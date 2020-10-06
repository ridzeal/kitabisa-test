SHELL := /bin/bash

## Golang
GOCMD=go
GORUN=$(GOCMD) run
export GOSUMDB=off

test:
	go test ./...

run-dev:
	${GORUN} main.go ${ARGS}

debug:
	${GORUN} debug.go ${ARGS}