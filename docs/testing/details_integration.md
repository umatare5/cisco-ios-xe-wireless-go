# 🌐 Integration Test Details

## Purpose

Validate live controller behavior & schema stability.

## Preconditions

`WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` exported; token is base64 of `user:pass`.

## Pattern

```go
c := tests.TestClient(t)
ctx := tests.TestContext(t)
resp, err := c.General().Oper(ctx)
if err != nil { t.Fatalf("general oper: %v", err) }
```

## Data Files

Saved under `<service>/test_data/` for: regression, offline dev, drift detection.

## Timeouts

Short deadlines (≤30s). Avoid global cancellations.

## Failure Handling

Log context, do not panic. Continue other services.

## Selective Runs

```bash
go test -run TestGeneralIntegration ./general
```

## Updating Fixtures

Manually inspect before commit; remove volatile fields if needed.

## Drift Detection

Schema change causes unmarshal errors → investigate YANG diff.

## Security

Never print token; redact on failure logs.

## Next

See `../security/` for auth/TLS guidance.
