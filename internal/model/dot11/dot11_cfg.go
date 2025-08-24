// Package model provides data models for 802.11 configuration data.
package model

// Dot11Cfg  represents the 802.11 configuration data.
type Dot11Cfg struct {
	Dot11CfgData Dot11CfgData `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data"`
}

// Dot11CfgConfiguredCountries  represents the configured countries.
type Dot11CfgConfiguredCountries struct {
	ConfiguredCountries *Dot11ConfiguredCountries `json:"Cisco-IOS-XE-wireless-dot11-cfg:configured-countries,omitempty"`
}

// Dot11CfgDot11Entries  represents the 802.11 entries.
type Dot11CfgDot11Entries struct {
	Dot11Entries *Dot11Entries `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-entries,omitempty"`
}

// Dot11CfgDot11acMcsEntries  represents the 802.11ac MCS entries.
type Dot11CfgDot11acMcsEntries struct {
	Dot11acMcsEntries *Dot11acMcsEntries `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11ac-mcs-entries,omitempty"`
}

// Dot11CfgFilter  represents the filtered 802.11 configuration data.
type Dot11CfgFilter struct {
	ConfiguredCountry []Dot11ConfiguredCountry `json:"Cisco-IOS-XE-wireless-dot11-cfg:configured-country,omitempty"`
	Dot11Entry        []Dot11Entry             `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11-entry,omitempty"`
	Dot11acMcsEntry   []Dot11acMcsEntry        `json:"Cisco-IOS-XE-wireless-dot11-cfg:dot11ac-mcs-entry,omitempty"`
}

type Dot11CfgData struct {
	ConfiguredCountries *Dot11ConfiguredCountries `json:"configured-countries,omitempty"`
	Dot11Entries        *Dot11Entries             `json:"dot11-entries,omitempty"`
	Dot11acMcsEntries   *Dot11acMcsEntries        `json:"dot11ac-mcs-entries,omitempty"`
}

type Dot11ConfiguredCountries struct {
	ConfiguredCountry []Dot11ConfiguredCountry `json:"configured-country"`
}

type Dot11ConfiguredCountry struct {
	CountryCode string `json:"country-code"`
}

type Dot11Entries struct {
	Dot11Entry []Dot11Entry `json:"dot11-entry"`
}

type Dot11Entry struct {
	Band                  string        `json:"band"`
	VoiceAdmCtrlSupport   *bool         `json:"voice-adm-ctrl-support,omitempty"`
	Dot11axCfg            *Dot11axCfg   `json:"dot11ax-cfg,omitempty"`
	AmpduEntries          *AmpduEntries `json:"ampdu-entries,omitempty"`
	AmsduxMaxSubframes    *int          `json:"amsdux-max-subframes,omitempty"`
	AmsduxMaxLength       *int          `json:"amsdux-max-length,omitempty"`
	Dot11HT40Support      *bool         `json:"dot11-ht40-support,omitempty"`
	Dot11HT40Intolerant   *bool         `json:"dot11-ht40-intolerant,omitempty"`
	Dot11HT40Above        *bool         `json:"dot11-ht40-above,omitempty"`
	Dot11HT40Below        *bool         `json:"dot11-ht40-below,omitempty"`
	OptimalLoadBalancing  *bool         `json:"optimal-load-balancing,omitempty"`
	BandSelectEnabled     *bool         `json:"band-select-enabled,omitempty"`
	BandSelectCycleCount  *int          `json:"band-select-cycle-count,omitempty"`
	BandSelectCycleThresh *int          `json:"band-select-cycle-thresh,omitempty"`
	BandSelectExpiredTime *int          `json:"band-select-expired-time,omitempty"`
}

type Dot11axCfg struct {
	HeBssColor *bool `json:"he-bss-color,omitempty"`
}

type AmpduEntries struct {
	AmpduEntry []AmpduEntry `json:"ampdu-entry"`
}

type AmpduEntry struct {
	Index                    int    `json:"index"`
	Apf80211nAmpduTxPriority string `json:"apf-80211n-ampdu-tx-priority"`
}

type Dot11acMcsEntries struct {
	Dot11acMcsEntry []Dot11acMcsEntry `json:"dot11ac-mcs-entry"`
}

type Dot11acMcsEntry struct {
	SpatialStream int    `json:"spatial-stream"`
	Index         string `json:"index"`
}
