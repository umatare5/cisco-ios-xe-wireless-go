// Package location provides data models for location operational data.
package location

import "time"

// LocationOper represents the response structure for location operational endpoint (HTTP 204).
type LocationOper struct {
	LocationOperData *LocationOperData `json:"Cisco-IOS-XE-wireless-location-oper:location-oper-data,omitempty"`
}

// LocationOperData represents location operational data container.
type LocationOperData struct {
	// Radio measurements for specific AP hearing from specific client
	LocationRssiMeasurements []LocationRssiMeasurement `json:"location-rssi-measurements,omitempty"` // (YANG: IOS-XE 17.12.1+)
}

// LocationLocationRssiMeasurements represents the response structure for location RSSI measurements endpoint (HTTP 204).
type LocationLocationRssiMeasurements struct {
	LocationRssiMeasurements []LocationRssiMeasurement `json:"Cisco-IOS-XE-wireless-location-oper:location-rssi-measurements,omitempty"`
}

// LocationRssiMeasurement represents radio measurements for AP-client pair.
type LocationRssiMeasurement struct {
	// Wireless client MAC address (key)
	ClientMacAddr string `json:"client-mac-addr"` // (YANG: IOS-XE 17.12.1+)
	// AP MAC address which heard the client (key)
	LradAddr string `json:"lrad-addr"` // (YANG: IOS-XE 17.12.1+)
	// List of radio measurements per AP slot (max 4 elements)
	RadioMeasurements []RadioMeasurement `json:"radio-measurements,omitempty"` // (YANG: IOS-XE 17.12.1+)
}

// RadioMeasurement represents RSSI measurements for an AP radio slot.
type RadioMeasurement struct {
	// Radio band type enum
	Band *string `json:"band,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Source of radio measurement enum
	Type *string `json:"type,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// RSSI samples for antennas A and B
	Sample *RssiSample `json:"sample,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Signal over noise ratio
	Snr *int8 `json:"snr,omitempty"` // (YANG: IOS-XE 17.12.1+)
}

// RssiSample represents RSSI values per antenna and timestamp.
type RssiSample struct {
	// RSSI value of antenna A
	RssiValueA *int8 `json:"rssi-value-a,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// RSSI value of antenna B
	RssiValueB *int8 `json:"rssi-value-b,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Last updated time for RSSI
	RssiTimestamp *time.Time `json:"rssi-timestamp,omitempty"` // (YANG: IOS-XE 17.12.1+)
}
