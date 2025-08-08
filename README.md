# 📗 Cisco IOS-XE Wireless Go SDK (Cisco Catalyst 9800 RESTCONF)

**High-quality, strongly-typed Go client for the Cisco Catalyst 9800 (IOS-XE 17.12) RESTCONF API.**

- **🎯 Type-Safe Operations**: Strongly-typed data structures from YANG models
- **🏗️ Clean Architecture**: Three-layer architecture with domain-specific services
- **🔧 Developer-Friendly**: Unified client access with intuitive service patterns
- **🧩 Low Boilerplate**: Internal generic helper eliminates repetitive GET code
- **📊 Comprehensive Coverage**: 25+ functional domains

![GitHub Tag](https://img.shields.io/github/v/tag/umatare5/cisco-ios-xe-wireless-go?label=Latest%20version)
[![Test and Build](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml/badge.svg?branch=main)](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/umatare5/cisco-ios-xe-wireless-go)](https://goreportcard.com/report/github.com/umatare5/cisco-ios-xe-wireless-go)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/10969/badge)](https://www.bestpractices.dev/projects/10969)
[![Go Reference](https://pkg.go.dev/badge/umatare5/cisco-ios-xe-wireless-go.svg)](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/umatare5/cisco-ios-xe-wireless-go/blob/main/LICENSE)
[![Published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/umatare5/cisco-ios-xe-wireless-go)

---

## ✨ Overview

This library provides a production-grade, idiomatic Go interface to the Cisco Catalyst 9800 RESTCONF API. It focuses on **clarity**, **maintainability**, and **consistency**—removing boilerplate while preserving explicit, readable intent.

## 🎯 Compatibility

Cisco Catalyst 9800 running Cisco IOS-XE **17.12.x**

## 📦 Installation

```bash
go get github.com/umatare5/cisco-ios-xe-wireless-go
```

## 🚀 Quick Start

### 🔑 Authentication

Create a Base64 token from your WNC credentials:

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

> [!CAUTION]
> `WithInsecureSkipVerify(true)` disables TLS verification. Use only in development environments with self-signed certificates.

## 🏗️ Architecture

Three clean layers:

1. **Unified Client Layer** – single entrypoint exposing each service.
2. **Domain Service Layer** – lightweight service methods (thin wrappers) calling shared helpers.
3. **Generated Type Layer** – YANG-derived struct definitions ensuring type fidelity.

### ♻️ Boilerplate Reduction with `core.Get`

Most simple GET endpoints now use an internal generic helper:

```go
// internal/core/get.go
func Get[T any](ctx context.Context, c *Client, endpoint string) (*T, error) {
    var out T
    return &out, c.Do(ctx, http.MethodGet, endpoint, &out)
}
```

Service methods become concise and consistent:

```go
func (s Service) Oper(ctx context.Context) (*model.RrmOperResponse, error) {
    return core.Get[model.RrmOperResponse](ctx, s.c, RrmOperEndpoint)
}
```

This improves readability and reduces maintenance surface without hiding logic.

## 📋 Available Services

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

<details>
<summary>View all 25+ services</summary>

- **APF**: Application Policy Framework
- **AWIPS**: Advanced Weather Interactive Processing System
- **CTS**: Cisco TrustSec
- **Dot11**: 802.11 Wireless Standards
- **Dot15**: 802.15 Standards
- **Flex**: FlexConnect
- **Geolocation**: Location Services
- **Hyperlocation**: High-Precision Location
- **LISP**: Locator/Identifier Separation Protocol
- **Location**: Location Services
- **Mcast**: Multicast
- **Mdns**: Multicast DNS
- **Mesh**: Wireless Mesh
- **NMSP**: Network Mobility Services Protocol
- **Radio**: Radio Management
- **RF**: Radio Frequency Management
- **RFID**: RFID Tracking
- **Site**: Site-based Configuration

</details>

### 🎯 Service Usage Patterns

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

## 🔧 Configuration Options

| Option | Type | Description |
| ------ | ---- | ----------- |
| `WithTimeout(d)` | `time.Duration` | HTTP timeout (default: 15s) |
| `WithInsecureSkipVerify(bool)` | `bool` | Skip TLS verification |
| `WithLogger(l)` | `*slog.Logger` | Custom logger |

## 📖 Documentation

- **[API Reference](./docs/API_REFERENCE.md)**: Complete API documentation
- **[Testing Guide](./docs/TESTING.md)**: Testing strategies and coverage
- **[Script Reference](./docs/SCRIPT_REFERENCE.md)**: Development and debugging scripts

## 🧪 Testing

```bash
# Run all tests with coverage
make test-coverage

# Unit tests only
make test-unit

# Integration tests (requires WNC environment)
make test-integration
```

## 🛠️ Development

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

1. Fork and create feature branch
2. Follow existing patterns and conventions
3. Add comprehensive tests (≥98% coverage required)
4. Update documentation
5. Submit pull request

## 🚀 Release Process

1. Update `VERSION` file
2. Submit pull request
3. Automated tag creation on merge
4. Manual release via GitHub Actions

## 📄 License

[MIT License](./LICENSE)
