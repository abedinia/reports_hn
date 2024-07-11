APP_NAME := reports_hn_app
CMD_PATH := cmd/main.go
BUILD_PATH := build
BINARY_NAME := $(BUILD_PATH)/$(APP_NAME)
COVERAGE_PATH := coverage


all: build


build:
	@echo "Building the application..."
	@go build -o $(BINARY_NAME) $(CMD_PATH)


run:
	@echo "Running the application..."
	@go run $(CMD_PATH) report


seed:
	@echo "Seeding the database..."
	@go run $(CMD_PATH) seed


test:
	@echo "Running tests..."
	@go test ./...


test-cover:
	@echo "Running tests with coverage..."
	@mkdir -p $(COVERAGE_PATH)
	@go test ./... -coverprofile=$(COVERAGE_PATH)/coverage.out
	@go tool cover -html=$(COVERAGE_PATH)/coverage.out -o $(COVERAGE_PATH)/coverage.html


clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_PATH)
	@rm -rf $(COVERAGE_PATH)


help:
	@echo "Makefile commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  seed        - Seed the database"
	@echo "  test        - Run tests"
	@echo "  test-cover  - Run tests with coverage"
	@echo "  clean       - Clean build artifacts"
	@echo "  help        - Display this help message"