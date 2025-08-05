---
description: Bash Shell Scripts Instructions
applyTo: "**/*.sh"
---

# GitHub Copilot Agent Mode – Bash Shell Scripts Instructions

This file contains specific instructions for creating and maintaining shell scripts in this repository.

Copilot **MUST** comply with all instructions described in this document when editing or creating any bash shell scripts in this repository.

However, when there are conflicts between this document and `general.instructions.md`, **ALWAYS** prioritize the instructions in `general.instructions.md`.

---

## 📣 General Requirements

All shell scripts MUST:

1. **Follow Google Shell Style Guide**: Use standard conventions for variable naming, function definition, and code structure
2. **Include shellcheck compliance**: All scripts must pass shellcheck linting without warnings
3. **Use strict error handling**: Include `set -euo pipefail` at the beginning of every script
4. **Support verbose output**: Include `--verbose` flag where appropriate for debugging
5. **Implement proper help**: Every script must have a `--help` option with clear usage instructions
6. **Handle edge cases**: Validate inputs and provide meaningful error messages
7. **Use consistent formatting**: 120-character line limit, consistent indentation (2 spaces)
8. **Use argc for argument parsing**: All scripts should use sigoden/argc for clean CLI interfaces
9. **Implement modular architecture**: Use script-specific library directories with help.sh, output.sh, and core.sh patterns
10. **Include predicate functions**: Use predicate functions for improved conditional statement readability

---

## � Architecture Patterns

### Entry Point Script (argc-based)
```bash
#!/usr/bin/env bash
# @meta author @umatare5
# @meta version 1.0.0
# @meta description Brief description of script functionality
#
# @flag --verbose         Enable verbose output
# @flag --no-color        Disable colored output
# @flag --insecure        Skip SSL certificate verification for WNC connections
# @arg target             Target parameter description

set -euo pipefail

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# shellcheck source=lib/common/common.sh
source "${SCRIPT_DIR}/lib/common/common.sh"

SOURCE_WNC_LIBRARIES "${SCRIPT_DIR}"

# shellcheck source=lib/script_name/help.sh
source "${SCRIPT_DIR}/lib/script_name/help.sh"
# shellcheck source=lib/script_name/output.sh
source "${SCRIPT_DIR}/lib/script_name/output.sh"
# shellcheck source=lib/script_name/core.sh
source "${SCRIPT_DIR}/lib/script_name/core.sh"

main() {
    if is_verbose_enabled; then
        echo_info "Starting script execution..."
    fi

    execute_script_logic "$@"
}

main "$@"
```

### Library Directory Structure
```
scripts/lib/script_name/
├── help.sh     # Help content and documentation functions
├── output.sh   # Output formatting and display functions
├── core.sh     # Main business logic implementation
```

### Function Design and Naming Conventions
- **Global functions**: Use UPPERCASE names (e.g., `VALIDATE_ARGC_YANG_ENVIRONMENT()`, `SOURCE_WNC_LIBRARIES()`)
- **Local/private functions**: Use lowercase names (e.g., `execute_script_logic()`, `process_response()`)
- Keep functions between 10-20 lines ideally
- Include error handling within functions
- Use `readonly` for constants
- Use `local` for function-scoped variables

### Predicate Functions
All scripts should use predicate functions for conditional statements:
```bash
# Good: Using predicate functions
is_verbose_enabled() { [[ "${argc_verbose:-0}" == "1" ]]; }
is_no_color_enabled() { [[ "${argc_no_color:-0}" == "1" ]]; }
is_insecure_enabled() { [[ "${argc_insecure:-0}" == "1" ]]; }

# Usage in conditionals
if is_verbose_enabled; then
    echo_info "Verbose mode enabled"
fi

if is_insecure_enabled; then
    curl_args+=("--insecure")
fi
```

### Library Management
- Common utilities go in `scripts/lib/common/common.sh`
- YANG-specific utilities go in `scripts/lib/common/yang_common.sh`
- argc-specific utilities go in `scripts/lib/common/argc_common.sh`
- Testing utilities go in `scripts/lib/common/testing.sh`
- Use consistent source patterns with shellcheck directives
- Load libraries using `SOURCE_WNC_LIBRARIES()` function

