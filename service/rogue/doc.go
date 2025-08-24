// Package rogue provides rogue detection operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with rogue detection operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for rogue AP and client monitoring,
// security threat analysis, and RLDP (Rogue Location Discovery Protocol) statistics.
//
// The package leverages the YANG model: Cisco-IOS-XE-wireless-rogue-oper for complete rogue detection monitoring.
//
// # Main Functions
//
// The Rogue service provides one main category of operations:
//
// 1. GetOper - Operational Data
//
// # Additional Functions
//
// The Rogue service also provides specialized data access:
//
//   - GetStats - Rogue statistics
//   - GetData - Rogue data
//   - GetClientData - Rogue client data
//   - GetRldpStats - RLDP (Rogue Location Discovery Protocol) statistics
//
// # Filter Functions
//
// This package provides 2 filter functions for precise data retrieval:
//
// Operational Filters:
//   - GetOperByRogueAddress - Filter by rogue address
//   - GetOperByRogueClientAddress - Filter by rogue client address
package rogue
