package rrm

// CiscoIOSXEWirelessRRMCfg represents RRM configuration response data.
type CiscoIOSXEWirelessRRMCfg struct {
	RRMCfgData struct {
		Rrms             *RRMs             `json:"rrms,omitempty"`                // RRM configuration (Live: IOS-XE 17.12.5)
		RRMMgrCfgEntries *RRMMgrCfgEntries `json:"rrm-mgr-cfg-entries,omitempty"` // Configuration related to RRM Algorithms (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"` // All fields related rrm feature (Live: IOS-XE 17.12.5)
}

// CiscoIOSXEWirelessRRMCfgRrms wraps the Rrms child struct for direct access (Live: IOS-XE 17.12.5).
type CiscoIOSXEWirelessRRMCfgRrms struct {
	RRMs RRMs `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrms"`
}

// CiscoIOSXEWirelessRRMCfgRRMMgrCfgEntries wraps the RRMMgrCfgEntries child struct for direct access (Live: IOS-XE 17.12.5).
type CiscoIOSXEWirelessRRMCfgRRMMgrCfgEntries struct {
	RRMMgrCfgEntries RRMMgrCfgEntries `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-mgr-cfg-entries"`
}

// RRMs represents RRM configurations by band.
type RRMs struct {
	RRM []RRMByBand `json:"rrm,omitempty"` // All rrm grouping algorithm related configurations (Live: IOS-XE 17.12.5)
}

// RRMByBand represents RRM configuration for a specific band.
type RRMByBand struct {
	Band string     `json:"band"`          // Key to st_rrm table, indicates band of configurations (Live: IOS-XE 17.12.5)
	RRM  *RRMConfig `json:"rrm,omitempty"` // All the basic rrm algorithms configurations (Live: IOS-XE 17.12.5)
}

// RRMConfig represents RRM configuration settings.
type RRMConfig struct {
	RoamingEn           bool   `json:"roaming-en"`                     // Optimized roaming mode enable/disable (Live: IOS-XE 17.12.5)
	DataRateThreshold   string `json:"data-rate-threshold"`            // Data rate threshold for 802.11 Optimized Roaming (Live: IOS-XE 17.12.5)
	MeasurementInterval *int   `json:"measurement-interval,omitempty"` // How often signal strength measurements at each AP (Live: IOS-XE 17.12.5)
}

// RRMMgrCfgEntries represents RRM manager configuration entries.
type RRMMgrCfgEntries struct {
	RRMMgrCfgEntry []RRMMgrCfgEntry `json:"rrm-mgr-cfg-entry,omitempty"` // A list of entries related to RRM Algorithms (Live: IOS-XE 17.12.5)
}

// RRMMgrCfgEntry represents a single RRM manager configuration entry.
type RRMMgrCfgEntry struct {
	Band string `json:"band"` // Key to st_rrm_mgr table, indicates band of configurations (Live: IOS-XE 17.12.5)
}
