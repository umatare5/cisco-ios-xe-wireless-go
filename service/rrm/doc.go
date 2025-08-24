// Package rrm provides Radio Resource Management (RRM) configuration and operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with RRM configuration and operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for radio frequency management,
// RF optimization, spectrum analysis, and emulation operations across wireless infrastructures.
//
// The package leverages the YANG models: Cisco-IOS-XE-wireless-rrm-cfg, Cisco-IOS-XE-wireless-rrm-oper, Cisco-IOS-XE-wireless-rrm-global-oper, and Cisco-IOS-XE-wireless-rrm-emul-oper for complete RRM management.
//
// # Main Functions
//
// The RRM service provides four main categories of operations:
//
// 1. GetCfg - Configuration Data
// 2. GetOper - Operational Data
// 3. GetGlobalOper - Global Operational Data
// 4. GetEmulOper - Emulation Operational Data
//
// # Filter Functions
//
// This package provides 13+ filter functions for precise data retrieval:
//
// Configuration Filters:
//   - GetCfgByBand - Filter by band
//   - GetCfgByMgrBand - Filter by manager band
//
// Operational Filters:
//   - GetOperByWtpMacAndRadioSlot - Filter by WTP MAC and radio slot ID
//   - GetOperByPhyType - Filter by PHY type
//   - GetOperByDeviceId - Filter by device ID
//
// Global Operational Filters:
//   - GetGlobalOperByPhyType - Filter by PHY type
//   - GetGlobalOperByChannelPhyType - Filter by PHY type for channel parameters
//   - GetGlobalOperByBandId - Filter by band ID
//   - GetGlobalOperByApMac - Filter by AP MAC address
//   - GetGlobalOperByWtpMacAndRadioSlot - Filter by WTP MAC and radio slot ID (2.4GHz)
//   - GetGlobalOperBy5GWtpMacAndRadioSlot - Filter by WTP MAC and radio slot ID (5GHz)
//   - GetGlobalOperBy6GhzWtpMacAndRadioSlot - Filter by WTP MAC and radio slot ID (6GHz)
package rrm
