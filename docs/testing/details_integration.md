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

## 🔽 Additional (Collapsed)

<details><summary>Fixtures & practices</summary>

Data files: `<service>/test_data/` (regression, drift detection). Timeouts: short (≤30s); avoid globals. Failures: log context, continue. Selective runs: `go test -run TestGeneralIntegration ./general`. Update fixtures: inspect, strip volatile fields. Drift: unmarshal errors may signal YANG revision changes. Security: never print token.

See `../security/` for auth/TLS guidance.

</details>
