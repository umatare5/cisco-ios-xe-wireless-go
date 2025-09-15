package routes

// Site Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for site configuration
// and operational data based on Cisco-IOS-XE-wireless-site YANG models.

// Site Configuration Paths.
const (
	// SiteCfgPath provides the path for site configuration data.
	SiteCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"

	// SiteTagsPath provides the path for site tags.
	SiteTagsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tags"

	// SiteTagByNamePath provides the path for site tag by name.
	SiteTagByNamePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tags/site-tag"

	// SiteTagConfigsPath provides the path for site tag configurations.
	SiteTagConfigsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs"

	// APProfilesPath provides the path for AP profiles (site context).
	APProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/ap-cfg-profiles"
)

// Site Query Paths.
const (
	// SiteTagConfigQueryPath provides the path for querying site tag config by tag name.
	SiteTagConfigQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config"
)
