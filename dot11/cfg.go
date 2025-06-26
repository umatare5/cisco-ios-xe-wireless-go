// Package dot11 provides 802.11 configuration functionality for the Cisco Wireless Network Controller API.
package dot11

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// Dot11CfgBasePath defines the base path for 802.11 configuration endpoints
	Dot11CfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data"
	// Dot11CfgEndpoint retrieves complete 802.11 configuration data
	Dot11CfgEndpoint = Dot11CfgBasePath
	// ConfiguredCountriesEndpoint retrieves configured country codes
	ConfiguredCountriesEndpoint = Dot11CfgBasePath + "/configured-countries"
	// Dot11acMcsEntriesEndpoint retrieves 802.11ac MCS entries
	Dot11acMcsEntriesEndpoint = Dot11CfgBasePath + "/dot11ac-mcs-entries"
	// Dot11EntriesEndpoint retrieves 802.11 band entries
	Dot11EntriesEndpoint = Dot11CfgBasePath + "/dot11-entries"
)

// Dot11CfgResponse represents the complete 802.11 configuration response
type Dot11CfgResponse struct {
	CiscoIOSXEWirelessDot11CfgDot11CfgData struct {
		ConfiguredCountries ConfiguredCountries `json:"configured-countries"`
		Dot11acMcsEntries   Dot11acMcsEntries   `json:"dot11ac-mcs-entries"`
		Dot11Entries        Dot11Entries        `json:"dot11-entries"`
	} `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data"`
}

// Dot11ConfiguredCountriesResponse represents the configured countries response
type Dot11ConfiguredCountriesResponse struct {
	ConfiguredCountries ConfiguredCountries `json:"Cisco-IOS-XE-wireless-dot11-cfg:configured-countries"`
}

// Dot11acMcsEntriesResponse represents the 802.11ac MCS entries response
type Dot11acMcsEntriesResponse struct {
	Dot11acMcsEntries Dot11acMcsEntries `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11ac-mcs-entries"`
}

// Dot11EntriesResponse represents the 802.11 entries response
type Dot11EntriesResponse struct {
	Dot11Entries Dot11Entries `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-entries"`
}

// ConfiguredCountries contains configured country codes for regulatory compliance
type ConfiguredCountries struct {
	ConfiguredCountry []ConfiguredCountry `json:"configured-country"`
}

// ConfiguredCountry represents a configured country code entry
type ConfiguredCountry struct {
	CountryCode string `json:"country-code"` // ISO country code
}

// Dot11acMcsEntries contains 802.11ac Modulation and Coding Scheme entries
type Dot11acMcsEntries struct {
	Dot11acMcsEntry []Dot11acMcsEntry `json:"dot11ac-mcs-entry"`
}

// Dot11acMcsEntry represents an 802.11ac MCS configuration entry
type Dot11acMcsEntry struct {
	SpatialStream int    `json:"spatial-stream"` // Number of spatial streams
	Index         string `json:"index"`          // MCS index
}

// Dot11Entries contains 802.11 band configuration entries
type Dot11Entries struct {
	Dot11Entry []Dot11Entry `json:"dot11-entry"`
}

// Dot11Entry represents 802.11 band-specific configuration
type Dot11Entry struct {
	Band                string `json:"band"`                             // Radio band (2.4GHz, 5GHz, 6GHz)
	VoiceAdmCtrlSupport bool   `json:"voice-adm-ctrl-support,omitempty"` // Voice admission control support
	Dot11axCfg          *struct {
		HeBssColor bool `json:"he-bss-color"` // High Efficiency BSS Color configuration
	} `json:"dot11ax-cfg,omitempty"`
	AmpduEntries *struct {
		AmpduEntry []struct {
			Index                    int    `json:"index"`                        // AMPDU entry index
			Apf80211nAmpduTxPriority string `json:"apf-80211n-ampdu-tx-priority"` // AMPDU transmission priority
		} `json:"ampdu-entry"`
	} `json:"ampdu-entries,omitempty"`
	AmpduTxScheduler *struct{} `json:"ampdu-tx-scheduler,omitempty"`
	AmsduEntries     *struct {
		AmsduEntry []struct {
			Index                    int    `json:"index"`                        // AMSDU entry index
			Apf80211nAmsduTxPriority string `json:"apf-80211n-amsdu-tx-priority"` // AMSDU transmission priority
		} `json:"amsdu-entry"`
	} `json:"amsdu-entries,omitempty"`
	MediaStreamBandInfoCfg *struct{} `json:"media-stream-band-info-cfg,omitempty"`
	SpectrumCfg            *struct {
		RrmEdEnable bool `json:"rrm-ed-enable,omitempty"` // RRM energy detection enable
	} `json:"spectrum-cfg,omitempty"`
	Dot11axMcsEntries *struct {
		Dot11axMcsEntry []struct {
			SpatialStream int    `json:"spatial-stream"` // Number of spatial streams
			Index         string `json:"index"`          // MCS index for 802.11ax
		} `json:"dot11ax-mcs-entry"`
	} `json:"dot11ax-mcs-entries,omitempty"`
}

// GetDot11Cfg retrieves complete 802.11 configuration data.
func GetDot11Cfg(client *wnc.Client, ctx context.Context) (*Dot11CfgResponse, error) {
	var data Dot11CfgResponse
	if err := client.SendAPIRequest(ctx, Dot11CfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDot11ConfiguredCountries retrieves configured country codes.
func GetDot11ConfiguredCountries(client *wnc.Client, ctx context.Context) (*Dot11ConfiguredCountriesResponse, error) {
	var data Dot11ConfiguredCountriesResponse
	if err := client.SendAPIRequest(ctx, ConfiguredCountriesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDot11acMcsEntries retrieves 802.11ac MCS configuration entries.
func GetDot11acMcsEntries(client *wnc.Client, ctx context.Context) (*Dot11acMcsEntriesResponse, error) {
	var data Dot11acMcsEntriesResponse
	if err := client.SendAPIRequest(ctx, Dot11acMcsEntriesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDot11Entries retrieves 802.11 band configuration entries.
func GetDot11Entries(client *wnc.Client, ctx context.Context) (*Dot11EntriesResponse, error) {
	var data Dot11EntriesResponse
	if err := client.SendAPIRequest(ctx, Dot11EntriesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
