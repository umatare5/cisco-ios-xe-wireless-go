# 📗 Cisco IOS-XE Wireless Go SDK (Catalyst 9800 RESTCONF)

Strongly‑typed, dependency‑free Go library for the Cisco Catalyst 9800 (IOS‑XE 17.12) RESTCONF API.

| Key Property | Summary |
|--------------|---------|
| Scope | Read‑only (GET) access to 25+ wireless domains |
| Stability | No external deps, stdlib only |
| Accuracy | Structs mirror Cisco YANG (17.12) |
| Coverage Policy | ≥99% total (hard gate) |
| Architecture | Unified client → thin services → typed models |
| Boilerplate Reduction | Internal generic `core.Get[T]` |
| Security | TLS on by default; explicit opt‑out for dev only |

![GitHub Tag](https://img.shields.io/github/v/tag/umatare5/cisco-ios-xe-wireless-go?label=Latest%20version)
[![Test and Build](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml/badge.svg?branch=main)](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/umatare5/cisco-ios-xe-wireless-go)](https://goreportcard.com/report/github.com/umatare5/cisco-ios-xe-wireless-go)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/10969/badge)](https://www.bestpractices.dev/projects/10969)
[![Go Reference](https://pkg.go.dev/badge/umatare5/cisco-ios-xe-wireless-go.svg)](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/umatare5/cisco-ios-xe-wireless-go/blob/main/LICENSE)
[![Published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/umatare5/cisco-ios-xe-wireless-go)

---

## ✨ Overview

Purpose: Provide an idiomatic, testable Go interface to Catalyst 9800 wireless controller RESTCONF resources with minimal ceremony and maximum type safety.

Compatibility: Cisco Catalyst 9800 on IOS‑XE **17.12.x** (YANG set `vendor/cisco/xe/17121`).

## 📦 Installation

```bash
go get github.com/umatare5/cisco-ios-xe-wireless-go
```

## 🚀 Quick Start

### 🔑 Credentials

Generate Base64 token (`username:password`):

```bash
echo -n "admin:password" | base64
# YWRtaW46cGFzc3dvcmQ=
```

### 🔧 Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "time"

    wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

func main() {
    // Create unified client
    client, err := wnc.NewClient("192.168.1.100", "YWRtaW46cGFzc3dvcmQ=",
        wnc.WithTimeout(30*time.Second))
    if err != nil {
        panic(err)
    }

    ctx := context.Background()

    // Access domain services directly from client
    generalData, err := client.General().Oper(ctx)
    if err != nil {
        panic(err)
    }

    afcData, err := client.AFC().Oper(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Controller: %s\n", generalData.GeneralOperData.ControllerDetail.Model)
    fmt.Printf("AFC Status: %s\n", afcData.AfcOperData.Ewlc.Mode)
}
```

### ⚙️ Advanced Configuration

```go
import (
    "log/slog"
    wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

client, err := wnc.NewClient("192.168.1.100", "YWRtaW46cGFzc3dvcmQ=",
    wnc.WithTimeout(30*time.Second),
    wnc.WithInsecureSkipVerify(true), // Development only
    wnc.WithLogger(slog.Default()),
)
```

> [!WARNING]
> `WithInsecureSkipVerify(true)` disables TLS certificate verification. Use **only** for local dev with self‑signed certs.

## 🏗️ Architecture

Layers:

1. Unified client: handles transport, TLS, retry (if configured).
2. Service layer: thin, typed methods mapping to RESTCONF endpoints.
3. Model layer: YANG‑aligned structs (pointer fields for optional nodes).

### ♻️ Generic GET Helper

Simple read operations use a generic helper to remove repetition:

```go
// internal/core/get.go
func Get[T any](ctx context.Context, c *Client, endpoint string) (*T, error) {
    var out T
    return &out, c.Do(ctx, http.MethodGet, endpoint, &out)
}
```

Service method (example):

```go
func (s Service) Oper(ctx context.Context) (*model.RrmOperResponse, error) {
    return core.Get[model.RrmOperResponse](ctx, s.c, RrmOperEndpoint)
}
```

This keeps public methods predictable while preserving clarity.

## 📋 Services

| Service | Description | Key Methods |
| ------- | ----------- | ----------- |
| `AFC()` | Automated Frequency Coordination | `Oper()`, `CloudOper()`, `APResp()`, `CloudStats()` |
| `AP()` | Access Point Management | `Oper()`, `Cfg()`, `Stats()` |
| `General()` | Controller Configuration | `Oper()`, `Cfg()`, `MgmtIntfData()` |
| `WLAN()` | Wireless LAN Configuration | `Cfg()`, `ProfileCfg()`, `GlobalCfg()` |
| `RRM()` | Radio Resource Management | `GlobalOper()`, `RadioOper()` |
| `Rogue()` | Rogue Detection | `Oper()`, `Stats()`, `Data()`, `ClientData()` |
| `Client()` | Wireless Client Management | `Oper()`, `Stats()`, `RedirectAp()` |
| `BLE()` | Bluetooth Low Energy | `Cfg()`, `GlobalCfg()` |
| `Fabric()` | SD-Access Fabric | `GlobalCfg()` |
| `Mobility()` | Client Mobility | `Data()`, `GlobalCfg()` |

<details><summary>Full list (expand)</summary>

APF · AWIPS · CTS · Dot11 · Dot15 · Flex · Geolocation · Hyperlocation · LISP · Location · Mcast · Mdns · Mesh · NMSP · Radio · RF · RFID · Site plus the tabled domains above.

</details>

### 🎯 Usage Patterns

```go
// Configuration data
wlanConfig, _ := client.WLAN().Cfg(ctx)
apConfig, _ := client.AP().Cfg(ctx)

// Operational data
apOper, _ := client.AP().Oper(ctx)
clientOper, _ := client.Client().Oper(ctx)

// Statistics
rogueStats, _ := client.Rogue().Stats(ctx)
afcStats, _ := client.AFC().CloudStats(ctx)

// Global operations
rrmGlobal, _ := client.RRM().GlobalOper(ctx)
mobilityGlobal, _ := client.Mobility().GlobalCfg(ctx)
```

## 🔧 Client Options

| Option | Type | Description |
| ------ | ---- | ----------- |
| `WithTimeout(d)` | `time.Duration` | HTTP timeout (default: 15s) |
| `WithInsecureSkipVerify(bool)` | `bool` | Skip TLS verification |
| `WithLogger(l)` | `*slog.Logger` | Custom logger |

## 📖 Documentation

| Doc | Purpose |
|-----|---------|
| [`API_REFERENCE.md`](./docs/API_REFERENCE.md) | Service & model usage |
| [`TESTING.md`](./docs/TESTING.md) | Unit / integration / coverage policy |
| [`SCRIPT_REFERENCE.md`](./docs/SCRIPT_REFERENCE.md) | YANG helper scripts |
| [`SECURITY.md`](./docs/SECURITY.md) | TLS, credential hygiene |
| [`CONTRIBUTING.md`](./CONTRIBUTING.md) | Workflow & quality rules |

## 🧪 Testing

```bash
# Run all tests with coverage
make test-coverage

# Unit tests only
make test-unit

# Integration tests (requires WNC environment)
make test-integration
```

> [!NOTE]
> `make test-unit` and `make test-integration` automatically run `lint` first.

## 🛠️ Development Tasks

```bash
# Install dependencies
make deps

# Lint code
make lint

# List available YANG models
make yang-list

# Get model details
make yang-model MODEL=Cisco-IOS-XE-wireless-general-oper
```

## 🤝 Contributing

See [`CONTRIBUTING.md`](./CONTRIBUTING.md). Key points:

* Maintain ≥99% total coverage (no regressions).
* No new external dependencies.
* Follow existing error wrapping style and service method shape.
* Update docs + tests together.

## 🚀 Release (Outline)

1. Update `VERSION` file
2. Submit pull request
3. Automated tag creation on merge
4. Manual release via GitHub Actions

## 📊 Coverage Artifact

Human‑readable combined coverage snapshot: `coverage/report.out` (committed). Regenerate via `make test-coverage`.

## 📄 License

MIT – see [`LICENSE`](./LICENSE).
