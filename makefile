# Variables
APP_NAME = gotam

GO_FILES = $(wildcard ./gotam/*.go)
BUILD_DIR := bin

# Default target
all: build

# Build targets
build: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME): $(GO_FILES)
	@echo Building $(APP_NAME)...
	@go build -o $(BUILD_DIR)/$(APP_NAME).exe cmd/main.go
	@echo Build completed.

# Create necessary directories
$(BUILD_DIR):
	@if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)

$(BIN_DIR):
	@if not exist $(BIN_DIR) mkdir $(BIN_DIR)

# Test target
test:
	@echo Running Go tests...
	@go test -v -p 4 ./tests/...

# Run the application
run: build
	@echo Running the application...
	@$(BUILD_DIR)/$(APP_NAME).exe

# Format code
fmt:
	@echo Formatting Go code...
	@gofmt -w $(GO_FILES)

# Linting (static code analysis)
lint:
	@echo Linting Go code...
	@golint ./...

# Dependency management
deps:
	@echo Installing Go dependencies...
	@go mod tidy

# Clean up build files
clean:
	@echo Cleaning Go build files...
	@if exist $(BUILD_DIR) rmdir /S /Q $(BUILD_DIR)

# Run formatting, linting, and tests
check: fmt lint test
	@echo All checks passed!

# Print help message
help:
	@echo Usage: make [target]
	@echo.
	@echo Targets:
	@echo   all      - Build the application
	@echo   build    - Build the Go application
	@echo   test     - Run tests
	@echo   fmt      - Format Go code
	@echo   run      - Build and run the Go application
	@echo   clean    - Clean up build artifacts
	@echo   check    - Run 'fmt', 'lint' and 'test'
	@echo   deps     - Install dependencies
	@echo   lint     - Lint Go code
