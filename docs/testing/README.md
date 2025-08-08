# 🧪 Testing Overview

Deterministic tests enforce ≥99% coverage.

## Types

| Kind | Scope | Live Controller |
|------|-------|-----------------|
| Unit | Pure logic | No |
| Integration | RESTCONF GET | Yes |
| Coverage | Aggregated | Mixed |

## Env

| Var | Meaning |
|-----|---------|
| `WNC_CONTROLLER` | Host/IP |
| `WNC_ACCESS_TOKEN` | Base64 creds |

Missing vars = hard fail.

## Commands

| Make | Purpose |
|------|---------|
| `test-unit` | Unit tests + lint |
| `test-integration` | Live tests (lint) |
| `test-coverage` | Merged coverage -> `coverage/report.out` |
| `test-coverage-html` | HTML view |

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

## Data Capture

Integration saves JSON under each service `test_data/`.

## Coverage

Gate: reject <99% total. Inspect:

```bash
go tool cover -func=./tmp/coverage.out | grep -v 100.0%
```

## Sequence

1. Lint & unit
2. Export env
3. Integration
4. Coverage HTML

## Errors

Always wrap: `fmt.Errorf("context: %w", err)`.

## Context

Use per test deadline contexts; avoid global timeouts.

## Storage

Only commit deterministic fixtures; exclude volatile data.

## Next

See `details_unit.md`, `details_integration.md` for deeper patterns.
