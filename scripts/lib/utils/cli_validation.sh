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
        "bash") echo "Shell interpreter" ;;
        "bc") echo "Basic calculator for coverage calculations" ;;
        "curl") echo "HTTP client for API requests" ;;
        "date") echo "Date/time operations" ;;
        "go") echo "Go toolchain for building and testing" ;;
        "gotestsum") echo "Enhanced Go test runner" ;;
        "golangci-lint") echo "Go linting tool" ;;
        "grep") echo "Text pattern matching" ;;
        "awk") echo "Text processing" ;;
        "sed") echo "Stream editor" ;;
        "head") echo "Display first lines of files" ;;
        "tail") echo "Display last lines of files" ;;
        "cat") echo "Display file contents" ;;
        "cut") echo "Extract columns from text" ;;
        "tr") echo "Character translation" ;;
        "sort") echo "Sort lines of text" ;;
        "uniq") echo "Report or filter unique lines" ;;
        "wc") echo "Word/line/character count" ;;
        "find") echo "Find files and directories" ;;
        "xargs") echo "Execute commands from standard input" ;;
        "stat") echo "Display file statistics" ;;
        "dirname") echo "Extract directory from path" ;;
        "basename") echo "Extract filename from path" ;;
        "mkdir") echo "Create directories" ;;
        "rm") echo "Remove files and directories" ;;
        "open") echo "Open files (macOS)" ;;
        "command") echo "Check command availability" ;;
        "type") echo "Display command type" ;;
        "which") echo "Locate command" ;;
        "cd") echo "Change directory" ;;
        "pwd") echo "Print working directory" ;;
        "echo") echo "Display text" ;;
        "printf") echo "Format and print text" ;;
        "env") echo "Set environment variables" ;;
        "export") echo "Set environment variables" ;;
        "set") echo "Set shell options" ;;
        "source") echo "Execute shell script" ;;
        "eval") echo "Evaluate and execute command" ;;
        *) echo "Unknown CLI tool" ;;
    esac
}

# Get CLI tools for a category
get_cli_tools_for_category() {
    local category="${1:-}"
    case "$category" in
        "essential") echo "go" ;;  # Only Go for essential tools
        "go_tools") echo "go gotestsum golangci-lint" ;;
        "http_client") echo "curl" ;;
        "text_processing") echo "" ;;  # Remove standard text tools
        "system_tools") echo "" ;;     # Remove standard system tools
        "optional") echo "open bc" ;;  # Keep macOS-specific and optional tools
        *) echo "" ;;
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

    echo "$tools" | tr ' ' '\n' | sort -u | tr '\n' ' '
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
        echo "✗ $cli_name: $(get_cli_description "$cli_name")" >&2
        return 1
    else
        echo "⚠ $cli_name: $(get_cli_description "$cli_name") (optional)" >&2
        return 0
    fi
}

# Validate multiple CLI tools
validate_cli_tools() {
    local tools=("$@")
    local missing_count=0
    local total_count=${#tools[@]}

    if [[ $total_count -eq 0 ]]; then
        echo "No CLI tools to validate" >&2
        return 0
    fi

    echo "Validating CLI tools..." >&2

    for tool in "${tools[@]}"; do
        if ! validate_cli_tool "$tool" "true"; then
            ((missing_count++))
        fi
    done

    if [[ $missing_count -eq 0 ]]; then
        echo "✓ All $total_count CLI tools are available" >&2
        return 0
    else
        echo "✗ $missing_count out of $total_count CLI tools are missing" >&2
        return 1
    fi
}

# Show installation instructions for missing CLI tools
show_cli_installation_instructions() {
    local missing_tools=("$@")

    if [[ ${#missing_tools[@]} -eq 0 ]]; then
        return 0
    fi

    echo "" >&2
    echo "Installation instructions for missing CLI tools:" >&2
    echo "" >&2

    for tool in "${missing_tools[@]}"; do
        case "$tool" in
            "go")
                echo "  $tool: https://golang.org/dl/" >&2
                ;;
            "gotestsum")
                echo "  $tool: go install gotest.tools/gotestsum@latest" >&2
                ;;
            "golangci-lint")
                echo "  $tool: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" >&2
                ;;
            "curl")
                echo "  $tool: Pre-installed on most systems, or brew install curl" >&2
                ;;
            "bc")
                echo "  $tool: brew install bc (macOS) or sudo apt-get install bc (Ubuntu)" >&2
                ;;
            *)
                echo "  $tool: $(get_cli_description "$tool") - usually pre-installed" >&2
                ;;
        esac
    done

    echo "" >&2
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

    echo "Validating CLI tools (level: $level)..." >&2

    for tool in "${tools[@]}"; do
        if is_cli_available "$tool"; then
            echo "✓ $tool" >&2
        else
            echo "✗ $tool" >&2
            missing_tools+=("$tool")
        fi
    done

    local missing_count=${#missing_tools[@]}
    local total_count=${#tools[@]}

    if [[ $missing_count -eq 0 ]]; then
        echo "" >&2
        echo "✓ All $total_count required CLI tools are available" >&2
        return 0
    else
        echo "" >&2
        echo "✗ $missing_count out of $total_count CLI tools are missing" >&2

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
        echo "" >&2
        echo "Please install the missing CLI tools and try again." >&2
        exit 1
    fi
}

# Export functions for use in other scripts
export -f is_cli_available
export -f get_cli_description
export -f validate_cli_tool
export -f validate_cli_tools
export -f validate_required_cli_tools
export -f require_cli_tools
