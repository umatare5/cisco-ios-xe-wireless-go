// Package wlan provides Wireless LAN (WLAN) configuration and operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with WLAN configuration and global operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for wireless LAN management,
// security policy configuration, and WLAN operational monitoring across wireless infrastructures.
//
// The package leverages the YANG models: Cisco-IOS-XE-wireless-wlan-cfg and Cisco-IOS-XE-wireless-wlan-global-oper for complete WLAN management.
//
// # Main Functions
//
// The WLAN service provides two main categories of operations:
//
// 1. GetCfg - Configuration Data
// 2. GetGlobalOper - Global Operational Data
//
// # Additional Functions
//
// The WLAN service also provides specialized data access:
//
//   - GetCfgEntries - WLAN configuration entries
//   - GetPolicies - WLAN policies
//   - GetPolicyListEntries - Policy list entries
//   - GetWirelessAaaPolicyConfigs - Wireless AAA policy configurations
//
// # Filter Functions
//
// This package provides 4 filter functions for precise data retrieval:
//
// Configuration Filters:
//   - GetCfgByProfileName - Filter by WLAN profile name
//   - GetCfgByID - Filter by WLAN ID
//   - GetPoliciesByPolicyProfileName - Filter by policy profile name
//
// Global Operational Filters:
//   - GetGlobalOperByWlanProfile - Filter by WLAN profile
//
// # WLAN Management Operations
//
// The service provides management operations for WLANs:
//
// **Status Control:**
//   - EnableWLAN(ctx, wlanID): Enables a WLAN by its ID
//   - DisableWLAN(ctx, wlanID): Disables a WLAN by its ID
//   - UpdateWLANStatus(ctx, wlanID, enabled): Generic status update
//
// These operations modify the WLAN configuration to control the
// administrative state. Disabling a WLAN prevents clients from
// associating while preserving the configuration.
package wlan
