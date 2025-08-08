# 📚 API Reference

This document provides a comprehensive reference for all services and methods available in the Cisco IOS-XE Wireless Go SDK.

## 🏗️ Architecture Overview

The library follows a **unified client architecture** where all domain services are accessed through a single client instance. This approach provides a clean, consistent interface for all wireless controller operations.

> Boilerplate for simple GET endpoints is minimized internally using a generic helper `core.Get[T]`, keeping service methods concise and uniform without sacrificing clarity.

```go
import wnc "github.com/umatare5/cisco-ios-xe-wireless-go"

// Create unified client
client, err := wnc.NewClient(host, token, options...)

// Access domain services directly
afcData, err := client.AFC().Oper(ctx)
apData, err := client.AP().Oper(ctx)
generalData, err := client.General().Oper(ctx)
```

## 🌟 Unified Client

### Creating a Client

```go
package main

import (
    "context"
    "log"
    "time"

    wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

func main() {
    client, err := wnc.NewClient(
        "your-controller-host",
        "your-api-token",
        wnc.WithTimeout(30*time.Second),
        wnc.WithInsecureSkipVerify(true),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Use client for API operations
    ctx := context.Background()
    data, err := client.General().Oper(ctx)
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }

    log.Printf("Device: %s", *data.GeneralOperData.EwlcGeneralOper.DeviceName)
}
```

### 🔧 Client Options

<details>
<summary>View all available configuration options</summary>

| Option | Description |
|--------|-------------|
| `WithTimeout(duration)` | Set request timeout |
| `WithInsecureSkipVerify(bool)` | Skip TLS certificate verification |
| `WithDebug(bool)` | Enable debug logging |
| `WithRetry(count, delay)` | Configure retry behavior |

</details>

## 📡 Domain Services

### AFC - Automated Frequency Coordination

Access AFC operations for 6 GHz band coordination.

```go
afc := client.AFC()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.AfcOperResponse` | AFC operational data |
| `APResp(ctx)` | `*model.AfcOperEwlcAfcApRespResponse` | Per-AP AFC responses |

### AP - Access Point Management

Comprehensive access point configuration and operational data.

```go
ap := client.AP()
```

#### Configuration Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.ApCfgResponse` | Complete AP configuration |
| `TagSourcePriorityConfigs(ctx)` | `*model.TagSourcePriorityConfigs` | Tag priority configurations |
| `ApTags(ctx)` | `*model.ApCfgApTagsResponse` | AP tag configurations |

#### Operational Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.ApOperResponse` | Complete AP operational data |
| `GlobalOper(ctx)` | `*model.ApGlobalOperResponse` | Global AP operational data |
| `ApJoinStats(ctx)` | `*model.ApJoinStatsResponse` | AP join statistics |
| `ApHistory(ctx)` | `*model.ApHistoryResponse` | AP history data |
| `EwlcApStats(ctx)` | `*model.EwlcApStatsResponse` | EWLC AP statistics |

### APF - Application Policy Framework

Application policy configuration and enforcement.

```go
apf := client.APF()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.ApfCfgResponse` | APF configuration |

### AWIPS - Advanced Weather Interactive Processing System

Weather-related processing system operations.

```go
awips := client.AWIPS()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.AwipsOperResponse` | AWIPS operational data |

### BLE - Bluetooth Low Energy

BLE beacon and asset tracking operations.

```go
ble := client.BLE()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.BleOperResponse` | BLE operational data |

### Client - Wireless Client Management

Comprehensive wireless client tracking and statistics.

```go
clientSvc := client.Client()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.ClientOperResponse` | Complete client operational data |
| `CommonOperData(ctx)` | `*model.ClientOperCommonOperDataResponse` | Common client operational data |
| `Dot11OperData(ctx)` | `*model.ClientOperDot11OperDataResponse` | 802.11 operational data |
| `MobilityOperData(ctx)` | `*model.ClientOperMobilityOperDataResponse` | Mobility operational data |
| `SiOperData(ctx)` | `*model.ClientOperSiOperDataResponse` | SI operational data |
| `Dot11Stats(ctx)` | `*model.ClientOperDot11StatsResponse` | 802.11 statistics |
| `TrafficStats(ctx)` | `*model.ClientOperTrafficStatsResponse` | Traffic statistics |

### CTS - Cisco TrustSec

Security group tagging and policy enforcement.

```go
cts := client.CTS()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.CtsCfgResponse` | CTS configuration |

