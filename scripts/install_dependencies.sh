#!/usr/bin/env bash

# Installs the four Go tools this repository builds and tests with, then tidies and
# downloads the module. Takes no arguments; every tool is installed at @latest.

set -Eeuo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR
readonly PROJECT_ROOT="${SCRIPT_DIR%/*}"

# shellcheck source=lib/log.sh
source "${SCRIPT_DIR}/lib/log.sh"
# shellcheck source=lib/env.sh
source "${SCRIPT_DIR}/lib/env.sh"

# _install <binary> <module path> <version>
_install() {
    local -r bin="$1" module="$2" version="$3"

    if command_exists "$bin"; then
        printf '✓ %s\n' "$bin is already installed"
        return 0
    fi

    printf 'Installing %s@%s...\n' "$bin" "$version"
    go install "${module}@${version}"
    command_exists "$bin" || fail "Failed to install $bin"
    printf '✓ %s\n' "$bin installed successfully"
}

# The gitleaks module still declares the former org in its path, so go install rejects
# github.com/gitleaks.
_install_tools() {
    printf '%s\n' "Installing development dependencies..."
    _install golangci-lint github.com/golangci/golangci-lint/v2/cmd/golangci-lint latest
    _install goreleaser github.com/goreleaser/goreleaser/v2 latest
    _install gotestsum gotest.tools/gotestsum latest
    _install gitleaks github.com/zricethezav/gitleaks/v8 latest
    printf '✓ %s\n' "All development dependencies installed!"
}

_assert_module() {
    local -r project_root="$1"

    dir_exists "$project_root" || fail "Invalid project directory: $project_root"
    file_exists "$project_root/go.mod" ||
        fail "No go.mod found in $project_root" "Run 'go mod init' to initialize a Go module"
    command_exists go || fail "Go is not installed or not in PATH"
    info "Using Go version: $(go version 2>/dev/null | cut -d' ' -f3)"
}

_sync_module() {
    format_step_message '1' 'Downloading dependencies...'
    go mod tidy 2>/dev/null || fail "Failed to tidy dependencies"
    success "Dependencies tidied"
    go mod download 2>/dev/null || fail "Failed to download dependencies"
    success "Dependencies downloaded"
    printf '\n'
}

# printf '%s' withholds the trailing newline, so the main module's own line is not counted
# and a dependency-free module reports 0 rather than 1.
_count_direct() {
    local listed
    listed=$(go list -mod=readonly -m -f '{{if not .Indirect}}{{.Path}}{{end}}' all 2>/dev/null)
    printf '%s' "$listed" | sed '/^$/d' | wc -l | tr -d ' '
}

_count_total() {
    go list -mod=readonly -m all 2>/dev/null | tail -n +2 | wc -l | tr -d ' '
}

_summary() {
    format_step_message '✓' 'Dependencies management completed'
    info "Direct dependencies: $(_count_direct)"
    info "Total dependencies: $(_count_total)"
}

# Entrypoint: install tools, then sync the module
main() {
    validate_required_cli_tools "minimal"
    wnc_banner "Cisco WNC Dependencies" "Module Management"
    _assert_module "$PROJECT_ROOT"
    cd "$PROJECT_ROOT" || fail "Cannot enter $PROJECT_ROOT"

    _install_tools
    _sync_module
    _summary
}

main "$@"
