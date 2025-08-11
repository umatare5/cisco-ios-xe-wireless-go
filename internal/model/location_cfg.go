package model

// LocationCfgResponse represents the response structure for Location configuration data.
type LocationCfgResponse struct {
	LocationCfgData LocationCfgData `json:"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"`
}

// LocationCfgData contains Location configuration data
type LocationCfgData struct {
	LocationConfig LocationConfig `json:"location-config"`
}

// LocationConfig represents location services configuration
type LocationConfig struct {
	RfidEnabled           bool   `json:"rfid-enabled"`
	Playout               bool   `json:"playout"`
	CalibrationModel      string `json:"calibration-model"`
	BluetoothEnabled      bool   `json:"bluetooth-enabled"`
	RogueLocation         bool   `json:"rogue-location"`
	NotificationThreshold int    `json:"notification-threshold"`
	HyperlocationEnabled  bool   `json:"hyperlocation-enabled"`
}
