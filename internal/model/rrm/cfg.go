package rrm

// RrmCfg represents RRM configuration response data.
type RrmCfg struct {
	CiscoIOSXEWirelessRrmCfgRrmCfgData RrmCfgData `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"` // All fields related rrm feature (Live: IOS-XE 17.12.5)
}

// RrmCfgData represents RRM configuration data container.
type RrmCfgData struct {
	Rrms             *Rrms             `json:"rrms,omitempty"`                // RRM configuration (Live: IOS-XE 17.12.5)
	RrmMgrCfgEntries *RrmMgrCfgEntries `json:"rrm-mgr-cfg-entries,omitempty"` // Configuration related to RRM Algorithms (Live: IOS-XE 17.12.5)
}

// Rrms represents RRM configurations by band.
type Rrms struct {
	Rrm []RrmByBand `json:"rrm,omitempty"` // All rrm grouping algorithm related configurations (Live: IOS-XE 17.12.5)
}

// RrmByBand represents RRM configuration for a specific band.
type RrmByBand struct {
	Band string     `json:"band"`          // Key to st_rrm table, indicates band of configurations (Live: IOS-XE 17.12.5)
	Rrm  *RrmConfig `json:"rrm,omitempty"` // All the basic rrm algorithms configurations (Live: IOS-XE 17.12.5)
}

// RrmConfig represents RRM configuration settings.
type RrmConfig struct {
	RoamingEn           bool   `json:"roaming-en"`                     // Optimized roaming mode enable/disable (Live: IOS-XE 17.12.5)
	DataRateThreshold   string `json:"data-rate-threshold"`            // Data rate threshold for 802.11 Optimized Roaming (Live: IOS-XE 17.12.5)
	MeasurementInterval *int   `json:"measurement-interval,omitempty"` // How often signal strength measurements at each AP (Live: IOS-XE 17.12.5)
}

// RrmMgrCfgEntries represents RRM manager configuration entries.
type RrmMgrCfgEntries struct {
	RrmMgrCfgEntry []RrmMgrCfgEntry `json:"rrm-mgr-cfg-entry,omitempty"` // A list of entries related to RRM Algorithms (Live: IOS-XE 17.12.5)
}

// RrmMgrCfgEntry represents a single RRM manager configuration entry.
type RrmMgrCfgEntry struct {
	Band string `json:"band"` // Key to st_rrm_mgr table, indicates band of configurations (Live: IOS-XE 17.12.5)
}