### Dot11 - 802.11 Wireless Standards

802.11 wireless standard configuration.

```go
dot11 := client.Dot11()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.Dot11CfgResponse` | 802.11 configuration |

### Dot15 - 802.15 Standards

802.15 standard configuration for IoT and sensor networks.

```go
dot15 := client.Dot15()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.Dot15CfgResponse` | 802.15 configuration |

### Fabric - SD-Access Fabric

Software-Defined Access fabric operations.

```go
fabric := client.Fabric()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.FabricCfgResponse` | Fabric configuration |

### Flex - FlexConnect

FlexConnect local switching and CAPWAP operations.

```go
flex := client.Flex()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.FlexCfgResponse` | FlexConnect configuration |

### General - General Controller Operations

Core controller configuration and operational data.

```go
general := client.General()
```

#### Operational Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.GeneralOperResponse` | General operational data |
| `MgmtIntfData(ctx)` | `*model.GeneralOperMgmtIntfDataResponse` | Management interface data |

#### Configuration Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.GeneralCfgResponse` | General configuration |
| `MewlcConfig(ctx)` | `*model.GeneralCfgMewlcConfigResponse` | MEWLC configuration |
| `CountryConfigs(ctx)` | `*model.GeneralCfgCountryConfigsResponse` | Country configurations |

### Geolocation - Location Services

Geographic positioning and location mapping.

```go
geolocation := client.Geolocation()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.GeolocationOperResponse` | Geolocation operational data |

### Hyperlocation - High-Precision Location

Enhanced location services with sub-meter accuracy.

```go
hyperlocation := client.Hyperlocation()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.HyperlocationOperResponse` | Hyperlocation operational data |
| `Profiles(ctx)` | `*model.HyperlocationProfilesResponse` | Hyperlocation profiles |

### LISP - Locator/Identifier Separation Protocol

LISP protocol operations for mobility and routing.

```go
lisp := client.LISP()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.LispOperResponse` | LISP operational data |

### Location - Location Services Configuration

Location services configuration and management.

```go
location := client.Location()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.LocationCfgResponse` | Location configuration |

### Mcast - Multicast Services

Multicast traffic management and IGMP operations.

```go
mcast := client.Mcast()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.McastOperResponse` | Multicast operational data |
| `FlexMediastreamClientSummary(ctx)` | `*model.McastOperFlexMediastreamClientSummaryResponse` | FlexConnect mediastream summary |
| `VlanL2MgidOp(ctx)` | `*model.McastOperVlanL2MgidOpResponse` | VLAN L2 multicast group operations |

### mDNS - Multicast DNS

Multicast DNS service discovery and statistics.

```go
mdns := client.Mdns()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.MdnsOperResponse` | mDNS operational data |
| `GlobalStats(ctx)` | `*model.MdnsGlobalStatsResponse` | Global mDNS statistics |
| `WlanStats(ctx)` | `*model.MdnsWlanStatsResponse` | Per-WLAN mDNS statistics |

### Mesh - Mesh Networking

Wireless mesh networking configuration and operations.

```go
mesh := client.Mesh()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.MeshOperResponse` | Mesh operational data |
| `Cfg(ctx)` | `*model.MeshCfgResponse` | Mesh configuration |

### Mobility - Client Mobility

Inter-controller client mobility and roaming.

```go
mobility := client.Mobility()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.MobilityOperResponse` | Mobility operational data |

### NMSP - Network Mobility Services Protocol

Location services protocol for third-party integration.

```go
nmsp := client.NMSP()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.NmspOperResponse` | NMSP operational data |
| `ClientRegistration(ctx)` | `*model.NmspClientRegistrationResponse` | Client registration data |
| `CmxConnection(ctx)` | `*model.NmspCmxConnectionResponse` | CMX connection data |
| `CmxCloudInfo(ctx)` | `*model.NmspCmxCloudInfoResponse` | CMX cloud information |

### Radio - Radio Management

Radio hardware configuration and control.

```go
radio := client.Radio()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.RadioCfgResponse` | Radio configuration |

### RF - Radio Frequency Management

RF profile and parameter management.

```go
rf := client.RF()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.RfCfgResponse` | RF configuration |

### RFID - Radio Frequency Identification

RFID tag tracking and asset management.

```go
rfid := client.RFID()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.RfidCfgResponse` | RFID configuration |

### Rogue - Rogue AP Detection

