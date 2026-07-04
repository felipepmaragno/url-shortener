GO_FILES := $(shell find . -name '*.go' -not -path './.git/*')

.PHONY: fmt fmt-check test test-race vet run validate

fmt:
	gofmt -w $(GO_FILES)

fmt-check:
	@test -z "$$(gofmt -l $(GO_FILES))"

test:
	go test ./...

test-race:
	go test -race ./...

vet:
	go vet ./...

run:
	go run ./cmd/urlshortener

validate: fmt-check vet test test-race
