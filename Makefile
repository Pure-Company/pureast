# Makefile for pureast

# Variables
GO          ?= go
GOFLAGS     ?=
LDFLAGS     ?= -s -w
BIN_DIR     := bin
MODULE      := github.com/Pure-Company/pureast

# Binaries to build (one per cmd/ subdirectory)
BINARIES    := pureast funcsig pureast-mcp

# Default target
.PHONY: all
all: build

## build: Build all binaries into ./bin
.PHONY: build
build: $(addprefix $(BIN_DIR)/,$(BINARIES))

$(BIN_DIR)/%: cmd/%
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GOFLAGS) -ldflags='$(LDFLAGS)' -o $@ ./$<

## install: Install binaries to $GOPATH/bin
.PHONY: install
install:
	$(GO) install $(GOFLAGS) -ldflags='$(LDFLAGS)' ./cmd/...

## test: Run tests
.PHONY: test
test:
	$(GO) test $(GOFLAGS) ./...

## test-race: Run tests with race detector
.PHONY: test-race
test-race:
	$(GO) test $(GOFLAGS) -race ./...

## test-cover: Run tests with coverage report
.PHONY: test-cover
test-cover:
	$(GO) test $(GOFLAGS) -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

## fmt: Format Go code
.PHONY: fmt
fmt:
	$(GO) fmt ./...

## vet: Run go vet
.PHONY: vet
vet:
	$(GO) vet ./...

## tidy: Tidy go.mod and go.sum
.PHONY: tidy
tidy:
	$(GO) mod tidy

## lint: Run golangci-lint (requires golangci-lint installed)
.PHONY: lint
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed; skipping"; \
	fi

## check: Run fmt, vet, and tests
.PHONY: check
check: fmt vet test

## clean: Remove build artifacts
.PHONY: clean
clean:
	rm -rf $(BIN_DIR) coverage.out coverage.html

## run-pureast: Build and run the pureast CLI (pass ARGS="...")
.PHONY: run-pureast
run-pureast: $(BIN_DIR)/pureast
	./$(BIN_DIR)/pureast $(ARGS)

## run-funcsig: Build and run funcsig (pass ARGS="...")
.PHONY: run-funcsig
run-funcsig: $(BIN_DIR)/funcsig
	./$(BIN_DIR)/funcsig $(ARGS)

## run-mcp: Build and run pureast-mcp (pass ARGS="...")
.PHONY: run-mcp
run-mcp: $(BIN_DIR)/pureast-mcp
	./$(BIN_DIR)/pureast-mcp $(ARGS)

## help: Show this help
.PHONY: help
help:
	@echo "Targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'
