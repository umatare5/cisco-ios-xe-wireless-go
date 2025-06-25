# 🌐 API Reference

This document provides an overview of the API functions available in the WNC Go client library.

- **Functional Organization**: The API is organized by functional areas such as Access Points, Clients and others.
- **Context-Aware Requests**: All API functions require a `context.Context` as the first parameter for proper request lifecycle management.
- **Structured Responses and Errors**: Each API returns structured data with comprehensive error handling based on original responses.

To know the meaning of the response from APIs, please refer to the [YANG Models and Platform Capabilities for Cisco IOS XE 17.12.1](https://github.com/YangModels/yang/tree/main/vendor/cisco/xe/17121#readme).

> [!Note]
> Currently, all APIs do not support filtering with keys. This feature will be implemented in the next release `v0.2.0`.

## Core Functions

| Function         | Parameters                                                 | Return Type        | Description                             |
| ---------------- | ---------------------------------------------------------- | ------------------ | --------------------------------------- |
| `NewClient`      | `controller, accessToken string, options ...ClientOption`  | `(*Client, error)` | Creates a new WNC client with options   |
| `SendAPIRequest` | `ctx context.Context, endpoint string, result interface{}` | `error`            | Sends API request to specified endpoint |

> [!Note]
> Currently, this library supports only `GET` requests. To send other type of requests, please use `SendAPIRequest()`.

## 📡 Access Point Operations

| Function                       | Return Type                                       | Description                                         |
| ------------------------------ | ------------------------------------------------- | --------------------------------------------------- |
| `GetApOper`                    | `(*ApOperResponse, error)`                        | Retrieves operational data for all access points    |
| `GetApRadioNeighbor`           | `(*ApOperApRadioNeighborResponse, error)`         | Gets radio neighbor information                     |
| `GetApRadioOperData`           | `(*ApOperRadioOperDataResponse, error)`           | Gets radio operational statistics and configuration |
| `GetApRadioResetStats`         | `(*ApOperRadioResetStatsResponse, error)`         | Gets radio reset statistics                         |
| `GetApQosClientData`           | `(*ApOperQosClientDataResponse, error)`           | Gets QoS client data information                    |
| `GetApCapwapData`              | `(*ApOperCapwapDataResponse, error)`              | Fetches CAPWAP tunnel information                   |
| `GetApNameMacMap`              | `(*ApOperApNameMacMapResponse, error)`            | Gets AP name to MAC address mapping                 |
| `GetApWtpSlotWlanStats`        | `(*ApOperWtpSlotWlanStatsResponse, error)`        | Gets WTP slot WLAN statistics                       |
| `GetApEthernetMacWtpMacMap`    | `(*ApOperEthernetMacWtpMacMapResponse, error)`    | Gets Ethernet MAC to WTP MAC mapping                |
| `GetApRadioOperStats`          | `(*ApOperRadioOperStatsResponse, error)`          | Gets radio operational statistics                   |
| `GetApEthernetIfStats`         | `(*ApOperEthernetIfStatsResponse, error)`         | Gets Ethernet interface statistics                  |
| `GetApEwlcWncdStats`           | `(*ApOperEwlcWncdStatsResponse, error)`           | Gets EWLC WNCD statistics                           |
| `GetApIoxOperData`             | `(*ApOperApIoxOperDataResponse, error)`           | Gets IOX operational data                           |
| `GetApQosGlobalStats`          | `(*ApOperQosGlobalStatsResponse, error)`          | Gets global QoS statistics                          |
| `GetApOperData`                | `(*ApOperOperDataResponse, error)`                | Gets AP operational data                            |
| `GetApRlanOper`                | `(*ApOperRlanOperResponse, error)`                | Gets RLAN operational data                          |
| `GetApEwlcMewlcPredownloadRec` | `(*ApOperEwlcMewlcPredownloadRecResponse, error)` | Gets EWLC MEWLC predownload records                 |
| `GetApCdpCacheData`            | `(*ApOperCdpCacheDataResponse, error)`            | Gets CDP cache data                                 |
| `GetApLldpNeigh`               | `(*ApOperLldpNeighResponse, error)`               | Retrieves LLDP neighbor discovery data              |
| `GetApTpCertInfo`              | `(*ApOperTpCertInfoResponse, error)`              | Gets TP certificate information                     |
| `GetApDiscData`                | `(*ApOperDiscDataResponse, error)`                | Gets discovery data                                 |
| `GetApCapwapPkts`              | `(*ApOperCapwapPktsResponse, error)`              | Gets CAPWAP packet statistics                       |
| `GetApCountryOper`             | `(*ApOperCountryOperResponse, error)`             | Gets country operational data                       |
| `GetApSuppCountryOper`         | `(*ApOperSuppCountryOperResponse, error)`         | Gets supported country operational data             |
| `GetApNhGlobalData`            | `(*ApOperApNhGlobalDataResponse, error)`          | Gets AP NH global data                              |
| `GetApImagePrepareLocation`    | `(*ApOperApImagePrepareLocationResponse, error)`  | Gets AP image prepare location                      |
| `GetApImageActiveLocation`     | `(*ApOperApImageActiveLocationResponse, error)`   | Gets AP image active location                       |

## 🌐 Access Point Global Operations

| Function                   | Return Type                                           | Description                          |
| -------------------------- | ----------------------------------------------------- | ------------------------------------ |
| `GetApGlobalOper`          | `(*ApGlobalOperResponse, error)`                      | Gets global AP operational data      |
| `GetApHistory`             | `(*ApGlobalOperApHistoryResponse, error)`             | Gets AP history data                 |
| `GetApEwlcApStats`         | `(*ApGlobalOperEwlcApStatsResponse, error)`           | Gets EWLC AP statistics              |
| `GetApImgPredownloadStats` | `(*ApGlobalOperApImgPredownloadStatsResponse, error)` | Gets AP image predownload statistics |
| `GetApJoinStats`           | `(*ApGlobalOperApJoinStatsResponse, error)`           | Gets AP join statistics              |
| `GetApWlanClientStats`     | `(*ApGlobalOperWlanClientStatsResponse, error)`       | Gets WLAN client statistics          |
| `GetApEmltdJoinCountStat`  | `(*ApGlobalOperEmltdJoinCountStatResponse, error)`    | Gets EMLATED join count statistics   |

## ⚙️ Access Point Configuration

| Function                        | Return Type                                       | Description                                |
| ------------------------------- | ------------------------------------------------- | ------------------------------------------ |
| `GetApCfg`                      | `(*ApCfgResponse, error)`                         | Gets AP configuration                      |
| `GetTagSourcePriorityConfigs`   | `(*TagSourcePriorityConfigs, error)`              | Gets tag source priority configurations    |
| `GetApTagSourcePriorityConfigs` | `(*ApCfgTagSourcePriorityConfigsResponse, error)` | Gets AP tag source priority configurations |
| `GetApApTags`                   | `(*ApCfgApTagsResponse, error)`                   | Gets AP tags                               |

## 👥 Client Operations

| Function                         | Return Type                                     | Description                                 |
| -------------------------------- | ----------------------------------------------- | ------------------------------------------- |
| `GetClientOper`                  | `(*ClientOperResponse, error)`                  | Gets operational data for connected clients |
| `GetClientOperCommonOperData`    | `(*ClientOperCommonOperDataResponse, error)`    | Gets common client operational data         |
| `GetClientOperDot11OperData`     | `(*ClientOperDot11OperDataResponse, error)`     | Gets 802.11 client operational data         |
| `GetClientOperMobilityOperData`  | `(*ClientOperMobilityOperDataResponse, error)`  | Gets mobility operational data              |
| `GetClientOperMmIfClientStats`   | `(*ClientOperMmIfClientStatsResponse, error)`   | Gets MM interface client statistics         |
| `GetClientOperMmIfClientHistory` | `(*ClientOperMmIfClientHistoryResponse, error)` | Gets MM interface client history            |
| `GetClientOperTrafficStats`      | `(*ClientOperTrafficStatsResponse, error)`      | Gets client traffic statistics              |
| `GetClientOperPolicyData`        | `(*ClientOperPolicyDataResponse, error)`        | Gets client policy data                     |
| `GetClientOperSisfDbMac`         | `(*ClientOperSisfDbMacResponse, error)`         | Gets SISF database MAC data                 |
| `GetClientOperDcInfo`            | `(*ClientOperDcInfoResponse, error)`            | Gets DC information                         |

## 👥 Client Global Operations

| Function                    | Return Type                               | Description                                     |
| --------------------------- | ----------------------------------------- | ----------------------------------------------- |
| `GetClientGlobalOper`       | `(*ClientGlobalOperResponse, error)`      | Retrieves global client statistics and counters |
| `GetClientLiveStats`        | `(*ClientLiveStatsResponse, error)`       | Gets real-time client statistics                |
| `GetClientGlobalStatsData`  | `(*ClientGlobalStatsDataResponse, error)` | Retrieves global client statistics data         |
| `GetClientStats`            | `(*ClientStatsResponse, error)`           | Gets client statistics                          |
| `GetClientDot11Stats`       | `(*ClientDot11StatsResponse, error)`      | Gets 802.11 client statistics                   |
| `GetClientLatencyStats`     | `(*ClientLatencyStatsResponse, error)`    | Gets client latency statistics                  |
| `GetClientSmWebauthStats`   | `(*SmWebauthStatsResponse, error)`        | Gets SM web authentication statistics           |
| `GetClientDot1XGlobalStats` | `(*Dot1XGlobalStatsResponse, error)`      | Gets 802.1X global statistics                   |
| `GetClientExclusionStats`   | `(*ClientExclusionStatsResponse, error)`  | Gets client exclusion statistics                |
| `GetClientSmDeviceCount`    | `(*SmDeviceCountResponse, error)`         | Gets SM device count                            |
| `GetClientTofStats`         | `(*TofStatsResponse, error)`              | Gets time-of-flight statistics                  |

## 🌐 WLAN Configuration & Operations

| Function                      | Return Type                                  | Description                             |
| ----------------------------- | -------------------------------------------- | --------------------------------------- |
| `GetWlanCfg`                  | `(*WlanCfgResponse, error)`                  | Fetches WLAN configuration settings     |
| `GetWlanCfgEntries`           | `(*WlanCfgEntriesResponse, error)`           | Gets WLAN configuration entries         |
| `GetWlanPolicies`             | `(*WlanPoliciesResponse, error)`             | Gets WLAN policies                      |
| `GetPolicyListEntries`        | `(*PolicyListEntriesResponse, error)`        | Gets policy list entries                |
| `GetWirelessAaaPolicyConfigs` | `(*WirelessAaaPolicyConfigsResponse, error)` | Gets wireless AAA policy configurations |
| `GetWlanGlobalOper`           | `(*WlanGlobalOperResponse, error)`           | Gets global WLAN operational statistics |
| `GetWlanGlobalOperWlanInfo`   | `(*WlanGlobalOperWlanInfoResponse, error)`   | Gets global WLAN information            |

## 📡 Radio Resource Management (RRM)

| Function                 | Return Type                             | Description                              |
| ------------------------ | --------------------------------------- | ---------------------------------------- |
| `GetRrmOper`             | `(*RrmOperResponse, error)`             | Retrieves RRM operational data per radio |
| `GetApAutoRfDot11Data`   | `(*ApAutoRfDot11DataResponse, error)`   | Gets auto RF 802.11 data                 |
| `GetApDot11RadarData`    | `(*ApDot11RadarDataResponse, error)`    | Gets 802.11 radar data                   |
| `GetApDot11SpectrumData` | `(*ApDot11SpectrumDataResponse, error)` | Gets 802.11 spectrum data                |
| `GetRrmMeasurement`      | `(*RrmMeasurementResponse, error)`      | Gets RRM measurement data                |
| `GetRadioSlot`           | `(*RadioSlotResponse, error)`           | Gets radio slot data                     |
| `GetMainData`            | `(*MainDataResponse, error)`            | Gets main data                           |
| `GetSpectrumDeviceTable` | `(*SpectrumDeviceTableResponse, error)` | Gets spectrum device table               |
| `GetSpectrumAqTable`     | `(*SpectrumAqTableResponse, error)`     | Gets spectrum AQ table                   |
| `GetRegDomainOper`       | `(*RegDomainOperResponse, error)`       | Gets regulatory domain operational data  |

## 📡 RRM Global Operations

| Function                             | Return Type                                | Description                                  |
| ------------------------------------ | ------------------------------------------ | -------------------------------------------- |
| `GetRrmGlobalOper`                   | `(*RrmGlobalOperResponse, error)`          | Gets global RRM statistics and configuration |
| `GetRrmGlobalOneShotCounters`        | `(*RrmOneShotCountersResponse, error)`     | Gets RRM one-shot counters                   |
| `GetRrmGlobalChannelParams`          | `(*RrmChannelParamsResponse, error)`       | Gets RRM channel parameters                  |
| `GetRrmGlobalSpectrumAqWorstTable`   | `(*SpectrumAqWorstTableResponse, error)`   | Gets spectrum AQ worst table                 |
| `GetRrmGlobalRadioOperData24G`       | `(*RadioOperData24GResponse, error)`       | Gets 2.4GHz radio operational data           |
| `GetRrmGlobalRadioOperData5G`        | `(*RadioOperData5GResponse, error)`        | Gets 5GHz radio operational data             |
| `GetRrmGlobalRadioOperData6G`        | `(*RadioOperData5GResponse, error)`        | Gets 6GHz radio operational data             |
| `GetRrmGlobalSpectrumBandConfigData` | `(*SpectrumBandConfigDataResponse, error)` | Gets spectrum band configuration data        |
| `GetRrmGlobalRadioOperDataDualband`  | `(*RadioOperDataDualbandResponse, error)`  | Gets dual-band radio operational data        |
| `GetRrmGlobalClientData`             | `(*RrmClientDataResponse, error)`          | Gets RRM client data                         |
| `GetRrmGlobalFraStats`               | `(*RrmFraStatsResponse, error)`            | Gets RRM FRA statistics                      |
| `GetRrmGlobalCoverage`               | `(*RrmCoverageResponse, error)`            | Gets RRM coverage data                       |

## ⚙️ RRM Configuration

| Function              | Return Type                          | Description                            |
| --------------------- | ------------------------------------ | -------------------------------------- |
| `GetRrmCfg`           | `(*RrmCfgResponse, error)`           | Gets RRM configuration                 |
| `GetRrmRrms`          | `(*RrmRrmsResponse, error)`          | Gets RRM entries                       |
| `GetRrmMgrCfgEntries` | `(*RrmMgrCfgEntriesResponse, error)` | Gets RRM manager configuration entries |

## 🔧 RRM Emulation Operations

| Function                | Return Type                                | Description                         |
| ----------------------- | ------------------------------------------ | ----------------------------------- |
| `GetRrmEmulOper`        | `(*RrmEmulOperResponse, error)`            | Gets RRM emulation operational data |
| `GetRrmEmulRrmFraStats` | `(*RrmEmulOperRrmFraStatsResponse, error)` | Gets RRM emulation FRA statistics   |

## 📻 Radio & RF Configuration

| Function                     | Return Type                                 | Description                     |
| ---------------------------- | ------------------------------------------- | ------------------------------- |
| `GetRadioCfg`                | `(*RadioCfgResponse, error)`                | Gets radio configuration        |
| `GetRadioProfiles`           | `(*RadioProfilesResponse, error)`           | Gets radio profiles             |
| `GetRfCfg`                   | `(*RfCfgResponse, error)`                   | Gets RF configuration           |
| `GetRfMultiBssidProfiles`    | `(*MultiBssidProfilesResponse, error)`      | Gets multi-BSSID profiles       |
| `GetRfAtfPolicies`           | `(*AtfPoliciesResponse, error)`             | Gets ATF policies               |
| `GetRfTags`                  | `(*RfTagsResponse, error)`                  | Gets RF tags                    |
| `GetRfProfiles`              | `(*RfProfilesResponse, error)`              | Gets RF profiles                |
| `GetRfProfileDefaultEntries` | `(*RfProfileDefaultEntriesResponse, error)` | Gets RF profile default entries |

## 🔧 General Configuration & Management

| Function                            | Return Type                                 | Description                              |
| ----------------------------------- | ------------------------------------------- | ---------------------------------------- |
| `GetGeneralCfg`                     | `(*GeneralCfgResponse, error)`              | Fetches general controller configuration |
| `GetGeneralOper`                    | `(*GeneralOperResponse, error)`             | Gets general operational data            |
| `GetGeneralOperMgmtIntfData`        | `(*GeneralOperMgmtIntfDataResponse, error)` | Gets WLC management interface data       |
| `GetGeneralMewlcConfig`             | `(*MewlcConfigResponse, error)`             | Gets MEWLC configuration                 |
| `GetGeneralCacConfig`               | `(*CacConfigResponse, error)`               | Gets CAC configuration                   |
| `GetGeneralMfp`                     | `(*MfpResponse, error)`                     | Gets MFP configuration                   |
| `GetGeneralFipsCfg`                 | `(*FipsCfgResponse, error)`                 | Gets FIPS configuration                  |
| `GetGeneralWsaApClientEvent`        | `(*WsaApClientEventResponse, error)`        | Gets WSA AP client event data            |
| `GetGeneralSimL3InterfaceCacheData` | `(*SimL3InterfaceCacheDataResponse, error)` | Gets SIM L3 interface cache data         |
| `GetGeneralWlcManagementData`       | `(*WlcManagementDataResponse, error)`       | Gets WLC management data                 |
| `GetGeneralLaginfo`                 | `(*LaginfoResponse, error)`                 | Gets LAG information                     |
| `GetGeneralMulticastConfig`         | `(*MulticastConfigResponse, error)`         | Gets multicast configuration             |
| `GetGeneralFeatureUsageCfg`         | `(*FeatureUsageCfgResponse, error)`         | Gets feature usage configuration         |
| `GetGeneralThresholdWarnCfg`        | `(*ThresholdWarnCfgResponse, error)`        | Gets threshold warning configuration     |
| `GetGeneralApLocRangingCfg`         | `(*ApLocRangingCfgResponse, error)`         | Gets AP location ranging configuration   |
| `GetGeneralGeolocationCfg`          | `(*GeolocationCfgResponse, error)`          | Gets geolocation configuration           |

## 🏢 Site Configuration

| Function               | Return Type                           | Description                         |
| ---------------------- | ------------------------------------- | ----------------------------------- |
| `GetSiteCfg`           | `(*SiteCfgResponse, error)`           | Gets site configuration             |
| `GetSiteApCfgProfiles` | `(*SiteApCfgProfilesResponse, error)` | Gets site AP configuration profiles |
| `GetSiteTagConfigs`    | `(*SiteTagConfigsResponse, error)`    | Gets site tag configurations        |

## 🌍 Mobility Operations

| Function                        | Return Type                                 | Description                             |
| ------------------------------- | ------------------------------------------- | --------------------------------------- |
| `GetMobilityOper`               | `(*MobilityOperResponse, error)`            | Gets mobility operational data          |
| `GetMobilityMmIfGlobalStats`    | `(*MmIfGlobalStatsResponse, error)`         | Gets MM interface global statistics     |
| `GetMobilityMmIfGlobalMsgStats` | `(*MmIfGlobalMsgStatsResponse, error)`      | Gets MM interface global message stats  |
| `GetMobilityGlobalStats`        | `(*MobilityGlobalStatsResponse, error)`     | Gets mobility global statistics         |
| `GetMobilityMmGlobalData`       | `(*MmGlobalDataResponse, error)`            | Gets MM global data                     |
| `GetMobilityGlobalMsgStats`     | `(*MobilityGlobalMsgStatsResponse, error)`  | Gets mobility global message statistics |
| `GetMobilityClientData`         | `(*MobilityClientDataResponse, error)`      | Gets mobility client data               |
| `GetMobilityApCache`            | `(*ApCacheResponse, error)`                 | Gets mobility AP cache                  |
| `GetMobilityApPeerList`         | `(*ApPeerListResponse, error)`              | Gets mobility AP peer list              |
| `GetMobilityClientStats`        | `(*MobilityClientStatsResponse, error)`     | Gets mobility client statistics         |
| `GetMobilityWlanClientLimit`    | `(*WlanClientLimitResponse, error)`         | Gets WLAN client limit data             |
| `GetMobilityGlobalDTLSStats`    | `(*MobilityGlobalDTLSStatsResponse, error)` | Gets mobility global DTLS statistics    |

## 🗺️ Location & Geolocation

| Function                          | Return Type                                      | Description                         |
| --------------------------------- | ------------------------------------------------ | ----------------------------------- |
| `GetLocationCfg`                  | `(*LocationCfgResponse, error)`                  | Gets location configuration         |
| `GetLocationNmspConfig`           | `(*LocationNmspConfigResponse, error)`           | Gets location NMSP configuration    |
| `GetGeolocationOper`              | `(*GeolocationOperResponse, error)`              | Gets geolocation operational data   |
| `GetGeolocationOperApGeoLocStats` | `(*GeolocationOperApGeoLocStatsResponse, error)` | Gets AP geolocation statistics      |
| `GetHyperlocationOper`            | `(*HyperlocationOperResponse, error)`            | Gets hyperlocation operational data |
| `GetHyperlocationProfiles`        | `(*HyperlocationProfilesResponse, error)`        | Gets hyperlocation profiles         |

## 🕸️ Mesh Operations & Configuration

| Function             | Return Type                         | Description                       |
| -------------------- | ----------------------------------- | --------------------------------- |
| `GetMeshCfg`         | `(*MeshCfgResponse, error)`         | Gets mesh configuration           |
| `GetMesh`            | `(*MeshResponse, error)`            | Gets mesh data                    |
| `GetMeshProfiles`    | `(*MeshProfilesResponse, error)`    | Gets mesh profiles                |
| `GetMeshGlobalOper`  | `(*MeshGlobalOperResponse, error)`  | Gets mesh global operational data |
| `GetMeshGlobalStats` | `(*MeshGlobalStatsResponse, error)` | Gets mesh global statistics       |
| `GetMeshApTreeData`  | `(*MeshApTreeDataResponse, error)`  | Gets mesh AP tree data            |

## 🔍 Network Management & Monitoring

| Function                    | Return Type                                | Description                        |
| --------------------------- | ------------------------------------------ | ---------------------------------- |
| `GetNmspOper`               | `(*NmspOperResponse, error)`               | Gets NMSP operational data         |
| `GetNmspClientRegistration` | `(*NmspClientRegistrationResponse, error)` | Gets NMSP client registration data |
| `GetNmspCmxConnection`      | `(*NmspCmxConnectionResponse, error)`      | Gets NMSP CMX connection data      |
| `GetNmspCmxCloudInfo`       | `(*NmspCmxCloudInfoResponse, error)`       | Gets NMSP CMX cloud information    |
| `GetMdnsOper`               | `(*MdnsOperResponse, error)`               | Gets mDNS operational data         |
| `GetMdnsGlobalStats`        | `(*MdnsGlobalStatsResponse, error)`        | Gets mDNS global statistics        |
| `GetMdnsWlanStats`          | `(*MdnsWlanStatsResponse, error)`          | Gets mDNS WLAN statistics          |

## 🚨 Rogue Detection & Security

| Function             | Return Type                         | Description                 |
| -------------------- | ----------------------------------- | --------------------------- |
| `GetRogueOper`       | `(*RogueOperResponse, error)`       | Gets rogue operational data |
| `GetRogueStats`      | `(*RogueStatsResponse, error)`      | Gets rogue statistics       |
| `GetRogueData`       | `(*RogueDataResponse, error)`       | Gets rogue data             |
| `GetRogueClientData` | `(*RogueClientDataResponse, error)` | Gets rogue client data      |
| `GetRldpStats`       | `(*RldpStatsResponse, error)`       | Gets RLDP statistics        |

## 🌐 Multicast Operations

| Function                               | Return Type                                               | Description                          |
| -------------------------------------- | --------------------------------------------------------- | ------------------------------------ |
| `GetMcastOper`                         | `(*McastOperResponse, error)`                             | Gets multicast operational data      |
| `GetMcastFlexMediastreamClientSummary` | `(*McastOperFlexMediastreamClientSummaryResponse, error)` | Gets flex mediastream client summary |
| `GetMcastVlanL2MgidOp`                 | `(*McastOperVlanL2MgidOpResponse, error)`                 | Gets VLAN L2 MGID operational data   |

## 🔗 LISP Agent Operations

| Function                  | Return Type                              | Description                       |
| ------------------------- | ---------------------------------------- | --------------------------------- |
| `GetLispAgentOper`        | `(*LispAgentOperResponse, error)`        | Gets LISP agent operational data  |
| `GetLispAgentMemoryStats` | `(*LispAgentMemoryStatsResponse, error)` | Gets LISP agent memory statistics |
| `GetLispWlcCapabilities`  | `(*LispWlcCapabilitiesResponse, error)`  | Gets LISP WLC capabilities        |
| `GetLispApCapabilities`   | `(*LispApCapabilitiesResponse, error)`   | Gets LISP AP capabilities         |

## 🔐 Security & Policy

| Function                 | Return Type                             | Description                        |
| ------------------------ | --------------------------------------- | ---------------------------------- |
| `GetCtsSxpCfg`           | `(*CtsSxpCfgResponse, error)`           | Gets CTS SXP configuration         |
| `GetCtsSxpConfiguration` | `(*CtsSxpConfigurationResponse, error)` | Gets CTS SXP configuration details |

## 📊 AFC (Automated Frequency Coordination)

| Function              | Return Type                                   | Description                     |
| --------------------- | --------------------------------------------- | ------------------------------- |
| `GetAfcOper`          | `(*AfcOperResponse, error)`                   | Gets AFC operational data       |
| `GetAfcEwlcAfcApResp` | `(*AfcOperEwlcAfcApRespResponse, error)`      | Gets AFC EWLC AP response data  |
| `GetAfcCloudOper`     | `(*AfcCloudOperResponse, error)`              | Gets AFC cloud operational data |
| `GetAfcCloudStats`    | `(*AfcCloudOperAfcCloudStatsResponse, error)` | Gets AFC cloud statistics       |

## 📡 Wireless Intrusion Prevention (AWIPS)

| Function               | Return Type                               | Description                   |
| ---------------------- | ----------------------------------------- | ----------------------------- |
| `GetAwipsOper`         | `(*AwipsOperResponse, error)`             | Gets AWIPS operational data   |
| `GetAwipsPerApInfo`    | `(*AwipsOperPerApInfoResponse, error)`    | Gets AWIPS per-AP information |
| `GetAwipsDwldStatus`   | `(*AwipsOperDwldStatusResponse, error)`   | Gets AWIPS download status    |
| `GetAwipsApDwldStatus` | `(*AwipsOperApDwldStatusResponse, error)` | Gets AWIPS AP download status |

## 📶 Bluetooth Low Energy & Location

| Function             | Return Type                         | Description                   |
| -------------------- | ----------------------------------- | ----------------------------- |
| `GetBleLtxOper`      | `(*BleLtxOperResponse, error)`      | Gets BLE LTX operational data |
| `GetBleLtxApAntenna` | `(*BleLtxApAntennaResponse, error)` | Gets BLE LTX AP antenna data  |
| `GetBleLtxAp`        | `(*BleLtxApResponse, error)`        | Gets BLE LTX AP data          |

## 🏷️ Advanced Configuration

| Function                      | Return Type                                  | Description                         |
| ----------------------------- | -------------------------------------------- | ----------------------------------- |
| `GetApfCfg`                   | `(*ApfCfgResponse, error)`                   | Gets APF configuration              |
| `GetApf`                      | `(*ApfCfgApfResponse, error)`                | Gets APF data                       |
| `GetDot11Cfg`                 | `(*Dot11CfgResponse, error)`                 | Gets 802.11 configuration           |
| `GetDot11ConfiguredCountries` | `(*Dot11ConfiguredCountriesResponse, error)` | Gets 802.11 configured countries    |
| `GetDot11acMcsEntries`        | `(*Dot11acMcsEntriesResponse, error)`        | Gets 802.11ac MCS entries           |
| `GetDot11Entries`             | `(*Dot11EntriesResponse, error)`             | Gets 802.11 entries                 |
| `GetDot15Cfg`                 | `(*Dot15CfgResponse, error)`                 | Gets 802.15 configuration           |
| `GetDot15GlobalConfig`        | `(*Dot15GlobalConfigResponse, error)`        | Gets 802.15 global configuration    |
| `GetFlexCfg`                  | `(*FlexCfgResponse, error)`                  | Gets FlexConnect configuration      |
| `GetFlexCfgData`              | `(*FlexCfgDataResponse, error)`              | Gets FlexConnect configuration data |
| `GetFabricCfg`                | `(*FabricCfgResponse, error)`                | Gets fabric configuration           |
| `GetFabricControlplaneNames`  | `(*FabricControlplaneNamesResponse, error)`  | Gets fabric control plane names     |
| `GetFabric`                   | `(*FabricResponse, error)`                   | Gets fabric data                    |
| `GetRfidCfg`                  | `(*RfidCfgResponse, error)`                  | Gets RFID configuration             |

## 🚨 Error Types

Understanding error types helps you implement proper error handling and debugging in your applications.

| Error Type            | Constant                  | Description                           | HTTP Status |
| --------------------- | ------------------------- | ------------------------------------- | ----------- |
| `APIError`            | -                         | API-specific errors with status codes | Various     |
| `AuthenticationError` | `ErrAuthenticationFailed` | Invalid credentials or token          | 401         |
| `AccessError`         | `ErrAccessForbidden`      | Insufficient permissions              | 403         |
| `NotFoundError`       | `ErrResourceNotFound`     | Requested resource not found          | 404         |
| `ConfigError`         | `ErrInvalidConfiguration` | Invalid client configuration          | -           |
| `TimeoutError`        | `ErrRequestTimeout`       | Request timeout exceeded              | -           |

> [!TIP]
> Use type assertions or `errors.As()` to handle specific error types and provide appropriate user feedback or retry logic.
