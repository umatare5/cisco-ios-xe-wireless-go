// Package geolocation provides geolocation operational data functionality for the Cisco Wireless Network Controller API.
package geolocation

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// GeolocationOperBasePath defines the base path for geolocation operational data endpoints.
	GeolocationOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"
	// GeolocationOperEndpoint defines the endpoint for geolocation operational data.
	GeolocationOperEndpoint = GeolocationOperBasePath
	// GeolocationApGeoLocStatsEndpoint defines the endpoint for AP geolocation statistics.
	GeolocationApGeoLocStatsEndpoint = GeolocationOperBasePath + "/ap-geo-loc-stats"
)

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

// ApGeoLocStats represents AP geolocation statistics including counts of different positioning methods.
type ApGeoLocStats struct {
	NumApGnss               int    `json:"num-ap-gnss"`
	NumApManHeight          int    `json:"num-ap-man-height"`
	NumApDerived            int    `json:"num-ap-derived"`
	LastDerivationTimestamp string `json:"last-derivation-timestamp"`
}

// GetGeolocationOper retrieves geolocation operational data.
func GetGeolocationOper(client *wnc.Client, ctx context.Context) (*GeolocationOperResponse, error) {
	var data GeolocationOperResponse
	if err := client.SendAPIRequest(ctx, GeolocationOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetGeolocationOperApGeoLocStats retrieves AP geolocation statistics.
func GetGeolocationOperApGeoLocStats(client *wnc.Client, ctx context.Context) (*GeolocationOperApGeoLocStatsResponse, error) {
	var data GeolocationOperApGeoLocStatsResponse
	if err := client.SendAPIRequest(ctx, GeolocationApGeoLocStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
