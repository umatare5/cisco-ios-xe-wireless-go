package location

import "time"

// LocationOper represents the response structure for location operational endpoint (HTTP 204).
type LocationOper struct {
	CiscoIOSXEWirelessLocationOperData *struct {
		LocationRssiMeasurements []LocationRssiMeasurement `json:"location-rssi-measurements,omitempty"` // Radio measurements for specific AP hearing client (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-location-oper:location-oper-data,omitempty"` // Location operational data (YANG: IOS-XE 17.12.1)
}

// LocationLocationRssiMeasurements represents the response structure for location RSSI measurements endpoint (HTTP 204).
type LocationLocationRssiMeasurements struct {
	LocationRssiMeasurements []LocationRssiMeasurement `json:"Cisco-IOS-XE-wireless-location-oper:location-rssi-measurements,omitempty"`
}

// LocationRssiMeasurement represents radio measurements for AP-client pair.
type LocationRssiMeasurement struct {
	ClientMacAddr     string             `json:"client-mac-addr"`              // Wireless client MAC address　(YANG: IOS-XE 17.12.1)
	LradAddr          string             `json:"lrad-addr"`                    // AP MAC address which heard the client　(YANG: IOS-XE 17.12.1)
	RadioMeasurements []RadioMeasurement `json:"radio-measurements,omitempty"` // List of radio measurement per AP slot　(YANG: IOS-XE 17.12.1)
}

// RadioMeasurement represents RSSI measurements for an AP radio slot.
type RadioMeasurement struct {
	Band   *string     `json:"band,omitempty"`   // Radio band type　(YANG: IOS-XE 17.12.1)
	Type   *string     `json:"type,omitempty"`   // Source of the radio measurement　(YANG: IOS-XE 17.12.1)
	Sample *RssiSample `json:"sample,omitempty"` // RSSI samples for antennas A and B　(YANG: IOS-XE 17.12.1)
	Snr    *int8       `json:"snr,omitempty"`    // Signal over noise ratio　(YANG: IOS-XE 17.12.1)
}

// RssiSample represents RSSI values per antenna and timestamp.
type RssiSample struct {
	RssiValueA    *int8      `json:"rssi-value-a,omitempty"`   // RSSI value of antenna A　(YANG: IOS-XE 17.12.1)
	RssiValueB    *int8      `json:"rssi-value-b,omitempty"`   // RSSI value of antenna B　(YANG: IOS-XE 17.12.1)
	RssiTimestamp *time.Time `json:"rssi-timestamp,omitempty"` // Last updated time for the RSSI　(YANG: IOS-XE 17.12.1)
}
