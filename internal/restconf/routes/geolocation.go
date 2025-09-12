package routes

// Geolocation Operational Paths
//
// These constants define the RESTCONF API paths for geolocation operational
// data based on Cisco-IOS-XE-wireless-geolocation-oper YANG model.

// Geolocation Operational Paths.
const (
	// GeolocationOperPath defines the path for geolocation operational data.
	GeolocationOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"

	// GeolocationApGeoLocStatsPath defines the path for AP geolocation statistics.
	GeolocationApGeoLocStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-stats"
)
