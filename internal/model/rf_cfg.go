package model

// RfCfgResponse represents the response structure for RF configuration data.
type RfCfgResponse struct {
	RfCfgData RfCfgData `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"`
}

// RfCfgData contains RF configuration data
type RfCfgData struct {
	RfProfiles []RfProfile `json:"rf-profiles"`
}

// RfProfile represents RF profile configuration
type RfProfile struct {
	ProfileName           string `json:"profile-name"`
	MinTxPower            int    `json:"min-tx-power"`
	MaxTxPower            int    `json:"max-tx-power"`
	DataRates             []int  `json:"data-rates"`
	Beamforming           bool   `json:"beamforming"`
	BandSelectEnabled     bool   `json:"band-select-enabled"`
	CoverageHoleDetection bool   `json:"coverage-hole-detection"`
}
