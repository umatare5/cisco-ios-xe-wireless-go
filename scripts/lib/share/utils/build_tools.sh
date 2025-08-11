#!/usr/bin/env bash

# Cisco WNC Build Tools Library
# Functions for managing Go build tools and linting operations

# Default tool versions (idempotent)
: "${DEFAULT_GOLANGCI_LINT_VERSION:=latest}"
readonly DEFAULT_GOLANGCI_LINT_VERSION
: "${DEFAULT_GORELEASER_VERSION:=latest}"
readonly DEFAULT_GORELEASER_VERSION
: "${DEFAULT_GOTESTSUM_VERSION:=latest}"
readonly DEFAULT_GOTESTSUM_VERSION

# Run linting operations for Go code, shell scripts, and markdown files
run_lints() {
    local exit_code=0

    printf 'Running linting operations...\n' >&2

    lint_go || exit_code=1
    lint_shell || exit_code=1
    lint_markdown || exit_code=1

    show_linting_summary "$exit_code"
    return $exit_code
}

# Run Go code linting with golangci-lint or go vet fallback
lint_go() {
    printf '\n--- Go Code Linting (golangci-lint) ---\n' >&2

    if command -v golangci-lint >/dev/null 2>&1; then
        golangci-lint run --timeout=10m
    else
        printf 'Warning: golangci-lint not found, falling back to go vet\n' >&2
        go vet ./...
    fi
}

# Run shell script linting with shellcheck
lint_shell() {
    printf '\n--- Shell Script Linting (shellcheck) ---\n' >&2

    if ! command -v shellcheck >/dev/null 2>&1; then
        printf 'Warning: shellcheck not found, skipping shell script linting\n' >&2
        return 0
    fi

    local shell_files
    shell_files=$(find . -type f -name "*.sh" | grep -v vendor/ | grep -v ".git/" | head -20)

    if [[ -z "$shell_files" ]]; then
        printf 'No shell script files found to lint\n' >&2
        return 0
    fi

    echo "$shell_files" | xargs shellcheck
}

# Run markdown linting with markdownlint-cli2
lint_markdown() {
    printf '\n--- Markdown Linting (markdownlint-cli2) ---\n' >&2

    if ! command -v markdownlint-cli2 >/dev/null 2>&1; then
        printf 'Warning: markdownlint-cli2 not found, skipping markdown linting\n' >&2
        return 0
    fi

    markdownlint-cli2
}

# Show linting summary based on exit code
show_linting_summary() {
    local exit_code="${1:-0}"

    printf '\n--- Linting Summary ---\n' >&2

    if [[ "$exit_code" -ne 0 ]]; then
        printf 'Linting failed - please fix the issues above\n' >&2
        printf '\n' >&2
        return
    fi

    printf 'All linting checks passed successfully\n' >&2
    printf '\n' >&2
}

# Check if a command exists in the system PATH
command_exists() {
    local cmd="$1"
    command -v "$cmd" >/dev/null 2>&1
}

# Install golangci-lint Go linter tool with specified version
install_golangci_lint() {
    local version="${1:-$DEFAULT_GOLANGCI_LINT_VERSION}"

    if command_exists golangci-lint; then
        printf '✓ %s\n' "golangci-lint is already installed"
        return 0
    fi

    printf 'Installing golangci-lint@%s...\n' "${version}"
    go install "github.com/golangci/golangci-lint/cmd/golangci-lint@${version}"

    if ! command_exists golangci-lint; then
        printf '✗ %s\n' "Failed to install golangci-lint" >&2
        return 1
    fi

    printf '✓ %s\n' "golangci-lint installed successfully"
}

# Install goreleaser Go release automation tool with specified version
install_goreleaser() {
    local version="${1:-$DEFAULT_GORELEASER_VERSION}"

    if command_exists goreleaser; then
        printf '✓ %s\n' "goreleaser is already installed"
        return 0
    fi

    printf 'Installing goreleaser@%s...\n' "${version}"
    go install "github.com/goreleaser/goreleaser@${version}"

    if ! command_exists goreleaser; then
        printf '✗ %s\n' "Failed to install goreleaser" >&2
        return 1
    fi

    printf '✓ %s\n' "goreleaser installed successfully"
}

# Install gotestsum Go test runner tool with pretty output formatting
install_gotestsum() {
    local version="${1:-$DEFAULT_GOTESTSUM_VERSION}"

    if command_exists gotestsum; then
        printf '✓ %s\n' "gotestsum is already installed"
        return 0
    fi

    printf 'Installing gotestsum@%s...\n' "${version}"
    go install "gotest.tools/gotestsum@${version}"

    if ! command_exists gotestsum; then
        printf '✗ %s\n' "Failed to install gotestsum" >&2
        return 1
    fi

    printf '✓ %s\n' "gotestsum installed successfully"
}

# Install all required development dependencies for the Go project
install_dev_dependencies() {
    printf '%s\n' "Installing development dependencies..."

    # Install each tool with default versions
    install_golangci_lint "$DEFAULT_GOLANGCI_LINT_VERSION"
    install_goreleaser "$DEFAULT_GORELEASER_VERSION"
    install_gotestsum "$DEFAULT_GOTESTSUM_VERSION"

    printf '✓ %s\n' "All development dependencies installed!"
}

# Clean build artifacts including coverage files, tmp directory, and Go caches
clean_build_artifacts() {
    local project_root="${1:-.}"

    printf '%s\n' "Cleaning build artifacts..."

    # Remove coverage files
    rm -f "$project_root/coverage.out"

    # Remove temporary directory
    rm -rf "$project_root/tmp"

    # Clean Go cache and test cache
    (cd "$project_root" && go clean -cache -testcache)

    printf '✓ %s\n' "Build artifacts cleaned"
}

# Check if tmp directory exists
has_tmp_directory() {
    local project_root="$1"
    local tmp_dir="$project_root/tmp"
    [[ -d "$tmp_dir" ]]
}

# Ensure tmp directory exists in project root for temporary artifacts
ensure_tmp_directory() {
    local project_root="${1:-.}"
    local tmp_dir="$project_root/tmp"

    if has_tmp_directory "$project_root"; then
        return
    fi

    mkdir -p "$tmp_dir"
    printf '✓ Created tmp directory: %s\n' "$tmp_dir"
}

# Validate Go module integrity by checking go.mod exists and running go mod verify
validate_go_module() {
    local project_root="${1:-.}"

    printf '%s\n' "Validating Go module..."

    if [[ ! -f "$project_root/go.mod" ]]; then
        printf '✗ go.mod not found in %s\n' "$project_root" >&2
        return 1
    fi

    (cd "$project_root" && go mod verify)
    printf '✓ %s\n' "Go module validated"
}

# Verify build compiles successfully by running go build on all packages
verify_build() {
    local project_root="${1:-.}"

    printf '%s\n' "Verifying build..."
    (cd "$project_root" && go build ./...)
    printf '✓ %s\n' "Build verification passed"
}
