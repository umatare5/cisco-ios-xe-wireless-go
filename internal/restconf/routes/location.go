package routes

// Location Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for location configuration
// and operational data based on Cisco-IOS-XE-wireless-location YANG models.

// Location Configuration Paths.
const (
	// LocationCfgPath provides the path for retrieving all location configuration data.
	LocationCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"

	// LocationOperatorLocationsPath provides the path for retrieving location profiles.
	LocationOperatorLocationsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/operator-locations"

	// LocationNmspConfigPath provides the path for retrieving location servers.
	LocationNmspConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/nmsp-config"

	// LocationPath provides the path for retrieving location settings (not supported).
	LocationPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/location"
)

// Location Operational Paths.
const (
	// LocationOperPath provides the path for retrieving location operational data.
	LocationOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-location-oper:location-oper-data"

	// LocationRssiMeasurementsPath provides the path for retrieving location statistics (not supported).
	LocationRssiMeasurementsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-location-oper:location-oper-data/location-rssi-measurements"
)