Rogue access point and client detection system.

```go
rogue := client.Rogue()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.RogueOperResponse` | Rogue operational data |
| `Stats(ctx)` | `*model.RogueStatsResponse` | Rogue statistics |
| `Data(ctx)` | `*model.RogueDataResponse` | Rogue detection data |
| `ApSummary(ctx)` | `*model.RogueApSummaryResponse` | Rogue AP summary |
| `ClientSummary(ctx)` | `*model.RogueClientSummaryResponse` | Rogue client summary |

### RRM - Radio Resource Management

Automated RF optimization and interference mitigation.

```go
rrm := client.RRM()
```

#### Configuration Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.RrmCfgResponse` | RRM configuration |

#### Operational Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.RrmOperResponse` | RRM operational data |
| `GlobalOper(ctx)` | `*model.RrmGlobalOperResponse` | Global RRM operational data |

### Site - Site Management

Site configuration and hierarchical management.

```go
site := client.Site()
```

#### Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Oper(ctx)` | `*model.SiteOperResponse` | Site operational data |

### WLAN - Wireless LAN Configuration

WLAN service set configuration and policy management.

```go
wlan := client.WLAN()
```

#### Configuration Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `Cfg(ctx)` | `*model.WlanCfgResponse` | Complete WLAN configuration |
| `CfgEntries(ctx)` | `*model.WlanCfgEntriesResponse` | WLAN configuration entries |
| `Policies(ctx)` | `*model.WlanPoliciesResponse` | WLAN policies |
| `PolicyListEntries(ctx)` | `*model.PolicyListEntriesResponse` | Policy list entries |
| `WirelessAaaPolicyConfigs(ctx)` | `*model.WirelessAaaPolicyConfigsResponse` | AAA policy configurations |

#### Operational Methods

| Method | Return Type | Description |
|--------|-------------|-------------|
| `GlobalOper(ctx)` | `*model.WlanGlobalOperResponse` | Global WLAN operational data |

## 🎯 Usage Patterns

### Basic Operations

```go
ctx := context.Background()

// Get general controller information
general, err := client.General().Oper(ctx)
if err != nil {
    log.Printf("Error: %v", err)
    return
}

fmt.Printf("Controller: %s\n", *general.GeneralOperData.EwlcGeneralOper.DeviceName)
fmt.Printf("Version: %s\n", *general.GeneralOperData.EwlcGeneralOper.SoftwareVersion)
```

### Error Handling

```go
data, err := client.AP().Oper(ctx)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "timeout"):
        log.Printf("Request timed out: %v", err)
    case strings.Contains(err.Error(), "404"):
        log.Printf("Endpoint not found: %v", err)
    case strings.Contains(err.Error(), "401"):
        log.Printf("Authentication failed: %v", err)
    default:
        log.Printf("API error: %v", err)
    }
    return
}
```

### Context with Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

data, err := client.Client().Oper(ctx)
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        log.Printf("Request timed out")
    } else {
        log.Printf("Request failed: %v", err)
    }
    return
}
```

### Concurrent Requests

```go
var wg sync.WaitGroup
var mu sync.Mutex
results := make(map[string]interface{})

// Fetch multiple service data concurrently
services := []string{"general", "ap", "wlan", "client"}
wg.Add(len(services))

for _, service := range services {
    go func(svc string) {
        defer wg.Done()
        var data interface{}
        var err error

        switch svc {
        case "general":
            data, err = client.General().Oper(ctx)
        case "ap":
            data, err = client.AP().Oper(ctx)
        case "wlan":
            data, err = client.WLAN().Cfg(ctx)
        case "client":
            data, err = client.Client().Oper(ctx)
        }

        if err == nil {
            mu.Lock()
            results[svc] = data
            mu.Unlock()
        } else {
            log.Printf("Failed to fetch %s data: %v", svc, err)
        }
    }(service)
}

wg.Wait()

// Process results
for service, data := range results {
    log.Printf("Successfully fetched %s data", service)
    // Process specific data types as needed
}
```

## 📚 Response Models

All API responses are **strongly typed** using Go structs that exactly match the Cisco IOS-XE RESTCONF API structure.

### Key Characteristics

- **YANG Model Compliance**: Structures follow Cisco YANG models exactly
- **JSON Marshaling**: All fields have proper JSON tags for automatic parsing
- **Null Safety**: Optional fields use pointers to handle nil values safely
- **Type Safety**: Compile-time type checking prevents runtime errors

### Example Response Structure

```go
type GeneralOperResponse struct {
    GeneralOperData *GeneralOperData `json:"Cisco-IOS-XE-wireless-general-oper:general-oper-data,omitempty"`
}

