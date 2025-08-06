package model

// Dot11CfgResponse represents the response structure for 802.11 configuration data.
type Dot11CfgResponse struct {
	Dot11CfgData Dot11CfgData `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data"`
}

// Dot11CfgData contains 802.11 configuration data
type Dot11CfgData struct {
	ConfiguredCountries Dot11ConfiguredCountries `json:"configured-countries"`
	Dot11Entries        Dot11Entries             `json:"dot11-entries"`
	Dot11acMcsEntries   Dot11acMcsEntries        `json:"dot11ac-mcs-entries"`
}

// Dot11ConfiguredCountries represents configured country settings
type Dot11ConfiguredCountries struct {
	ConfiguredCountry []Dot11ConfiguredCountry `json:"configured-country"`
}

// Dot11ConfiguredCountry represents a configured country
type Dot11ConfiguredCountry struct {
	CountryCode string `json:"country-code"`
}

// Dot11Entries represents 802.11 entries
type Dot11Entries struct {
	Dot11Entry []Dot11Entry `json:"dot11-entry"`
}

// Dot11Entry represents a 802.11 entry configuration
type Dot11Entry struct {
	RadioType   string `json:"radio-type"`
	RadioPolicy string `json:"radio-policy"`
}

// Dot11acMcsEntries represents 802.11ac MCS entries
type Dot11acMcsEntries struct {
	Dot11acMcsEntry []Dot11acMcsEntry `json:"dot11ac-mcs-entry"`
}

// Dot11acMcsEntry represents a 802.11ac MCS entry
type Dot11acMcsEntry struct {
	SpatialStream int    `json:"spatial-stream"`
	Index         string `json:"index"`
}
