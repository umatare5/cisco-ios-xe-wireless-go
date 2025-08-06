# 📗 cisco-ios-xe-wireless-go - Go Library for C9800

![GitHub Tag](https://img.shields.io/github/v/tag/umatare5/cisco-ios-xe-wireless-go?label=Latest%20version)
[![Test and Build](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml/badge.svg?branch=main)](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/umatare5/cisco-ios-xe-wireless-go)](https://goreportcard.com/report/github.com/umatare5/cisco-ios-xe-wireless-go)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/10969/badge)](https://www.bestpractices.dev/projects/10969)
[![Go Reference](https://pkg.go.dev/badge/umatare5/cisco-ios-xe-wireless-go.svg)](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/umatare5/cisco-ios-xe-wireless-go/blob/main/LICENSE)
[![Published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/umatare5/cisco-ios-xe-wireless-go)

A Go library for interacting with Cisco Catalyst 9800 Wireless Network Controller.

- **🔧 Developer Friendly**: Transparent YANG model handling with all responses in JSON format
- **📊 Comprehensive Coverage**: Access most status information and metrics available from the WNC
- **🚀 Quick Integration**: Get started in minutes with simple configuration and clear examples
- **🎯 Type-Safe Operations**: Strongly-typed Go structs for all API interactions and responses
- **📖 Comprehensive Documentation**: Detailed API reference, testing guides, and best practices

## 📡 Supported Environment

Cisco Catalyst 9800 Wireless Network Controller running Cisco IOS-XE `17.12.x`.

## 📦 Installation

```bash
go get github.com/umatare5/cisco-ios-xe-wireless-go
```

## 🚀 Quick Start

### 🔑 Creating Basic Auth Token

You must create a Basic Auth token using your Cisco WNC credentials before using the client.

```bash
# Create token for username:password
echo -n "admin:your-password" | base64
# Output: YWRtaW46eW91ci1wYXNzd29yZA==
```

### 🔧 Basic Usage

Start with this simple example to verify your WNC connection and credentials using the new service-based API.

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
    "github.com/umatare5/cisco-ios-xe-wireless-go/afc"
    "github.com/umatare5/cisco-ios-xe-wireless-go/general"
)

func main() {
    // Create configuration
    config := wnc.Config{
        Controller:  "192.168.1.100",
        AccessToken: "YWRtaW46eW91ci1wYXNzd29yZA==",
        Timeout:     30 * time.Second,
    }

    // Create client
    client, err := wnc.New(config.Controller, config.AccessToken,
        wnc.WithTimeout(config.Timeout))
    if err != nil {
        fmt.Printf("Failed to create client: %v\n", err)
        return
    }

    // Create services using the core client
    afcService := afc.NewService(client)
    generalService := general.NewService(client)

    // Use service-based API
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Get AFC 6 GHz statistics
    stats, err := afcService.CloudOper(ctx)
    if err != nil {
        fmt.Printf("Failed to get AFC stats: %v\n", err)
        return
    }

    fmt.Printf("Successfully connected to WNC!\n")
    fmt.Printf("AFC status: %s\n", stats.Cisco.Response.Code)

    // Get general operational data
    operData, err := generalService.Oper(ctx)
    if err != nil {
        fmt.Printf("Failed to get operational data: %v\n", err)
        return
    }

    fmt.Printf("Controller model: %s\n",
        operData.GeneralOperData.ControllerDetail.Model)
}
```

### ⚙️ Advanced Configuration

Customize client behavior using configuration options to optimize for your specific environment and requirements.

```go
import (
    "log/slog"
    "time"

    "github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Create client with custom configuration
client, err := wnc.New("192.168.1.100", "YWRtaW46eW91ci1wYXNzd29yZA==",
    wnc.WithTimeout(30*time.Second),
    wnc.WithInsecureSkipVerify(true), // Only for development
    wnc.WithLogger(customLogger),
)
if err != nil {
    fmt.Printf("Failed to create client: %v\n", err)
    return
}
```

> [!CAUTION]
> The `WithInsecureSkipVerify(true)` option disables TLS certificate verification. This should only be used in development environments or when connecting to controllers with self-signed certificates. **Never use this option in production environments** as it compromises security.

### 📊 Custom Logging

The library supports structured logging using Go's standard `slog` package.

```go
import (
    "log/slog"
    "os"

    "github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelDebug,
}))

client, err := wnc.New("192.168.1.100", "YWRtaW46eW91ci1wYXNzd29yZA==",
    wnc.WithLogger(logger),
)
```

## ⚙️ Configuration Options

All configuration options are set using functional options in the `wnc.New()` constructor.

| Option                         | Type            | Description                                    |
| ------------------------------ | --------------- | ---------------------------------------------- |
| `WithTimeout(d)`               | `time.Duration` | HTTP request timeout (default: 15s)            |
| `WithInsecureSkipVerify(bool)` | `bool`          | Skips TLS certificate verification (dev only)  |
| `WithLogger(l)`                | `*slog.Logger`  | Custom structured logger instance              |

## 🔄 Service-Based API

This library uses a modern three-layer architecture with domain-specific services. Each functional domain provides a service with typed methods accessible via `client.<Domain>().<Method>()`.

### 📋 Available Services

```go
import (
    "github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
    "github.com/umatare5/cisco-ios-xe-wireless-go/afc"
    "github.com/umatare5/cisco-ios-xe-wireless-go/ap"
    "github.com/umatare5/cisco-ios-xe-wireless-go/general"
    "github.com/umatare5/cisco-ios-xe-wireless-go/rogue"
    "github.com/umatare5/cisco-ios-xe-wireless-go/rrm"
    "github.com/umatare5/cisco-ios-xe-wireless-go/wlan"
)

// Create client and services
client, _ := wnc.New("192.168.1.100", "YWRtaW46eW91ci1wYXNzd29yZA==")
afcService := afc.NewService(client)
apService := ap.NewService(client)
generalService := general.NewService(client)
rogueService := rogue.NewService(client)
rrmService := rrm.NewService(client)
wlanService := wlan.NewService(client)

// Use domain services with typed methods
afcOper, _ := afcService.Oper(ctx)                       // AFC operational data
afcStats, _ := afcService.CloudStats(ctx)                // AFC 6 GHz statistics
apCfg, _ := apService.Cfg(ctx)                           // AP configuration
apOper, _ := apService.Oper(ctx)                         // AP operational data
generalOper, _ := generalService.Oper(ctx)               // General operational data
generalCfg, _ := generalService.Cfg(ctx)                 // General configuration
rogueOper, _ := rogueService.Oper(ctx)                   // Rogue detection data
rrmGlobal, _ := rrmService.GlobalOper(ctx)               // RRM global operations
wlanCfg, _ := wlanService.Cfg(ctx)                       // WLAN configuration
```

### 🏗️ Service Architecture

All services follow a consistent pattern:

- **Service Creation**: `service.NewService(client)` pattern
- **Typed Methods**: Each method returns strongly-typed structs from `internal/model`
- **Context Support**: All methods accept `context.Context` for timeouts and cancellation
- **Error Handling**: Consistent error types including HTTP status details

### ⚠️ Breaking Changes and Deprecations

**Legacy helper functions were deprecated in v1.5 and will be removed in v2.0.0:**

```go
// ❌ Deprecated - will be removed in v2.0.0
apData, err := ap.GetApOper(ctx, client)
afcData, err := afc.GetAfcOper(ctx, client)

// ✅ Use service-based API instead
apService := ap.NewService(client)
afcService := afc.NewService(client)
apData, err := apService.Oper(ctx)
afcData, err := afcService.Oper(ctx)
```

**Large API interfaces are also deprecated:**

- `WirelessControllerAPI` - use `client.<Domain>()` methods
- `AccessPointAPI` - use `client.AP()` methods
- `AFCControllerAPI` - use `client.AFC()` methods
- All other `*API` interfaces

## 🌐 API Reference

The library provides a set of functions for interacting with all major Cisco Catalyst 9800 WNC subsystems. For detailed API documentation, please see **[API_REFERENCE.md](./docs/API_REFERENCE.md)**.

## 🧪 Testing

This library includes comprehensive unit and integration tests to ensure reliability and compatibility with Cisco Catalyst 9800 controllers. For detailed testing information, please see **[TESTING.md](./docs/TESTING.md)**.

## 🛠️ Debugging

This library includes the scripts that are useful for debugging and development. These scripts use `curl` to access WNC, so they don't depend on Go. For detailed scripts documentation, please refer to **[SCRIPT_REFERENCE.md](./docs/SCRIPT_REFERENCE.md)**.

## 🤝 Contributing

I welcome contributions to improve this library. Please follow these guidelines to ensure smooth collaboration.

1. **Fork the repository** and create a feature branch from `main`
2. **Make your changes** following existing code style and conventions
3. **Add comprehensive tests** for new functionality
4. **Update documentation** including README.md and code comments
5. **Ensure all tests pass** including unit and integration tests
6. **Submit a pull request** with a clear description of changes

### 🔧 Code Quality Standards

This project maintains high code quality standards:

- **Error Handling**: All client validation uses standardized error wrapping: `fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)`
- **Test Coverage**: Maintains ≥98% coverage for main codebase, ≥92% total project coverage
- **Mock Testing**: HTTP tests use full RESTCONF paths: `/restconf/data/[YANG-MODULE]:[CONTAINER]/[ENDPOINT]`
- **Code Consistency**: Follows established Go best practices and project conventions
- **Import Management**: Uses `goimports` for consistent import formatting

## 🚀 Release Process

To release a new version:

1. **Update the version** in the `VERSION` file
2. **Submit a pull request** with the updated `VERSION` file

Once merged, the GitHub Workflow will automatically:

- **Create and push a new tag** based on the `VERSION` file

After that, manual release via [GitHub Actions: release workflow](https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/release.yaml).

## 🙏 Acknowledgments

This code was developed with the assistance of **GitHub Copilot Agent Mode**. I extend our heartfelt gratitude to the global developer community who have contributed their knowledge and code to open source projects and public repositories.

## 📄 License

Please see the [LICENSE](./LICENSE) file for details.
