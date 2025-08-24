package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Rogue Detection Operational Endpoints
//
// These constants define the RESTCONF API endpoints for rogue detection
// operational data based on Cisco-IOS-XE-wireless-rogue-oper YANG model.

const (
	// RogueOperBasePath defines the base path for rogue operational data endpoints
	RogueOperBasePath = restconf.YANGModelPrefix + "rogue-oper:rogue-oper-data"
)

// Rogue Operational Endpoints
const (
	// RogueOperEndpoint defines the endpoint for rogue operational data
	RogueOperEndpoint = RogueOperBasePath

	// RogueStatsEndpoint defines the endpoint for rogue statistics
	RogueStatsEndpoint = RogueOperBasePath + "/rogue-stats"

	// RogueDataEndpoint defines the endpoint for rogue data
	RogueDataEndpoint = RogueOperBasePath + "/rogue-data"

	// RogueClientDataEndpoint defines the endpoint for rogue client data
	RogueClientDataEndpoint = RogueOperBasePath + "/rogue-client-data"

	// RldpStatsEndpoint defines the endpoint for RLDP statistics
	RldpStatsEndpoint = RogueOperBasePath + "/rldp-stats"
)
