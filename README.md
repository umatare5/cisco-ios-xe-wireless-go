<div align="center">

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="docs/assets/logo_dark.png" width="400px">
  <img src="docs/assets/logo.png" width="400px">
</picture>

  <h1>cisco-ios-xe-wireless-go</h1>

  <p>A Go SDK for interacting with Cisco Catalyst 9800 Wireless Network Controller.</p>

  <p>
    <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/umatare5/cisco-ios-xe-wireless-go?label=Latest%20version" />
    <a href="https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml"><img alt="Test and Build" src="https://github.com/umatare5/cisco-ios-xe-wireless-go/actions/workflows/go-test-build.yml/badge.svg?branch=main" /></a>
    <img alt="Test Coverage" src="docs/assets/coverage.svg" />
    <a href="https://goreportcard.com/report/github.com/umatare5/cisco-ios-xe-wireless-go"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/umatare5/cisco-ios-xe-wireless-go" /></a><br/>
    <a href="https://www.bestpractices.dev/projects/10969"><img alt="OpenSSF Best Practices" src="https://www.bestpractices.dev/projects/10969/badge" /></a>
    <a href="./LICENSE"><img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" /></a>
    <a href="https://developer.cisco.com/codeexchange/github/repo/umatare5/cisco-ios-xe-wireless-go"><img alt="Published" src="https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg" /></a>
  </p>

</div>

## ✨️ Key Features

- 🔧 **Developer-Friendly**: Seamless YANG model handling with responses consistently in JSON
- 🚀 **Fast Integration**: Start in minutes with straightforward setup and clear examples
- 📊 **Broad Coverage**: Access most configurations and statistics provided by the WNC
- 🎯 **Type-Safe Operations**: Strongly typed Go structs for reliable API calls and responses
- 📖 **Detailed Documentation**: Detailed API references, testing guides, and best practices via godoc

## 📡 Supported Environment

Cisco Catalyst 9800 Wireless Network Controller running on:

- **Cisco IOS-XE 17.12.x** - Verified on 17.12.6a
- **Cisco IOS-XE 17.15.x** - Verified on 17.15.4b
- **Cisco IOS-XE 17.18.x** - Verified on 17.18.1

## 📦 Installation

```bash
go get github.com/umatare5/cisco-ios-xe-wireless-go
```

## 🚀 Quick Start

You have to enable RESTCONF and HTTPS on the C9800 before using this SDK. Please see:

