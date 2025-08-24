package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// AP (Access Point) Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for access point configuration
// operations based on the Cisco-IOS-XE-wireless-ap-cfg YANG model.

const (
	// APCfgBasePath defines the base path for access point configuration endpoints
	APCfgBasePath = restconf.YANGModelPrefix + "ap-cfg:ap-cfg-data"

	// APOperBasePath defines the base path for access point operational endpoints
	APOperBasePath = restconf.YANGModelPrefix + "access-point-oper:access-point-oper-data"

	// APGlobalOperBasePath defines the base path for global access point operational endpoints
	APGlobalOperBasePath = restconf.YANGModelPrefix + "ap-global-oper:ap-global-oper-data"
)

// AP Configuration Endpoints
const (
	// APCfgEndpoint retrieves complete access point configuration data
	APCfgEndpoint = APCfgBasePath
)

// AP Global Operational Endpoints
const (
	// APGlobalOperEndpoint retrieves complete AP global operational data
	APGlobalOperEndpoint = APGlobalOperBasePath

	// EwlcAPStatsEndpoint retrieves EWLC AP statistics
	EwlcAPStatsEndpoint = APGlobalOperBasePath + "/ewlc-ap-stats"
)

// AP Operational Endpoints
const (
	// APOperEndpoint retrieves complete access point operational data
	APOperEndpoint = APOperBasePath

	// RadioOperDataEndpoint retrieves radio operational data for access points
	RadioOperDataEndpoint = APOperBasePath + "/radio-oper-data"

	// CapwapDataEndpoint retrieves CAPWAP data for access points
	CapwapDataEndpoint = APOperBasePath + "/capwap-data"

	// NameMacMapEndpoint retrieves AP name to MAC address mapping
	NameMacMapEndpoint = APOperBasePath + "/ap-name-mac-map"
)

// AP RPC Endpoints
//
// These constants define RPC endpoints for access point operations such as
// administrative state changes and reset operations.

const (
	// SetAPSlotAdminStateRPC defines the RPC endpoint for setting AP slot (radio) administrative state
	SetAPSlotAdminStateRPC = "/restconf/operations/Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-slot-admin-state"

	// SetAPAdminStateRPC defines the RPC endpoint for setting AP administrative state
	SetAPAdminStateRPC = "/restconf/operations/Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-admin-state"

	// APResetRPC defines the RPC endpoint for AP reset operations
	APResetRPC = "/restconf/operations/Cisco-IOS-XE-wireless-access-point-cmd-rpc:ap-reset"
)
