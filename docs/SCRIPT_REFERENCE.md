# 🧾 Script Reference

See `docs/scripts/` for full detail.

| Topic | Path |
|-------|------|
| Overview & flags | `docs/scripts/README.md` |
| Common models | `docs/scripts/models.md` |
| Troubleshooting | `docs/scripts/troubleshooting.md` |

## ▶️ Examples

```bash
./scripts/list_yang_models.sh -k
./scripts/get_yang_model_details.sh -m Cisco-IOS-XE-wireless-general-oper -k
./scripts/get_yang_statement_details.sh -m Cisco-IOS-XE-wireless-general-oper -i general-oper-data -k
```

## ⚙️ Env

`WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` required (or flags). Missing ⇒ exit 1.

## ℹ️ Notes (Collapsed)

<details><summary>Runtime behavior</summary>

Hard fail on missing env. Standard banners via library. Shared helpers isolated under `scripts/lib` only.

</details>

## 🔗 Related

`docs/api/` · `docs/testing/` · `docs/security/`
