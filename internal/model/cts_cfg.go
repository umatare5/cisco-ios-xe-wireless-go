package model

// CtsCfgResponse represents the response structure for CTS configuration data.
type CtsCfgResponse struct {
	CtsCfgData CtsCfgData `json:"Cisco-IOS-XE-wireless-cts-cfg:cts-cfg-data"`
}

// CtsCfgData contains CTS configuration data
type CtsCfgData struct {
	CtsConfig CtsConfig `json:"cts-config"`
}

// CtsConfig represents CTS configuration settings
type CtsConfig struct {
	SxpEnabled         bool   `json:"sxp-enabled"`
	SxpDefaultPassword string `json:"sxp-default-password"`
	SgtPropagation     bool   `json:"sgt-propagation"`
	SxpSpeaker         bool   `json:"sxp-speaker"`
	SxpListener        bool   `json:"sxp-listener"`
}
