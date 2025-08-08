# 🤝 Contributing

Goal: keep the SDK minimal, deterministic, secure.

## 🎯 Scope

Read‑only wireless RESTCONF (IOS‑XE 17.12). Stdlib only. No write operations.

## 🔑 Core Rules

| Area         | Rule                                |
| ------------ | ----------------------------------- |
| Public API   | Add sparingly; include tests + docs |
| Dependencies | Stdlib only                         |
| Coverage     | ≥99% total (no regressions)         |
| Errors       | Wrap (`fmt.Errorf("ctx: %w", err)`) |
| Panics       | None in library code                |
| Globals      | Avoid mutable state                 |
| TLS          | Verify ON by default                |

## 🧪 Testing

| Kind        | Purpose           | Notes           |
| ----------- | ----------------- | --------------- |
| Unit        | Logic, validation | No env          |
| Integration | Live controller   | Needs env       |
| Coverage    | Gate ≥99%         | Fails CI < gate |

Helpers: `internal/tests`. New export ⇒ new tests.

## 🔁 Workflow

```bash
make lint
make test-unit
make test-integration   # needs WNC_* env
make test-coverage
```

## 🧩 Service Pattern

`Func(ctx) (*model.XResponse, error)`; basic GET uses `core.Get[T]`.

## 🚨 Error Messages

| Case        | Message Sketch                 |
| ----------- | ------------------------------ |
| Nil client  | `invalid client configuration` |
| Decode fail | `decode <domain> response: %w` |

No internal logging unless via user logger.

## 📦 Env Vars

| Var                | Meaning      |
| ------------------ | ------------ |
| `WNC_CONTROLLER`   | Host/IP      |
| `WNC_ACCESS_TOKEN` | Base64 creds |

Missing ⇒ fail fast.

## ✅ PR Checklist

Build passes • Lint clean • Tests green • ≥99% coverage • Docs updated • No debug leftovers.

## ✍️ Commits

Conventional Commits; subject ≤72 chars (imperative).

## ❌ Avoid

Third‑party deps, hardcoded creds, panics, untested exports, coverage drops, silent skips.

## 💡 Discuss First

Pagination helpers, retries, streaming endpoints, new service groups.

## 🔍 Support Requests

Provide Go version, controller version, repro snippet.

## 🔽 Additional (Collapsed)

<details><summary>Extended guidance</summary>

Docs: single H1 per file. Update impacted docs with API or env changes. Prefer focused PRs. Table tests keep naming consistent. Fail early on configuration errors. Use per‑call contexts. Keep examples minimal.

</details>

## 🙏 Thanks

Focused contributions are appreciated.
