// Package general provides general controller configuration and operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with general controller settings
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for system-wide configuration,
// operational data monitoring, and global controller management settings.
//
// The package leverages the YANG models: Cisco-IOS-XE-wireless-general-cfg and Cisco-IOS-XE-wireless-general-oper for complete general controller management.
//
// # Main Functions
//
// The General service provides two main categories of operations:
//
// 1. GetCfg - Configuration Data
// 2. GetOper - Operational Data
//
// # Additional Functions
//
// The General service also provides specialized data access:
//
//   - GetMgmtIntfData - Management interface operational data
//   - GetMewlcConfig - MEWLC configuration data
//   - GetCacConfig - CAC configuration data
//   - GetMfp - MFP (Management Frame Protection) configuration data
//   - GetFipsCfg - FIPS configuration data
//   - GetWsaApClientEvent - WSA AP client event configuration data
//   - GetSimL3InterfaceCacheData - SIM L3 interface cache data
//   - GetWlcManagementData - WLC management data
//   - GetLaginfo - LAG (Link Aggregation) information
//   - GetMulticastConfig - Multicast configuration data
//   - GetFeatureUsageCfg - Feature usage configuration data
//   - GetThresholdWarnCfg - Threshold warning configuration data
//   - GetApLocRangingCfg - AP location ranging configuration data
//   - GetGeolocationCfg - Geolocation configuration data
package general
