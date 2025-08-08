# 🤝 Contributing

Clear, minimal rules to keep the library stable, high‑quality, and predictable.

## 🔍 Scope

Library only (no CLI, no external deps). Focus: Read‑only RESTCONF access for Catalyst 9800 (IOS‑XE 17.12) wireless domains.

## ✅ Principles

| Aspect | Rule |
|--------|------|
| Public API | Add only when necessary & justified |
| Dependencies | Std lib only |
| Coverage | Maintain ≥99% total (no regressions) |
| Errors | `fmt.Errorf("context: %w", err)`; never lose root cause |
| Panics | None in library code |
| Globals | Avoid mutable global state |
| Determinism | Tests must be stable & repeatable |
| Security | TLS verify ON by default |

## 🧪 Testing

1. Unit tests for logic & validation
2. Integration tests for live controller paths
3. Coverage must remain ≥99% (enforced gate)
4. Add tests with any new exported symbol
5. Store any shared helpers under `internal/tests`

### Helper Patterns

Use existing harness helpers (`tests.TestClient`, `tests.SaveTestDataToFile`). Do not re‑implement ad‑hoc JSON writers.

## 🧱 Service Methods

All service methods follow: `Func(ctx context.Context) (*model.TypeResponse, error)`.

Use internal `core.Get[T]` for simple GET endpoints. Do not wrap additional logic unless required for validation.

## 🧯 Error Handling

| Case | Pattern |
|------|---------|
| Nil client | `return nil, fmt.Errorf("invalid client configuration: %w", ErrInvalidConfiguration)` |
| HTTP failure | Wrap with context path segment |
| JSON decode | `fmt.Errorf("decode <domain> response: %w", err)` |
| Input validation | Early return with clear message |

Never log inside library packages (unless a user‑provided logger is configured). Fail via returned errors only.

## 📝 Documentation

Update relevant markdown when:

* Adding/removing a service method
* Changing environment requirements
* Adjusting testing or coverage behavior

Single H1 per file; follow `.github/instructions/markdown.instructions.md`.

## 🧰 Make Targets

Run (in order) before opening a PR:

```bash
make lint
make test-unit
# (optional) export WNC_CONTROLLER / WNC_ACCESS_TOKEN
make test-integration
make test-coverage
```

`make test-unit` and `make test-integration` already invoke lint.

## 🎯 Commit Messages

Use Conventional Commits:

```text
feat(<scope>): add X
fix(<scope>): correct Y
refactor(<scope>): simplify Z
docs(<scope>): clarify usage
chore(ci): update workflow
```

Keep subject ≤72 chars, present tense imperative.

## 🔒 Environment Variables

| Variable | Purpose | Required for Integration |
|----------|---------|--------------------------|
| `WNC_CONTROLLER` | Controller host/IP | Yes |
| `WNC_ACCESS_TOKEN` | Base64 `user:pass` | Yes |

Missing any required variable causes a hard failure (no silent skip logic).

## 🚦 PR Checklist

* [ ] Code builds (`go build ./...`)
* [ ] Lint passes
* [ ] Tests pass (unit + integration where applicable)
* [ ] Coverage ≥99% total
* [ ] No new untested exports
* [ ] Docs updated
* [ ] No stray debug prints / TODO leftovers

## 🚫 Do Not

| Item | Reason |
|------|--------|
| Add third‑party modules | Maintain minimal attack surface & footprint |
| Embed credentials in tests | Security & reproducibility |
| Introduce panics | Library consumers must handle errors |
| Expand API without docs/tests | Consistency & discoverability |
| Lower coverage | Policy gate |

## 🧩 Future Enhancements (Propose First)

| Idea | Considerations |
|------|----------------|
| Retry policy tuning | Must remain opt‑in, predictable |
| Pagination helpers | Need demonstrable repetition before abstraction |
| Streaming endpoints | Validate RESTCONF support & stability |

Open an issue before large changes.

## 📫 Support

Issues + PRs via GitHub. Provide:

* Repro steps
* Go version
* Controller version
* Minimal code sample

## 🙏 Thanks

Maintainers value precise, incremental improvements—small, focused PRs are easiest to review.
