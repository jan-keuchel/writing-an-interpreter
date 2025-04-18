# Project settings
BINARY_NAME := jk
SRC_DIR := src
MAIN := $(SRC_DIR)/main.go

# Go tools
GO := go

# Default target
.PHONY: all
all: build

# Build the project
.PHONY: build
build:
	$(GO) build -o $(BINARY_NAME) $(MAIN)

# Run the project
.PHONY: run
run: build
	./$(BINARY_NAME)

# Run tests (if tests exist and use Go test framework)
.PHONY: test
test:
	$(GO) test ./...

# Clean binary and cache
.PHONY: clean
clean:
	$(GO) clean
	rm -f $(BINARY_NAME)

# Format all Go files
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Vet (lint) code
.PHONY: vet
vet:
	$(GO) vet ./...

# Tidy up go.mod
.PHONY: tidy
tidy:
	$(GO) mod tidy
