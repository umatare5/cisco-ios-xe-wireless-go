package rrm

// RrmCfg represents RRM configuration response data from WNC 17.12.5.
type RrmCfg struct {
	CiscoIOSXEWirelessRrmCfgRrmCfgData RrmCfgData `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"` // RRM configuration data container
}

// RrmCfgData represents RRM configuration data container from WNC 17.12.5.
type RrmCfgData struct {
	Rrms             *Rrms             `json:"rrms,omitempty"`                // RRM band configurations
	RrmMgrCfgEntries *RrmMgrCfgEntries `json:"rrm-mgr-cfg-entries,omitempty"` // RRM manager configuration entries
}

// Rrms represents RRM configurations by band.
type Rrms struct {
	Rrm []RrmByBand `json:"rrm,omitempty"` // RRM configurations for each band
}

// RrmByBand represents RRM configuration for a specific band from WNC 17.12.5.
type RrmByBand struct {
	Band string     `json:"band"`          // Radio band identifier
	Rrm  *RrmConfig `json:"rrm,omitempty"` // RRM configuration for this band
}

// RrmConfig represents RRM configuration settings from WNC 17.12.5.
type RrmConfig struct {
	RoamingEn           bool   `json:"roaming-en"`                     // Enable optimized roaming for band
	DataRateThreshold   string `json:"data-rate-threshold"`            // Data rate threshold for optimized roaming
	MeasurementInterval *int   `json:"measurement-interval,omitempty"` // RRM measurement interval in seconds
}

// RrmMgrCfgEntries represents RRM manager configuration entries.
type RrmMgrCfgEntries struct {
	RrmMgrCfgEntry []RrmMgrCfgEntry `json:"rrm-mgr-cfg-entry,omitempty"` // List of RRM manager configuration entries
}

// RrmMgrCfgEntry represents a single RRM manager configuration entry.
type RrmMgrCfgEntry struct {
	Band string `json:"band"` // Radio band for RRM manager configuration
}
