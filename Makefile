SHELL=/bin/bash

BINARY_NAME=label
MAIN_FILE=cmd/main.go
SOURCES=$(shell find . -name '*.go')
VERSION=$(shell git describe --tags --always)
FLAGS=-ldflags "-X main.Version=$(VERSION) -w"
TARGETS = bin/$(BINARY_NAME)_linux_amd64 bin/$(BINARY_NAME)_darwin_amd64 bin/$(BINARY_NAME)_darwin_arm64 bin/$(BINARY_NAME)_windows_amd64.exe

all: $(TARGETS)

pkg/frontend/build:
	cd pkg/frontend && bun install && bun run build

bin/$(BINARY_NAME)_linux_amd64: pkg/frontend/build $(SOURCES)
	GOOS=linux GOARCH=amd64 go build $(FLAGS) -o bin/$(BINARY_NAME)_linux_amd64 $(MAIN_FILE)

bin/$(BINARY_NAME)_darwin_amd64: pkg/frontend/build $(SOURCES)
	GOOS=darwin GOARCH=amd64 go build $(FLAGS) -o bin/$(BINARY_NAME)_darwin_amd64 $(MAIN_FILE)

bin/$(BINARY_NAME)_darwin_arm64: pkg/frontend/build $(SOURCES)
	GOOS=darwin GOARCH=arm64 go build $(FLAGS) -o bin/$(BINARY_NAME)_darwin_arm64 $(MAIN_FILE)

bin/$(BINARY_NAME)_windows_amd64.exe: pkg/frontend/build $(SOURCES)
	GOOS=windows GOARCH=amd64 go build $(FLAGS) -o bin/$(BINARY_NAME)_windows_amd64.exe $(MAIN_FILE)

