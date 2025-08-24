// Package mdns provides multicast DNS (mDNS) operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with mDNS operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for service discovery monitoring,
// Bonjour protocol analysis, and zero-configuration networking statistics.
//
// The package leverages the YANG model: Cisco-IOS-XE-wireless-mdns-oper for complete mDNS monitoring.
//
// # Main Functions
//
// The MDNS service provides one main category of operations:
//
// 1. GetOper - Operational Data
//
// # Additional Functions
//
// The MDNS service also provides specialized data access:
//
//   - GetGlobalStats - mDNS global statistics
//   - GetWlanStats - mDNS WLAN statistics
//
// # Filter Functions
//
// This package provides 1 filter function for precise data retrieval:
//
// Operational Filters:
//   - GetOperByWlanID - Filter by WLAN ID
package mdns
