.PHONY: build test clean install fmt vet lint

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

build:
	go build $(LDFLAGS) -o obsidian-mcp ./cmd/obsidian-mcp

test:
	go test ./...

test-verbose:
	go test -v ./...

test-cover:
	go test -cover ./...

test-race:
	go test -race -count=1 ./...

clean:
	rm -f obsidian-mcp

install: build
	mv obsidian-mcp $(GOPATH)/bin/

fmt:
	goimports -w .

vet:
	go vet ./...

lint: fmt vet
