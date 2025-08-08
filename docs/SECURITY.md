# 🔐 Security (Stub)

Full content moved to modular docs under `docs/security/`.

| Topic | File |
|-------|------|
| Overview & checklist | `docs/security/README.md` |
| Hardening practices | `docs/security/hardening.md` |
| Incident response | `docs/security/incident.md` |

## Defaults

- TLS verification ON (strict)
- Read‑only GET operations
- Context timeouts recommended

## Quick Auth Setup

```bash
export WNC_CONTROLLER=<host>
export WNC_ACCESS_TOKEN=$(echo -n 'user:pass' | base64)
```

## Dev Only

```go
wnc.NewClient(host, token, wnc.WithInsecureSkipVerify(true)) // avoid in prod
```

## Notes

- Never commit tokens
- Rotate credentials regularly
- Disable insecure skip in production

## See Also

`docs/api/` · `docs/testing/` · `docs/scripts/`
