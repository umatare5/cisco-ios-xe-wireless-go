# 📋 Scripts Reference

Central reference for all development scripts under `scripts/`. Each entry point focuses on a single concern and delegates shared logic to modular libraries in `scripts/lib/`. All scripts are: idempotent, side‑effect constrained, and exit non‑zero on failure.

| Domain       | Scripts (entry points)                                                                   |
| ------------ | ---------------------------------------------------------------------------------------- |
| YANG / Data  | `get_yang_models.sh`, `get_yang_model_details.sh`, `get_yang_statement_details.sh`       |
| Testing      | `test_unit.sh`, `test_integration.sh`, `test_coverage.sh`, `generate_coverage_report.sh` |
| Quality      | `lint.sh`, `pre_commit_hook.sh`                                                          |
| Dependencies | `install_dependencies.sh`                                                                |
| Artifacts    | `clean_artifacts.sh`                                                                     |
| Meta / Help  | `help.sh`                                                                                |

> [!NOTE]
> Environment variables `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` apply to scripts that contact a controller. CLI flags always override environment defaults.

## 🗂️ Structure

```text
scripts/
├── <command>.sh            # Thin entry point(s)
└── lib/                    # Reusable modules (loaded via bootstrap)
    ├── bootstrap.sh        # Loader + init
    ├── artifacts/          # Cleanup operations
    ├── coverage/           # Coverage + HTML generation
    ├── dependencies/       # Dependency install/update
    ├── lint/               # Lint operations
    ├── output/             # Banners + formatting helpers
    ├── testing/            # go test orchestration
    ├── utils/              # generic predicates (jq detection, etc.)
    ├── validation/         # git / branch protection helpers
    └── yang/               # RESTCONF + YANG data utilities
```

<details>
<summary>Library Initialization (Bootstrap Flow)</summary>

1. Entry script sources `lib/bootstrap.sh`.
2. Calls `init_wnc_libraries <script_dir> <module_dir>` selecting a feature module (e.g. `lib/testing`).
3. Shared predicates / formatting / validation become available in current shell.
4. Script invokes a single exported `run_*_operation` function.

</details>

## 🌐 Common Flags & Conventions

| Flag/Env           | Meaning                                  | Applies To              | Default |
| ------------------ | ---------------------------------------- | ----------------------- | ------- |
| `-c, --controller` | WNC host / IP                            | YANG / integration test | (env)   |
| `-t, --token`      | Base64 `user:pass`                       | YANG / integration test | (env)   |
| `-p, --protocol`   | `http` or `https`                        | YANG scripts            | https   |
| `-k, --insecure`   | Skip TLS verify                          | YANG scripts            | false   |
| `-v, --verbose`    | Verbose diagnostics                      | Most scripts            | false   |
| `--no-color`       | Disable ANSI colors                      | All                     | false   |
| `-h, --help`       | Show script specific help (argc powered) | All                     | -       |

<details>
<summary>Output Helper Functions (Validation / General)</summary>

| Function       | Purpose                      | Color Icon | Notes                             |
| -------------- | ---------------------------- | ---------- | --------------------------------- |
| `show_info`    | Informational message        | Cyan ℹ     | Respects `--no-color`             |
| `show_warning` | Non-fatal warning            | Yellow ⚠   | Stderr                            |
| `show_error`   | Error message                | Red ✗      | Stderr, used before non-zero exit |
| `show_success` | Success / completion notice  | Green ✓    |                                   |
| `wnc_banner_*` | Context banners (if present) | Mixed      | Optional, auto-detected           |

Usage (example):

```bash
show_info "Starting integration tests"
show_warning "Skipping optional slow test"
show_error "Controller unreachable"
show_success "All checks passed"
```

</details>

---

## 🧪 Testing Scripts

| Script                        | Focus                               | Core Operation Function                                   |
| ----------------------------- | ----------------------------------- | --------------------------------------------------------- |
| `test_unit.sh`                | Unit + table + fail-fast            | `run_unit_test_operation` / `run_coverage_test_operation` |
| `test_integration.sh`         | Live controller integration tests   | `run_integration_test_operation`                          |
| `test_coverage.sh`            | Unified coverage (unit+integration) | `run_coverage_test_operation`                             |
| `generate_coverage_report.sh` | HTML coverage export                | `run_coverage_html_operation`                             |

