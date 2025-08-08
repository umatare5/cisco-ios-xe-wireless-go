# �️ Script Reference (Stub)

Full content moved to modular docs under `docs/scripts/`.

| Topic | File |
|-------|------|
| Overview & flags | `docs/scripts/README.md` |
| Common models | `docs/scripts/models.md` |
| Troubleshooting | `docs/scripts/troubleshooting.md` |

## Quick Examples

```bash
./scripts/list_yang_models.sh -k
./scripts/get_yang_model_details.sh -m Cisco-IOS-XE-wireless-general-oper -k
./scripts/get_yang_statement_details.sh -m Cisco-IOS-XE-wireless-general-oper -i general-oper-data -k
```

## Environment

`WNC_CONTROLLER` and `WNC_ACCESS_TOKEN` must be set (or provided via flags). Missing → exit 1.

## Notes

- Scripts hard‑fail on missing required env vars.
- Standard banners: `scripts/lib/output/banner.sh`.
- Reusable helpers live only in `scripts/lib/`.

## See Also

`docs/api/` · `docs/testing/` · `docs/security/`
