SHELL=/bin/bash

BINARY_NAME=label
MAIN_FILE=main.go
SOURCES=$(shell find . -name '*.go')

all: bin/linux/amd64/$(BINARY_NAME) bin/darwin/amd64/$(BINARY_NAME) bin/darwin/arm64/$(BINARY_NAME) bin/windows/amd64/$(BINARY_NAME).exe

bin/linux/amd64/$(BINARY_NAME): $(SOURCES)
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/$(BINARY_NAME) $(MAIN_FILE)

bin/darwin/amd64/$(BINARY_NAME): $(SOURCES)
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/$(BINARY_NAME) $(MAIN_FILE)

bin/darwin/arm64/$(BINARY_NAME): $(SOURCES)
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm64/$(BINARY_NAME) $(MAIN_FILE)

bin/windows/amd64/$(BINARY_NAME).exe: $(SOURCES)
	GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/$(BINARY_NAME).exe $(MAIN_FILE)

