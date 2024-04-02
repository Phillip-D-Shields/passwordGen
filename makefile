# Makefile for building and analyzing your Cobra CLI

# Set the name of your Go application
APP_NAME := passwordGen

# Set the version of your Go application
VERSION := 1.0.0

# Set the build output directory
BUILD_DIR := build

# Set the installation directory
INSTALL_DIR := $(HOME)/go/bin # or any other directory in your PATH

# Build variables
BINARY_FILE := $(BUILD_DIR)/$(APP_NAME)

.PHONY: build run clean install

# Build the application binary
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BINARY_FILE)

# Run the application
run: build
	$(BINARY_FILE) password -l 20

# Clean up build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Install the application
install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_FILE) $(INSTALL_DIR)

# Uninstall the application
uninstall:
	rm -f $(INSTALL_DIR)/$(APP_NAME)