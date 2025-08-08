# 🔐 Security

Full detail in `docs/security/`.

| Topic | Path |
|-------|------|
| Overview & checklist | `docs/security/README.md` |
| Hardening | `docs/security/hardening.md` |
| Incident response | `docs/security/incident.md` |

## ✅ Defaults

TLS verify ON • Read‑only GET • Use context timeouts.

## 🔑 Auth

```bash
export WNC_CONTROLLER=<host>
export WNC_ACCESS_TOKEN=$(echo -n 'user:pass' | base64)
```

Dev only insecure:

```go
wnc.NewClient(host, token, wnc.WithInsecureSkipVerify(true))
```

## 🚫 Risks (Collapsed)

<details><summary>Avoid these</summary>

Committing tokens, stale creds, disabling TLS verify in prod, sharing tokens across envs.

</details>

## 🔗 Related

`docs/api/` · `docs/testing/` · `docs/scripts/`
