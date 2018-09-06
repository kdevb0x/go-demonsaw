# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
BINARY_NAME=$(shell basename "$(PWD)")

build:
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) -v

all: clean get build test

test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(GOBIN)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	$(GOBIN)/$(BINARY_NAME)

get:
	$(GOCMD) mod tidy
	$(GOGET) .

.PHONY: build test clean run get
