// Package client provides wireless client operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with wireless client operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for client monitoring,
// statistics retrieval, mobility tracking, and policy management across wireless infrastructures.
//
// The package leverages the YANG model: Cisco-IOS-XE-wireless-client-oper for complete client monitoring.
//
// # Main Functions
//
// The Client service provides two main categories of operations:
//
// 1. GetOper - Operational Data
//
// # Additional Functions
//
// The Client service also provides specialized data access:
//
//   - GetCommonOperData - Common client operational data
//   - GetDot11OperData - 802.11 client operational data
//   - GetMobilityOperData - Client mobility operational data
//   - GetMmIfClientStats - Mobility manager interface client statistics
//   - GetMmIfClientHistory - Mobility manager interface client history
//   - GetTrafficStats - Client traffic statistics
//   - GetPolicyData - Client policy data
//   - GetSisfDBMac - SISF database MAC information
//   - GetDcInfo - Discovery client information
//
// # Client Management Operations (Not supported in IOS-XE 17.12)
//
// Note: The following management operations are not available in IOS-XE 17.12
// as the required RPC endpoints are not implemented in this version.
// Support may be added in future IOS-XE releases.
//
// **Deauthentication Operations (Commented out):**
//   - DeauthenticateByMAC would force disconnect a client by MAC address
//   - DeauthenticateByIP would force disconnect a client by IP address
//   - DeauthenticateByUsername would force disconnect all clients with a specific username
//
// These operations would use Cisco IOS-XE RPC endpoints to execute administrative
// commands equivalent to the following CLI commands:
//   - "wireless client mac-address <mac> deauthenticate"
//   - "wireless client ip-address <ip> deauthenticate"
//   - "wireless client username <username> deauthenticate"
//
// # Filter Functions
//
// This package provides 7 filter functions for precise data retrieval:
//
// Operational Filters:
//   - GetOperByClientMac - Filter by client MAC address
//   - GetOperByApName - Filter by AP name
//   - GetOperByWlanId - Filter by WLAN ID
//   - GetOperByClientType - Filter by client type
//   - GetOperByCoState - Filter by client operational state
//   - GetOperByMsRadioType - Filter by radio type
//   - GetOperByUsername - Filter by username
package client
