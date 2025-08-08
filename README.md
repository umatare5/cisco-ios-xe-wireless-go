# 📗 Cisco IOS-XE Wireless Go SDK

Strongly typed, dependency‑free Go SDK for Cisco Catalyst 9800 (IOS‑XE 17.12) RESTCONF (read‑only GET).

| Aspect | Value |
|--------|-------|
| Scope | 25+ wireless domains (oper + cfg) |
| Deps | Stdlib only |
| Coverage | ≥99% enforced |
| Security | TLS verify on (opt‑out dev only) |
| Pattern | Unified client → thin services → YANG models |

Badges: version, tests, report card, Go ref, license.

## 🚀 Quick Start

```bash
go get github.com/umatare5/cisco-ios-xe-wireless-go
echo -n "admin:password" | base64   # make token
```

```go
client, _ := wnc.NewClient("192.168.1.10", "<base64>")
ctx := context.Background()
gen, _ := client.General().Oper(ctx)
_ = gen // use data
```

Options: `WithTimeout(d)`, `WithInsecureSkipVerify(true)` (dev), `WithLogger(l)`.

## 📚 Documentation Index

| Topic | Location |
|-------|----------|
| API (client, models, core services) | `docs/api/` |
| Extended services list | `docs/api/services_extended.md` |
| Testing (patterns, coverage) | `docs/testing/` |
| Scripts (YANG utilities) | `docs/scripts/` |
| Security (TLS, creds, network) | `docs/security/` |
| Contributing | `CONTRIBUTING.md` |

## 🛠 Commands

```bash
make lint            # static analysis
make test-unit       # unit (runs lint)
make test-integration# integration (env required)
make test-coverage   # merged + report.out
```

Env (integration): `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` (base64 user:pass).

## 🔧 Services (Sample)

`General()`, `AP()`, `WLAN()`, `Client()`, `RRM()`, `Rogue()`, `AFC()`, more: see extended services doc.

Generic GET helper removes boilerplate (internal `core.Get[T]`).

## ✅ Policies

| Policy | Rule |
|--------|------|
| Coverage | ≥99% total, no drop |
| Errors | Wrap: `fmt.Errorf("context: %w", err)` |
| Logging | Only via optional user logger |
| Panics | None in library code |
| Deps | No third‑party additions |

## 📊 Coverage Artifact

Committed summary: `coverage/report.out` (regen: `make test-coverage`).

## 📄 License

MIT. See `LICENSE`.
