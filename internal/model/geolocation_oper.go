// Package model contains generated response structures for the Cisco WNC API.
// This package is part of the three-layer architecture providing Generated Type separation.
package model

// Geolocation Operational Response Types

// GeolocationOperResponse represents the response structure for geolocation operational data.
type GeolocationOperResponse struct {
	CiscoIOSXEWirelessGeolocationOperGeolocationOperData struct {
		ApGeoLocStats ApGeoLocStats `json:"ap-geo-loc-stats"`
	} `json:"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"`
}

// GeolocationOperApGeoLocStatsResponse represents the response structure for AP geolocation statistics.
type GeolocationOperApGeoLocStatsResponse struct {
	ApGeoLocStats ApGeoLocStats `json:"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats"`
}

// Geolocation Supporting Types

// ApGeoLocStats represents AP geolocation statistics including counts of different positioning methods.
type ApGeoLocStats struct {
	NumApGnss               int    `json:"num-ap-gnss"`
	NumApManHeight          int    `json:"num-ap-man-height"`
	NumApDerived            int    `json:"num-ap-derived"`
	LastDerivationTimestamp string `json:"last-derivation-timestamp"`
}
