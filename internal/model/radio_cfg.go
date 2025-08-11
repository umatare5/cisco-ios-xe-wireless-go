package model

// RadioCfgResponse represents the response structure for Radio configuration data.
type RadioCfgResponse struct {
	RadioCfgData RadioCfgData `json:"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"`
}

// RadioCfgData contains Radio configuration data
type RadioCfgData struct {
	RadioConfigs []RadioConfig `json:"radio-configs"`
}

// RadioConfig represents radio configuration settings
type RadioConfig struct {
	RadioType         string `json:"radio-type"`
	RadioEnabled      bool   `json:"radio-enabled"`
	TxPowerLevel      int    `json:"tx-power-level"`
	ChannelAssignment string `json:"channel-assignment"`
	ChannelWidth      string `json:"channel-width"`
	AntennaGain       int    `json:"antenna-gain"`
	Beamforming       bool   `json:"beamforming"`
}
