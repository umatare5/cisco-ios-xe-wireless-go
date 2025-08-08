# 🧪 Testing

Extended docs: `docs/testing/`.

| Topic | Path |
|-------|------|
| Overview | `docs/testing/README.md` |
| Unit details | `docs/testing/details_unit.md` |
| Integration details | `docs/testing/details_integration.md` |

## 🎯 Gate

Coverage ≥99% (`make test-coverage`).

## ▶️ Commands

```bash
make test-unit
make test-integration  # needs env
```

## 🔐 Env (Integration)

`WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` (base64 `user:pass`). Missing ⇒ fail.

## 🔍 Notes (Collapsed)

<details><summary>Execution behavior</summary>

Lint precedes tests. Integration uses live GET only. Helpers in `internal/tests`.

</details>

## 🔗 Related

`docs/api/` · `docs/security/` · `docs/scripts/`
