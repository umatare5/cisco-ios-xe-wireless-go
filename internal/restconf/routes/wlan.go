package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// WLAN Configuration and Operational Endpoints
//
// These constants define the RESTCONF API endpoints for WLAN configuration
// and operational data based on Cisco-IOS-XE-wireless-wlan YANG models.

// WLAN Base Paths
const (
	// WlanCfgBasePath defines the base path for WLAN configuration endpoints
	WlanCfgBasePath = restconf.YANGModelPrefix + "wlan-cfg:wlan-cfg-data"

	// WlanGlobalOperBasePath defines the base path for WLAN global operational data endpoints
	WlanGlobalOperBasePath = restconf.YANGModelPrefix + "wlan-global-oper:wlan-global-oper-data"

	// WlanOperGlobalOperPath defines the path for WLAN operational global operation data
	WlanOperGlobalOperPath = restconf.YANGModelPrefix + "wlan-oper:wlan-oper/global-oper"
)

// WLAN Configuration Endpoints
const (
	// WlanCfgEndpoint defines complete WLAN configuration data endpoint
	WlanCfgEndpoint = WlanCfgBasePath

	// WlanCfgEntriesEndpoint defines WLAN configuration entries endpoint
	WlanCfgEntriesEndpoint = WlanCfgBasePath + "/wlan-cfg-entries"

	// WlanCfgEntryEndpoint defines individual WLAN configuration entry endpoint
	WlanCfgEntryEndpoint = WlanCfgEntriesEndpoint + "/wlan-cfg-entry"

	// WlanPoliciesEndpoint defines WLAN policies endpoint
	WlanPoliciesEndpoint = WlanCfgBasePath + "/wlan-policies"

	// WlanPolicyEntryEndpoint defines individual WLAN policy entry endpoint
	WlanPolicyEntryEndpoint = WlanPoliciesEndpoint + "/wlan-policy"

	// PolicyListEntriesEndpoint defines policy list entries endpoint
	PolicyListEntriesEndpoint = WlanCfgBasePath + "/policy-list-entries"

	// PolicyListEntryEndpoint defines individual policy list entry endpoint
	PolicyListEntryEndpoint = PolicyListEntriesEndpoint + "/policy-list-entry"

	// WirelessAaaPolicyConfigsEndpoint defines wireless AAA policy configurations endpoint
	WirelessAaaPolicyConfigsEndpoint = WlanCfgBasePath + "/wireless-aaa-policy-configs"
)

// WLAN Operational Endpoints
const (
	// WlanGlobalOperDataEndpoint defines WLAN global operational data endpoint
	WlanGlobalOperDataEndpoint = WlanGlobalOperBasePath

	// WlanProfilesEndpoint defines WLAN profiles endpoint for operational data
	WlanProfilesEndpoint = WlanOperGlobalOperPath + "/wlan-profiles"
)
