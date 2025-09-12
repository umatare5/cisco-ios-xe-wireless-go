package routes

// Multicast Operational Paths
//
// These constants define the RESTCONF API paths for multicast operational
// data based on Cisco-IOS-XE-wireless-mcast-oper YANG model.

// Multicast Operational Paths.
const (
	// McastOperPath provides the path for retrieving all multicast operational data.
	McastOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data"

	// McastFlexMediastreamPath provides the path for retrieving FlexConnect mediastream data.
	McastFlexMediastreamPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/flex-mediastream-client-summary"

	// McastVlanL2MgidPath provides the path for retrieving VLAN Layer 2 multicast group ID data.
	McastVlanL2MgidPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/vlan-l2-mgid-op"

	// McastFabricMediastreamPath provides the path for retrieving fabric mediastream client summary.
	McastFabricMediastreamPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/fabric-media-stream-client-summary"

	// McastMgidInfoPath provides the path for retrieving multicast MGID information.
	McastMgidInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/mcast-mgid-info"

	// McastMulticastOperDataPath provides the path for retrieving multicast operational data.
	McastMulticastOperDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/multicast-oper-data"

	// McastStatisticsPath provides the path for retrieving multicast statistics.
	McastStatisticsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/statistics"

	// McastGroupsPath provides the path for retrieving multicast groups.
	McastGroupsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/groups"
)
