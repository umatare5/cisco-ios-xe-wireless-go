package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// AFC (Automated Frequency Coordination) Operational Endpoints
//
// These constants define the RESTCONF API endpoints for AFC operational
// data based on Cisco-IOS-XE-wireless-afc YANG models.

// AFC Base Paths
const (
	// AFCOperBasePath defines the base path for AFC operational data endpoints
	AFCOperBasePath = restconf.YANGModelPrefix + "afc-oper:afc-oper-data"

	// AFCCloudOperBasePath defines the base path for AFC cloud operational data endpoints
	AFCCloudOperBasePath = restconf.YANGModelPrefix + "afc-cloud-oper:afc-cloud-oper-data"
)

// AFC Operational Endpoints
const (
	// AFCOperEndpoint retrieves overall AFC operational data
	AFCOperEndpoint = AFCOperBasePath

	// ApRespEndpoint retrieves per-AP AFC response data
	ApRespEndpoint = AFCOperBasePath + "/ewlc-afc-ap-resp"
)

// AFC Cloud Operational Endpoints
const (
	// CloudOperEndpoint retrieves AFC cloud operational data
	CloudOperEndpoint = AFCCloudOperBasePath

	// CloudStatsEndpoint retrieves AFC cloud statistics
	CloudStatsEndpoint = AFCCloudOperBasePath + "/afc-cloud-stats"
)
