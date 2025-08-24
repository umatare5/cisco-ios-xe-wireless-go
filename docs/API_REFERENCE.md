# 🌐 API Reference

This document provides an overview of the API functions available in the WNC Go client library.

> [!NOTE]
> All data shapes map directly to Cisco YANG models for IOS-XE 17.12.1. See official [YANG Models](https://github.com/YangModels/yang/tree/main/vendor/cisco/xe/17121#readme) for field semantics.

## 🏗️ Client Construction

Following is an example of how to create a WNC client in Go. The client is stateless and can be reused for multiple requests.

```go
import (
  "context"
  "time"

  wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

func example() error {
  client, err := wnc.NewClient("wnc1.example.internal", "<base64token>",
    wnc.WithTimeout(30*time.Second),
    // wnc.WithInsecureSkipVerify(true), // lab only
  )
  if err != nil { return err }

  ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
  defer cancel()
  _, _ = client.General().GetOper(ctx)
  return nil
}
```

### Options

| Option                      | Type            | Default                    | Description                              |
| --------------------------- | --------------- | -------------------------- | ---------------------------------------- |
| `WithTimeout(d)`            | `time.Duration` | `wnc.DefaultTimeout` (60s) | Sets HTTP request timeout; must be > 0.  |
| `WithInsecureSkipVerify(b)` | `bool`          | `false`                    | Skips TLS verify; use only in labs.      |
| `WithLogger(l)`             | `*slog.Logger`  | `slog.Default()`           | Sets structured logger; cannot be nil.   |
| `WithUserAgent(ua)`         | `string`        | `"wnc-go-client/1.0"`      | Custom User-Agent; may be ignored today. |

> [!TIP]
>
> Defaults: TLS verification is `on` and headers include `Accept: application/yang-data+json` and a default User-Agent.

## 🧭 Domain Services

Each `*wnc.Client` accessor returns a lightweight, stateless value‑receiver service.

| Accessor          | Package         | Example Call                          |
| ----------------- | --------------- | ------------------------------------- |
| `AP()`            | `ap`            | `client.AP().GetOper(ctx)`            |
| `Client()`        | `client`        | `client.Client().GetOper(ctx)`        |
| `General()`       | `general`       | `client.General().GetOper(ctx)`       |
| `WLAN()`          | `wlan`          | `client.WLAN().GetCfg(ctx)`           |
| `RRM()`           | `rrm`           | `client.RRM().GetOper(ctx)`           |
| `Radio()`         | `radio`         | `client.Radio().GetCfg(ctx)`          |
| `RF()`            | `rf`            | `client.RF().GetCfg(ctx)`             |
| `Mobility()`      | `mobility`      | `client.Mobility().GetOper(ctx)`      |
| `AFC()`           | `afc`           | `client.AFC().GetOper(ctx)`           |
| `Rogue()`         | `rogue`         | `client.Rogue().GetOper(ctx)`         |
| `Mcast()`         | `mcast`         | `client.Mcast().GetOper(ctx)`         |
| `MDNS()`          | `mdns`          | `client.MDNS().GetOper(ctx)`          |
| `Mesh()`          | `mesh`          | `client.Mesh().GetOper(ctx)`          |
| `Hyperlocation()` | `hyperlocation` | `client.Hyperlocation().GetOper(ctx)` |
| `Geolocation()`   | `geolocation`   | `client.Geolocation().GetOper(ctx)`   |
| `Location()`      | `location`      | `client.Location().GetCfg(ctx)`       |
| `Site()`          | `site`          | `client.Site().GetOper(ctx)`          |
| `BLE()`           | `ble`           | `client.BLE().GetOper(ctx)`           |
| `CTS()`           | `cts`           | `client.CTS().GetCfg(ctx)`            |
| `Dot11()`         | `dot11`         | `client.Dot11().GetCfg(ctx)`          |
| `Dot15()`         | `dot15`         | `client.Dot15().GetCfg(ctx)`          |
| `Flex()`          | `flex`          | `client.Flex().GetCfg(ctx)`           |
| `Fabric()`        | `fabric`        | `client.Fabric().GetCfg(ctx)`         |
| `APF()`           | `apf`           | `client.APF().GetCfg(ctx)`            |
| `AWIPS()`         | `awips`         | `client.AWIPS().GetOper(ctx)`         |
| `RFID()`          | `rfid`          | `client.RFID().GetCfg(ctx)`           |
| `LISP()`          | `lisp`          | `client.LISP().GetOper(ctx)`          |
| `NMSP()`          | `nmsp`          | `client.NMSP().GetOper(ctx)`          |

Please refer to the following list for the methods of each service:

<details><summary><strong>Access Point Service</strong></summary>

The AP service provides comprehensive access point management with extensive filtering capabilities.

**Base Methods:**

| Method                        | Description                            |
| ----------------------------- | -------------------------------------- |
| `GetCfg`                      | Access point configuration for all APs |
| `GetTagSourcePriorityConfigs` | Tag source priority config             |
| `GetApTags`                   | AP tag assignments                     |
| `GetOper`                     | AP operational root data               |
| `GetNameMacMap`               | AP name ↔ MAC list                     |
| `GetCapwapData`               | CAPWAP session data                    |
| `GetGlobalOper`               | Global AP operational data             |
| `GetGlobalOperApHistory`      | AP history records                     |
| `GetGlobalOperEwlcApStats`    | EWLC AP statistics                     |

**Configuration Filter Methods (GetCfg\*):**

| Method              | Description               |
| ------------------- | ------------------------- |
| `GetCfgByMAC`       | Filter by AP MAC address  |
| `GetCfgByPolicyTag` | Filter by policy tag name |
| `GetCfgBySiteTag`   | Filter by site tag name   |
| `GetCfgByRfTag`     | Filter by RF tag name     |
| `GetCfgByPriority`  | Filter by priority value  |
| `GetCfgByTagSrc`    | Filter by tag source      |

**Operational Filter Methods (GetOper\*):**

| Method                           | Description                    |
| -------------------------------- | ------------------------------ |
| `GetOperByWtpName`               | Filter by WTP name             |
| `GetOperByWtpMac`                | Filter by WTP MAC address      |
| `GetOperByEthMac`                | Filter by Ethernet MAC address |
| `GetOperByRadioID`               | Filter by radio ID             |
| `GetOperByApIPAddr`              | Filter by AP IP address        |
| `GetOperByPrimaryControllerName` | Filter by primary controller   |
| `GetOperByPowerType`             | Filter by power type           |
| `GetOperBySlotID`                | Filter by slot ID              |
| `GetOperByPrimaryChannel`        | Filter by primary channel      |

**Global Operational Filter Methods (GetGlobalOper\*):**

| Method                              | Description                    |
| ----------------------------------- | ------------------------------ |
| `GetGlobalOperByApName`             | Filter by AP name              |
| `GetGlobalOperByWtpMac`             | Filter by WTP MAC address      |
| `GetGlobalOperByEthernetMac`        | Filter by Ethernet MAC address |
| `GetGlobalOperByApIPAddr`           | Filter by AP IP address        |
| `GetGlobalOperByIsJoined`           | Filter by join status          |
| `GetGlobalOperByApDisconnectReason` | Filter by disconnect reason    |

**Examples:**

```go
// Get all APs
aps, err := client.AP().GetCfg(ctx)

// Get specific AP by MAC address
ap, err := client.AP().GetCfgByMAC(ctx, "aa:bb:cc:dd:ee:ff")

// Get operational data for specific WTP
operData, err := client.AP().GetOperByWtpName(ctx, "AP-01")

// Get global operational data by AP name
globalOper, err := client.AP().GetGlobalOperByApName(ctx, "AP-Lab-01")
```

</details>

<details><summary><strong>Client Service</strong></summary>

| Method                 | Description                 |
| ---------------------- | --------------------------- |
| `GetOper`              | Client operational overview |
| `GetCommonOperData`    | Common operational subset   |
| `GetDot11OperData`     | 802.11 specific data        |
| `GetMobilityOperData`  | Mobility related data       |
| `GetMmIfClientStats`   | MM IF client stats          |
| `GetMmIfClientHistory` | MM IF client history        |
| `GetTrafficStats`      | Traffic statistics          |
| `GetPolicyData`        | Policy association data     |
| `GetSisfDBMac`         | SISF DB MAC entries         |
| `GetDcInfo`            | Data center info            |
| `GetOperByClientMac`   | Filter by client MAC        |
| `GetOperByApName`      | Filter by AP name           |
| `GetOperByWlanId`      | Filter by WLAN ID           |
| `GetOperByClientType`  | Filter by client type       |
| `GetOperByCoState`     | Filter by client state      |
| `GetOperByMsRadioType` | Filter by radio type        |
| `GetOperByUsername`    | Filter by username          |

</details>

<details><summary><strong>WLAN Service</strong></summary>

| Method                           | Description                                    |
| -------------------------------- | ---------------------------------------------- |
| `GetCfg`                         | WLAN configuration                             |
| `GetCfgEntries`                  | Individual WLAN entries                        |
| `GetPolicies`                    | WLAN policies                                  |
| `GetPolicyListEntries`           | Policy list entries                            |
| `GetWirelessAaaPolicyConfigs`    | AAA policy configs                             |
| `GetGlobalOper`                  | Global WLAN operational data                   |
| `GetCfgByProfileName`            | Filter configuration by WLAN profile name      |
| `GetPoliciesByPolicyProfileName` | Filter policies by policy profile name         |
| `GetGlobalOperByWlanProfile`     | Filter global operational data by WLAN profile |

</details>

<details><summary><strong>RRM Service</strong></summary>

| Method          | Description                |
| --------------- | -------------------------- |
| `GetCfg`        | RRM configuration          |
| `GetOper`       | Per-radio RRM data         |
| `GetGlobalOper` | Global RRM data            |
| `GetEmulOper`   | Emulation operational data |

</details>

<details><summary><strong>General Service</strong></summary>

| Method                       | Description                |
| ---------------------------- | -------------------------- |
| `GetOper`                    | General operational root   |
| `GetMgmtIntfData`            | Management interface data  |
| `GetCfg`                     | General configuration      |
| `GetMewlcConfig`             | ME WLC config              |
| `GetCacConfig`               | CAC config                 |
| `GetMfp`                     | MFP config/data            |
| `GetFipsCfg`                 | FIPS config                |
| `GetWsaApClientEvent`        | AP client event log        |
| `GetSimL3InterfaceCacheData` | SIM L3 interface cache     |
| `GetWlcManagementData`       | Controller management data |
| `GetLaginfo`                 | LAG info                   |
| `GetMulticastConfig`         | Multicast config           |
| `GetFeatureUsageCfg`         | Feature usage              |
| `GetThresholdWarnCfg`        | Threshold warnings         |
| `GetApLocRangingCfg`         | AP location ranging config |
| `GetGeolocationCfg`          | Geolocation config         |

</details>

<details><summary><strong>Other Services</strong></summary>

| Package         | Key Methods                                                                                                               |
| --------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `radio`         | `GetCfg`, `GetCfgByName`                                                                                                  |
| `rf`            | `GetCfg`, `GetProfiles`                                                                                                   |
| `afc`           | `GetOper`, `GetAPResp`, `GetCloudOper`, `GetCloudStats`                                                                   |
| `rogue`         | `GetOper`, `GetStats`, `GetData`, `GetClientData`, `GetRldpStats`, `GetOperByRogueAddress`, `GetOperByRogueClientAddress` |
| `mcast`         | `GetOper`, `GetFlexMediastreamClientSummary`, `GetVlanL2MgidOp`, `GetOperByClientMAC`, `GetOperByVlanIndex`               |
| `mdns`          | `GetOper`, `GetGlobalStats`, `GetWlanStats`, `GetOperByWlanID`                                                            |
| `mesh`          | `GetOper`, `GetCfg`, `GetCfgByProfileName`                                                                                |
| `mobility`      | `GetOper`, `GetOperByClientMAC`, `GetOperByAPMac`                                                                         |
| `geolocation`   | `GetOper`, `GetApGeoLocStats`                                                                                             |
| `hyperlocation` | `GetOper`, `GetProfiles`                                                                                                  |
| `location`      | `GetCfg`                                                                                                                  |
| `site`          | `GetOper`                                                                                                                 |
| `ble`           | `GetOper`, `GetOperByApMac`, `GetOperByApMacSlotAntenna`                                                                  |
| `cts`           | `GetCfg`, `GetCfgBySxpProfileName`                                                                                        |
| `dot11`         | `GetCfg`, `GetCfgByCountryCode`, `GetCfgByBand`, `GetCfgBySpatialStreamAndIndex`                                          |
| `dot15`         | `GetCfg`                                                                                                                  |
| `flex`          | `GetCfg`, `GetCfgByPolicyName`                                                                                            |
| `fabric`        | `GetCfg`, `GetCfgByFabricProfileName`, `GetCfgByControlPlaneName`                                                         |
| `apf`           | `GetCfg` (Application Policy Framework - single global configuration object)                                              |
| `awips`         | `GetOper`, `GetOperByApMac`, `GetOperByApMacDownloadStatus`                                                               |
| `rfid`          | `GetCfg`                                                                                                                  |
| `lisp`          | `GetOper`                                                                                                                 |
| `nmsp`          | `GetOper`, `GetClientRegistration`, `GetCmxConnection`, `GetCmxCloudInfo`, `GetOperByClientID`                            |

</details>

## 📦 Response & Error Handling

Each `Get*` returns a JSON‑unmarshalled `*model.<Type>Response` and an error; list endpoints return a pointer to a slice.

| Sentinel                  | Meaning                           |
| ------------------------- | --------------------------------- |
| `ErrAuthenticationFailed` | 401 credentials/token invalid     |
| `ErrAccessForbidden`      | 403 authorization failure         |
| `ErrResourceNotFound`     | 404 on RESTCONF path              |
| `ErrInvalidConfiguration` | Client misuse during construction |
| `ErrRequestTimeout`       | Request exceeded timeout          |

Use `errors.Is(err, wnc.ErrAuthenticationFailed)` or `errors.As(err, *wnc.APIError)`.

## 🔁 Method Pattern

All domain methods internally call a generic helper equivalent to:

```go
func (s Service) GetOper(ctx context.Context) (*model.DomainOperResponse, error) {
    return core.Get[model.DomainOperResponse](ctx, s.c, endpointConst)
}
```

No retries, caching, pagination or filtering are currently implemented.

## ❓ FAQ

| Question                      | Answer                                                                          |
| ----------------------------- | ------------------------------------------------------------------------------- |
| Why only GET?                 | Focus on read use cases first; write ops require model validation & RPC design. |
| Why so many granular methods? | Mirrors YANG subtree granularity for explicit intent & smaller payloads.        |
| Are retries built-in?         | No. Implement externally to keep core deterministic.                            |