---

## 📣 argc Integration Patterns

### Meta Information
Every argc script must include:
```bash
# @meta author @umatare5
# @meta version 1.0.0
# @meta description Clear description of script functionality
```

### Common Flags and Arguments
```bash
# Standard flags for all scripts
# @flag --verbose         Enable verbose output
# @flag --no-color        Disable colored output

# YANG operation scripts
# @flag --insecure        Skip SSL certificate verification for WNC connections

# Optional arguments with validation
# @arg model              YANG model name (required for model-specific operations)
# @arg path               YANG path (optional, defaults to root)
```

### argc Variable Usage
argc generates variables with numeric values for flags:
- `argc_verbose` → "1" (enabled) or "0" (disabled)
- `argc_no_color` → "1" (enabled) or "0" (disabled)
- `argc_insecure` → "1" (enabled) or "0" (disabled)

**Critical**: Use predicate functions for flag checks, not direct string comparison.

---

## 📣 YANG Integration Patterns

Scripts that interact with YANG models and RESTCONF APIs must:

1. **Environment Validation**: Check for required `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN` variables
2. **SSL Handling**: Support `--insecure` flag for development environments
3. **Authentication**: Use proper Bearer token authentication
4. **Error Handling**: Handle HTTP errors gracefully with meaningful messages
5. **Response Processing**: Parse JSON responses appropriately
6. **macOS Compatibility**: Use global arrays instead of `local -n` for bash 3.x compatibility

### Environment Variables
- `WNC_CONTROLLER`: Wireless controller hostname/IP (required)
- `WNC_ACCESS_TOKEN`: Base64-encoded authentication token (required)

**Important**: The legacy `WNC_CONTROLLERS` variable is deprecated and should not be used.

### Example YANG Script Pattern
```bash
validate_yang_environment() {
    if [[ -z "${WNC_CONTROLLER:-}" ]]; then
        echo_error "WNC_CONTROLLER environment variable is required"
        exit 1
    fi

    if [[ -z "${WNC_ACCESS_TOKEN:-}" ]]; then
        echo_error "WNC_ACCESS_TOKEN environment variable is required"
        exit 1
    fi
}

# Global array for macOS bash compatibility
declare -a CURL_ARGS

prepare_curl_arguments() {
    CURL_ARGS=(
        "--silent"
        "--show-error"
        "--header" "Authorization: Bearer ${WNC_ACCESS_TOKEN}"
        "--header" "Accept: application/yang-data+json"
    )

    if is_insecure_enabled; then
        CURL_ARGS+=("--insecure")
    fi
}

execute_yang_request() {
    local -r url="$1"
    local -r output_file="$2"

    prepare_curl_arguments
    curl "${CURL_ARGS[@]}" "${url}" > "${output_file}"
}
```

---

## 📣 Testing and Quality

### Testing Requirements
- All scripts must pass `shellcheck` validation
- Integration tests must work with real WNC controller access
- Support dry-run modes for destructive operations
- Test with various input scenarios
- Environment variable validation for integration tests

### Quality Standards
- No shellcheck warnings
- Consistent exit codes (0 for success, non-zero for errors)
- Proper cleanup of temporary files using `./tmp` directory
- Resource cleanup on script termination
- File permissions set to 755 for executable scripts

### Testing Environment Setup
```bash
# Required for integration tests
export WNC_CONTROLLER=wnc1.example.internal
export WNC_ACCESS_TOKEN=YWRtaW46Y3l0WU43WVh4M2swc3piUnVhb1V1ZUx6

# Run integration tests
./scripts/run_integration_tests.sh --insecure --no-color
```

---

## 📣 Documentation

### Inline Documentation
- Include argc meta information for help generation
- Document complex logic with inline comments
- Use meaningful variable names that reduce need for comments
- Include examples in help content

