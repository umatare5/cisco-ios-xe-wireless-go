// Package dot11 provides data models for 802.11 configuration data.
package dot11

// Dot11Cfg represents the 802.11 configuration data container.
type Dot11Cfg struct {
	Dot11CfgData Dot11CfgData `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data"`
}

// Dot11CfgConfiguredCountries represents the configured countries wrapper.
type Dot11CfgConfiguredCountries struct {
	ConfiguredCountries *Dot11ConfiguredCountries `json:"Cisco-IOS-XE-wireless-dot11-cfg:configured-countries,omitempty"`
}

// Dot11CfgDot11Entries represents the 802.11 entries wrapper.
type Dot11CfgDot11Entries struct {
	Dot11Entries *Dot11Entries `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-entries,omitempty"`
}

// Dot11CfgDot11acMcsEntries represents the 802.11ac MCS entries wrapper.
type Dot11CfgDot11acMcsEntries struct {
	Dot11acMcsEntries *Dot11acMcsEntries `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11ac-mcs-entries,omitempty"`
}

// Dot11CfgFilter represents filtered 802.11 configuration data container.
type Dot11CfgFilter struct {
	ConfiguredCountry []Dot11ConfiguredCountry `json:"Cisco-IOS-XE-wireless-dot11-cfg:configured-country,omitempty"`
	Dot11Entry        []Dot11Entry             `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-entry,omitempty"`
	Dot11acMcsEntry   []Dot11acMcsEntry        `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11ac-mcs-entry,omitempty"`
}

// Dot11CfgData represents 802.11 feature related configuration data.
type Dot11CfgData struct {
	ConfiguredCountries *Dot11ConfiguredCountries `json:"configured-countries,omitempty"`
	Dot11Entries        *Dot11Entries             `json:"dot11-entries,omitempty"`
	Dot11acMcsEntries   *Dot11acMcsEntries        `json:"dot11ac-mcs-entries,omitempty"`
}

// Dot11ConfiguredCountries represents country code configuration container.
type Dot11ConfiguredCountries struct {
	ConfiguredCountry []Dot11ConfiguredCountry `json:"configured-country"`
}

// Dot11ConfiguredCountry represents individual configured country code entry.
type Dot11ConfiguredCountry struct {
	CountryCode string `json:"country-code"` // Configured country code identifier
}

// Dot11Entries represents 802.11 protocol related configuration container.
type Dot11Entries struct {
	Dot11Entry []Dot11Entry `json:"dot11-entry"`
}

// Dot11Entry represents configuration for specific 802.11 radio band.
type Dot11Entry struct {
	Band                   string                  `json:"band"`                                 // Radio band identifier (2.4GHz, 5GHz, 6GHz)
	VoiceAdmCtrlSupport    *bool                   `json:"voice-adm-ctrl-support,omitempty"`     // Voice admission control support
	Dot11axCfg             *Dot11axCfg             `json:"dot11ax-cfg,omitempty"`                // 802.11ax configuration parameters
	AmpduEntries           *AmpduEntries           `json:"ampdu-entries,omitempty"`              // AMPDU aggregation configuration entries
	AmpduTxScheduler       *AmpduTxScheduler       `json:"ampdu-tx-scheduler,omitempty"`         // AMPDU transmission scheduler configuration
	AmsduEntries           *AmsduEntries           `json:"amsdu-entries,omitempty"`              // AMSDU aggregation configuration entries
	MediaStreamBandInfoCfg *MediaStreamBandInfoCfg `json:"media-stream-band-info-cfg,omitempty"` // Media stream band information configuration
	SpectrumCfg            *SpectrumCfg            `json:"spectrum-cfg,omitempty"`               // Radio spectrum management configuration
	Dot11axMcsEntries      *Dot11axMcsEntries      `json:"dot11ax-mcs-entries,omitempty"`        // 802.11ax MCS configuration entries
}

// Dot11axCfg represents 802.11ax feature related configuration.
type Dot11axCfg struct {
	HeMbssidCap *bool `json:"he-mbssid-cap,omitempty"` // 802.11ax multi-BSSID capability support
	HeBssColor  *bool `json:"he-bss-color,omitempty"`  // 802.11ax BSS color feature support
}

// AmpduEntries represents AMPDU configuration entries container.
type AmpduEntries struct {
	AmpduEntry []AmpduEntry `json:"ampdu-entry"`
}

// AmpduEntry represents individual AMPDU configuration entry.
type AmpduEntry struct {
	Index                    int    `json:"index"`                        // AMPDU entry index identifier
	Apf80211nAmpduTxPriority string `json:"apf-80211n-ampdu-tx-priority"` // AMPDU transmission priority setting
}

// AmpduTxScheduler represents AMPDU transmission scheduler configuration.
type AmpduTxScheduler struct{}

// AmsduEntries represents AMSDU configuration entries container.
type AmsduEntries struct {
	AmsduEntry []AmsduEntry `json:"amsdu-entry"`
}

// AmsduEntry represents individual AMSDU configuration entry.
type AmsduEntry struct {
	Index                    int    `json:"index"`                        // AMSDU entry index identifier
	Apf80211nAmsduTxPriority string `json:"apf-80211n-amsdu-tx-priority"` // AMSDU transmission priority setting
}

// MediaStreamBandInfoCfg represents media stream band information configuration.
type MediaStreamBandInfoCfg struct{}

// SpectrumCfg represents radio spectrum configuration.
type SpectrumCfg struct {
	RrmEdEnable *bool `json:"rrm-ed-enable,omitempty"` // Radio resource management energy detection enable
}

// Dot11axMcsEntries represents 802.11ax MCS entries container.
type Dot11axMcsEntries struct {
	Dot11axMcsEntry []Dot11axMcsEntry `json:"dot11ax-mcs-entry"`
}

// Dot11axMcsEntry represents individual 802.11ax MCS configuration entry.
type Dot11axMcsEntry struct {
	SpatialStream int    `json:"spatial-stream"` // Spatial stream count for MCS configuration
	Index         string `json:"index"`          // MCS index identifier
}

// Dot11acMcsEntries represents 802.11ac MCS entries container.
type Dot11acMcsEntries struct {
	Dot11acMcsEntry []Dot11acMcsEntry `json:"dot11ac-mcs-entry"`
}

// Dot11acMcsEntry represents individual 802.11ac MCS configuration entry.
type Dot11acMcsEntry struct {
	SpatialStream int    `json:"spatial-stream"` // Spatial stream count for MCS configuration
	Index         string `json:"index"`          // MCS index identifier
}
