# === Tools ===
PROTOC ?= protoc

# === Directories ===
PROTO_DIR := proto
GOOGLEAPIS_DIR := ../googleapis # Path to cloned googleapis repo - https://github.com/googleapis/googleapis.git
OUT_DIR := internal/pb

# === Proto Files ===
PROTO_FILES := \
	$(PROTO_DIR)/v1/totle_tasks.proto

# === Targets ===

# Generate Go code from proto files
.PHONY: proto
proto:
	protoc \
	  -I proto \
	  -I $(GOOGLEAPIS_DIR) \
	  --go_out=internal/pb --go_opt=paths=source_relative \
	  --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
	  --connect-go_out=internal/pb --connect-go_opt=paths=source_relative \
	  ${PROTO_FILES}

# Test Go code
.PHONY: test
test:
	go clean -testcache; \
	go test ./... -v
