// Package model provides data models for location configuration data.
package model

// LocationCfg  represents the structure for Location configuration data.
type LocationCfg struct {
	LocationCfgData LocationCfgData `json:"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"`
}

// LocationCfgLocationConfig  represents the corresponding data structure.
type LocationCfgLocationConfig struct {
	LocationConfig LocationConfig `json:"Cisco-IOS-XE-wireless-location-cfg:location-config"`
}

type LocationCfgData struct {
	LocationConfig LocationConfig `json:"location-config"`
}

type LocationConfig struct {
	RfidEnabled           bool   `json:"rfid-enabled"`
	Playout               bool   `json:"playout"`
	CalibrationModel      string `json:"calibration-model"`
	BluetoothEnabled      bool   `json:"bluetooth-enabled"`
	RogueLocation         bool   `json:"rogue-location"`
	NotificationThreshold int    `json:"notification-threshold"`
	HyperlocationEnabled  bool   `json:"hyperlocation-enabled"`
}
