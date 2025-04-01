# Variables
APP_NAME := data-seal
BIN_DIR := ../bin
BUILD_PATH := cmd/main.go

build-data-seal:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)

build-data-seal-test:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)
