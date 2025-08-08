# 🌐 API Reference

High-level reference for the exported, user‑facing API of this Go library. Domain data access uses a consistent pattern: construct a unified client, pick a domain service accessor, call a `Get*` method with a `context.Context`.

## ⏰️ Quick Reference

| Aspect       | Rule                                                                                      |
| ------------ | ----------------------------------------------------------------------------------------- |
| Import       | `import wnc "github.com/umatare5/cisco-ios-xe-wireless-go"`                               |
| Construction | `client, err := wnc.NewClient(host, token, wnc.WithTimeout(30*time.Second))`              |
| Pattern      | `resp, err := client.AP().GetOper(ctx)`                                                   |
| Errors       | Wrap with `errors.Is / errors.As` using exported sentinels (`ErrAuthenticationFailed`, …) |
| HTTP Methods | Only `GET` is implemented (read‑only)                                                     |
| Filtering    | Not yet supported (future enhancement)                                                    |

> [!NOTE]
> All data shapes map directly to Cisco YANG models for IOS-XE 17.12.1. See official [YANG Models](https://github.com/YangModels/yang/tree/main/vendor/cisco/xe/17121#readme) for field semantics.

## 🧭 Domain Services

Each accessor on `*wnc.Client` returns a lightweight value receiver service. Services have no internal mutable state; re-use freely.

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

## 📦 Response & Error Handling

Every `Get*` method returns `(*model.<Type>Response, error)`, or pointer to slice for list wrappers, with JSON unmarshalled data. Errors:

| Sentinel                  | Meaning                       |
| ------------------------- | ----------------------------- |
| `ErrAuthenticationFailed` | 401 credentials/token invalid |
| `ErrAccessForbidden`      | 403 authorization failure     |
| `ErrResourceNotFound`     | 404 on RESTCONF path          |
| `ErrInvalidConfiguration` | Client misuse (construction)  |
| `ErrRequestTimeout`       | Request exceeded timeout      |

Use `errors.Is(err, wnc.ErrAuthenticationFailed)` or `errors.As(err, *wnc.APIError)`.

## 🔁 Method Pattern

All domain methods internally call a generic helper equivalent to:

```go
func (s Service) GetOper(ctx context.Context) (*model.DomainOperResponse, error) {
    return core.Get[model.DomainOperResponse](ctx, s.c, endpointConst)
}
```

No retries, caching, pagination or filtering are currently implemented.

## 📚 Domain Method Index

Low-priority exhaustive lists are collapsed below. Expand only what you need.

<details>
<summary><strong>Access Point (ap)</strong></summary>

| Method                        | Description                      |
| ----------------------------- | -------------------------------- |
| `GetCfg`                      | Access point configuration (all) |
| `GetTagSourcePriorityConfigs` | Tag source priority config       |
| `GetApTags`                   | AP tag assignments               |
| `GetOper`                     | AP operational root data         |
| `GetRadioNeighbor`            | Radio neighbor info              |
| `GetNameMacMap`               | AP name ↔ MAC list               |
| `GetCapwapData`               | CAPWAP session data              |
| `GetGlobalOper`               | Global AP operational data       |
| `GetHistory`                  | AP history records               |
| `GetEwlcApStats`              | EWLC AP statistics               |

</details>

<details>
<summary><strong>Client (client)</strong></summary>

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

</details>

<details>
<summary><strong>WLAN (wlan)</strong></summary>

| Method                        | Description                  |
| ----------------------------- | ---------------------------- |
| `GetCfg`                      | WLAN configuration           |
| `GetCfgEntries`               | Individual WLAN entries      |
| `GetPolicies`                 | WLAN policies                |
| `GetPolicyListEntries`        | Policy list entries          |
| `GetWirelessAaaPolicyConfigs` | AAA policy configs           |
| `GetGlobalOper`               | Global WLAN operational data |

</details>

<details>
<summary><strong>RRM (rrm)</strong></summary>

| Method          | Description                |
| --------------- | -------------------------- |
| `GetCfg`        | RRM configuration          |
| `GetOper`       | Per-radio RRM data         |
| `GetGlobalOper` | Global RRM data            |
| `GetEmulOper`   | Emulation operational data |

</details>

<details>
<summary><strong>General (general)</strong></summary>

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

<details>
<summary><strong>Other Domains (Grouped)</strong></summary>

| Package         | Key Methods                                                               |
| --------------- | ------------------------------------------------------------------------- |
| `radio`         | `GetCfg`                                                                  |
| `rf`            | `GetCfg`, `GetProfiles`                                                   |
| `afc`           | `GetOper`, `GetAPResp`, `GetCloudOper`, `GetCloudStats`                   |
| `rogue`         | `GetOper`, `GetStats`, `GetData`, `GetClientData`, `GetRldpStats`         |
| `mcast`         | `GetOper`, `GetFlexMediastreamClientSummary`, `GetVlanL2MgidOp`           |
| `mdns`          | `GetOper`, `GetGlobalStats`, `GetWlanStats`                               |
| `mesh`          | `GetOper`, `GetCfg`                                                       |
| `mobility`      | `GetOper`                                                                 |
| `geolocation`   | `GetOper`, `GetApGeoLocStats`                                             |
| `hyperlocation` | `GetOper`, `GetProfiles`                                                  |
| `location`      | `GetCfg`                                                                  |
| `site`          | `GetOper`                                                                 |
| `ble`           | `GetOper`                                                                 |
| `cts`           | `GetCfg`                                                                  |
| `dot11`         | `GetCfg`                                                                  |
| `dot15`         | `GetCfg`                                                                  |
| `flex`          | `GetCfg`                                                                  |
| `fabric`        | `GetCfg`                                                                  |
| `apf`           | `GetCfg`                                                                  |
| `awips`         | `GetOper`                                                                 |
| `rfid`          | `GetCfg`                                                                  |
| `lisp`          | `GetOper`                                                                 |
| `nmsp`          | `GetOper`, `GetClientRegistration`, `GetCmxConnection`, `GetCmxCloudInfo` |

</details>

## 🧪 Stability & Versioning

Current library version: `v0.2.0` (see `VERSION`). Breaking doc changes may occur until `v1.0.0`.

## ❓ FAQ

| Question                      | Answer                                                                          |
| ----------------------------- | ------------------------------------------------------------------------------- |
| Why only GET?                 | Focus on read use cases first; write ops require model validation & RPC design. |
| Why so many granular methods? | Mirrors YANG subtree granularity for explicit intent & smaller payloads.        |
| Are retries built-in?         | No. Implement externally to keep core deterministic.                            |
