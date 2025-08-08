# 🧪 Testing Overview

Deterministic tests enforce ≥99% coverage.

## Types

| Kind        | Scope        | Live Controller |
| ----------- | ------------ | --------------- |
| Unit        | Pure logic   | No              |
| Integration | RESTCONF GET | Yes             |
| Coverage    | Aggregated   | Mixed           |

## Env

| Var                | Meaning      |
| ------------------ | ------------ |
| `WNC_CONTROLLER`   | Host/IP      |
| `WNC_ACCESS_TOKEN` | Base64 creds |

Missing vars = hard fail.

## Commands

| Make                   | Purpose                                  |
| ---------------------- | ---------------------------------------- |
| `test-unit`            | Unit tests + lint                        |
| `test-integration`     | Live tests (lint)                        |
| `test-coverage`        | Merged coverage -> `coverage/report.out` |
| `test-coverage-report` | HTML view                                |

## Patterns

```go
ctx := tests.TestContext(t)
resp, err := client.AP().Oper(ctx)
```

Nil context test:

```go
var nilCtx context.Context
_, err := client.AP().Oper(nilCtx)
```

## 🔽 Additional (Collapsed)

<details><summary>Data & sequence</summary>

Fixtures: integration saves JSON under each service `test_data/`. Coverage gate: reject <99%. Inspect with `go tool cover -func=./tmp/coverage.out | grep -v 100.0%`. Sequence: lint/unit → export env → integration → coverage HTML. Wrap errors: `fmt.Errorf("context: %w", err)`. Use per test deadline contexts; avoid global timeouts. Only commit deterministic fixtures.

See `details_unit.md`, `details_integration.md` for deeper patterns.

</details>
