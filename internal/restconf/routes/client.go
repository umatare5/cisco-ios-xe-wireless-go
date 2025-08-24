package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Client Operational Endpoints
//
// These constants define the RESTCONF API endpoints for wireless client operational
// data based on Cisco-IOS-XE-wireless-client-oper YANG model.

const (
	// ClientOperBasePath defines the base path for client operational data endpoints
	ClientOperBasePath = restconf.YANGModelPrefix + "client-oper:client-oper-data"
)

// Client Operational Endpoints
const (
	// ClientOperEndpoint retrieves complete client operational data
	ClientOperEndpoint = ClientOperBasePath

	// CommonOperDataEndpoint retrieves common operational data for clients
	CommonOperDataEndpoint = ClientOperBasePath + "/common-oper-data"

	// Dot11OperDataEndpoint retrieves 802.11 operational data for clients
	Dot11OperDataEndpoint = ClientOperBasePath + "/dot11-oper-data"

	// MobilityOperDataEndpoint retrieves mobility operational data for clients
	MobilityOperDataEndpoint = ClientOperBasePath + "/mobility-oper-data"

	// MmIfClientStatsEndpoint retrieves mobility manager interface client statistics
	MmIfClientStatsEndpoint = ClientOperBasePath + "/mm-if-client-stats"

	// MmIfClientHistoryEndpoint retrieves mobility manager interface client history
	MmIfClientHistoryEndpoint = ClientOperBasePath + "/mm-if-client-history"

	// TrafficStatsEndpoint retrieves client traffic statistics
	TrafficStatsEndpoint = ClientOperBasePath + "/traffic-stats"

	// PolicyDataEndpoint retrieves client policy data
	PolicyDataEndpoint = ClientOperBasePath + "/policy-data"

	// SisfDBMacEndpoint retrieves SISF database MAC information
	SisfDBMacEndpoint = ClientOperBasePath + "/sisf-db-mac"

	// DcInfoEndpoint retrieves discovery client information
	DcInfoEndpoint = ClientOperBasePath + "/dc-info"
)
