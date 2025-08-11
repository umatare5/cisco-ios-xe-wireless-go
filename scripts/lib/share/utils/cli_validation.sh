#!/usr/bin/env bash
#
# CLI Validation Library
# Functions for validating required CLI tools before script execution
#

# Check if a single CLI tool is available
is_cli_available() {
    local cli_name="${1:-}"
    [[ -n "$cli_name" ]] && command -v "$cli_name" >/dev/null 2>&1
}

# Get CLI tool description
get_cli_description() {
    local cli_name="${1:-}"
    case "$cli_name" in
    "bash") printf '%s\n' "Shell interpreter" ;;
    "bc") printf '%s\n' "Basic calculator for coverage calculations" ;;
    "curl") printf '%s\n' "HTTP client for API requests" ;;
    "date") printf '%s\n' "Date/time operations" ;;
    "go") printf '%s\n' "Go toolchain for building and testing" ;;
    "gotestsum") printf '%s\n' "Enhanced Go test runner" ;;
    "golangci-lint") printf '%s\n' "Go linting tool" ;;
    "shellcheck") printf '%s\n' "Shell script linting tool" ;;
    "markdownlint-cli2") printf '%s\n' "Markdown linting tool" ;;
    "grep") printf '%s\n' "Text pattern matching" ;;
    "awk") printf '%s\n' "Text processing" ;;
    "sed") printf '%s\n' "Stream editor" ;;
    "head") printf '%s\n' "Display first lines of files" ;;
    "tail") printf '%s\n' "Display last lines of files" ;;
    "cat") printf '%s\n' "Display file contents" ;;
    "cut") printf '%s\n' "Extract columns from text" ;;
    "tr") printf '%s\n' "Character translation" ;;
    "sort") printf '%s\n' "Sort lines of text" ;;
    "uniq") printf '%s\n' "Report or filter unique lines" ;;
    "wc") printf '%s\n' "Word/line/character count" ;;
    "find") printf '%s\n' "Find files and directories" ;;
    "xargs") printf '%s\n' "Execute commands from standard input" ;;
    "stat") printf '%s\n' "Display file statistics" ;;
    "dirname") printf '%s\n' "Extract directory from path" ;;
    "basename") printf '%s\n' "Extract filename from path" ;;
    "mkdir") printf '%s\n' "Create directories" ;;
    "rm") printf '%s\n' "Remove files and directories" ;;
    "open") printf '%s\n' "Open files (macOS)" ;;
    "command") printf '%s\n' "Check command availability" ;;
    "type") printf '%s\n' "Display command type" ;;
    "which") printf '%s\n' "Locate command" ;;
    "cd") printf '%s\n' "Change directory" ;;
    "pwd") printf '%s\n' "Print working directory" ;;
    "echo") printf '%s\n' "Display text" ;;
    "printf") printf '%s\n' "Format and print text" ;;
    "env") printf '%s\n' "Set environment variables" ;;
    "export") printf '%s\n' "Set environment variables" ;;
    "set") printf '%s\n' "Set shell options" ;;
    "source") printf '%s\n' "Execute shell script" ;;
    "eval") printf '%s\n' "Evaluate and execute command" ;;
    *) printf '%s\n' "Unknown CLI tool" ;;
    esac
}

# Get CLI tools for a category
get_cli_tools_for_category() {
    local category="${1:-}"
    case "$category" in
    "essential") printf '%s\n' "go" ;;  # Only Go for essential tools
    "go_tools") printf '%s\n' "go gotestsum golangci-lint shellcheck markdownlint-cli2" ;;
    "http_client") printf '%s\n' "curl" ;;
    "text_processing") printf '%s\n' "" ;;  # Remove standard text tools
    "system_tools") printf '%s\n' "" ;;     # Remove standard system tools
    "optional") printf '%s\n' "open bc" ;;  # Keep macOS-specific and optional tools
    *) printf '%s\n' "" ;;
    esac
}

# Get CLI tools for a validation level
get_cli_tools_for_level() {
    local level="${1:-standard}"
    local categories tools=""

    case "$level" in
        "strict")
            categories="essential go_tools http_client optional"
            ;;
        "standard")
            categories="essential go_tools http_client"
            ;;
        "minimal")
            categories="essential"
            ;;
        *)
            categories="essential"
            ;;
    esac

    for category in $categories; do
        local category_tools
        category_tools=$(get_cli_tools_for_category "$category")
        tools="$tools $category_tools"
    done

    printf '%s\n' "$tools" | tr ' ' '\n' | sort -u | tr '\n' ' '
}

