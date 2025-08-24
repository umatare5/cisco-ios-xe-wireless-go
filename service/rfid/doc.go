// Package rfid provides Radio Frequency Identification (RFID) configuration and operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with RFID configuration and operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for RFID tag tracking,
// asset management, and identification system monitoring across wireless infrastructures.
//
// The package leverages the YANG models: Cisco-IOS-XE-wireless-rfid-cfg, Cisco-IOS-XE-wireless-rfid-oper, and Cisco-IOS-XE-wireless-rfid-global-oper for complete RFID management.
//
// # Main Functions
//
// The RFID service provides three main categories of operations:
//
// 1. GetCfg - Configuration Data
// 2. GetOper - Operational Data
// 3. GetGlobalOper - Global Operational Data
//
// # Filter Functions
//
// This package provides 3 filter functions for precise data retrieval:
//
// Operational Filters:
//   - GetOperByMac - Filter by RFID MAC address
//   - GetGlobalOperByMac - Filter by RFID MAC address for global operational data
//   - GetGlobalOperByRadioKey - Filter by radio key combination (MAC, AP MAC, slot)
package rfid
