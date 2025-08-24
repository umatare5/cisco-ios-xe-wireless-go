// Package geolocation provides geographic location services operational operations for Cisco IOS-XE wireless controllers.
//
// This package allows you to retrieve geolocation operational data from a Cisco Catalyst 9800 Wireless LAN Controller.
// It provides methods for accessing geographic positioning, location mapping, and spatial analytics across wireless infrastructures.
//
// # Main Features
//
// - Operational Data: GetOper(), GetOperApGeoLocStats()
//
// # Usage Example
//
//	// Create a client and access Geolocation service
//	client := wnc.NewClient("controller.example.com", "your-token")
//	geolocationService := client.Geolocation()
//
//	// Get geolocation operational data
//	oper, err := geolocationService.GetOper(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get AP geolocation statistics
//	stats, err := geolocationService.GetOperApGeoLocStats(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
// # KNOWN LIMITATIONS
//
// Based on testing with actual Cisco Catalyst 9800 controllers, this package provides
// access to available geolocation operational data. Some endpoints may return HTTP 404
// errors if geolocation features are not enabled or configured on the controller.
//
// The available functionality includes:
// - Basic geolocation operational data retrieval
// - AP geolocation statistics access
//
// # Error Handling
//
// Methods may return HTTP 404 errors for unavailable endpoints. This is expected behavior
// for endpoints not implemented in the target controller deployment or when geolocation
// features are not enabled.
//
// # Requirements
//
// - Cisco Catalyst 9800 Wireless LAN Controller
// - IOS-XE 17.12 or later
// - RESTCONF API access enabled
// - Valid authentication credentials
// - Geolocation features enabled on the controller (if required)
package geolocation
