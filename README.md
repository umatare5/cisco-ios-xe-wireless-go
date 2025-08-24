<h1 align="center">cisco-ios-xe-wireless-go - Go Library for C9800</h1>

<p align="center">
  <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/umatare5/cisco-ios-xe-wireless-go?label=Latest%20version" />
  <a href="https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml"><img alt="Test and Build" src="https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml/badge.svg?branch=main" /></a>
  <img alt="Test Coverage" src="docs/assets/coverage.svg" />
  <a href="https://goreportcard.com/report/github.com/umatare5/cisco-ios-xe-wireless-go"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/umatare5/cisco-ios-xe-wireless-go" /></a>
  <a href="https://www.bestpractices.dev/projects/10969"><img alt="OpenSSF Best Practices" src="https://www.bestpractices.dev/projects/10969/badge" /></a>
  <a href="https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main"><img alt="Go Reference" src="https://pkg.go.dev/badge/umatare5/cisco-ios-xe-wireless-go.svg" /></a>
  <a href="./LICENSE"><img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" /></a>
  <a href="https://developer.cisco.com/codeexchange/github/repo/umatare5/cisco-ios-xe-wireless-go"><img alt="Published" src="https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg" /></a>
</p>

<p align="center">A Go library for interacting with Cisco Catalyst 9800 Wireless Network Controller.</p>

## ✨️ Key Features

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

> [!NOTE]
> You have to enable RESTCONF and HTTPS on the C9800 before using this library. Please see:
>
> - [Cisco IOS XE 17.12 Programmability Configuration Guide — RESTCONF](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1712/b_1712_programmability_cg/m_1712_prog_restconf.html#id_70432)

### 1. Generate a Basic Auth token

Encode your controller credentials as Base64.

```bash
# username:password → Base64
echo -n "admin:your-password" | base64
# Output: YWRtaW46eW91ci1wYXNzd29yZA==
```

### 2. Create a sample application

Use your controller host and token to fetch AP operational data.

```go
package main

import wnc "github.com/umatare5/cisco-ios-xe-wireless-go"

func main() {
    // Load environment variables
    controller := os.Getenv("WNC_CONTROLLER")
    token := os.Getenv("WNC_ACCESS_TOKEN")

    // Create client
    client, err := wnc.NewClient(controller, token,
        wnc.WithTimeout(30*time.Second),
        wnc.WithInsecureSkipVerify(true), // remove for production
    )

    // Create simple context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    // Request AP operational data
    apData, err := client.AP().GetOper(ctx)
    if err != nil {
        fmt.Fprintln(os.Stderr, "AP oper request:", err)
        os.Exit(1)
    }

    // Print AP operational data
    fmt.Printf("Successfully connected! Found %d APs\n",
        len(apData.CiscoIOSXEWirelessAccessPointOperAccessPointOperData.OperData)
    )
}
```

> [!CAUTION]
> The `wnc.WithInsecureSkipVerify(true)` option disables TLS certificate verification. This should only be used in development environments or when connecting to controllers with self-signed certificates. **Never use this option in production environments** as it compromises security.

### 3. Run the application with environment variables

```bash
# Set environment variables
export WNC_CONTROLLER="wnc1.example.internal"
export WNC_ACCESS_TOKEN="YWRtaW46eW91ci1wYXNzd29yZA=="

# Run the application
go run main.go
```

## 🌐 API Reference

This library provides a client to interact with the Cisco Catalyst 9800 Wireless Network Controller's RESTCONF API.

### Client Options

There are several options to customize the client behavior:

| Option                      | Type            | Default                    | Description                              |
| --------------------------- | --------------- | -------------------------- | ---------------------------------------- |
| `WithTimeout(d)`            | `time.Duration` | `wnc.DefaultTimeout` (60s) | Sets HTTP request timeout; must be > 0.  |
| `WithInsecureSkipVerify(b)` | `bool`          | `false`                    | Skips TLS verify; use only in labs.      |
| `WithLogger(l)`             | `*slog.Logger`  | `slog.Default()`           | Sets structured logger; cannot be nil.   |
| `WithUserAgent(ua)`         | `string`        | `"wnc-go-client/1.0"`      | Custom User-Agent; may be ignored today. |

### Supported Service APIs

Please refer to **[https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main)**.

> [!TIP]
>
> - Most Set operations in AP service require either the AP’s **Radio MAC** (`radioMac`) or **Ethernet MAC** (`apMac`). If you want to specify the AP by hostname, you can use the following function to perform the conversion.
> - `wtpMac` refers to the **Radio MAC** as same as `radioMac`. `WTP` (Wireless Termination Point) is a term defined in [RFC 5415: CAPWAP Protocol Specification](https://datatracker.ietf.org/doc/html/rfc5415) that denotes an AP.

## 🔖 Usecases

Runnable examples are available:

### Count Associated APs

[`cmd/count_aps`](./cmd/count_aps) — count associated APs

```bash
❯ go run cmd/count_aps/main.go
Successfully connected! Found 2 APs
```

## Reload an AP

[`cmd/reload_ap`](./cmd/reload_ap) — manually reload an AP **(dangerous operation)**

## Reload a Controller

[`cmd/reload_controller`](./cmd/reload_controller) — manually reload the controller **(dangerous operation)**

## 🤝 Contributing

I welcome all kinds of contributions from the community! Please read the **[Contribution Guide](./CONTRIBUTING.md)** before submitting PRs or issues. For additional guidance, please also see the following documents:

- **📋 [Make Command Reference](./docs/MAKE_REFERENCE.md)** — Make targets and the usage
- **📜 [Scripts Reference](./docs/SCRIPT_REFERENCE.md)** — Per-script usage and sample outputs
- **🧪 [Testing Guide](./docs/TESTING.md)** — How to run unit and integration tests

## 🙏 Acknowledgments

This project was developed with the assistance of **GitHub Copilot Agent Mode**. Thanks to the global developer community who have contributed their knowledge and code to open source projects and public repositories.

## 📄 License

[MIT](./LICENSE)
