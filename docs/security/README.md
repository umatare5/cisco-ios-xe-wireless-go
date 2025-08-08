# 🔐 Security Overview

Defaults: TLS verify ON, read‑only GET operations, explicit context timeouts.

## Auth

Base64 of `user:pass` provided as token.

```bash
export WNC_CONTROLLER=<host>
export WNC_ACCESS_TOKEN=$(echo -n 'user:pass' | base64)
```

## TLS

Disable verify only in dev:

```go
c, _ := wnc.NewClient(host, token, wnc.WithInsecureSkipVerify(true))
```

## Secrets

| Practice | Rule |
|----------|------|
| Storage | Env / secret store only |
| Logging | Never log token |
| Rotation | Quarterly (min) |
| Separation | Distinct per env |

## Timeouts

Always set client or per‑call context deadlines.

## Anti‑Patterns

- Hardcoded tokens
- Reused prod creds in dev
- Global skip TLS

## 🔽 Additional (Collapsed)

<details><summary>Incident & checklist</summary>

Auth incident: revoke token → issue new → rotate env → audit logs. Network: restrict egress to controller:443. Checklist: TLS verify ON, token not in VCS, context deadlines used, minimum privileges.

See also: `hardening.md`, `incident.md` for depth.

</details>
