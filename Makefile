.PHONY: test
test-go:
	go test ./... -race

.PHONY: lint
lint:
	golangci-lint run