- [Cisco IOS XE 17.12 Programmability Configuration Guide — RESTCONF](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1712/b_1712_programmability_cg/m_1712_prog_restconf.html#id_70432)

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

import (
    "context"
    "fmt"
    "os"
    "time"

    wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

func main() {
    // Load environment variables
    controller := os.Getenv("WNC_CONTROLLER")
    token := os.Getenv("WNC_ACCESS_TOKEN")

    // Create client
    client, err := wnc.NewClient(controller, token,
        wnc.WithTimeout(30*time.Second),
        wnc.WithInsecureSkipVerify(true), // remove for production
    )
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create client: %v\n", err)
        os.Exit(1)
    }

    // Create simple context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    // Request AP operational data
    apData, err := client.AP().GetOperational(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "AP oper request failed: %v\n", err)
        os.Exit(1)
    }

    // Print AP operational data
    fmt.Printf("Successfully connected! Found %d APs\n",
        len(apData.CiscoIOSXEWirelessAPOperData.CAPWAPData))
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

# result: Successfully connected! Found 2 APs
```

## 🌐 API Reference

This SDK provides a client to interact with the Cisco Catalyst 9800 Wireless Network Controller's RESTCONF.

### Client Initialization

To create a new client, use the `wnc.NewClient` function with the controller address and access token.

| Parameter     | Type        | Description                            |
| ------------- | ----------- | -------------------------------------- |
| `controller`  | `string`    | The hostname or IP address of the WNC. |
| `accessToken` | `string`    | The Base64-encoded Basic Auth token.   |
| `options...`  | `...Option` | Optional client configuration options. |

### Client Options

There are several options to customize the client behavior.

| Option                      | Type            | Default             | Description                |
| --------------------------- | --------------- | ------------------- | -------------------------- |
| `WithTimeout(d)`            | `time.Duration` | `60s`               | Sets HTTP request timeout. |
| `WithInsecureSkipVerify(b)` | `bool`          | `false`             | Skips TLS verify.          |
| `WithLogger(l)`             | `*slog.Logger`  | `slog.Default()`    | Sets structured logger.    |
| `WithUserAgent(ua)`         | `string`        | `wnc-go-client/1.0` | Custom User-Agent.         |

### Supported Services

Please refer to the Go Reference for the complete reference.

<a href="https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main#section-documentation"><img alt="Go Reference" src="https://pkg.go.dev/badge/umatare5/cisco-ios-xe-wireless-go.svg" /></a>

The following table summarizes the supported service APIs and their capabilities.

**Legend:**

- ✅️ Supported
- 🟩 Partial Supported
- 🟨 Experimental Supported
- ⬜️ Not Supported

| API                                                                                                             | `GetOperational()` | `GetConfig()` | Other Functions | Notes                                                                                                                                              |
| --------------------------------------------------------------------------------------------------------------- | :----------------: | :-----------: | :-------------: | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| [`AFC()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/afc)                     |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`AP()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/ap)                       |        ✅️         |      ✅️      |       🟩        |                                                                                                                                                    |
| [`APF()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/apf)                     |        ⬜️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`AWIPS()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/awips)                 |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`BLE()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/ble)                     |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`Client()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/client)               |        ✅️         |      ⬜️      |       ⬜️       | Issue [#28](https://github.com/umatare5/cisco-ios-xe-wireless-go/issues/28), [#29](https://github.com/umatare5/cisco-ios-xe-wireless-go/issues/29) |
| [`Controller()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/controller)       |        ⬜️         |      ⬜️      |       🟩        |                                                                                                                                                    |
| [`CTS()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/cts)                     |        ⬜️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Dot11()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/dot11)                 |        ⬜️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Dot15()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/dot15)                 |        ⬜️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Fabric()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/fabric)               |        ⬜️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Flex()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/flex)                   |        ⬜️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`General()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/general)             |        ✅️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Geolocation()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/geolocation)     |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`Hyperlocation()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/hyperlocation) |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`LISP()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/lisp)                   |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`Location()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/location)           |        ✅️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Mcast()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/mcast)                 |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`MDNS()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/mdns)                   |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`Mesh()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/mesh)                   |        ✅️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Mobility()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/mobility)           |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`NMSP()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/nmsp)                   |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`Radio()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/radio)                 |        ⬜️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`RF()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/rf)                       |        ⬜️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`RFTag()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/rf)                    |        ⬜️         |      ⬜️      |       🟩        |                                                                                                                                                    |
| [`RFID()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/rfid)                   |        ✅️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Rogue()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/rogue)                 |        ✅️         |      ⬜️      |       ⬜️       |                                                                                                                                                    |
| [`RRM()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/rrm)                     |        ✅️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`Site()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/site)                   |        ✅️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`SiteTag()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/site)                |        ⬜️         |      ⬜️      |       🟩        |                                                                                                                                                    |
| [`Spaces()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/spaces)               |         🟨         |      ⬜️      |       ⬜️       | Requires 17.15 or later+                                                                                                                           |
| [`URWB()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/urwb)                   |         🟨         |      🟨       |       ⬜️       | Requires 17.18 or later+                                                                                                                           |
| [`WAT()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/wat)                     |        ⬜️         |      🟨       |       ⬜️       | Requires 17.18 or later+                                                                                                                           |
| [`WLAN()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/wlan)                   |        ✅️         |      ✅️      |       ⬜️       |                                                                                                                                                    |
| [`PolicyTag()`](https://pkg.go.dev/github.com/umatare5/cisco-ios-xe-wireless-go@main/service/wlan)              |        ⬜️         |      ⬜️      |       🟩        |                                                                                                                                                    |

> [!TIP]
>
> `wtpMac` is the same as `radioMac`. `WTP` (Wireless Termination Point), defined in [RFC 5415](https://datatracker.ietf.org/doc/html/rfc5415) denotes an AP.

## 🔖 Usecases

Runnable examples are available:

### List Operation

#### Usecase 1: List Associating APs

[`example/list_aps/main.go`](./example/list_aps/main.go) lists APs managed by the controller.

<details><summary><u>Click to show example</u></summary><p>

```bash
❯ go run example/list_aps/main.go

Successfully connected! Found 2 APs

AP Name           | MAC Address         | IP Address       | Status
------------------|---------------------|------------------|-----------------
TEST-AP01         | aa:bb:ff:dd:ee:a0   | 192.168.255.11   | registered
TEST-AP02         | aa:bb:ff:dd:ee:b0   | 192.168.255.12   | registered
```

</p></details>

#### Usecase 2: List Associating Clients

[`example/list_clients/main.go`](./example/list_clients/main.go) lists clients associating to wireless networks.

<details><summary><u>Click to show example</u></summary><p>

```bash
❯ go run example/list_clients/main.go

Successfully connected! Found 17 clients

MAC Address           | IP Address
----------------------|----------------
08:84:9d:92:47:00     | 192.168.0.84
2a:e3:42:8f:06:c8     | 192.168.0.89
40:23:43:3e:c5:bf     | 192.168.0.62
40:80:e1:6b:11:16     | 192.168.0.92
<snip>
```

</p></details>

#### Usecase 3: List WLANs and BSSIDs

[`example/list_wlans/main.go`](./example/list_wlans/main.go) lists WLANs and their BSSIDs.

<details><summary><u>Click to show example</u></summary><p>

```bash
❯ go run example/list_wlans/main.go

Successfully connected! Found 7 WLANs across all APs

AP Name           | AP MAC Address    | Slot | WLAN | BSSID             | SSID
------------------|-------------------|------|------|-------------------|-------------------------
TEST-AP01         | aa:bb:ff:dd:ee:a0 |    0 |    1 | aa:bb:ff:dd:ee:a1 | labo-wlan
TEST-AP01         | aa:bb:ff:dd:ee:a0 |    1 |    2 | aa:bb:ff:dd:ee:ad | labo-psk
TEST-AP01         | aa:bb:ff:dd:ee:a0 |    1 |    4 | aa:bb:ff:dd:ee:af | labo-tls
<snip>
```

</p></details>

#### Usecase 4: List AP Neighbors

[`example/list_neighbors/main.go`](./example/list_neighbors/main.go) lists neighboring APs detected by the APs.

<details><summary><u>Click to show example</u></summary><p>

```bash
❯ go run example/list_neighbors/main.go

Successfully connected! Found 11 AP neighbors

AP Name           | Slot | Neighbor BSSID    | Neighbor SSID          | RSSI  | Channel | Last Heard At
------------------|------|-------------------|------------------------|-------|---------|--------------------------
TEST-AP01         |    0 | d8:21:da:a2:30:f0 | Rogue-WiFi             |   -20 |      11 | 2025-09-12 20:24:57
TEST-AP01         |    0 | 08:10:86:bf:07:e3 | rogue-abcdef123-g      |   -62 |       4 | 2025-09-13 06:49:59
TEST-AP01         |    1 | 98:f1:99:c2:03:db | rogue-abcdef123        |   -64 |      36 | 2025-09-13 06:52:57
<snip>
```

</p></details>

### Destructive Operation

#### Usecase 1: Reload an AP

[`example/reload_ap/main.go`](./example/reload_ap/main.go) reloads a specified AP by its MAC address.

<details><summary><u>Click to show example</u></summary><p>

```bash
❯ go run example/reload_ap/main.go

=== Access Point Reload Tool ===
WARNING: This tool will restart access points causing service interruption!
Use only in controlled environments with proper authorization.

Target Controller: wnc1.example.internal
Enter AP MAC address (format: xx:xx:xx:xx:xx:xx or xx-xx-xx-xx-xx-xx): aa:bb:ff:dd:ee:a0
Target AP MAC: aa:bb:ff:dd:ee:a0
This will restart the specified Access Point(s). Type 'YES' to confirm: YES

✓ WNC client created successfully
Executing AP reload for MAC aa:bb:ff:dd:ee:a0
WARNING: AP will become unavailable and disconnect all clients during restart...

✓ AP reload command sent successfully for MAC: aa:bb:ff:dd:ee:a0
Note: AP is now restarting and will be temporarily unavailable
Clients will need to reconnect after AP restart completes
```

</p></details>

#### Usecase 2: Reload a Controller

[`example/reload_controller/main.go`](./example/reload_controller/main.go) reloads the entire wireless controller.

<details><summary><u>Click to show example</u></summary><p>

```bash
❯ go run ./example/reload_controller/main.go

=== WNC Controller Reload Tool ===
WARNING: This tool will restart the wireless controller!
Use only in controlled environments with proper authorization.

Target Controller: wnc1.example.internal

This will restart the WNC controller. Type 'YES' to confirm: YES

✓ WNC client created successfully
Executing controller reload with reason: Manual reload via CLI tool at 2025-09-06T13:11:50+09:00
WARNING: Controller will become unavailable during restart...

✓ Controller reload command sent successfully
Note: Controller is now restarting and will be temporarily unavailable
Wait for controller to complete restart before attempting reconnection
```

</p></details>

## 📦 Used By

- [cisco-wnc-exporter](https://github.com/umatare5/cisco-wnc-exporter) - Prometheus exporter for Cisco C9800 Wireless Network Controller metrics ([v0.3.6](https://github.com/umatare5/cisco-ios-xe-wireless-go/releases/tag/v0.3.6))

## 🤝 Contributing

Please read the **[Contribution Guide](./CONTRIBUTING.md)** before submitting PRs and issues and also see the following documents:

- **📋 [Make Command Reference](./docs/MAKE_REFERENCE.md)** — Make targets and the usage
- **📜 [Scripts Reference](./docs/SCRIPT_REFERENCE.md)** — Per-script usage and sample outputs
- **🧪 [Testing Guide](./docs/TESTING.md)** — How to run unit and integration tests

## 🙏 Acknowledgments

I maintain this project with the help of **GitHub Copilot Coding Agent**, and I'm grateful to the global developer community for their contributions to open source projects and public repositories.

## 📄 License

[MIT](./LICENSE)
