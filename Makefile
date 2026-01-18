# === Targets ===

# Generate Go code from proto files
.PHONY: proto
proto:
	buf generate proto

# Test Go code
.PHONY: test
test:
	go clean -testcache; \
	go test ./... -v
