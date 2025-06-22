PROJECT_ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
PROJECT_NAME := $(shell basename $(PROJECT_ROOT))
PROJECT_VERSION := $(shell git describe --exact-match --tags HEAD 2>/dev/null || git rev-parse --short HEAD)$(shell git diff --quiet || echo '-dirty')

# Go
GO ?= go
GOMOBILE ?= gomobile
GOVERSION ?= $(shell go env GOVERSION)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPRIVATE ?=

# Artifacts
TARGET_DIR ?= $(PROJECT_ROOT)/build

# Build options
PROFILE ?= release
LDFLAGS_debug :=
LDFLAGS_release := -s -w

# ------------------------------------------------------------------------------
# Build Targets
# ------------------------------------------------------------------------------

.PHONY: build
build:

# ------------------------------------------------------------------------------
# Development Tools
# ------------------------------------------------------------------------------

.PHONY: env
env:
	@echo "PROJECT_ROOT: $(PROJECT_ROOT)"
	@echo "PROJECT_NAME: $(PROJECT_NAME)"
	@echo "PROJECT_VERSION: $(PROJECT_VERSION)"
	@echo "GO: $(GO)"
	@echo "GOVERSION: $(GOVERSION)"
	@echo "GOOS: $(GOOS)"
	@echo "GOARCH: $(GOARCH)"
	@echo "PROFILE: $(PROFILE)"

.PHONY: init
init:
	GOPRIVATE=$(GOPRIVATE) $(GO) mod download

# ------------------------------------------------------------------------------
# Utilities
# ------------------------------------------------------------------------------

.PHONY: generate
generate:
	go generate ./...

.PHONY: clean
clean:
	rm -rf $(TARGET_DIR)

.PHONY: test
test:
	go clean -testcache
	go test `go list ./... | grep -v /generated/`

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: fmt
fmt:
	go fmt ./...
