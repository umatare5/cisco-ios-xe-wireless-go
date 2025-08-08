# 🤝 Contributing

Concise rules to keep the SDK stable and high quality.

## Scope

Read‑only wireless RESTCONF (IOS‑XE 17.12). Stdlib only.

## Principles

| Aspect | Rule |
|--------|------|
| Public API | Add only with strong justification |
| Deps | Std lib only |
| Coverage | ≥99% total (no regressions) |
| Errors | `fmt.Errorf("context: %w", err)` |
| Panics | None in library code |
| Globals | Avoid mutable state |
| Security | TLS verify ON by default |

## Testing

1. Unit (logic/validation)
2. Integration (live controller)
3. Gate: ≥99% coverage
4. Add tests for any new export
5. Helpers live in `internal/tests`

## Services

Signature: `Func(ctx) (*model.XResponse, error)`; simple GET → `core.Get[T]`.

## Errors

| Case | Pattern |
|------|---------|
| Nil client | `invalid client configuration` |
| JSON decode | `decode <domain> response: %w` |

Never log internally unless via provided logger.

## Docs

Update affected markdown on API/env/test changes (single H1 rule).

## Workflow

```bash
make lint
make test-unit
make test-integration  # with env
make test-coverage
```

## Commits

Conventional Commits (`feat:`, `fix:`, etc.) ≤72 char subject.

## Env Vars

| Var | Purpose |
|-----|---------|
| `WNC_CONTROLLER` | Host/IP |
| `WNC_ACCESS_TOKEN` | Base64 creds |

Missing → fail fast.

## PR Checklist

Build, lint, tests, ≥99% coverage, docs updated, no stray debug.

## Avoid

Third‑party deps, hardcoded creds, panics, untested exports, coverage drops.

## Ideas (Discuss First)

Retry tuning, pagination helpers, streaming endpoints.

## Support

Include repro, Go version, controller version, minimal snippet.

## Thanks

Prefer small focused PRs.
