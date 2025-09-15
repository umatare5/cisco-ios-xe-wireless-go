---
description: "Bash Shell Script & Bootstrap System Instructions"
applyTo: "**/*.sh"
---

# Bash Shell **Scripts** & Bootstrap System Instructions

GitHub Copilot **MUST** follow these instructions when generating or modifying Bash scripts in this repository.

## Scope & Metadata

- **Last Updated**: 2025-08-11
- **Precedence**: 1. `copilot-instructions.md` (Global) → 2. `bash-umatare5.instructions.md` (This)
- **Compatibility**: POSIX-leaning / macOS bash 3.x compatible (`local -n` not allowed)
- **Style Base**: [Google Shell Style Guide](https://google.github.io/styleguide/shellguide.html)
- **Goal**: High-quality, maintainable, user-friendly scripts with a robust bootstrap system and strict dependency management.

## 1. Architecture

- **RS-001 (MUST)** Create entry points directly under `scripts/`.
- **RS-002 (MUST)** Under `lib`, create a directory per command/module; place the module entry in `core.sh` (formerly `run.sh`).
- **RS-003 (MUST)** Keep entry scripts minimal; implementation resides in module libraries.
- **RS-004 (SHOULD)** `core.sh` focuses on `run_*`/`execute_*` operations; split other responsibilities into dedicated files.
- **RS-005 (MUST)** Store cross-module shared functions under `lib/share`.
- **RS-006 (MUST)** Organize `lib/share` by responsibility (e.g., `argc/`, `utils/`, `http/`, `output/`).
- **RS-007 (MUST)** Entry scripts must source `lib/bootstrap.sh` **first**.
- **RS-008 (MUST)** Call chain flows **top-down** only: `scripts/<cmd>.sh → lib/<module>/core.sh → lib/share/...`.
- **RS-009 (MUST NOT)** Do not cross-reference siblings (`lib/<moduleA>/*` → `lib/<moduleB>/*`). Use `lib/share/*` instead.
- **RS-010 (MUST NOT)** Do not reference **upward** once moved down (e.g., `lib/<module>/core.sh → lib/bootstrap.sh`).

## 1a. Bootstrap System

- **BS-001 (MUST)** Provide a single bootstrap loader: `lib/bootstrap.sh`.
- **BS-002 (MUST)** Bootstrap is sourced **only** by entry scripts in `scripts/`.
- **BS-003 (MUST NOT)** Libraries under `lib/` or `lib/share/` must **never** source `bootstrap.sh`.
- **BS-004 (MUST)** Use `init_wnc_libraries "<module_dir>"` (or repo-specific `init_*_libraries`) for module loading.
- **BS-005 (MUST)** Define a stable `LIB_ROOT` (e.g., `WNC_LIB_ROOT`) used as the **anchor path** for dynamic discovery.
- **BS-006 (SHOULD)** Loading order guideline: **core → utils → http → testing → output → module-specific**.
- **BS-007 (MUST)** Implement **dynamic library loading** in bootstrap via a private loader (e.g., `_load_share_libraries()`).
- **BS-008 (MUST)** Use **intelligent module detection**: iterate directories and source all `.sh` files guarded by include-guards.

**Example (bootstrap snippet)**

```bash
# In lib/bootstrap.sh
_load_share_libraries() {
  local share_dir="${WNC_LIB_ROOT}/share"
  local lib_dir lib_file
  for lib_dir in "${share_dir}"/*; do
    [[ -d "${lib_dir}" ]] || continue
    for lib_file in "${lib_dir}"/*.sh; do
      [[ -f "${lib_file}" ]] || continue
      # Each file must have its own include guard (see Appendix B)
      # shellcheck source=/dev/null
      source "${lib_file}"
    done
  done
}
```

## 2. Coding Rules

- **SH-001 (MUST)** Prioritize human maintainability and readability.
- **SH-002 (MUST)** First line: `#!/usr/bin/env bash`.
- **SH-003 (MUST)** Strict mode (entry only): `set -euo pipefail` (add custom `IFS` only when word-splitting is required).
- **SH-004 (MUST)** Indentation: **4 spaces**.
- **SH-005 (MUST)** Max **120 chars** per logical line.
- **SH-006 (MUST)** Functions target **10–20 lines**.
- **SH-007 (MUST)** Hard max **25 lines** (structured `case` may exceed).
- **SH-008 (MUST)** One-line, clear purpose comment above each function.
- **SH-009 (MUST)** Naming: functions `lowercase_snake_case`; globals `readonly UPPERCASE_SNAKE_CASE`; locals `lowercase_snake_case`.
- **SH-010 (MUST)** Minimize globals; prefer `local`.
- **SH-011 (MUST)** Extract shared logic into helpers and `source` them from `lib/share`.
- **SH-012 (MUST)** Favor portability; avoid OS/locale-specific behavior; do not use `local -n` (bash 3.x).
- **SH-013 (MUST)** Define defaults as `readonly` constants; allow overrides via **CLI flags** or **ENV**.
- **SH-014 (MUST)** Diagnostics follow editor settings (e.g., ShellCheck). Do **not** introduce new flags; prefer fixes over inline disables.
- **SH-015 (MUST)** Libraries must **not** alter `set -euo pipefail` or `IFS`.

## 2a. Library Behavior

- **LB-001 (MUST)** Per SH-015, libraries must not change strict mode/IFS.
- **LB-002 (MUST)** Do not install global `trap`s; expose trap handlers as functions.
- **LB-003 (SHOULD)** Route output via logging utilities; avoid raw `echo`.
- **LB-004 (MUST)** Do not perform final argument parsing in libraries; provide validators/handlers only.
- **LB-005 (SHOULD)** Mask BSD/GNU differences in shared utils (e.g., `stat`, `mktemp`).

## 3. Arguments & Help (argc)

- **AR-001 (MUST)** Use argc metadata (`@flag`, `@arg`, `@option`).
- **AR-002 (MUST)** Provide `@meta description`; rely on autogenerated help.
- **AR-003 (MUST)** Support short/long flags and ENV overrides.
- **AR-004 (SHOULD)** Declare defaults via argc `[default: …]`; mirror to `readonly` when needed.
- **AR-005 (SHOULD)** If a boolean defaults to `true`, add a negative flag (e.g., `--no-race`) or flip default to `false`.
- **AR-006 (SHOULD)** Support `--` passthrough to underlying tools.
- **AR-007 (MUST)** Resolution order: **argc → ENV (UPPERCASE) → defaults**.
- **AR-008 (SHOULD)** Library `handle_*`/`validate_*` return values via stdout; log via `log_*`.
- **AR-009 (MUST)** Use exit code **2** for argument errors (a helper may be used but is not required).

## 4. Separation of Concerns & Validation

- **SC-001 (MUST)** Separate: parsing → env setup → core logic → output.
- **SC-002 (MUST)** Validate inputs/required ENV early.
- **SC-003 (MUST)** Verify dependencies before execution; fail fast when missing.

## 4a. Required CLI Tools (Profiles)

- **CLI-001 (MUST)** Entry scripts call `validate_required_cli_tools "<profile>"` at startup.
- **CLI-002 (MUST)** Profiles are defined in the repo (e.g., `standard`, `testing`, `coverage`).
- **CLI-003 (SHOULD)** On failure, list missing tools and hint at installation.

## 4b. Flag Resolution (Sample)

```bash
# Precedence: argc > ENV > defaults
is_verbose_enabled() {
  if [[ "${argc_verbose:-}" == 1 ]]; then return 0; fi
  [[ "${VERBOSE:-false}" == "true" ]]
}

is_race_enabled() {
  if [[ "${argc_race:-}" == 1 ]]; then return 0; fi
  [[ "${RACE_FLAG:-true}" == "true" ]]
}

is_html_enabled() {
  if [[ "${argc_html:-}" == 1 ]]; then return 0; fi
  [[ "${HTML_COVERAGE:-false}" == "true" ]]
}
```

## 5. Module Dependency Management

- **MD-001 (MUST)** Dependency hierarchy: **core → utils → domain-specific → module**.
- **MD-002 (MUST)** `lib/share/` is dependency-agnostic.
- **MD-003 (MUST)** `lib/<module>/` may depend on `lib/share/`.
- **MD-004 (MUST NOT)** No cross-module deps between `lib/<module1>/` and `lib/<module2>/`.
- **MD-005 (SHOULD)** Use ENV for inter-module communication when needed.
- **MD-006 (MUST)** Document module dependencies in file headers.

## 5a. Circular Reference Prevention

- **CR-001 (MUST)** Libraries must not source the bootstrap loader.

- **CR-002 (MUST)** Implement include guards in **all** libraries:

  ```bash
  if [[ -n "${MODULE_NAME_LOADED:-}" ]]; then
    return 0
  fi
  readonly MODULE_NAME_LOADED=1
  ```

- **CR-003 (MUST)** Set the `*_LOADED` guard (`readonly`) after the check.

- **CR-004 (MUST NOT)** Do not create cycles where A sources B and B sources A.

- **CR-005 (SHOULD)** Prefer dependency injection (pass function refs/params) instead of direct sourcing.

## 6. Library Loading Optimization

- **LO-001 (MUST)** Implement intelligent loading order in bootstrap: **critical → optional → module**.
- **LO-002 (MUST)** Use include guards to prevent duplicate loading and performance issues.
- **LO-003 (SHOULD)** Group loaders: `source_core_libraries`, `source_utils_libraries`, etc.
- **LO-004 (MUST)** Validate file existence before sourcing: `[[ -f "$lib_file" ]]`.
- **LO-005 (SHOULD)** Lazy-load expensive operations when possible.

## 7. Output / UX (Logging)

- **UX-001 (MUST)** Use `log_info` / `log_success` / `log_warn` / `log_error` (no ad-hoc color prints).
- **UX-002 (MUST)** Color by default; fallback to no-color when `is_no_color_enabled` or `NO_COLOR`/`CI` is set.
- **UX-003 (MUST)** Streams: `log_warn`/`log_error` → **stderr**; `log_info`/`log_success` → **stdout**.
- **UX-004 (MUST)** Use `printf` (do not use `echo -e`).
- **UX-005 (SHOULD)** Symbols assume UTF-8; provide ASCII fallbacks when colorless.
- **UX-006 (SHOULD)** Provide `show_*_banner` and `show_*_help`; do not fail if helpers are absent.
- **UX-007 (SHOULD)** Provide `show_status <context>` and gate details with `is_verbose_enabled`.
- **UX-008 (MUST NOT)** Do not hardcode ANSI codes outside helper.
- **UX-009 (SHOULD)** For generated artifacts, include OS-specific open tips (macOS: `open`; Linux: `xdg-open`).

**Sample**

```bash
log_info() {
  if is_no_color_enabled; then
    printf "Info: %s\n" "$*"
    return 0
  fi
  printf "\033[36mℹ Info:\033[0m %s\n" "$*"
}
```

## 8. Conditionals

- **PR-001 (MUST)** Do not reference `argc_*` directly in conditionals; wrap them with predicate helper.
- **PR-002 (MUST)** Prefer early returns.

## 8a. Conditional Statement Style

- **CS-001 (MUST)** Do **not** use short-circuit forms like `[[ cond ]] && cmd` or `[[ cond ]] || cmd`.
- **CS-002 (MUST)** Use explicit `if ... then ... fi`, including early returns.

**Good**

```bash
if [[ ! -d "$module_dir" ]]; then
  return 0
fi
```

**Bad**

```bash
[[ -d "$module_dir" ]] || return 0
```

## 9. Networking (API calls)

- **NET-001 (MUST)** Execute `curl` via shared HTTP helpers with `--fail` and `--show-error`.
- **NET-002 (MUST)** Example headers: `Authorization: Bearer ${API_TOKEN}`, `Accept: ${API_ACCEPT:-application/json}`.
- **NET-003 (MUST)** Validate required ENV early.
- **NET-004 (SHOULD)** Reuse a common CURL args builder to avoid duplication.

**Validation**

```bash
validate_env() {
  if [[ -z "${API_BASE_URL:-}" ]]; then log_error "API_BASE_URL required"; exit 1; fi
  if [[ -z "${API_TOKEN:-}"    ]]; then log_error "API_TOKEN required"; exit 1; fi
}
```

## 10. Temporary Files / Cleanup

- **TMP-001 (MUST)** Create temporary files under `./tmp`.
- **TMP-003 (SHOULD)** Use portable `mktemp` templates (e.g., `mktemp "${TMPDIR:-/tmp}/${PROJECT_SLUG:-proj}.XXXXXX"`).

## 11. Error Handling & Exit Codes

- **EX-001 (MUST)** `0` = success.
- **EX-002 (MUST)** On error, write to **stderr** and exit non-zero.
- **EX-003 (MUST)** Exit codes: `1` = runtime/env error, `2` = argument/usage error.
- **EX-004 (MUST)** Validator functions may exit immediately with appropriate code on failure.
- **EX-005 (SHOULD)** Using numeric codes directly is fine (`exit 2`); a helper may be used.
- **EX-006 (SHOULD)** For unknown options, print a usage hint to **stderr**.

**Samples**

```bash
requires_argument() {
  if [[ -n "${2:-}" ]]; then return 0; fi
  log_error "$1 requires an argument"
  exit 2
}

unknown_option() {
  log_error "Unknown option: $1"
  printf "Use --help for usage\n" >&2
  exit 2
}

validate_project_directory() {
  local project_root="$1"
  if is_directory_valid "$project_root"; then return 0; fi
  log_error "Directory not found: $project_root"
  exit 1
}
```

## 12. Testing & Diagnostics

- **TV-001 (MUST)** Validate bootstrap integrity before main operations.
- **TV-002 (MUST)** Test circular-reference prevention (e.g., with `bash -x`).
- **TV-003 (SHOULD)** Implement module isolation tests.
- **TV-004 (MUST)** Validate required functions are available after bootstrap.
- **TV-005 (SHOULD)** Provide diagnostic commands for troubleshooting library loading.

## 13. Function Design Patterns

- **FD-001 (MUST)** Prefix private functions with `_`.
- **FD-002 (MUST)** Use arrays and loops to remove duplicated `source`/checks.
- **FD-003 (SHOULD)** Manage conditional loads with `required_scripts[]` and `optional_scripts[]`.
- **FD-004 (MUST NOT)** Avoid short-circuit forms; use explicit `if`.

## 14. Function Definition Order

- **FO-001 (MUST)** Order: (1) globals/init → (2) private (`_`) → (3) public (by feature) → (4) main → (5) exports.
- **FO-002 (SHOULD)** In private funcs: **predicates → workers/operations**.

## 15. Array Usage Patterns

- **AU-001 (MUST)** Define filenames in arrays and loop over them for loads.
- **AU-002 (SHOULD)** Use `priority_scripts[]` for ordered execution.
- **AU-003 (SHOULD)** Track `skip_scripts[]` and decide with predicates.

**Standard pattern**

```bash
local required_scripts=("script1.sh" "script2.sh")
local optional_scripts=("optional1.sh" "optional2.sh")

for script_file in "${required_scripts[@]}"; do
  if [[ -f "$script_file" ]]; then
    source "$script_file"
  else
    log_error "Missing required script: $script_file"; exit 1
  fi
done

for script_file in "${optional_scripts[@]}"; do
  if [[ -f "$script_file" ]]; then
    source "$script_file"
  fi
done
```

## 16. Module Loading Patterns

- **ML-001 (MUST)** Load in this order within a module: `predicate.sh` → `banner.sh` → `help.sh` → other features → `core.sh`.
- **ML-002 (SHOULD)** Implement stages as private loaders: `_load_priority_scripts`, `_load_other_scripts`, `_load_core_script`.

**Example**

```bash
_load_priority_scripts() {
  local module_dir="$1"
  local priority_scripts=("predicate.sh" "banner.sh" "help.sh")
  local script path
  for script in "${priority_scripts[@]}"; do
    path="${module_dir}/${script}"
    if [[ -f "$path" ]]; then
      source "$path"
    else
      log_error "Missing priority script: $path"; exit 1
    fi
  done
}

_load_other_scripts() {
  local module_dir="$1"
  local others=("client.sh" "operations.sh")
  local script path
  for script in "${others[@]}"; do
    path="${module_dir}/${script}"
    if [[ -f "$path" ]]; then
      source "$path"
    fi
  done
}

_load_core_script() {
  local module_dir="$1"
  local path="${module_dir}/core.sh"
  if [[ -f "$path" ]]; then
    source "$path"
  else
    log_error "Missing core script: $path"; exit 1
  fi
}
```

## Appendix A: Recommended Directory Layout

```
scripts/
├── lint.sh
├── fmt.sh
└── lib
    ├── bootstrap.sh
    ├── lint
    │   ├── core.sh          # module entry (formerly run.sh)
    │   ├── operations.sh
    │   ├── predicate.sh
    │   ├── banner.sh
    │   └── help.sh
    ├── fmt
    │   ├── core.sh
    │   └── predicate.sh
    └── share
        ├── argc/
        ├── core/            # constants, predicates, argument helpers
        ├── http/
        ├── output/          # log.sh defines log_* family
        ├── utils/
        ├── artifacts/       # artifact & cleanup helpers
        ├── coverage/        # HTML coverage generators
        ├── dependencies/    # CLI/dep validation
        ├── testing/         # test orchestration helpers
        ├── validation/      # repo/policy validation
        └── yang/            # domain-specific (example)
```

## Appendix B: Include-Guard Pattern

```bash
# lib/foo/core.sh
if [[ -n "${MOD_FOO_CORE_LOADED:-}" ]]; then
  return 0
fi
readonly MOD_FOO_CORE_LOADED=1
# module code follows...
```

## Appendix C: Entry Script Skeleton (argc + bootstrap + core.sh)

```bash
#!/usr/bin/env bash
# @meta version 1.0.0
# @meta author "@umatare5"
# @describe Lint Script - Code linting with golangci-lint

# @option -p --project <DIR>           Project root directory [default: .]
# @flag   -v --verbose                 Enable verbose output
# @flag      --fix                     Automatically fix issues where possible
# @option    --config <FILE>           Custom golangci-lint config file path
# @flag      --no-color                Disable colored output

set -euo pipefail
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"; readonly SCRIPT_DIR
source "${SCRIPT_DIR}/lib/bootstrap.sh"

# Initialize module libraries
init_wnc_libraries "${SCRIPT_DIR}/lib/lint"
source "${SCRIPT_DIR}/lib/lint/core.sh"

# Validate required CLI tools
validate_required_cli_tools "standard"

main() {
  local project_root="${argc_project:-.}"
  run_lint_operation "$project_root"
}

# Evaluate argc then run
eval "$(argc --argc-eval "$0" "$@")"
main "$@" || exit $?
```
