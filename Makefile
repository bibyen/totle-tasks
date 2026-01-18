# Test Go code
.PHONY: test
test:
	go clean -testcache; \
	go test ./... -v