### Help Message Patterns
Using argc's automatic help generation:
```bash
# @meta description List available YANG models from the wireless controller
# @flag --verbose         Enable verbose output for debugging
# @flag --no-color        Disable colored output
# @flag --insecure        Skip SSL certificate verification
# @example                ./scripts/list_yang_models.sh --verbose --insecure
```

---

## 📣 Error Handling

### Standard Error Patterns
```bash
# Environment validation
validate_environment() {
    if [[ -z "${WNC_CONTROLLER:-}" ]]; then
        echo_error "WNC_CONTROLLER environment variable is required"
        show_environment_help
        exit 1
    fi
}

# File validation
validate_file() {
    local -r file_path="$1"
    if [[ ! -f "$file_path" ]]; then
        echo_error "File '$file_path' not found"
        exit 1
    fi
}

# Command validation
validate_dependencies() {
    local -a missing_commands=()

    for cmd in curl jq; do
        if ! command -v "$cmd" >/dev/null 2>&1; then
            missing_commands+=("$cmd")
        fi
    done

    if [[ ${#missing_commands[@]} -gt 0 ]]; then
        echo_error "Missing required commands: ${missing_commands[*]}"
        exit 1
    fi
}
```

### Cleanup Patterns
```bash
# Temporary file cleanup using ./tmp directory
cleanup() {
    if [[ -n "${temp_file:-}" ]] && [[ -f "$temp_file" ]]; then
        rm -f "$temp_file"
    fi
}
trap cleanup EXIT

# Create temporary files in ./tmp
temp_file="./tmp/script_temp_$(date +%s).json"
mkdir -p "./tmp"
```

---

## 📣 Output and Formatting

### Consistent Output Functions
```bash
echo_info() { [[ "${GITHUB_ACTIONS:-false}" != "true" ]] && echo -e "\033[36mℹ Info:\033[0m $*" || echo "Info: $*"; }
echo_success() { [[ "${GITHUB_ACTIONS:-false}" != "true" ]] && echo -e "\033[32m✓ Success:\033[0m $*" || echo "Success: $*"; }
echo_warning() { [[ "${GITHUB_ACTIONS:-false}" != "true" ]] && echo -e "\033[33m⚠ Warning:\033[0m $*" || echo "Warning: $*"; }
echo_error() { [[ "${GITHUB_ACTIONS:-false}" != "true" ]] && echo -e "\033[31m✗ Error:\033[0m $*" >&2 || echo "Error: $*" >&2; }
```

### Color Support
- Respect `--no-color` flag across all scripts
- Support GitHub Actions environment detection
- Use consistent color scheme (info=cyan, success=green, warning=yellow, error=red)

---

## 📣 Performance Considerations

- Use built-in commands where possible instead of external tools
- Minimize subprocess creation in loops
- Use appropriate data structures (arrays vs strings)
- Consider memory usage for large file processing
- Store temporary files in `./tmp` directory for easy cleanup

---

## 📣 Security Guidelines

- Never hardcode credentials or sensitive data
- Validate all inputs to prevent injection attacks
- Use appropriate file permissions (755 for executables)
- Sanitize environment variables before use
- Use secure temporary file creation patterns in `./tmp` directory
- Support SSL certificate verification by default, allow bypass with `--insecure` flag

---

## 📣 macOS Compatibility Notes

### bash Version Considerations
macOS ships with bash 3.x which has limitations:
- No support for `local -n` (nameref variables)
- Limited associative array support
- Different behavior for some built-in commands

### Workarounds
```bash
# Instead of local -n (not supported in bash 3.x)
declare -a CURL_ARGS  # Use global arrays

# Array passing patterns
prepare_curl_arguments() {
    CURL_ARGS=("--silent" "--show-error")
    # Modify global array instead of nameref
}
```

---

## 📣 Integration with Makefile

Scripts should be designed to work both:
1. **Standalone**: Direct execution with argc argument parsing
2. **Via Makefile**: Called from `make` targets with appropriate arguments

Example Makefile integration:
```makefile
test-integration:
	./scripts/run_integration_tests.sh --no-color

lint:
	./scripts/lint_code.sh --no-color

deps:
	./scripts/install_dependencies.sh --verbose
```

