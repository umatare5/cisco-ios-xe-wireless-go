package routes

// Client Operational Paths
//
// These constants define the RESTCONF API paths for wireless client operational
// data based on Cisco-IOS-XE-wireless-client-oper YANG model.

// Client Operational Paths.
const (
	// ClientOperPath retrieves complete client operational data.
	ClientOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data"

	// ClientCommonOperDataPath retrieves common operational data for clients.
	ClientCommonOperDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data"

	// ClientDot11OperDataPath retrieves 802.11 operational data for clients.
	ClientDot11OperDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/dot11-oper-data"

	// ClientMobilityOperDataPath retrieves mobility operational data for clients.
	ClientMobilityOperDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mobility-oper-data"

	// ClientMmIfClientStatsPath retrieves mobility manager interface client statistics.
	ClientMmIfClientStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-stats"

	// ClientMmIfClientHistoryPath retrieves mobility manager interface client history.
	ClientMmIfClientHistoryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-history"

	// ClientTrafficStatsPath retrieves client traffic statistics.
	ClientTrafficStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/traffic-stats"

	// ClientPolicyDataPath retrieves client policy data.
	ClientPolicyDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/policy-data"

	// ClientSisfDBMacPath retrieves SISF database MAC information.
	ClientSisfDBMacPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/sisf-db-mac"

	// ClientDcInfoPath retrieves discovery client information.
	ClientDcInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/dc-info"
)

// Client RPC Operations
//
// Based on Cisco-IOS-XE-wireless-client-rpc YANG model specifications.

// Client RPC Operations.
const (
	// ClientDeauthRPC deauthenticates one wireless client. The module declares it from revision
	// 2024-03-01, so 17.12 answers 400 rather than 404 — the path is absent, not the resource.
	ClientDeauthRPC = RESTCONFOperationsPath + "/Cisco-IOS-XE-wireless-client-rpc:apf-ms-delete-all"
)
