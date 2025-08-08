# 🛠 Troubleshooting Scripts

| Symptom | Cause | Action |
|---------|-------|--------|
| Missing vars | Not exported | Export `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` |
| 401 Unauthorized | Bad token | Recreate base64 `user:pass` |
| TLS failure | Cert chain | Install CA or use `-k` (dev) |
| Empty model list | Non‑wireless version | Verify IOS‑XE 17.12 |
| Slow responses | Network/latency | Retry with `-v` for timing |
| Truncated output | Pipe/pager | Redirect to file or use raw |

## Verbose

Add `-v` to show request targets.

## Exit Codes

0 success, 1 validation/auth, 2 network, 3 parse.

## Redaction

Scripts never echo token; sanitize logs.
