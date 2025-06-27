# 📗 cisco-ios-xe-wireless-go - Go Library for C9800

![GitHub Tag](https://img.shields.io/github/v/tag/umatare5/cisco-ios-xe-wireless-go?label=Latest%20version)
[![Go Reference](https://pkg.go.dev/badge/umatare5/cisco-ios-xe-wireless-go.svg)](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/umatare5/cisco-ios-xe-wireless-go)](https://goreportcard.com/report/github.com/umatare5/cisco-ios-xe-wireless-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/umatare5/cisco-ios-xe-wireless-go/blob/main/LICENSE)

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

Start with this simple example to verify your WNC connection and credentials.

```go
package main

import (
    "context"
    "fmt"
    "time"

    wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

func main() {
    // Create configuration
    config := wnc.Config{
        Controller:  "192.168.1.100",
        AccessToken: "YWRtaW46eW91ci1wYXNzd29yZA==",
        Timeout:     30 * time.Second,
    }

    // Create client with configuration
    client, err := wnc.NewClient(config)
    if err != nil {
        fmt.Printf("Failed to create client: %v\n", err)
        return
    }

    // Get AP operational data with context timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    apData, err := client.GetApOper(ctx)
    if err != nil {
        fmt.Printf("Failed to get AP data: %v\n", err)
        return
    }

    fmt.Printf("Successfully connected! Found %d APs\n", len(apData.CiscoIOSXEWirelessAccessPointOperAccessPointOperData.OperData))
}
```

### ⚙️ Advanced Configuration

Customize client behavior using configuration options to optimize for your specific environment and requirements.

```go
import (
    "log/slog"
    "time"

    wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// Create client with custom configuration
config := wnc.Config{
    Controller:         "192.168.1.100",
    AccessToken:        "YWRtaW46eW91ci1wYXNzd29yZA==",
    Timeout:            30 * time.Second,
    InsecureSkipVerify: true, // Only for development
}

client, err := wnc.NewClient(config)
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

    wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelDebug,
}))

config := wnc.Config{
    Controller:  "192.168.1.100",
    AccessToken: "YWRtaW46eW91ci1wYXNzd29yZA==",
    Logger:      logger,
}

client, err := wnc.NewClient(config)
```

## ⚙️ Configuration Options

All configuration options are set in the `Config` struct during client creation.

| Field                | Type            | Description                                    |
| -------------------- | --------------- | ---------------------------------------------- |
| `Controller`         | `string`        | Hostname or IP address of the WNC (required)   |
| `AccessToken`        | `string`        | Authentication token for API access (required) |
| `Timeout`            | `time.Duration` | HTTP request timeout (default: 15s)            |
| `InsecureSkipVerify` | `bool`          | Skips TLS certificate verification (dev only)  |
| `Logger`             | `*slog.Logger`  | Custom structured logger instance              |

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
