# Makefile for Effective Gin Project

APP_NAME = effective-gin
VERSION ?= 0.0.1

swag:
		swag init --generalInfo ./cmd/main.go --output ./docs

build: golangci-lint swag
		go build -ldflags "-X 'effective-gin/internal/handlers.Version=${VERSION}'\
		-X 'effective-gin/internal/handlers.BuildCommit=$(shell git rev-parse --short HEAD)'\
		-X 'effective-gin/internal/handlers.BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')'" \
		-o ${APP_NAME} cmd/main.go
		echo "Build completed. Executable: ${APP_NAME}"

golangci-lint:
		golangci-lint run ./...

fmt:
		go fmt ./...

test:
		go test -v ./...

check: golangci-lint fmt test
		echo "Check completed."

run: build
		echo "Running ${APP_NAME}..."
		GIN_MODE=debug ./${APP_NAME}

clean:
		echo "Cleaning build artifacts..."
		go clean
		rm -f ${APP_NAME}
		rm -rf ./docs
		echo "Clean completed."

.PHONY: swag golangci-lint fmt test check build run clean