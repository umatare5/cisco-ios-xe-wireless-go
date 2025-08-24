// Package dot11 provides 802.11 wireless standard configuration operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with 802.11 wireless protocol configuration
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for wireless standard management,
// frequency band configuration, MCS settings, and country-specific regulatory compliance.
//
// The package leverages the YANG model: Cisco-IOS-XE-wireless-dot11-cfg for complete 802.11 configuration management.
//
// # Main Functions
//
// The Dot11 service provides one main category of operations:
//
// 1. GetCfg - Configuration Data
//
// # Filter Functions
//
// This package provides 3 filter functions for precise data retrieval:
//
// Configuration Filters:
//   - GetCfgByCountryCode - Filter by country code
//   - GetCfgByBand - Filter by frequency band
//   - GetCfgBySpatialStreamAndIndex - Filter by spatial stream and index for 802.11ac MCS
package dot11
