# Variables
BUF_BIN := $(shell which buf)
PROTO_SRC := proto

# === Targets ===

.PHONY: all lint format generate clean help

## lint: Run buf lint to check for AIP/Buf style violations
lint:
	@echo "?? Running buf lint..."
	@$(BUF_BIN) lint

## format: Automatically format proto files to standard style
format:
	@echo "?? Formatting proto files..."
	@$(BUF_BIN) format -w

## generate: Generate Go code from proto files
proto: lint
	@echo "?? Generating code..."
	@$(BUF_BIN) generate

# Test Go code
.PHONY: test
test:
	go clean -testcache; \
	go test ./... -v
