---
description: Bash Shell Scripting Instructions
applyTo: "**/*.sh"
---

# GitHub Copilot Agent Mode – Bash Shell Scripting Instructions

This repository's shell scripts are for collecting and analyzing YANG models from Cisco Wireless Network Controllers (WNC).

These scripts help users discover available YANG models, retrieve model details, and support development, testing, validation, and automation tasks for Cisco Catalyst 9800 controllers.

These tools are intended both for direct use and as companions to the Go-based RESTCONF API client library implemented in this repository.

All scripts are designed to be **clear, robust, easy to use, highly readable, and maintainable**, and to follow established best practices for shell scripting and documentation, with a strong emphasis on overall script quality, readability, and maintainability.

Copilot **MUST** comply with all instructions described in this document when editing or creating scripts within the `scripts` directory.

---

## 🎯 Primary Goal

Develop high-quality, maintainable, and user-friendly Bash scripts that adhere to all specified best practices.

---

## 🛠️ Scripting Practices & Style

Apply the following style and best practices to all shell scripts:

- **Follow Standard Shell Style:**
  Use the [Google Shell Style Guide](https://google.github.io/styleguide/shellguide.html).

- **Style & Linting:**
  Ensure all scripts pass `shellcheck`.

- **Functions:**
  Keep functions **ideally between 10 to 20 lines**.

- **Line Length:**
  Keep line length **up to 120 characters**.

- **KISS Principle:**
  Keep implementations simple and avoid unnecessary complexity.

- **DRY Principle:**
  Factor out reusable, unexported helper functions.

- **Robust Script Settings:**
  Always start scripts with `set -euo pipefail`.

- **Modularization:**
  Source and reuse logic from helper scripts (e.g., `validation.sh`, `dependencies.sh`).

- **Defaults & Overrides:**
  Define default settings as `readonly` variables and allow overrides via CLI arguments or environment variables.

- **Comprehensive Argument Parsing:**
  Support both short/long option flags and ensure user-friendly `--help` and usage output.

- **Environment Variable Support:**
  Allow authentication and configuration via environment variables.

- **Dependency Checks:**
  Check that all required dependencies are installed before script execution.

- **Early Validation:**
  Validate user inputs early at the start of the script.

- **Error Handling & Exit:**
  On error, output a clear message and exit immediately.

- **Separation of Concerns:**
  Organize code into dedicated functions for argument parsing, environment setup, main logic, and output formatting.

- **Usage Documentation:**
  Provide clear usage/help text, including examples and variable descriptions.

- **POSIX & Portability:**
  Use POSIX-compliant syntax and avoid OS/locale-specific behavior.

- **User Feedback:**
  Give progress and completion feedback for a better user experience.

- **Additional Practices:**

  - Use clear, explicit, and consistent names.
  - Prefer constants over hardcoded values.
  - Prefer early returns, minimize deep nesting and loops.
  - Only write necessary, non-redundant comments.
