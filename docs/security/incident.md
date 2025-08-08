# 🚨 Incident Response

## Auth Compromise

1. Revoke token
2. Issue new
3. Update env / stores
4. Force restart dependents
5. Audit access logs

## Suspicious Traffic

Isolate host → verify TLS certs → check for replay attempts.

## Data Drift

Schema errors: diff YANG revision; update expectations.

## Logging Review

Search for anomalies:

- Surge in 401/403
- Latency spikes
- Repeated timeouts

## Postmortem

Document root cause, containment, follow‑ups (rotation cadence, monitoring gaps).

## Communication

Limit sensitive detail; share indicators not credentials.

## 🔽 Additional (Collapsed)

<details><summary>Preparation tips</summary>

Pre-create rotation procedure doc. Maintain minimal log retention consistent with policy. Conduct token rotation drills quarterly.

</details>
