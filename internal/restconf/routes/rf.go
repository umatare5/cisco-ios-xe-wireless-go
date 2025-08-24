package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// RF (Radio Frequency) Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for RF configuration
// operations based on the Cisco-IOS-XE-wireless-rf-cfg YANG model.

const (
	// RFCfgBasePath defines the base path for RF configuration endpoints
	RFCfgBasePath = restconf.YANGModelPrefix + "rf-cfg:rf-cfg-data"
)

// RF Configuration Endpoints
const (
	// RFCfgEndpoint retrieves complete RF configuration data
	RFCfgEndpoint = RFCfgBasePath

	// RfProfilesEndpoint retrieves RF profile configurations
	RfProfilesEndpoint = RFCfgBasePath + "/rf-profiles"

	// RfTagsEndpoint retrieves RF tag configurations
	RfTagsEndpoint = RFCfgBasePath + "/rf-tags"

	// RfTagEndpoint retrieves a specific RF tag configuration by name
	RfTagEndpoint = RfTagsEndpoint + "/rf-tag"

	// AtfPoliciesEndpoint retrieves ATF policy configurations
	AtfPoliciesEndpoint = RFCfgBasePath + "/atf-policies"

	// MultiBssidProfilesEndpoint retrieves Multi-BSSID profile configurations
	MultiBssidProfilesEndpoint = RFCfgBasePath + "/multi-bssid-profiles"

	// RfProfileDefaultEntriesEndpoint retrieves RF profile default entry configurations
	RfProfileDefaultEntriesEndpoint = RFCfgBasePath + "/rf-profile-default-entries"
)
