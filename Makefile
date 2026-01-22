.PHONY: build test clean install

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

build:
	go build $(LDFLAGS) -o mcp-obsidian ./cmd/mcp-obsidian

test:
	go test ./...

test-verbose:
	go test -v ./...

test-cover:
	go test -cover ./...

clean:
	rm -f mcp-obsidian

install: build
	mv mcp-obsidian $(GOPATH)/bin/

fmt:
	goimports -w .

vet:
	go vet ./...

lint: fmt vet
