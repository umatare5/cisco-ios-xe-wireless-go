package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// General Configuration and Operational Endpoints
//
// These constants define the RESTCONF API endpoints for general wireless
// configuration and operational data based on Cisco-IOS-XE-wireless-general YANG models.

// General Base Paths
const (
	// GeneralOperBasePath defines the base path for general operational endpoints
	GeneralOperBasePath = restconf.YANGModelPrefix + "general-oper:general-oper-data"

	// GeneralCfgBasePath defines the base path for general configuration endpoints
	GeneralCfgBasePath = restconf.YANGModelPrefix + "general-cfg:general-cfg-data"
)

// General Operational Endpoints
const (
	// GeneralOperEndpoint defines the endpoint for general operational data
	GeneralOperEndpoint = GeneralOperBasePath

	// MgmtIntfDataEndpoint defines the endpoint for management interface data
	MgmtIntfDataEndpoint = GeneralOperBasePath + "/mgmt-intf-data"
)

// General Configuration Endpoints
const (
	// MewlcConfigEndpoint defines the endpoint for MEWLC configuration
	MewlcConfigEndpoint = GeneralCfgBasePath + "/mewlc-config"

	// CacConfigEndpoint defines the endpoint for CAC configuration
	CacConfigEndpoint = GeneralCfgBasePath + "/cac-config"

	// MfpEndpoint defines the endpoint for MFP configuration
	MfpEndpoint = GeneralCfgBasePath + "/mfp"

	// FipsCfgEndpoint defines the endpoint for FIPS configuration
	FipsCfgEndpoint = GeneralCfgBasePath + "/fips-cfg"

	// WsaApClientEventEndpoint defines the endpoint for WSA AP client event configuration
	WsaApClientEventEndpoint = GeneralCfgBasePath + "/wsa-ap-client-event"

	// SimL3InterfaceCacheDataEndpoint defines the endpoint for SIM L3 interface cache data
	SimL3InterfaceCacheDataEndpoint = GeneralCfgBasePath + "/sim-l3-interface-cache-data"

	// WlcManagementDataEndpoint defines the endpoint for WLC management data
	WlcManagementDataEndpoint = GeneralCfgBasePath + "/wlc-management-data"

	// LaginfoEndpoint defines the endpoint for LAG information
	LaginfoEndpoint = GeneralCfgBasePath + "/laginfo"

	// MulticastConfigEndpoint defines the endpoint for multicast configuration
	MulticastConfigEndpoint = GeneralCfgBasePath + "/multicast-config"

	// FeatureUsageCfgEndpoint defines the endpoint for feature usage configuration
	FeatureUsageCfgEndpoint = GeneralCfgBasePath + "/feature-usage-cfg"

	// ThresholdWarnCfgEndpoint defines the endpoint for threshold warning configuration
	ThresholdWarnCfgEndpoint = GeneralCfgBasePath + "/threshold-warn-cfg"

	// ApLocRangingCfgEndpoint defines the endpoint for AP location ranging configuration
	ApLocRangingCfgEndpoint = GeneralCfgBasePath + "/ap-loc-ranging-cfg"

	// GeolocationCfgEndpoint defines the endpoint for geolocation configuration
	GeolocationCfgEndpoint = GeneralCfgBasePath + "/geolocation-cfg"
)
