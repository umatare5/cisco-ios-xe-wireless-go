package model

// Dot15CfgResponse represents the response structure for 802.15 configuration data.
type Dot15CfgResponse struct {
	Dot15CfgData Dot15CfgData `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"`
}

// Dot15CfgData contains 802.15 configuration data
type Dot15CfgData struct {
	Dot15Config Dot15Config `json:"dot15-config"`
}

// Dot15Config represents 802.15 configuration settings
type Dot15Config struct {
	Enabled       bool   `json:"enabled"`
	Channel       int    `json:"channel"`
	TxPower       int    `json:"tx-power"`
	NetworkID     string `json:"network-id"`
	SecurityLevel int    `json:"security-level"`
}
