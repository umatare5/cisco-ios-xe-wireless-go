# 🛡 Hardening

## Client Construction

```go
c, _ := wnc.NewClient(host, token, wnc.WithTimeout(30*time.Second))
```

Avoid debug logging in prod. Prefer JSON structured logs.

## TLS

Pin CA where possible; never blanket skip verify in prod.

## Token Handling

Rotate and scope tokens. Keep out of crash dumps.

## Context

Short deadlines prevent resource leaks.

## Logging

Redact sensitive identifiers; no bearer values.

## Dependency Hygiene

Std lib only; review Go patch releases routinely.

## YANG Drift

Unexpected unmarshal errors may indicate model changes → investigate before workaround.

## Validation

Fail fast on missing env; surface config errors early.
