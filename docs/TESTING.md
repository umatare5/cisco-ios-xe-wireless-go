# 🧪 Testing

This guide explains how to run unit and integration tests, manage test data, and generate coverage reports.

> [!NOTE]
> Integration tests require an accessible Cisco C9800 and these variables: `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN`.

## 🎯 Scope & layers

The suite is deterministic and read‑only. It validates serialization, service wiring, and live data shapes.

| Layer        | Purpose                                | Trigger                 |
| ------------ | -------------------------------------- | ----------------------- |
| Unit         | JSON ↔ Go struct fidelity              | `make test-unit`        |
| Table-driven | Multiple scenario assertions           | `make test-unit`        |
| Fail-fast    | Immediate surface of unexpected errors | `make test-unit`        |
| Integration  | Live controller data shape sanity      | `make test-integration` |

## 🧰 Prerequisites

Tools and environment required to run unit and integration tests.

### Unit / Table / Fail-fast

Run locally without a controller to validate structs, scenarios, and early failures.

| Requirement | Version | Notes                    |
| ----------- | ------- | ------------------------ |
| Go          | 1.24+   | Uses only stdlib testing |
| make        | Latest  | Convenience targets      |

### Integration

Requires a reachable controller and credentials:

| Variable           | Description        | Example                 |
| ------------------ | ------------------ | ----------------------- |
| `WNC_CONTROLLER`   | Controller host/IP | `wnc1.example.internal` |
| `WNC_ACCESS_TOKEN` | Base64 `user:pass` | `YWRtaW46cGFzc3dvcmQ=`  |

<details><summary>Environment setup</summary>

```bash
export WNC_CONTROLLER="wnc1.example.internal"
export WNC_ACCESS_TOKEN="<base64 user:pass>"
```

</details>

> [!CAUTION]
> Never commit real tokens or `.env` files; see Security → Token Handling.

## 🚀 Running tests

Run unit and integration suites via Make targets for consistent, reproducible results.

Primary Make targets:

| Command                     | Description                          |
| --------------------------- | ------------------------------------ |
| `make test-unit`            | Run unit + table + fail-fast suites  |
| `make test-integration`     | Run integration (skips if env unset) |
| `make test-coverage`        | Run all tests and write coverage     |
| `make test-coverage-report` | Generate HTML coverage report        |

> [!NOTE]
> Lint runs automatically where configured; see Make and Scripts references.

## �️ Test data artifacts

Integration tests persist controller responses to support regression and offline inspection.

| Aspect   | Detail                                 |
| -------- | -------------------------------------- |
| Location | `*/test_data/*.json`                   |
| Format   | Raw controller JSON (pretty if stable) |
| Use      | Schema drift and offline analysis      |

<details><summary>Example layout</summary>

```text
test_data/
├── ap_oper_response.json
├── client_oper_response.json
├── general_cfg_response.json
└── rrm_global_oper_response.json
```

</details>

## 📈 Coverage

Generate coverage summaries and an HTML report to assess tested code paths.

### Reports

| Target  | Command                     | Notes                           |
| ------- | --------------------------- | ------------------------------- |
| Summary | `make test-coverage`        | Writes `./tmp/coverage.out`     |
| HTML    | `make test-coverage-report` | Generates `./tmp/coverage.html` |

> [!NOTE]
> CI may publish a coverage badge from `coverage/report.out` when present.

## 💡 Tips

Practical guidance to keep test runs fast, reliable, and easy to debug.

1. Run unit first (`make test-unit`).
2. Export env vars only when needed for integration.
3. Use `grep` over `test_data/` to spot schema drift.
4. Keep JSON fixtures minimal to aid diffs.
5. Fail fast: add explicit error checks in new tests.

> [!TIP]
> Unit tests can run on CI without a controller; integration can be opt‑in.

## 🧩 Troubleshooting

Common issues and concise fixes for failed or flaky test runs.

- Missing env vars: set `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN` for integration.
- Unreachable controller: verify DNS or use `wnc1.example.internal` or an IP.
- TLS errors: see Security → TLS Verification and avoid disabling checks in prod.
- Auth failures: ensure the token is Base64 of `user:pass` and not expired.
- Flaky tests: re-run with verbose logs; isolate by package using `go test ./pkg`.