<details>
<summary>`test_unit.sh`</summary>

| Flag             | Description                         | Default |
| ---------------- | ----------------------------------- | ------- |
| `-p, --project`  | Project root                        | `.`     |
| `-s, --short`    | Skip long tests                     | off     |
| `-c, --coverage` | Produce coverage file (`./tmp/...`) | off     |
| `-t, --timeout`  | Test timeout                        | 30s     |
| `-v, --verbose`  | Verbose test output                 | off     |
| `--no-color`     | Disable color                       | off     |

Example:

```bash
./scripts/test_unit.sh -c -v
```

</details>

<details>
<summary>`test_integration.sh`</summary>

| Flag               | Description         | Default |
| ------------------ | ------------------- | ------- |
| `-p, --project`    | Project root        | `.`     |
| `--package`        | Package pattern     | `./...` |
| `--check-env-only` | Validate env & exit | off     |
| `-t, --timeout`    | Timeout             | 10m     |
| `--race`           | Race detector       | on      |
| `-v, --verbose`    | Verbose             | off     |
| `--no-color`       | Disable color       | off     |

Environment required: `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN`.

Example:

```bash
export WNC_CONTROLLER=wnc1.example.internal
export WNC_ACCESS_TOKEN=YWRtaW46...
./scripts/test_integration.sh -v
```

</details>

<details>
<summary>`test_coverage.sh`</summary>

| Flag            | Description        | Default              |
| --------------- | ------------------ | -------------------- |
| `-p, --project` | Project root       | `.`                  |
| `-o, --output`  | Coverage file path | `./tmp/coverage.out` |
| `-s, --short`   | Skip long tests    | off                  |
| `-t, --timeout` | Timeout            | 30s                  |
| `-v, --verbose` | Verbose output     | off                  |
| `--no-color`    | Disable color      | off                  |

Example:

```bash
./scripts/test_coverage.sh -o ./tmp/all.out
```

</details>

<details>
<summary>`generate_coverage_report.sh`</summary>

| Flag            | Description         | Default               |
| --------------- | ------------------- | --------------------- |
| `-p, --project` | Project root        | `.`                   |
| `-i, --input`   | Coverage input file | `./tmp/coverage.out`  |
| `-o, --output`  | HTML output file    | `./tmp/coverage.html` |
| `-v, --verbose` | Verbose output      | off                   |
| `--no-color`    | Disable color       | off                   |

Example:

```bash
./scripts/generate_coverage_report.sh -i ./tmp/coverage.out -o ./tmp/report.html
```

</details>

---

## 📦 Dependency & Artifacts

| Script                    | Purpose                               | Key Ops Function             |
| ------------------------- | ------------------------------------- | ---------------------------- |
| `install_dependencies.sh` | Install / update dev tools            | `run_dependencies_operation` |
| `clean_artifacts.sh`      | Remove caches / temp / coverage files | `run_artifacts_operation`    |

<details>
<summary>`install_dependencies.sh`</summary>

| Flag / Option     | Description                      | Default |
| ----------------- | -------------------------------- | ------- |
| `-p, --project`   | Project root                     | `.`     |
| `--golangci-lint` | golangci-lint version            | latest  |
| `--gotestsum`     | gotestsum version                | latest  |
| `-u, --update`    | Update dependencies (Go modules) | off     |
| `-c, --clean`     | Clean module cache first         | off     |
| `--force`         | Force reinstall                  | off     |
| `--download-only` | Download modules only            | off     |
| `--verify`        | `go mod verify` after install    | off     |
| `-v, --verbose`   | Verbose output                   | off     |
| `--no-color`      | Disable color                    | off     |

Example:

```bash
./scripts/install_dependencies.sh --golangci-lint v1.60.0 --verify -v
```

</details>

<details>
<summary>`clean_artifacts.sh`</summary>

