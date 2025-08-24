package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// mDNS (Multicast DNS) Operational Endpoints
//
// These constants define the RESTCONF API endpoints for mDNS operational
// data based on Cisco-IOS-XE-wireless-mdns-oper YANG model.

const (
	// MDNSOperBasePath defines the base path for mDNS operational data endpoints
	MDNSOperBasePath = restconf.YANGModelPrefix + "mdns-oper:mdns-oper-data"
)

// mDNS Operational Endpoints
const (
	// MDNSOperEndpoint provides the base endpoint for mDNS operational data
	MDNSOperEndpoint = MDNSOperBasePath

	// MDNSGlobalStatsEndpoint provides the endpoint for mDNS global statistics
	MDNSGlobalStatsEndpoint = MDNSOperBasePath + "/mdns-global-stats"

	// MDNSWlanStatsEndpoint provides the endpoint for mDNS WLAN statistics
	MDNSWlanStatsEndpoint = MDNSOperBasePath + "/mdns-wlan-stats"
)
