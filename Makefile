SHELL := /bin/bash
version=$(shell cat VERSION)
LDFLAGS=-ldflags "-X main.AppVersion=$(version)"
format_output=$(shell gofmt -l .)

.PHONY: all
all: clean build

clean:
	rm -f github-issue-schedule

build: lint-check unit-test
	go build -o github-issue-schedule $(LDFLAGS) ./cmd

unit-test:
	CGO_ENABLED=0 go test -v ./...

lint-check:
	@[ "$(format_output)" == "" ] || exit -1

format:
	go fmt ./...
