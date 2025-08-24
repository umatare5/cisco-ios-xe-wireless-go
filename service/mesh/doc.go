// Package mesh provides wireless mesh networking configuration and operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with mesh networking configuration and operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for mesh topology management,
// backhaul configuration, and mesh network monitoring across wireless infrastructures.
//
// The package leverages the YANG models: Cisco-IOS-XE-wireless-mesh-cfg and Cisco-IOS-XE-wireless-mesh-oper for complete mesh management.
//
// # Main Functions
//
// The Mesh service provides two main categories of operations:
//
// 1. GetCfg - Configuration Data
// 2. GetOper - Operational Data
//
// # Filter Functions
//
// This package provides 1 filter function for precise data retrieval:
//
// Configuration Filters:
//   - GetCfgByProfileName - Filter by mesh profile name
package mesh
