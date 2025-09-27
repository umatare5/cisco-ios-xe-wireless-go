package location

import "time"

// CiscoIOSXEWirelessLocationOper represents the response structure for location operational endpoint (HTTP 204).
type CiscoIOSXEWirelessLocationOper struct {
	CiscoIOSXEWirelessLocationOperData *struct {
		LocationRSSIMeasurements []LocationRSSIMeasurement `json:"location-rssi-measurements,omitempty"` // Radio measurements for specific AP hearing client (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-location-oper:location-oper-data,omitempty"` // Location operational data (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessLocationLocationRSSIMeasurements represents the response structure for location RSSI measurements endpoint (HTTP 204).
type CiscoIOSXEWirelessLocationLocationRSSIMeasurements struct {
	LocationRSSIMeasurements []LocationRSSIMeasurement `json:"Cisco-IOS-XE-wireless-location-oper:location-rssi-measurements,omitempty"`
}

// LocationRSSIMeasurement represents radio measurements for AP-client pair.
type LocationRSSIMeasurement struct {
	ClientMACAddr     string             `json:"client-mac-addr"`              // Wireless client MAC address (YANG: IOS-XE 17.12.1)
	LradAddr          string             `json:"lrad-addr"`                    // AP MAC address which heard the client (YANG: IOS-XE 17.12.1)
	RadioMeasurements []RadioMeasurement `json:"radio-measurements,omitempty"` // List of radio measurement per AP slot (YANG: IOS-XE 17.12.1)
}

// RadioMeasurement represents RSSI measurements for an AP radio slot.
type RadioMeasurement struct {
	Band   *string     `json:"band,omitempty"`   // Radio band type (YANG: IOS-XE 17.12.1)
	Type   *string     `json:"type,omitempty"`   // Source of the radio measurement (YANG: IOS-XE 17.12.1)
	Sample *RSSISample `json:"sample,omitempty"` // RSSI samples for antennas A and B (YANG: IOS-XE 17.12.1)
	SNR    *int8       `json:"snr,omitempty"`    // Signal over noise ratio (YANG: IOS-XE 17.12.1)
}

// RSSISample represents RSSI values per antenna and timestamp.
type RSSISample struct {
	RSSIValueA    *int8      `json:"rssi-value-a,omitempty"`   // RSSI value of antenna A (YANG: IOS-XE 17.12.1)
	RSSIValueB    *int8      `json:"rssi-value-b,omitempty"`   // RSSI value of antenna B (YANG: IOS-XE 17.12.1)
	RSSITimestamp *time.Time `json:"rssi-timestamp,omitempty"` // Last updated time for the RSSI (YANG: IOS-XE 17.12.1)
}
