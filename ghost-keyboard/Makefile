# Variables
APP_NAME := ghost-keyboard
BIN_DIR := bin
BUILD_PATH := $(wildcard *.go)

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)
