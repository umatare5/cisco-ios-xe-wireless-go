#!/usr/bin/env bash

# Development help for this repository. Self-contained: it prints text and nothing else, so
# it needs no lib/ at all. Its stdout is quoted verbatim in docs/SCRIPT_REFERENCE.md.

set -Eeuo pipefail

main() {
    printf '%s\n' "Cisco WNC Development Scripts"
    printf '%s\n' "-------------------------------"
    cat <<'EOF'

USAGE:
    make <target>                   # Use Makefile targets (recommended)
    ./scripts/<script>.sh [options] # Use scripts directly

COMMON DEVELOPMENT TARGETS:
    help                Show this help message
    deps                Install development dependencies
    lint                Run code linting tools
    build               Verify build compilation
    test-unit           Run unit tests only
    test-integration    Run integration tests (requires environment)
    test-unit-coverage  Run unit tests with coverage analysis

ENVIRONMENT VARIABLES:
    WNC_CONTROLLER      Controller hostname/IP for integration tests
    WNC_ACCESS_TOKEN    Base64 encoded credentials for integration tests

EXAMPLES:
    # Basic development workflow
    make deps               # Install dependencies
    make lint               # Check code quality
    make test-unit          # Run unit tests
    make test-unit-coverage # Run unit tests with coverage
    make build              # Verify compilation

    # Integration testing (requires environment setup)
    export WNC_CONTROLLER="<controller-host-or-ip>"
    export WNC_ACCESS_TOKEN="<base64-username:password>"
    make test-integration

SCRIPT DETAILS:
    Every script takes no arguments.

    Available scripts:
    - install_dependencies.sh Install Go dependencies
    - lint.sh                Run golangci-lint
    - test_unit.sh           Run unit tests
    - test_coverage.sh       Run unit tests with coverage and an HTML report
    - test_integration.sh    Run integration tests

PROJECT STRUCTURE:
    scripts/                Script directory
    +-- <command>.sh       One entry point per Make target
    +-- lib/               Sourced explicitly by the entry points
        +-- log.sh         Colour, log lines and the banner
        +-- env.sh         Toolchain and project-directory checks
        +-- gotest.sh      The gotestsum runs behind the test targets
EOF
}

main "$@"