type GeneralOperData struct {
    EwlcGeneralOper *EwlcGeneralOper `json:"ewlc-general-oper,omitempty"`
}

type EwlcGeneralOper struct {
    DeviceName       *string `json:"device-name,omitempty"`
    SoftwareVersion  *string `json:"software-version,omitempty"`
    UptimeSeconds    *uint64 `json:"uptime-seconds,omitempty"`
    MaxApSupported   *uint32 `json:"max-ap-supported,omitempty"`
    // ... additional fields
}
```

### Working with Response Data

```go
// Safe field access with nil checking
data, err := client.General().Oper(ctx)
if err != nil {
    return err
}

if data.GeneralOperData != nil &&
   data.GeneralOperData.EwlcGeneralOper != nil {

    if name := data.GeneralOperData.EwlcGeneralOper.DeviceName; name != nil {
        fmt.Printf("Device Name: %s\n", *name)
    }

    if version := data.GeneralOperData.EwlcGeneralOper.SoftwareVersion; version != nil {
        fmt.Printf("Software Version: %s\n", *version)
    }
}
```

## 🔒 Authentication & Security

### Bearer Token Authentication

The library uses **bearer token authentication** for secure API access:

```go
client, err := wnc.NewClient(
    "https://controller.example.com",
    "your-bearer-token-here",
    wnc.WithTimeout(30*time.Second),
)
```

### TLS Configuration

#### Production Environment

```go
client, err := wnc.NewClient(
    "https://controller.example.com",
    "your-bearer-token",
    wnc.WithInsecureSkipVerify(false), // Verify certificates
    wnc.WithTimeout(30*time.Second),
)
```

#### Development/Testing Environment

```go
client, err := wnc.NewClient(
    "https://controller.example.com",
    "your-bearer-token",
    wnc.WithInsecureSkipVerify(true), // Skip certificate verification
    wnc.WithTimeout(30*time.Second),
)
```

## 🚀 Best Practices

### 1. Context Management

Always use `context.Context` for timeout and cancellation control:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

data, err := client.General().Oper(ctx)
```

### 2. Error Handling

Check and handle all returned errors appropriately:

```go
data, err := client.AP().Oper(ctx)
if err != nil {
    log.Printf("API call failed: %v", err)
    return fmt.Errorf("failed to get AP data: %w", err)
}
```

### 3. Client Reuse

Create one client instance and reuse it across your application:

```go
// Good: Single client instance
var client *wnc.Client

func init() {
    var err error
    client, err = wnc.NewClient(host, token, options...)
    if err != nil {
        log.Fatal(err)
    }
}

func getAPData(ctx context.Context) (*model.ApOperResponse, error) {
    return client.AP().Oper(ctx)
}
```

### 4. Graceful Degradation

Handle partial failures gracefully in concurrent operations:

```go
func fetchAllData(ctx context.Context) error {
    var errors []error

    if _, err := client.General().Oper(ctx); err != nil {
        errors = append(errors, fmt.Errorf("general data: %w", err))
    }

    if _, err := client.AP().Oper(ctx); err != nil {
        errors = append(errors, fmt.Errorf("AP data: %w", err))
    }

    if len(errors) > 0 {
        // Log errors but continue with available data
        for _, err := range errors {
            log.Printf("Warning: %v", err)
        }
    }

    return nil
}
```

### 5. Resource Management

Use defer statements for proper cleanup:

```go
func processData(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel() // Ensure context is cancelled

    data, err := client.Client().Oper(ctx)
    if err != nil {
        return err
    }

    // Process data...
    return nil
}
```

## 📖 Additional Resources

- **YANG Models**: [Cisco XE 17121 YANG Models](https://github.com/YangModels/yang/tree/main/vendor/cisco/xe/17121)
- **RESTCONF Guide**: [Cisco IOS-XE RESTCONF Documentation](https://developer.cisco.com/docs/ios-xe/)
- **Go Context**: [Go Context Package Documentation](https://pkg.go.dev/context)
- **Cisco Catalyst 9800**: [Official Product Documentation](https://www.cisco.com/c/en/us/support/wireless/catalyst-9800-series-wireless-controllers/tsd-products-support-series-home.html)
