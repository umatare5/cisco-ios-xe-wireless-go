// Package nmsp provides Network Mobility Services Protocol (NMSP) operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with NMSP operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for location services monitoring,
// client tracking, CMX integration, and network mobility analytics across wireless infrastructures.
//
// The package leverages the YANG model: Cisco-IOS-XE-wireless-nmsp-oper for complete NMSP monitoring.
//
// # Main Functions
//
// The NMSP service provides one main category of operations:
//
// 1. GetOper - Operational Data
//
// # Additional Functions
//
// The NMSP service also provides specialized data access:
//
//   - GetClientRegistration - NMSP client registration data
//   - GetCmxConnection - NMSP CMX connection data
//   - GetCmxCloudInfo - NMSP CMX cloud information
//
// # Filter Functions
//
// This package provides 1 filter function for precise data retrieval:
//
// Operational Filters:
//   - GetOperByClientID - Filter by client registration ID
package nmsp
