
SHA := $(shell git rev-parse --short=8 HEAD)
GITVERSION := $(shell git describe --long --all)
BUILDDATE := $(shell date -Iseconds)
VERSION := $(or ${VERSION},$(shell git describe --tags --exact-match 2> /dev/null || git symbolic-ref -q --short HEAD || git rev-parse --short HEAD))
LOCALBIN ?= $(shell pwd)/bin

all: proto generate test

release: proto generate test

.PHONY: proto
proto: protolint
	$(MAKE) -C go clean
	$(MAKE) -C proto protoc

.PHONY: protolint
protolint:
	$(MAKE) -C proto protolint

.PHONY: generate
generate:
	$(MAKE) -C generate generate

.PHONY: test
test:
	$(MAKE) -C go test