---

## 📣 Practical Examples

### YANG Model Operations
```bash
# List all YANG models
./scripts/list_yang_models.sh --verbose --insecure

# Get specific model details
./scripts/get_yang_model_details.sh --insecure "Cisco-IOS-XE-wireless-ap-cfg"

# Fetch model with output
./scripts/fetch_yang_model_details.sh --insecure --output "ap_cfg.json" "Cisco-IOS-XE-wireless-ap-cfg"
```

### Development Workflow
```bash
# Install dependencies
./scripts/install_dependencies.sh --verbose

# Run tests
./scripts/run_unit_tests.sh --no-color
./scripts/run_integration_tests.sh --insecure --no-color

# Generate coverage
./scripts/run_coverage_tests.sh --html --no-color

# Lint code
./scripts/lint_code.sh --no-color

# Clean artifacts
./scripts/clean_artifacts.sh --verbose
```

---

## 📣 Critical Lessons Learned

### argc Predicate Functions
- argc flags generate "1"/"0" strings, not "true"/"false"
- Always use predicate functions: `[[ "${argc_verbose:-0}" == "1" ]]`
- Default missing argc variables to "0" for safety

### macOS bash Compatibility
- Avoid `local -n` syntax (not supported in bash 3.x)
- Use global arrays for complex data passing
- Test scripts on macOS environment

### Environment Variable Consistency
- Use `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN` consistently
- Deprecate legacy `WNC_CONTROLLERS` variable
- Validate environment early in script execution

### Modular Architecture Benefits
- Script-specific libraries improve maintainability
- Shared common libraries reduce code duplication
- Clear separation of concerns (help/output/core)
- Easier testing and debugging

### SSL Certificate Handling
- Support `--insecure` flag for development environments
- Default to secure connections in production
- Clear documentation about certificate requirements

---

## 📣 Future Considerations

- Monitor argc project for updates and new features
- Consider migration to newer bash versions as macOS updates
- Evaluate additional YANG operation patterns as use cases evolve
- Maintain backward compatibility while improving functionality

## 🏗️ Architecture Patterns

### Script Organization
- **Entry Points:** Lightweight scripts that load libraries and call main operations
- **Library Structure:** Use `scripts/lib/<module>/` subdirectories with:
  - `help.sh` - Help and usage functions
  - `output.sh` - Output formatting and display
  - `core.sh` - Core business logic
- **Common Libraries:** Shared functionality in `scripts/lib/common/`

### Function Naming Conventions
- **Global Functions:** UPPERCASE (e.g., `SOURCE_WNC_LIBRARIES`, `INIT_SCRIPT_ENVIRONMENT`)
- **Local Functions:** lowercase (e.g., `validate_input`, `format_output`)
- **Predicate Functions:** `is_*` pattern for boolean checks (e.g., `is_verbose_enabled`)

### Variable Scoping
- **Principle:** Function-level scope by default
- **Global Variables:** Only for multi-function usage
- **argc Variables:** Use with proper shellcheck disable comments
- **Local Variables:** Declare within functions with `local`

---

## 🔧 Modern Tooling Integration

