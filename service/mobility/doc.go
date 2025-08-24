// Package mobility provides wireless client mobility operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with client mobility operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for client roaming monitoring,
// mobility group management, and seamless handoff tracking across wireless infrastructures.
//
// The package leverages the YANG model: Cisco-IOS-XE-wireless-mobility-oper for complete mobility monitoring.
//
// # Main Functions
//
// The Mobility service provides one main category of operations:
//
// 1. GetOper - Operational Data
//
// # Filter Functions
//
// This package provides 2 filter functions for precise data retrieval:
//
// Operational Filters:
//   - GetOperByClientMAC - Filter by client MAC address
//   - GetOperByAPMac - Filter by Access Point MAC address
package mobility
