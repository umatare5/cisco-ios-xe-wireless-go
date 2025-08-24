// Package model provides type definitions for Cisco IOS-XE wireless controller operations.
package model

// Geolocation Operational Response Types

// GeolocationOper  represents the geolocation operational data.
type GeolocationOper struct {
	CiscoIOSXEWirelessGeolocationOperGeolocationOperData struct {
		ApGeoLocStats ApGeoLocStats `json:"ap-geo-loc-stats"`
	} `json:"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"`
}

// GeolocationOperApGeoLocStats  represents the AP geolocation statistics.
type GeolocationOperApGeoLocStats struct {
	ApGeoLocStats ApGeoLocStats `json:"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats"`
}

// Geolocation Supporting Types

type ApGeoLocStats struct {
	NumApGnss               int    `json:"num-ap-gnss"`
	NumApManHeight          int    `json:"num-ap-man-height"`
	NumApDerived            int    `json:"num-ap-derived"`
	LastDerivationTimestamp string `json:"last-derivation-timestamp"`
}
