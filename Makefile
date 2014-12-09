BIN = go-tugboat

all: clean build test

setup:
	go get github.com/tools/godep
	go get github.com/golang/lint/golint

test:
	go test $(TESTFLAGS) ./...

build:
	go build -o build/$(BIN)

run: build
	./build/$(BIN)

clean:
	rm -f build/$(BIN)
	go clean

lint:
	golint ./...

vet:
	go vet ./...

.PHONY: setup test build run clean lint vet coverage