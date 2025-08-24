package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Multicast Operational Endpoints
//
// These constants define the RESTCONF API endpoints for multicast operational
// data based on Cisco-IOS-XE-wireless-mcast-oper YANG model.

const (
	// McastOperBasePath defines the base path for Multicast operational data endpoints
	McastOperBasePath = restconf.YANGModelPrefix + "mcast-oper:mcast-oper-data"
)

// Multicast Operational Endpoints
const (
	// McastOperEndpoint provides the endpoint for retrieving all multicast operational data
	McastOperEndpoint = McastOperBasePath

	// McastOperFlexMediastreamEndpoint provides the endpoint for retrieving FlexConnect mediastream data
	McastOperFlexMediastreamEndpoint = McastOperBasePath + "/flex-mediastream-client-summary"

	// McastOperVlanL2MgidEndpoint provides the endpoint for retrieving VLAN Layer 2 multicast group ID data
	McastOperVlanL2MgidEndpoint = McastOperBasePath + "/vlan-l2-mgid-op"

	// McastOperStatisticsEndpoint provides the endpoint for retrieving multicast statistics
	McastOperStatisticsEndpoint = McastOperBasePath + "/statistics"

	// McastOperGroupsEndpoint provides the endpoint for retrieving multicast groups
	McastOperGroupsEndpoint = McastOperBasePath + "/groups"
)
