// Package mcast provides multicast operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with multicast operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for multicast group monitoring,
// IGMP statistics, and multicast traffic analysis across wireless infrastructures.
//
// The package leverages the YANG model: Cisco-IOS-XE-wireless-mcast-oper for complete multicast monitoring.
//
// # Main Functions
//
// The Mcast service provides one main category of operations:
//
// 1. GetOper - Operational Data
//
// # Additional Functions
//
// The Mcast service also provides specialized data access:
//
//   - GetFlexMediastreamClientSummary - FlexConnect mediastream client summary data
//   - GetVlanL2MgidOp - VLAN Layer 2 multicast group ID operational data
//
// # Filter Functions
//
// This package provides 2 filter functions for precise data retrieval:
//
// Operational Filters:
//   - GetOperByClientMAC - Filter by client MAC address
//   - GetOperByVlanIndex - Filter by VLAN index
package mcast
