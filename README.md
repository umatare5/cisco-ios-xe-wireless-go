# 📗 Cisco IOS-XE Wireless Go SDK

Dependency‑free (stdlib only), strongly typed Go SDK for Cisco Catalyst 9800 (IOS‑XE 17.12) wireless RESTCONF (GET only).

| Aspect           | Value                      |
| ---------------- | -------------------------- |
| Domains          | 25+ (oper + cfg)           |
| Dependencies     | Stdlib only                |
| Coverage Gate    | ≥99%                       |
| Security Default | TLS verify ON              |
| Layering         | Client → Services → Models |

## 🚀 Quick Start

```bash
go get github.com/umatare5/cisco-ios-xe-wireless-go
export WNC_CONTROLLER=<controller-host-or-ip>
export WNC_ACCESS_TOKEN=$(echo -n 'admin:password' | base64)
```

```go
client, err := wnc.NewClient(
	os.Getenv("WNC_CONTROLLER"),
	os.Getenv("WNC_ACCESS_TOKEN"),
	wnc.WithTimeout(30*time.Second),
)
if err != nil { log.Fatal(err) }
ctx := context.Background()
gen, err := client.General().Oper(ctx)
if err != nil { log.Fatal(err) }
_ = gen
```

Options: `WithTimeout(d)`, `WithInsecureSkipVerify(true)` (dev only), `WithLogger(l)`.

## 📚 Navigation

| Topic                  | Location                        |
| ---------------------- | ------------------------------- |
| API overview           | `docs/api/README.md`            |
| Core services          | `docs/api/services_core.md`     |
| Extended services      | `docs/api/services_extended.md` |
| Testing & coverage     | `docs/testing/`                 |
| Security               | `docs/security/`                |
| Scripts / YANG tooling | `docs/scripts/`                 |
| Contributing           | `CONTRIBUTING.md`               |

## 🛠 Build & Test

```bash
make lint             # static analysis
make test-unit        # unit (runs lint)
make test-integration # live (needs env)
make test-coverage    # merged coverage -> coverage/report.out
```

Integration requires: `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` (base64 `user:pass`).

## 🔧 Service Call Pattern

Each service exposes `Oper(ctx)` (plus domain specific variants) returning typed YANG‑aligned structs. Shared HTTP + decode lives in internal `core.Get[T]`.

## ✅ Policies

| Area        | Rule                                      |
| ----------- | ----------------------------------------- |
| Errors      | Wrap context (`fmt.Errorf("x: %w", err)`) |
| Panics      | None in library code                      |
| Logging     | Only via user supplied logger             |
| Third‑Party | Forbidden (stdlib only)                   |
| Coverage    | ≥99% before merge                         |

## 📊 Coverage

Artifact: `coverage/report.out` (`make test-coverage`).

## 🔍 Additional (Collapsed)

<details><summary>Design & tips</summary>

YANG optional leaves → pointer fields: always nil‑check. Prefer one long‑lived client (connection reuse). Use context deadlines for every call. Avoid disabling TLS verify outside development.

</details>

## 📄 License

MIT (see `LICENSE`).
