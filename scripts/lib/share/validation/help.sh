#!/usr/bin/env bash

# Cisco WNC Pre-commit Validation - Help Functions
# Provides help and documentation functionality for pre-commit validation

# Display main help content for pre-commit validation
_show_pre_commit_main_help() {
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
    ./scripts/pre_commit_hook.sh [--verbose] [--no-color]

    # As Git hook (install)
    ln -sf ../../scripts/pre_commit_hook.sh .git/hooks/pre-commit

EXAMPLES:
    # Basic validation
    ./scripts/pre_commit_hook.sh

    # Verbose output
    ./scripts/pre_commit_hook.sh --verbose

    # No color output
    ./scripts/pre_commit_hook.sh --no-color
EOF
}

# Display emergency bypass instructions
_show_pre_commit_emergency_help() {
    cat << 'EOF'

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

# Display comprehensive help documentation for pre-commit validation
show_pre_commit_help() {
    _show_pre_commit_main_help
    _show_pre_commit_emergency_help
}

# Display recommended Git workflow for feature branch development
show_workflow_guidance() {
    info "Recommended workflow:"
    info "1. Create a feature branch:"
    info "   git checkout -b feature/your-feature-name"
    printf '\n'
    info "2. Make your changes and commit:"
    info "   git add ."
    info "   git commit -m \"your commit message\""
    printf '\n'
    info "3. Push the feature branch:"
    info "   git push origin feature/your-feature-name"
    printf '\n'
    info "4. Create a pull request for review"
}

# Display emergency bypass procedures for urgent main branch commits
show_emergency_bypass() {
    warn "If you need to make an emergency fix to main, you can:"
    warn "1. Temporarily disable this hook:"
    warn "   mv .git/hooks/pre-commit .git/hooks/pre-commit.disabled"
    warn "2. Make your emergency commit"
    warn "3. Re-enable the hook:"
    warn "   mv .git/hooks/pre-commit.disabled .git/hooks/pre-commit"
}

# Display instructions for installing pre-commit hook
show_installation_help() {
    info "To install as Git pre-commit hook:"
    info "   ln -sf ../../scripts/pre_commit_hook.sh .git/hooks/pre-commit"
    printf '\n'
    info "To test without installing:"
    info "   ./scripts/pre_commit_hook.sh"
}
