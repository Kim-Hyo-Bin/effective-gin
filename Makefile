# Makefile for Effective Gin Project

APP_NAME = effective-gin
VERSION ?= 0.0.1

swag:
		swag init --generalInfo ./cmd/main.go --output ./docs

build: swag
		go build -ldflags "-X 'effective-gin/internal/handlers.Version=${VERSION}'\
		-X 'effective-gin/internal/handlers.BuildCommit=$(shell git rev-parse --short HEAD)'\
		-X 'effective-gin/internal/handlers.BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')'" \
		-o ${APP_NAME} cmd/main.go
		echo "Build completed. Executable: ${APP_NAME}"

run: build
		echo "Running ${APP_NAME}..."
		GIN_MODE=debug ./${APP_NAME}

clean:
		echo "Cleaning build artifacts..."
		go clean
		rm -f ${APP_NAME}
		rm -rf ./docs
		echo "Clean completed."

.PHONY: swag build run clean