# Makefile for Effective Gin Project

APP_NAME = effective-gin
VERSION ?= $(shell git describe --tags --abbrev=0 2> /dev/null || echo "unknown")

swag:
		swag init --output ./docs --dir ./cmd,./internal/handlers

golangci-lint:
		golangci-lint run ./...

fmt:
		go fmt ./...

test:
		go test -v ./...

build: swag golangci-lint fmt test
		go build -ldflags "-X 'effective-gin/internal/handlers.Version=${VERSION}'\
		-X 'effective-gin/internal/handlers.BuildCommit=$(shell git rev-parse --short HEAD)'\
		-X 'effective-gin/internal/handlers.BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')'" \
		-o ${APP_NAME} cmd/main.go
		echo "Build completed. Executable: ${APP_NAME}"

clean:
		echo "Cleaning build artifacts..."
		go clean
		rm -f ${APP_NAME}
		rm -rf ./docs
		echo "Clean completed."

check: build clean
		echo "Check completed."

run: build
		echo "Running ${APP_NAME}..."
		GIN_MODE=debug ./${APP_NAME}

.PHONY: swag golangci-lint fmt test build clean check run
