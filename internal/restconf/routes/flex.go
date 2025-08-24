package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// FlexConnect Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for FlexConnect configuration
// based on Cisco-IOS-XE-wireless-flex-cfg YANG model.

const (
	// FlexCfgBasePath defines the base path for FlexConnect configuration endpoints
	FlexCfgBasePath = restconf.YANGModelPrefix + "flex-cfg:flex-cfg-data"
)

// FlexConnect Configuration Endpoints
const (
	// FlexCfgEndpoint defines the endpoint for FlexConnect configuration data
	FlexCfgEndpoint = FlexCfgBasePath

	// FlexCfgPolicyEntriesEndpoint defines the endpoint for FlexConnect policy entries
	FlexCfgPolicyEntriesEndpoint = FlexCfgBasePath + "/flex-policy-entries"
)
