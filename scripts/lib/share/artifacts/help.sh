#!/usr/bin/env bash

# Artifacts Cleanup Help Functions (usage text and examples)
# Provides help and documentation functionality for artifacts cleanup

show_artifacts_help() {
    cat << 'EOF'
Cisco WNC Artifacts Cleanup - Clean build artifacts, temporary files, and caches

USAGE:
    clean_artifacts.sh [OPTIONS]

OPTIONS:
    -p, --project <DIR>         Project root directory (default: .)
    -v, --verbose               Enable verbose output
        --go-cache              Clean Go build cache
        --go-modules            Clean Go module cache
        --temp-files            Clean temporary files (./tmp)
        --test-files            Clean test artifacts (.test binaries, coverage files)
        --all                   Clean all artifacts (default)
        --dry-run               Show what would be cleaned without actually cleaning
        --no-color              Disable colored output
    -h, --help                  Show this help message

EXAMPLES:
    # Clean all artifacts (default)
    clean_artifacts.sh

    # Clean only Go cache
    clean_artifacts.sh --go-cache

    # Clean with verbose output
    clean_artifacts.sh --verbose

    # Show what would be cleaned
    clean_artifacts.sh --dry-run

    # Clean specific project
    clean_artifacts.sh --project ./my-project

CLEANUP TARGETS:
    Go Build Cache        $(go env GOCACHE)
    Go Module Cache       $(go env GOMODCACHE)
    Temporary Files       ./tmp/
    Test Binaries         ./**/*.test
    Coverage Files        ./**/*.out, ./**/*.html

This script safely removes build artifacts and temporary files
to free up disk space and ensure clean builds.
EOF
}
