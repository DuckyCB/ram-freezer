# Variables
APP_NAME := project-manager
BIN_DIR := ../bin
BUILD_PATH := cmd/*.go

build-project-manager:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)
