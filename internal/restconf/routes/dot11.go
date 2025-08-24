package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// 802.11 Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for 802.11 standard
// configuration based on Cisco-IOS-XE-wireless-dot11-cfg YANG model.

const (
	// Dot11CfgBasePath defines the base path for 802.11 configuration endpoints
	Dot11CfgBasePath = restconf.YANGModelPrefix + "dot11-cfg:dot11-cfg-data"
)

// 802.11 Configuration Endpoints
const (
	// Dot11CfgEndpoint retrieves complete 802.11 configuration data
	Dot11CfgEndpoint = Dot11CfgBasePath

	// Dot11CfgConfiguredCountriesEndpoint retrieves configured countries
	Dot11CfgConfiguredCountriesEndpoint = Dot11CfgBasePath + "/configured-countries"

	// Dot11CfgDot11EntriesEndpoint retrieves 802.11 entries
	Dot11CfgDot11EntriesEndpoint = Dot11CfgBasePath + "/dot11-entries"
)
