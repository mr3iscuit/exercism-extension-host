.PHONY: build run install clean help

# Binary name
BINARY_NAME=extension-host

# Build the application
build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) ./main.go

# Run the application
run: build
	@echo "Running..."
	@./$(BINARY_NAME)

# Install the Chrome native messaging host
install: build
	@echo "Installing Chrome native messaging host..."
	@./$(BINARY_NAME) install

# Clean build files
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@go clean

# Show help
help:
	@echo "Available commands:"
	@echo "  make build    - Build the application"
	@echo "  make run      - Build and run the application"
	@echo "  make install  - Build and install the Chrome native messaging host"
	@echo "  make clean    - Clean build files"
	@echo "  make help     - Show this help message" 