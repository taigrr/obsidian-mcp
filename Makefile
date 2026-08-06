.PHONY: build test clean install fmt vet lint

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"
GOBIN ?= $(shell go env GOPATH)/bin

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
	install -d $(GOBIN)
	install -m 0755 obsidian-mcp $(GOBIN)/obsidian-mcp

fmt:
	goimports -w .

vet:
	go vet ./...

lint: fmt vet
