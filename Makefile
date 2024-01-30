SHELL=/bin/bash

BINARY_NAME=label
MAIN_FILE=main.go
SOURCES=$(shell find . -name '*.go')
VERSION=$(shell git describe --tags --always)
FLAGS=-ldflags "-X main.Version=$(VERSION)"

all: bin/linux/amd64/$(BINARY_NAME) bin/darwin/amd64/$(BINARY_NAME) bin/darwin/arm64/$(BINARY_NAME) bin/windows/amd64/$(BINARY_NAME).exe

bin/linux/amd64/$(BINARY_NAME): $(SOURCES)
	GOOS=linux GOARCH=amd64 go build $(FLAGS) -o bin/linux/amd64/$(BINARY_NAME) $(MAIN_FILE)

bin/darwin/amd64/$(BINARY_NAME): $(SOURCES)
	GOOS=darwin GOARCH=amd64 go build $(FLAGS) -o bin/darwin/amd64/$(BINARY_NAME) $(MAIN_FILE)

bin/darwin/arm64/$(BINARY_NAME): $(SOURCES)
	GOOS=darwin GOARCH=arm64 go build $(FLAGS) -o bin/darwin/arm64/$(BINARY_NAME) $(MAIN_FILE)

bin/windows/amd64/$(BINARY_NAME).exe: $(SOURCES)
	GOOS=windows GOARCH=amd64 go build $(FLAGS) -o bin/windows/amd64/$(BINARY_NAME).exe $(MAIN_FILE)

