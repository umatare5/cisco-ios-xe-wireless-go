#!/usr/bin/env bash

# Cisco WNC Build Tools Library
# Functions for managing Go build tools and dependencies

# Prevent double sourcing
if [[ -n "${WNC_BUILD_TOOLS_LOADED:-}" ]]; then
    return 0
fi
readonly WNC_BUILD_TOOLS_LOADED=1

set -euo pipefail

# Default tool versions
readonly DEFAULT_GOLANGCI_LINT_VERSION="latest"
readonly DEFAULT_GORELEASER_VERSION="latest"
readonly DEFAULT_GOTESTSUM_VERSION="latest"

# Function to check if a command exists
command_exists() {
    local cmd="$1"
    command -v "$cmd" >/dev/null 2>&1
}

# Function to install golangci-lint
install_golangci_lint() {
    local version="${1:-$DEFAULT_GOLANGCI_LINT_VERSION}"

    if command_exists golangci-lint; then
        echo "✓ golangci-lint is already installed"
        return 0
    fi

    echo "Installing golangci-lint@${version}..."
    go install "github.com/golangci/golangci-lint/cmd/golangci-lint@${version}"

    if command_exists golangci-lint; then
        echo "✓ golangci-lint installed successfully"
    else
        echo "✗ Failed to install golangci-lint" >&2
        return 1
    fi
}

# Function to install goreleaser
install_goreleaser() {
    local version="${1:-$DEFAULT_GORELEASER_VERSION}"

    if command_exists goreleaser; then
        echo "✓ goreleaser is already installed"
        return 0
    fi

    echo "Installing goreleaser@${version}..."
    go install "github.com/goreleaser/goreleaser@${version}"

    if command_exists goreleaser; then
        echo "✓ goreleaser installed successfully"
    else
        echo "✗ Failed to install goreleaser" >&2
        return 1
    fi
}

# Function to install gotestsum
install_gotestsum() {
    local version="${1:-$DEFAULT_GOTESTSUM_VERSION}"

    if command_exists gotestsum; then
        echo "✓ gotestsum is already installed"
        return 0
    fi

    echo "Installing gotestsum@${version}..."
    go install "gotest.tools/gotestsum@${version}"

    if command_exists gotestsum; then
        echo "✓ gotestsum installed successfully"
    else
        echo "✗ Failed to install gotestsum" >&2
        return 1
    fi
}

# Function to install all development dependencies
install_dev_dependencies() {
    echo "Installing development dependencies..."

    # Install each tool with default versions
    install_golangci_lint "$DEFAULT_GOLANGCI_LINT_VERSION"
    install_goreleaser "$DEFAULT_GORELEASER_VERSION"
    install_gotestsum "$DEFAULT_GOTESTSUM_VERSION"

    echo "✓ All development dependencies installed!"
}

# Function to run linting
run_linting() {
    local project_root="${1:-.}"

    echo "Running linting..."

    if command_exists golangci-lint; then
        echo "Using golangci-lint for comprehensive linting..."
        (cd "$project_root" && golangci-lint run ./...)
        echo "✓ golangci-lint passed"
    else
        echo "golangci-lint not found, using go vet instead..."
        (cd "$project_root" && go vet ./...)
        echo "✓ go vet passed"
    fi
}

# Function to clean build artifacts
clean_build_artifacts() {
    local project_root="${1:-.}"

    echo "Cleaning build artifacts..."

    # Remove coverage files
    rm -f "$project_root/coverage.out"

    # Remove temporary directory
    rm -rf "$project_root/tmp"

    # Clean Go cache and test cache
    (cd "$project_root" && go clean -cache -testcache)

    echo "✓ Build artifacts cleaned"
}

# Function to ensure tmp directory exists
ensure_tmp_directory() {
    local project_root="${1:-.}"
    local tmp_dir="$project_root/tmp"

    if [[ ! -d "$tmp_dir" ]]; then
        mkdir -p "$tmp_dir"
        echo "✓ Created tmp directory: $tmp_dir"
    fi
}

# Function to validate Go module
validate_go_module() {
    local project_root="${1:-.}"

    echo "Validating Go module..."

    if [[ ! -f "$project_root/go.mod" ]]; then
        echo "✗ go.mod not found in $project_root" >&2
        return 1
    fi

    (cd "$project_root" && go mod verify)
    echo "✓ Go module validated"
}

# Function to run go build to verify code compiles
verify_build() {
    local project_root="${1:-.}"

    echo "Verifying build..."
    (cd "$project_root" && go build ./...)
    echo "✓ Build verification passed"
}
