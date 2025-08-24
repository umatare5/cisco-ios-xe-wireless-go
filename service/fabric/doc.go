// Package fabric provides SD-Access fabric configuration operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with Software-Defined Access fabric configuration
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for fabric profile management,
// overlay network configuration, and policy enforcement settings.
//
// The package leverages the YANG model: Cisco-IOS-XE-wireless-fabric-cfg for complete fabric configuration management.
//
// # Main Functions
//
// The Fabric service provides one main category of operations:
//
// 1. GetCfg - Configuration Data
//
// # Filter Functions
//
// This package provides 2 filter functions for precise data retrieval:
//
// Configuration Filters:
//   - GetCfgByFabricProfileName - Filter by fabric profile name
//   - GetCfgByControlPlaneName - Filter by control plane name
package fabric
