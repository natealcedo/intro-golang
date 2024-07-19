.PHONY: build clean test run install

APP_NAME = intro-golang
BIN = bin/$(APP_NAME)

run: build
	@./$(BIN)

install:
	@go mod tidy

build:
	@go build -o $(BIN)

clean:
	@rm -rf bin/*

test:
	@go test -v ./...