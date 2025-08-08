# 🧪 Testing (Stub)

Full guide moved to modular docs under `docs/testing/`.

| Topic | File |
|-------|------|
| Overview | `docs/testing/README.md` |
| Unit details | `docs/testing/details_unit.md` |
| Integration details | `docs/testing/details_integration.md` |

## Coverage Gate

Total ≥99% enforced (`make test-coverage`).

## Quick Run

```bash
make test-unit
make test-integration  # needs env
```

## Env Vars (Integration)

`WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` (base64 `user:pass`). Missing → fail.

## Notes

- Lint runs before test targets.
- Integration: live RESTCONF GET only.
- Shared helpers in `internal/tests`.

## See Also

`docs/api/` · `docs/security/` · `docs/scripts/`
