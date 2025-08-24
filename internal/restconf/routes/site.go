package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Site Configuration and Operational Endpoints
//
// These constants define the RESTCONF API endpoints for site configuration
// and operational data based on Cisco-IOS-XE-wireless-site YANG models.

const (
	// SiteCfgBasePath defines the base path for site configuration endpoints
	SiteCfgBasePath = restconf.YANGModelPrefix + "site-cfg:site-cfg-data"

	// SiteOperBasePath defines the base path for site operational endpoints
	SiteOperBasePath = restconf.YANGModelPrefix + "site-oper-data:site-oper-data"
)

// Site Configuration Endpoints
const (
	// SiteCfgEndpoint retrieves complete site configuration data
	SiteCfgEndpoint = SiteCfgBasePath

	// APCfgProfilesEndpoint retrieves AP configuration profiles
	APCfgProfilesEndpoint = SiteCfgBasePath + "/ap-cfg-profiles"

	// SiteTagConfigsEndpoint retrieves site tag configurations
	SiteTagConfigsEndpoint = SiteCfgBasePath + "/site-tag-configs"

	// SiteTagConfigEndpoint retrieves a specific site tag configuration by name
	SiteTagConfigEndpoint = SiteTagConfigsEndpoint + "/site-tag-config"

	// SiteListEntriesEndpoint retrieves site list entries
	SiteListEntriesEndpoint = SiteCfgBasePath + "/site-list-entries"

	// APPacketCaptureProfilesEndpoint retrieves AP packet capture profiles
	APPacketCaptureProfilesEndpoint = SiteCfgBasePath + "/ap-packet-capture-profiles"

	// APTraceProfilesEndpoint retrieves AP trace profiles
	APTraceProfilesEndpoint = SiteCfgBasePath + "/ap-trace-profiles"

	// APPrimingProfilesEndpoint retrieves AP priming profiles
	APPrimingProfilesEndpoint = SiteCfgBasePath + "/ap-priming-profiles"
)

// Site Operational Endpoints
const (
	// SiteOperEndpoint retrieves site operational data
	SiteOperEndpoint = SiteOperBasePath
)
