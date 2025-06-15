.PHONY: build run install clean help chrome-install

# Binary name
BINARY_NAME=exercism-extension-host
INSTALL_DIR=/opt/biscuit/exercism
MANIFEST_NAME=com.biscuit.extensions.exercism.json
CHROME_HOSTS_DIR=$(HOME)/.config/google-chrome/NativeMessagingHosts

# Build the application
build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) ./main.go

# Run the application
run: build
	@echo "Running..."
	@./$(BINARY_NAME)

# Install both the application and Chrome host manifest
install: build chrome-install
	@echo "Installing application to $(INSTALL_DIR)..."
	@sudo mkdir -p $(INSTALL_DIR)
	@sudo cp $(BINARY_NAME) $(INSTALL_DIR)/
	@sudo cp chrome-host.json $(INSTALL_DIR)/$(MANIFEST_NAME)
	@echo "Installing Chrome native messaging host..."
	@$(INSTALL_DIR)/$(BINARY_NAME) install
	@echo "Installation complete! App installed to $(INSTALL_DIR) and Chrome host manifest is installed."

# Install Chrome host manifest
chrome-install:
	@echo "Installing Chrome host manifest..."
	@mkdir -p $(CHROME_HOSTS_DIR)
	@cp chrome-host.json $(CHROME_HOSTS_DIR)/$(MANIFEST_NAME)
	@echo "Chrome host manifest installed at $(CHROME_HOSTS_DIR)/$(MANIFEST_NAME)"

# Clean build files
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@go clean

# Show help
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run           - Build and run the application"
	@echo "  make install       - Build, install the app to $(INSTALL_DIR), and install Chrome host manifest"
	@echo "  make chrome-install - Install only the Chrome host manifest"
	@echo "  make clean         - Clean build files"
	@echo "  make help          - Show this help message" 