### argc Integration
- **Argument Parsing:** Use [sigoden/argc](https://github.com/sigoden/argc) for clean CLI definitions
- **Metadata Format:**
  ```bash
  # @meta version 1.0.0
  # @meta author @umatare5
  # @meta description "Script description"
  ```
- **Variable Handling:** Assign argc variables to local variables with shellcheck disable:
  ```bash
  # shellcheck disable=SC2154  # argc variables are generated by argc
  local model_name="${argc_model_name}"
  ```

### macOS Compatibility
- **Array Handling:** Avoid `local -n` syntax, use global arrays instead
- **Shell Features:** Test compatibility with macOS bash (version 3.x+)
- **File Permissions:** Maintain 755 for all executable scripts

---

## 🎨 Standards & Quality

### Error Handling
- **Unified Messages:** Use `show_error` function for consistent error reporting
- **Early Exit:** Fail fast with clear error descriptions
- **Context:** Provide actionable error messages with suggestions

### Code Quality
- **Shellcheck Compliance:** Zero warnings required
- **Lint Integration:** All scripts must pass `shellcheck` validation
- **Testing:** Verify functionality with unit tests where applicable

### Documentation
- **Inline Help:** argc metadata provides automatic help generation
- **Usage Examples:** Include practical examples in help text
- **Environment Variables:** Document required and optional variables

---

## 🌐 WNC Integration Patterns

### Authentication
- **Environment Variables:** Support `WNC_CONTROLLERS` format: `host:base64token`
- **Token Handling:** Extract and validate authentication tokens securely
- **Error Context:** Provide clear authentication failure messages

### YANG Operations
- **Model Validation:** Validate YANG model names (minimum 3 characters)
- **Controller Connectivity:** Test connection before operations
- **Response Handling:** Parse JSON responses with proper error checking
- **Output Formats:** Support multiple output formats (pretty, raw, file)

### Common Libraries
- **yang_common.sh:** Shared YANG operation functions
- **argc_common.sh:** argc-specific shared functionality
- **output_formatter.sh:** Consistent output formatting

---

## 📋 Development Workflow

### Script Creation Checklist
1. **Start with argc metadata** - Define version, author, description
2. **Add argument definitions** - Use @option, @flag, @arg annotations
3. **Implement library loading** - Source common libraries
4. **Add predicate functions** - For readable conditionals
5. **Write main logic** - Keep functions 10-20 lines
6. **Test thoroughly** - Verify with shellcheck and unit tests

### Common Pitfalls
- **Unused Variables:** Ensure all argc variables are used in main()
- **Array Compatibility:** Use global arrays instead of `local -n`
- **Function Scope:** Follow uppercase/lowercase naming conventions
- **Error Messages:** Use unified error reporting functions

---

## 💡 Practical Examples

### YANG Model Operations
```bash
# List available models
./scripts/list_yang_models.sh --insecure --verbose

# Get specific model details
./scripts/get_yang_model_details.sh \
  --model "Cisco-IOS-XE-wireless-access-point-oper" \
  --insecure

# Fetch statement details
./scripts/get_yang_statement_details.sh \
  --model "Cisco-IOS-XE-wireless-wlan-cfg" \
  --statement "wlan-cfg-data" \
  --insecure
```

### Environment Setup
```bash
# Single controller with token
export WNC_CONTROLLERS=wnc1.example.internal:YWRtaW46Y3l0WU43WVh4M2swc3piUnVhb1V1ZUx6

# Alternative: Individual variables
export WNC_CONTROLLER=wnc1.example.internal
export WNC_ACCESS_TOKEN=YWRtaW46Y3l0WU43WVh4M2swc3piUnVhb1V1ZUx6
```

### Real YANG Models (Examples)
- `Cisco-IOS-XE-wireless-access-point-oper` - AP operational data
- `Cisco-IOS-XE-wireless-wlan-cfg` - WLAN configuration
- `Cisco-IOS-XE-wireless-rrm-oper` - RRM operational data
- `Cisco-IOS-XE-wireless-client-oper` - Client operational data

---

## 🚨 Critical Lessons Learned

### macOS Bash Limitations
- **Avoid `local -n`:** Use global arrays for reference passing
- **Array Declaration:** Always declare global arrays at script/function start
- **Version Compatibility:** Test with bash 3.x+ for macOS compatibility

### Shellcheck Integration
- **SC2034 Warnings:** Ensure all declared variables are used
- **argc Variables:** Use appropriate disable comments for generated variables
- **Exit Code 0:** Achieve zero warnings before deployment

### argc Best Practices
- **Variable Assignment:** Always assign argc variables to local variables
- **Help Generation:** argc metadata provides automatic help text
- **Error Context:** Provide meaningful error messages for missing arguments

### Library Architecture
- **Modular Design:** Separate help, output, and core logic
- **Consistent Naming:** Follow global/local function naming conventions
- **Error Propagation:** Use consistent error handling patterns

```
