package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// AWIPS (Advanced Wireless Intrusion Prevention System) Operational Endpoints
//
// These constants define the RESTCONF API endpoints for AWIPS operational
// data based on Cisco-IOS-XE-wireless-awips-oper YANG model.

const (
	// AWIPSOperBasePath defines the base path for AWIPS operational data endpoints
	AWIPSOperBasePath = restconf.YANGModelPrefix + "awips-oper:awips-oper-data"
)

// AWIPS Operational Endpoints
const (
	// AWIPSOperEndpoint retrieves complete AWIPS operational data
	AWIPSOperEndpoint = AWIPSOperBasePath
)
