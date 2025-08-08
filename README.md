# 📗 Cisco IOS-XE Wireless Go SDK

Dependency‑free (stdlib), strongly typed Go SDK for Cisco Catalyst 9800 (IOS‑XE 17.12) wireless RESTCONF (GET only).

| Aspect | Value |
|--------|-------|
| Domains | 25+ (oper + cfg) |
| Deps | Stdlib only |
| Coverage | ≥99% gate |
| Security | TLS verify ON (dev opt‑out) |
| Design | Client → Services → Models |

## 🚀 Quick Start (60s)

```bash
go get github.com/umatare5/cisco-ios-xe-wireless-go
export WNC_CONTROLLER=192.168.1.10
export WNC_ACCESS_TOKEN=$(echo -n 'admin:password' | base64)
```

```go
client, err := wnc.NewClient(os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN"), wnc.WithTimeout(30*time.Second))
if err != nil { log.Fatal(err) }
ctx := context.Background()
gen, err := client.General().Oper(ctx)
if err != nil { log.Fatal(err) }
_ = gen
```

Options: `WithTimeout(d)`, `WithInsecureSkipVerify(true)` (dev only), `WithLogger(l)`.

## 📚 Index

| Topic | Path |
|-------|------|
| API overview | `docs/api/README.md` |
| Core services | `docs/api/services_core.md` |
| Extended services | `docs/api/services_extended.md` |
| Testing & coverage | `docs/testing/` |
| Security | `docs/security/` |
| Scripts (YANG tooling) | `docs/scripts/` |
| Contributing | `CONTRIBUTING.md` |

## 🛠 Make Targets

```bash
make lint            # static analysis
make test-unit       # unit (runs lint)
make test-integration # live (needs env)
make test-coverage   # merged + report.out
```

Required env (integration): `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` (base64 `user:pass`).

## 🔧 Service Pattern

All services expose simple `Oper(ctx)` (or domain specific) returning typed YANG‑aligned structs; internal `core.Get[T]` centralizes HTTP + decode.

## ✅ Core Policies

| Area | Rule |
|------|------|
| Errors | Wrap with context (`fmt.Errorf("x: %w", err)`) |
| Panics | None in library code |
| Logging | Only via user‑supplied logger |
| Deps | No third‑party libs |
| Coverage | ≥99% enforced before merge |

## 📊 Coverage

Artifact: `coverage/report.out` (regenerate: `make test-coverage`).

## 🔍 More (Collapsed)

<details><summary>Advanced notes</summary>

YANG models: optional leaves are pointers; always nil‑check. Use per‑call contexts for cancellation. Reuse a single client instance. For deeper service list or testing patterns see referenced docs.

</details>

## 📄 License

MIT (see `LICENSE`).
