# Cisco IOS-XE Wireless Go Scripts

[![Script Status](https://img.shields.io/badge/status-validated-green)](scripts/)
[![Architecture](https://img.shields.io/badge/architecture-modular-blue)](scripts/lib/)
[![Shell Standard](https://img.shields.io/badge/shell-bash%204.0%2B-blue)](https://www.gnu.org/software/bash/)

Comprehensive development and operation scripts for the Cisco IOS-XE Wireless Go SDK. This collection provides standardized tools for building, testing, linting, and interacting with Cisco WNC (Wireless Network Controller) systems.

## 🚀 Quick Start

### Prerequisites

- **Go 1.19+**: For building and testing
- **Bash 4.0+**: For script execution
- **curl**: For RESTCONF API operations
- **jq**: For JSON processing
- **WNC Environment**: Controller access for YANG operations

### Environment Setup

```bash
# Required for YANG operations
export WNC_ACCESS_TOKEN="YWRtaW46Y3l0WU43WVh4M2swc3piUnVhb1V1ZUx6"  # Base64 encoded credentials
export WNC_CONTROLLER="wnc1.example.internal"                          # Controller hostname/IP

# Verify environment
./scripts/show_help.sh
```

### Basic Usage

```bash
# Install dependencies
./scripts/install_dependencies.sh

# Run code quality checks
./scripts/lint_code.sh

# Execute test suites
./scripts/run_unit_tests.sh --short
./scripts/run_coverage_tests.sh --short

# List available YANG models
./scripts/list_yang_models.sh --insecure

# Clean build artifacts
./scripts/clean_artifacts.sh --dry-run
```

## 📋 Script Reference

### Development Scripts

| Script | Purpose | Key Options |
|--------|---------|-------------|
| `install_dependencies.sh` | Manage Go dependencies | `--clean`, `--update`, `--verify` |
| `lint_code.sh` | Code quality analysis | `--fix`, `--config <file>` |
| `clean_artifacts.sh` | Clean build artifacts | `--dry-run`, `--force`, `--all` |

### Testing Scripts

| Script | Purpose | Key Options |
|--------|---------|-------------|
| `run_unit_tests.sh` | Execute unit tests | `--short`, `--coverage`, `--verbose` |
| `run_integration_tests.sh` | Execute integration tests | `--check-env-only`, `--timeout <duration>` |
| `run_coverage_tests.sh` | Generate coverage reports | `--short`, `--html` |
| `generate_coverage_html.sh` | Generate HTML coverage | `--input <file>`, `--output <file>` |

### YANG Operations Scripts

| Script | Purpose | Key Options |
|--------|---------|-------------|
| `list_yang_models.sh` | List available YANG models | `--insecure`, `--format <json\|xml>` |
| `get_yang_model_details.sh` | Get specific model details | `--insecure`, `--verbose` |
| `get_yang_statement_details.sh` | Get statement details | `--insecure`, `--verbose` |
| `fetch_yang_model_details.sh` | Fetch model from controller | `--output <file>`, `--raw` |
| `fetch_yang_statement_details.sh` | Fetch statement from controller | `--statement <name>`, `--output <file>` |

### Utility Scripts

| Script | Purpose | Key Options |
|--------|---------|-------------|
| `show_help.sh` | Display comprehensive help | `--no-color` |

## 🛠 Advanced Usage

### Development Workflow

```bash
# Complete development cycle
./scripts/install_dependencies.sh --clean    # Fresh dependency install
./scripts/lint_code.sh --fix                 # Auto-fix code issues
./scripts/run_unit_tests.sh --coverage       # Test with coverage
./scripts/run_integration_tests.sh           # Full integration tests
./scripts/clean_artifacts.sh                 # Clean up
```

### YANG Model Development

```bash
# Explore available models
./scripts/list_yang_models.sh --insecure --format json | jq '.[]'

# Get detailed model information
./scripts/get_yang_model_details.sh --insecure wireless-access-point

# Fetch model data from controller
./scripts/fetch_yang_model_details.sh --output model.json wireless-client

# Get statement details
./scripts/get_yang_statement_details.sh --insecure ieee802-dot11 module
```

### Configuration & Testing

```bash
# Test environment validation
./scripts/run_integration_tests.sh --check-env-only

# Coverage analysis with HTML output
./scripts/run_coverage_tests.sh --short
./scripts/generate_coverage_html.sh --input ./tmp/coverage.out

# Dependency management
./scripts/install_dependencies.sh --update --verify
```

## 🏗 Architecture

### Modular Design

```text
scripts/
├── *.sh                          # Entry point scripts (lightweight)
├── lib/                          # Shared library modules
│   ├── common/                   # Common utilities & predicates
│   │   ├── common.sh            # Library loader & predicates
│   │   ├── constants.sh         # Global constants
│   │   ├── validation.sh        # Input validation
│   │   └── yang_common.sh       # YANG-specific utilities
│   ├── yang_operations/         # YANG/RESTCONF operations
│   │   ├── core.sh              # Core YANG functionality
│   │   ├── help.sh              # YANG help messages
│   │   └── output.sh            # YANG output formatting
│   ├── testing/                 # Test execution utilities
│   ├── lint_code/               # Code linting functions
│   ├── dependencies/            # Dependency management
│   └── artifacts/               # Cleanup operations
└── README.md                    # This documentation
```

### Key Features

- **argc Integration**: Modern CLI argument parsing with [sigoden/argc](https://github.com/sigoden/argc)
- **Predicate Functions**: Readable conditional logic (`is_verbose_enabled()`, `is_insecure_enabled()`)
- **Error Handling**: Comprehensive error checking with `set -euo pipefail`
- **Consistent Output**: Standardized formatting with color support
- **SSL Flexibility**: Configurable certificate verification for development environments

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `WNC_ACCESS_TOKEN` | Base64 encoded credentials for WNC | Required for YANG ops |
| `WNC_CONTROLLER` | WNC controller hostname/IP | Required for YANG ops |

### Script Options

All scripts support common options:

- `--help`: Display detailed help information
- `--version`: Show script version
- `--verbose`: Enable detailed output
- `--no-color`: Disable colored output

YANG scripts additionally support:

- `--insecure`: Skip TLS certificate verification
- `--controller <host>`: Override default controller
- `--token <token>`: Override default credentials

## 📊 Testing & Coverage

### Expected Test Results

| Test Type | Expected Outcome | Notes |
|-----------|------------------|-------|
| Unit Tests | ✅ All Pass | ~10s execution time |
| Integration Tests | ✅ All Pass | 6GHz tests may skip (expected) |
| Coverage Tests | ✅ 70%+ Coverage | Target: 75%+ overall |

### Known Limitations

- **6GHz Support**: Some controllers may not support 6GHz operations (tests will skip)
- **SSL Certificates**: Development environments may require `--insecure` flag
- **Rate Limiting**: Rapid API calls may encounter temporary throttling

## 🐛 Troubleshooting

### Common Issues

#### Environment Setup

```bash
# Missing environment variables
export WNC_ACCESS_TOKEN="your_base64_token"
export WNC_CONTROLLER="your_controller_hostname"

# Test connectivity
./scripts/list_yang_models.sh --insecure --no-color
```

#### SSL Certificate Issues

```bash
# For development environments
./scripts/list_yang_models.sh --insecure
```

#### Permission Issues

```bash
# Ensure scripts are executable
chmod +x scripts/*.sh
```

#### Dependency Issues

```bash
# Clean and reinstall
./scripts/install_dependencies.sh --clean --update
```

## 📚 Additional Resources

- **API Reference**: See `docs/API_REFERENCE.md`
- **Testing Guide**: See `docs/TESTING.md`
- **Security Notes**: See `docs/SECURITY.md`
- **Development Instructions**: See `.github/instructions/scripts.instructions.md`

## 🏷 Version Information

- **Script Version**: 1.0.0
- **Author**: @umatare5
- **License**: See project LICENSE file
- **Compatibility**: Cisco IOS-XE 17.12+ Wireless Controllers

---

For detailed script-specific options and examples, run any script with `--help`:

```bash
./scripts/<script_name>.sh --help
```

*This project uses a modular script architecture designed for maintainability, consistency, and ease of use across all development operations.*
