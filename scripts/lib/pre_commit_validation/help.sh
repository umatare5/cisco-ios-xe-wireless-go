#!/usr/bin/env bash

# Cisco WNC Pre-commit Validation - Help Functions
# Help content and documentation functions

set -euo pipefail

# Simple output functions for help (avoiding color conflicts)
help_info() {
    printf "ℹ Info: %s\n" "$*"
}

help_warning() {
    printf "⚠ Warning: %s\n" "$*" >&2
}

# Show detailed help for pre-commit validation
show_pre_commit_help() {
    cat << 'EOF'
Pre-commit Validation - Prevent Direct Pushes to Main Branch

DESCRIPTION:
    This script validates commits before they are made, preventing direct
    commits to main/master branches and enforcing proper development workflow
    through feature branches.

BEHAVIOR:
    - Blocks commits to 'main' or 'master' branches
    - Provides helpful workflow guidance
    - Allows commits on all other branches
    - Can be used as a Git pre-commit hook

USAGE:
    # Direct execution
    ./scripts/validate_pre_commit.sh [--verbose] [--no-color]

    # As Git hook (install)
    ln -sf ../../scripts/validate_pre_commit.sh .git/hooks/pre-commit

EXAMPLES:
    # Basic validation
    ./scripts/validate_pre_commit.sh

    # Verbose output
    ./scripts/validate_pre_commit.sh --verbose

    # No color output
    ./scripts/validate_pre_commit.sh --no-color

EMERGENCY BYPASS:
    If you need to make an emergency commit to main:

    # Temporarily disable hook
    mv .git/hooks/pre-commit .git/hooks/pre-commit.disabled

    # Make emergency commit
    git commit -m "emergency: fix critical issue"

    # Re-enable hook
    mv .git/hooks/pre-commit.disabled .git/hooks/pre-commit

ENVIRONMENT VARIABLES:
    NO_COLOR=1     Disable colored output (alternative to --no-color)

EXIT CODES:
    0    Success - commit allowed
    1    Error - commit blocked

EOF
}

# Show workflow guidance
show_workflow_guidance() {
    help_info "Recommended workflow:"
    help_info "1. Create a feature branch:"
    help_info "   git checkout -b feature/your-feature-name"
    echo
    help_info "2. Make your changes and commit:"
    help_info "   git add ."
    help_info "   git commit -m \"your commit message\""
    echo
    help_info "3. Push the feature branch:"
    help_info "   git push origin feature/your-feature-name"
    echo
    help_info "4. Create a pull request for review"
}

# Show emergency bypass instructions
show_emergency_bypass() {
    help_warning "If you need to make an emergency fix to main, you can:"
    help_warning "1. Temporarily disable this hook:"
    help_warning "   mv .git/hooks/pre-commit .git/hooks/pre-commit.disabled"
    help_warning "2. Make your emergency commit"
    help_warning "3. Re-enable the hook:"
    help_warning "   mv .git/hooks/pre-commit.disabled .git/hooks/pre-commit"
}

# Show installation instructions
show_installation_help() {
    help_info "To install as Git pre-commit hook:"
    help_info "   ln -sf ../../scripts/validate_pre_commit.sh .git/hooks/pre-commit"
    echo
    help_info "To test without installing:"
    help_info "   ./scripts/validate_pre_commit.sh"
}