# Validate a single CLI tool
validate_cli_tool() {
    local cli_name="${1:-}"
    local required="${2:-true}"

    if [[ -z "$cli_name" ]]; then
        return 1
    fi

    if is_cli_available "$cli_name"; then
        return 0
    fi

    # CLI tool not found
    if [[ "$required" == "true" ]]; then
    printf '✗ %s: %s\n' "$cli_name" "$(get_cli_description "$cli_name")" >&2
        return 1
    else
    printf '⚠ %s: %s (optional)\n' "$cli_name" "$(get_cli_description "$cli_name")" >&2
        return 0
    fi
}

# Validate multiple CLI tools
validate_cli_tools() {
    local tools=("$@")
    local missing_count=0
    local total_count=${#tools[@]}

    if [[ $total_count -eq 0 ]]; then
        printf '%s\n' "No CLI tools to validate" >&2
        return 0
    fi

    printf '%s\n' "Validating CLI tools..." >&2

    for tool in "${tools[@]}"; do
        if ! validate_cli_tool "$tool" "true"; then
            ((missing_count++))
        fi
    done

    if [[ $missing_count -eq 0 ]]; then
        printf '✓ All %d CLI tools are available\n' "$total_count" >&2
        return 0
    else
        printf '✗ %d out of %d CLI tools are missing\n' "$missing_count" "$total_count" >&2
        return 1
    fi
}

# Show installation instructions for missing CLI tools
show_cli_installation_instructions() {
    local missing_tools=("$@")

    if [[ ${#missing_tools[@]} -eq 0 ]]; then
        return 0
    fi

    printf '\n' >&2
    printf '%s\n' "Installation instructions for missing CLI tools:" >&2
    printf '\n' >&2

    for tool in "${missing_tools[@]}"; do
        case "$tool" in
            "go")
                printf '  %s: %s\n' "$tool" "https://golang.org/dl/" >&2
                ;;
            "gotestsum")
                printf '  %s: %s\n' "$tool" "go install gotest.tools/gotestsum@latest" >&2
                ;;
            "golangci-lint")
                printf '  %s: %s\n' "$tool" "go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" >&2
                ;;
            "shellcheck")
                printf '  %s: %s\n' "$tool" "brew install shellcheck (macOS) or sudo apt-get install shellcheck (Ubuntu)" >&2
                ;;
            "markdownlint-cli2")
                printf '  %s: %s\n' "$tool" "npm install -g markdownlint-cli2" >&2
                ;;
            "curl")
                printf '  %s: %s\n' "$tool" "Pre-installed on most systems, or brew install curl" >&2
                ;;
            "bc")
                printf '  %s: %s\n' "$tool" "brew install bc (macOS) or sudo apt-get install bc (Ubuntu)" >&2
                ;;
            *)
                printf '  %s: %s - usually pre-installed\n' "$tool" "$(get_cli_description "$tool")" >&2
                ;;
        esac
    done

    printf '\n' >&2
}

# Main validation function with detailed reporting
validate_required_cli_tools() {
    local level="${1:-standard}"
    local show_instructions="${2:-true}"

    local tools_str
    tools_str=$(get_cli_tools_for_level "$level")

    # Convert to array
    local tools missing_tools
    IFS=' ' read -ra tools <<< "$tools_str"
    missing_tools=()

    printf 'Validating CLI tools (level: %s)...\n' "$level" >&2

    for tool in "${tools[@]}"; do
        if is_cli_available "$tool"; then
            printf '✓ %s\n' "$tool" >&2
        else
            printf '✗ %s\n' "$tool" >&2
            missing_tools+=("$tool")
        fi
    done

    local missing_count=${#missing_tools[@]}
    local total_count=${#tools[@]}

    if [[ $missing_count -eq 0 ]]; then
    printf '\n' >&2
    printf '✓ All %d required CLI tools are available\n' "$total_count" >&2
        return 0
    else
    printf '\n' >&2
    printf '✗ %d out of %d CLI tools are missing\n' "$missing_count" "$total_count" >&2

        if [[ "$show_instructions" == "true" ]]; then
            show_cli_installation_instructions "${missing_tools[@]}"
        fi

        return 1
    fi
}

# Quick validation function for use in scripts
require_cli_tools() {
    local level="${1:-standard}"

    if ! validate_required_cli_tools "$level" "true" >/dev/null 2>&1; then
        validate_required_cli_tools "$level" "true"
        printf '\n' >&2
        printf '%s\n' "Please install the missing CLI tools and try again." >&2
        exit 1
    fi
}

# Export for child scripts
export is_cli_available get_cli_description validate_cli_tool validate_cli_tools validate_required_cli_tools require_cli_tools
