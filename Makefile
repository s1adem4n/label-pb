SHELL=/bin/bash

BINARY_NAME=label
MAIN_FILE=cmd/main.go
SOURCES=$(shell find . -name '*.go')
VERSION=$(shell git describe --tags --always)
FLAGS=-ldflags "-X main.Version=$(VERSION) -w"
TARGETS = build/$(BINARY_NAME)_linux_amd64 build/$(BINARY_NAME)_darwin_amd64 build/$(BINARY_NAME)_darwin_arm64 build/$(BINARY_NAME)_windows_amd64.exe

.PHONY: all

all: $(TARGETS)

frontend/build:
	cd frontend && bun install && bun run build

build/$(BINARY_NAME)_linux_amd64: frontend/build $(SOURCES)
	GOOS=linux GOARCH=amd64 go build $(FLAGS) -o build/$(BINARY_NAME)_linux_amd64 $(MAIN_FILE)

build/$(BINARY_NAME)_darwin_amd64: frontend/build $(SOURCES)
	GOOS=darwin GOARCH=amd64 go build $(FLAGS) -o build/$(BINARY_NAME)_darwin_amd64 $(MAIN_FILE)

build/$(BINARY_NAME)_darwin_arm64: frontend/build $(SOURCES)
	GOOS=darwin GOARCH=arm64 go build $(FLAGS) -o build/$(BINARY_NAME)_darwin_arm64 $(MAIN_FILE)

build/$(BINARY_NAME)_windows_amd64.exe: frontend/build $(SOURCES)
	GOOS=windows GOARCH=amd64 go build $(FLAGS) -o build/$(BINARY_NAME)_windows_amd64.exe $(MAIN_FILE)
