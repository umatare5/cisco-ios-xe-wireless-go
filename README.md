<h1 align="center">📗 cisco-ios-xe-wireless-go - Go Library for C9800</h1>

<p align="center">
    <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/umatare5/cisco-ios-xe-wireless-go?label=Latest%20version" />
    <a href="https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml">
        <img alt="Test and Build" src="https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml/badge.svg?branch=main" />
    </a>
    <img alt="Test Coverage" src="docs/assets/coverage.svg" />
    <a href="https://goreportcard.com/report/github.com/umatare5/cisco-ios-xe-wireless-go">
        <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/umatare5/cisco-ios-xe-wireless-go" />
    </a>
    <a href="https://www.bestpractices.dev/projects/10969">
        <img alt="OpenSSF Best Practices" src="https://www.bestpractices.dev/projects/10969/badge" />
    </a>
    <a href="https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go">
        <img alt="Go Reference" src="https://pkg.go.dev/badge/umatare5/cisco-ios-xe-wireless-go.svg" />
    </a>
    <a href="./LICENSE">
        <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
    </a>
    <a href="https://developer.cisco.com/codeexchange/github/repo/umatare5/cisco-ios-xe-wireless-go">
        <img alt="Published" src="https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg" />
    </a>
  
</p>

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

### 1. Create your basic auth token

```bash
# username:password → Base64
echo -n "admin:your-password" | base64
# Output: YWRtaW46eW91ci1wYXNzd29yZA==
```

### 2. Write and run a sample code

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
    apData, err := client.AP().Oper(ctx)
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

> [!NOTE]
> Runnable examples are available:
>
> - **Minimal**: [`examples/minimal`](./examples/minimal) — create a client and call a single endpoint
> - **Advanced**: [`examples/advanced`](./examples/advanced) — multi-service workflow with logging and context

## 📚 Documentation

- 🌐 **[API Reference](./docs/API_REFERENCE.md)** — All available functions and supported endpoints
- 💉 **[Testing Guide](./docs/TESTING.md)** — How to run unit and integration tests
- 📜 **[Script Reference](./docs/SCRIPT_REFERENCE.md)** — Standalone `curl` scripts for debugging

> [!NOTE]
> Runnable examples are available:
>
> - **Minimal**: [`examples/minimal`](./examples/minimal) — create a client and call a single endpoint
> - **Advanced**: [`examples/advanced`](./examples/advanced) — multi-service workflow with logging and context

## 🤝 Contributing

Please read the **[Contributing Guide](./CONTRIBUTING.md)** before submitting PRs or issues. I welcome contributions from the community. 👍️

## 🙇‍♂️ Agreement

Current library is in development. Breaking changes may occur until `v1.0.0`.

The remaining tasks to reach to `v1.0.0` are available in the **[Milestone: 1.0.0](https://github.com/umatare5/cisco-ios-xe-wireless-go/milestone/1)**.

## 🙏 Acknowledgments

This project was developed with the assistance of **GitHub Copilot Agent Mode**. Thanks to the global open-source community for knowledge, tools, and inspiration.

## 📄 License

[MIT License](./LICENSE)
