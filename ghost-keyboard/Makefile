# Variables
APP_NAME := ghost-keyboard
BIN_DIR := ../bin
BUILD_PATH := cmd/main.go

build-ghost-keyboard:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)

build-ghost-keyboard-test:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)
