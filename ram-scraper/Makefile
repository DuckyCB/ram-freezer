# Variables
APP_NAME := ram-scraper-validation
BIN_DIR := ../bin
BUILD_PATH := raspberry/*.go

build-ram-scraper-validation:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)

build-ram-scraper-validation-test:
	@echo "Building test $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)