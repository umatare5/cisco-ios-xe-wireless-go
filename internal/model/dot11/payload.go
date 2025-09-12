// Package dot11 provides dot11-specific payloads for the Cisco IOS-XE Wireless Network Controller API.
// These structures are compatible with IOS-XE 17.12.1 YANG models and include 17.18.1+ extensions.
package dot11

// Dot11ConfigByCountryPayload represents payload for retrieving dot11 configuration by country code.
type Dot11ConfigByCountryPayload struct {
	CountryCode string `json:"country-code"`
}

// Dot11ConfigByBandPayload represents payload for retrieving dot11 configuration by band.
type Dot11ConfigByBandPayload struct {
	Band string `json:"band"`
}

// Dot11AcMcsConfigPayload represents payload for retrieving 802.11ac MCS configuration.
type Dot11AcMcsConfigPayload struct {
	SpatialStream int    `json:"spatial-stream"`
	Index         string `json:"index"`
}

// Dot11axMcsConfigPayload represents payload for retrieving 802.11ax MCS configuration.
type Dot11axMcsConfigPayload struct {
	SpatialStream int    `json:"spatial-stream"`
	Index         string `json:"index"`
}
