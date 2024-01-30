SHELL=/bin/bash

BINARY_NAME=label
MAIN_FILE=main.go
SOURCES=$(shell find . -name '*.go')
VERSION=$(shell git describe --tags --always)
FLAGS=-ldflags "-X main.Version=$(VERSION)"
TARGETS = bin/$(BINARY_NAME)_linux_amd64 bin/$(BINARY_NAME)_darwin_amd64 bin/$(BINARY_NAME)_darwin_arm64 bin/$(BINARY_NAME)_windows_amd64

all: frontend/build $(TARGETS)

frontend/build:
	cd frontend && bun install && bun run build

bin/$(BINARY_NAME)_linux_amd64: $(SOURCES)
	GOOS=linux GOARCH=amd64 go build $(FLAGS) -o bin/$(BINARY_NAME)_linux_amd64 $(MAIN_FILE)

bin/$(BINARY_NAME)_darwin_amd64: $(SOURCES)
	GOOS=darwin GOARCH=amd64 go build $(FLAGS) -o bin/$(BINARY_NAME)_darwin_amd64 $(MAIN_FILE)

bin/$(BINARY_NAME)_darwin_arm64: $(SOURCES)
	GOOS=darwin GOARCH=arm64 go build $(FLAGS) -o bin/$(BINARY_NAME)_darwin_arm64 $(MAIN_FILE)

bin/$(BINARY_NAME)_windows_amd64: $(SOURCES)
	GOOS=windows GOARCH=amd64 go build $(FLAGS) -o bin/$(BINARY_NAME)_windows_amd64 $(MAIN_FILE)

