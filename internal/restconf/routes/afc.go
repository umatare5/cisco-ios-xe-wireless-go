package routes

// AFC (Automated Frequency Coordination) Operational Paths
//
// These constants define the RESTCONF API paths for AFC operational
// data based on Cisco-IOS-XE-wireless-afc YANG models.

// AFC Operational Paths.
const (
	// AFCOperPath retrieves overall AFC operational data.
	AFCOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data"

	// AFCEwlcAFCApRespPath retrieves per-AP AFC response data.
	AFCEwlcAFCApRespPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp"

	// AFCEwlcAFCApReqPath retrieves per-AP AFC request data.
	AFCEwlcAFCApReqPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-req"
)

// AFC Cloud Operational Paths.
const (
	// AFCCloudOperPath retrieves AFC cloud operational data.
	AFCCloudOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data"

	// AFCAfcCloudStatsPath retrieves AFC cloud statistics.
	AFCAfcCloudStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data/afc-cloud-stats"
)
