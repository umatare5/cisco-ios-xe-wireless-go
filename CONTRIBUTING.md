# 🤝 Contributing

Keep the SDK minimal, deterministic, secure.

## 🎯 Scope

Read‑only wireless RESTCONF (IOS‑XE 17.12). Stdlib only. No write ops.

## 🔑 Core Rules

| Area | Rule |
|------|------|
| Public API | Add sparingly; require tests & docs |
| Dependencies | Stdlib only |
| Coverage | ≥99% total (no regressions) |
| Errors | Wrap: `fmt.Errorf("ctx: %w", err)` |
| Panics | None in library code |
| Globals | Avoid mutable state |
| TLS | Verify ON by default |

## 🧪 Testing

| Kind | Purpose | Notes |
|------|---------|-------|
| Unit | Logic, validation | No env |
| Integration | Live controller | Needs env |
| Coverage | Gate ≥99% | Fails CI < gate |

Helpers: `internal/tests`. New export ⇒ new tests.

## 🔁 Workflow

```bash
make lint
make test-unit
make test-integration   # needs WNC_* env
make test-coverage
```

## 🧩 Services Pattern

`Func(ctx) (*model.XResponse, error)`; basic GET uses `core.Get[T]`.

## 🚨 Errors

| Case | Message Sketch |
|------|----------------|
| Nil client | `invalid client configuration` |
| Decode fail | `decode <domain> response: %w` |

No internal logging unless via user logger.

## 📦 Env Vars

| Var | Meaning |
|-----|---------|
| `WNC_CONTROLLER` | Host/IP |
| `WNC_ACCESS_TOKEN` | Base64 creds |

Missing ⇒ fail fast.

## ✅ PR Checklist

Build passes • Lint clean • Tests (unit+integration) green • ≥99% coverage • Docs updated • No stray debug.

## ✍️ Commits

Conventional Commits; subject ≤72 chars (imperative).

## ❌ Avoid

Third‑party deps, hardcoded creds, panics, untested exports, coverage drops, silent skips.

## 💡 Discuss First

Pagination helpers, retries, streaming endpoints, new service groups.

## 🔍 Support Requests

Provide Go version, controller version, repro snippet.

## 🔽 Extra (Collapsed)

<details><summary>Extended guidance</summary>

Documentation: single H1 per file. Update impacted docs with API or env changes. Prefer focused PRs. Table tests keep naming consistent. Fail early on configuration errors.

</details>

## 🙏 Thanks

Small focused contributions are appreciated.
