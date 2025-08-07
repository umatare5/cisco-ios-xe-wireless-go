# WNC Scripts Library Architecture

This document describes the new modular library architecture for Cisco WNC scripts.

## 🏗️ Directory Structure

```
scripts/lib/
├── bootstrap.sh           # Unified library loader
├── core/                  # Core functionality
│   ├── common.sh         # Basic predicates and validation
│   ├── constants.sh      # Project constants and defaults
│   ├── argc_common.sh    # argc utility functions
│   └── argument_parsing.sh # Command-line argument utilities
├── utils/                # General utilities
│   ├── validation.sh     # Input validation functions
│   ├── file_utils.sh     # File manipulation utilities
│   ├── cli_validation.sh # CLI tool validation
│   ├── build_tools.sh    # Build and compilation tools
│   └── dependencies.sh   # Dependency management
├── network/              # Network and HTTP operations
│   ├── http_client.sh    # HTTP client functionality
│   ├── authentication.sh # Authentication handling
│   └── yang_common.sh    # YANG/RESTCONF operations
├── output/               # Output formatting
│   └── output_formatter.sh # Output formatting utilities
└── modules/              # Feature-specific modules
    ├── testing/          # Test operations
    ├── lint_code/        # Code linting
    ├── dependencies/     # Dependency management
    ├── yang_operations/  # YANG model operations
    └── artifacts/        # Build artifact management
```

## 🚀 Bootstrap System

### Usage Patterns

#### Full Initialization
```bash
# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize with module-specific libraries
init_wnc_libraries "$SCRIPT_DIR" "${SCRIPT_DIR}/lib/testing"
```

#### Lightweight Initialization
```bash
# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Basic initialization (core + utils only)
init_wnc_basic
```

#### Network-Focused Initialization
```bash
# Source bootstrap library
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Network initialization (core + network + output)
init_wnc_network
```

### Initialization Functions

| Function | Libraries Loaded | Use Case |
|----------|------------------|----------|
| `init_wnc_libraries(script_dir, module_dir)` | All + module-specific | Full-featured scripts |
| `init_wnc_basic()` | core + utils | Simple utility scripts |
| `init_wnc_network()` | core + network + output | Network/API scripts |

## 📚 Library Categories

### Core Libraries
- **common.sh**: Basic predicate functions (`is_enabled`, `is_verbose_enabled`, etc.)
- **constants.sh**: Project-wide constants and default values
- **argc_common.sh**: argc-specific utility functions
- **argument_parsing.sh**: Command-line argument processing

### Utility Libraries
- **validation.sh**: Input validation and format checking
- **file_utils.sh**: File system operations and utilities
- **cli_validation.sh**: CLI tool presence and version checking
- **build_tools.sh**: Build system integration
- **dependencies.sh**: Package and dependency management

### Network Libraries
- **http_client.sh**: HTTP/HTTPS request handling
- **authentication.sh**: Authentication token management
- **yang_common.sh**: YANG model and RESTCONF API operations

### Output Libraries
- **output_formatter.sh**: Consistent output formatting across scripts

## 🔧 Migration Guide

### Old Pattern (Deprecated)
```bash
source "${SCRIPT_DIR}/lib/common/common.sh"
init_script_libraries "$SCRIPT_DIR" "$MODULE_DIR"
```

### New Pattern (Recommended)
```bash
source "${SCRIPT_DIR}/lib/bootstrap.sh"
init_wnc_libraries "$SCRIPT_DIR" "$MODULE_DIR"
```

## 💡 Benefits

1. **Clear Separation of Concerns**: Each namespace has a specific purpose
2. **Reduced Coupling**: Libraries only load what they need
3. **Improved Maintainability**: Easier to locate and update functionality
4. **Better Performance**: Selective loading reduces startup time
5. **Enhanced Readability**: Clear dependency relationships

## 🧪 Testing

All scripts using the new bootstrap system are backward compatible and maintain existing functionality while benefiting from the improved architecture.

## 🚨 Deprecated Features

- `init_script_libraries()`: Use `init_wnc_libraries()` instead
- `source_wnc_libraries()`: Use bootstrap functions instead
- Direct library sourcing: Use bootstrap initialization instead

Deprecated functions will show warnings in verbose mode and will be removed in future versions.
