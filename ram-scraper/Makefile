# Variables
APP_NAME := ram-scraper
BIN_DIR := ../bin
BUILD_PATH := cmd/*.go

build-ram-scraper:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)

build-ram-scraper-test:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(BUILD_PATH)