| Flag / Option   | Description                         | Default |
| --------------- | ----------------------------------- | ------- |
| `-p, --project` | Project root                        | `.`     |
| `-f, --force`   | Force deletion (no prompt)          | off     |
| `--go-cache`    | Clean Go build cache                | off     |
| `--go-modules`  | Clean module cache                  | off     |
| `--temp-files`  | Clean `./tmp` directory             | off     |
| `--test-files`  | Remove coverage / test binaries     | off     |
| `--all`         | Clean everything (default behavior) | on      |
| `--dry-run`     | Show actions only                   | off     |
| `-v, --verbose` | Verbose output                      | off     |
| `--no-color`    | Disable color                       | off     |

Example:

```bash
./scripts/clean_artifacts.sh --go-cache --temp-files --test-files -v
```

</details>

---

## 🔍 YANG & Data Retrieval

| Script                          | Purpose                                | Function Pattern                       |
| ------------------------------- | -------------------------------------- | -------------------------------------- |
| `get_yang_models.sh`            | List available YANG modules            | `format_yang_models_pretty`            |
| `get_yang_model_details.sh`     | Fetch single YANG module definition    | `format_yang_model_details_pretty`     |
| `get_yang_statement_details.sh` | Fetch specific data subtree (RESTCONF) | `format_yang_statement_details_pretty` |

<details>
<summary>Shared Behaviors</summary>

| Feature         | Description                                         |
| --------------- | --------------------------------------------------- |
| Protocol select | `--protocol http\|https` (default https)            |
| TLS bypass      | `--insecure` (lab only)                             |
| Verbose         | Prints constructed URL + raw response section       |
| Color toggle    | `--no-color` disables ANSI styling                  |
| Output modes    | JSON (default), XML (`-f xml`), raw passthrough     |
| Filtering       | Model list filters `Cisco-IOS-XE-wireless*` entries |

</details>

<details>
<summary>Example: List Models</summary>

```bash
export WNC_CONTROLLER=wnc1.example.internal
export WNC_ACCESS_TOKEN=YWRtaW46...
./scripts/get_yang_models.sh -k
```

</details>

<details>
<summary>Example: Model Definition (Raw)</summary>

```bash
./scripts/get_yang_model_details.sh -c wnc1.example.internal -r -k > ap-oper.yang
```

</details>

<details>
<summary>Example: Statement Data (Pretty)</summary>

```bash
./scripts/get_yang_statement_details.sh -c wnc1.example.internal \
  Cisco-IOS-XE-wireless-access-point-oper access-point-oper-data -k
```

</details>

---

## 🧹 Quality & Validation

| Script               | Purpose                           | Notes                 |
| -------------------- | --------------------------------- | --------------------- |
| `lint.sh`            | Run `golangci-lint` over codebase | Supports `--fix`      |
| `pre_commit_hook.sh` | Branch / staging guard + guidance | Uses `show_*` helpers |
| `help.sh`            | Aggregate usage summary           | No flags              |

<details>
<summary>`lint.sh` Flags</summary>

| Flag / Option   | Description              | Default |
| --------------- | ------------------------ | ------- |
| `-p, --project` | Project root             | `.`     |
| `--config`      | Custom config file       | (auto)  |
| `--fix`         | Auto-fix eligible issues | off     |
| `-v, --verbose` | Verbose linter output    | off     |
| `--no-color`    | Disable color            | off     |

</details>

<details>
<summary>Pre-Commit Validation (Behavior)</summary>

| Check                    | Action on Fail                       |
| ------------------------ | ------------------------------------ |
| Main branch protection   | Block direct commit, show guidance   |
| Empty staging area       | Warn and exit                        |
| (Future) lint/test gates | Can be integrated into hook pipeline |

</details>

---

## 🛠️ Tips & Best Practices

1. Use Make targets (`make test-unit`) for concise workflows.
2. Prefer explicit flags in CI; rely on env vars locally.
3. Keep coverage artifacts under `./tmp` (gitignored) for cleanliness.
4. Avoid `--insecure` outside lab / PoC scenarios.
5. Use `--check-env-only` in automation to assert readiness fast.

> [!TIP]
> Combine: `./scripts/test_coverage.sh && ./scripts/generate_coverage_report.sh` for fast HTML overview.

---

**Back to:** [API Reference](API_REFERENCE.md) | [Security](SECURITY.md)
