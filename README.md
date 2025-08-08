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

### Basic Usage

Start with this simple example to verify your WNC connection and credentials.

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

## 🌐 API Reference

The library provides a set of functions for interacting with all major Cisco Catalyst 9800 WNC subsystems. For detailed API documentation, please see **[API_REFERENCE.md](./docs/API_REFERENCE.md)**.

## 🤝 Contributing

We welcome contributions to this project! Please see our **[CONTRIBUTING.md](./docs/CONTRIBUTING.md)** for guidelines on how to contribute.

## 🙏 Acknowledgments

This code was developed with the assistance of **GitHub Copilot Agent Mode**. I extend our heartfelt gratitude to the global developer community who have contributed their knowledge and code to open source projects and public repositories.

## 📄 License

Please see the [LICENSE](./LICENSE) file for details.
