package rf

// RFOper represents RF operational data response structure.
type RFOper struct {
	RFOperData struct {
		ApAutoRFDot11Data *ApAutoRFDot11Data `json:"ap-auto-rf-dot11-data,omitempty"` // Auto RF data for 802.11 radios (Live: IOS-XE 17.12.5)
		ApDot11RadarData  *ApDot11RadarData  `json:"ap-dot11-radar-data,omitempty"`   // Radar detection data for 802.11 radios (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"` // RRM operational data container (Live: IOS-XE 17.12.5)
}

// ApAutoRFDot11Data represents auto RF data collection for 802.11 radios.
type ApAutoRFDot11Data struct {
	ApAutoRFList []ApAutoRF `json:"ap-auto-rf-dot11-data"` // List of auto RF data for AP radios (Live: IOS-XE 17.12.5)
}

// ApAutoRF represents auto RF information for specific AP radio.
type ApAutoRF struct {
	WtpMAC            string             `json:"wtp-mac"`                       // Access point MAC address (Live: IOS-XE 17.12.5)
	RadioSlotID       int                `json:"radio-slot-id"`                 // Radio slot identifier (Live: IOS-XE 17.12.5)
	NeighborRadioInfo *NeighborRadioInfo `json:"neighbor-radio-info,omitempty"` // Neighbor radio information data (Live: IOS-XE 17.12.5)
}

// NeighborRadioInfo represents neighbor radio information collection.
type NeighborRadioInfo struct {
	NeighborRadioList []NeighborRadioList `json:"neighbor-radio-list"` // List of neighbor radio entries (Live: IOS-XE 17.12.5)
}

// NeighborRadioList represents neighbor radio entry container.
type NeighborRadioList struct {
	NeighborRadioInfo *NeighborRadioDetail `json:"neighbor-radio-info,omitempty"` // Neighbor radio detail data (Live: IOS-XE 17.12.5)
}

// NeighborRadioDetail represents detailed neighbor radio information.
type NeighborRadioDetail struct {
	NeighborRadioMAC    string `json:"neighbor-radio-mac"`     // Neighbor radio MAC address (Live: IOS-XE 17.12.5)
	NeighborRadioSlotID int    `json:"neighbor-radio-slot-id"` // Neighbor radio slot identifier (Live: IOS-XE 17.12.5)
	RSSI                int    `json:"rssi"`                   // Received Signal Strength in dBm (Live: IOS-XE 17.12.5)
	SNR                 int    `json:"snr"`                    // Signal to Noise Ratio in dB (Live: IOS-XE 17.12.5)
	Channel             int    `json:"channel"`                // Operating channel number (Live: IOS-XE 17.12.5)
	Power               int    `json:"power"`                  // Transmit power in dBm (Live: IOS-XE 17.12.5)
	GroupLeaderIP       string `json:"group-leader-ip"`        // RRM group leader IP address (Live: IOS-XE 17.12.5)
	ChanWidth           string `json:"chan-width"`             // Channel width configuration (Live: IOS-XE 17.12.5)
	SensorCovered       bool   `json:"sensor-covered"`         // Sensor coverage status flag (Live: IOS-XE 17.12.5)
}

// ApDot11RadarData represents radar detection data collection for 802.11 radios.
type ApDot11RadarData struct {
	ApRadarList []ApRadar `json:"ap-dot11-radar-data"` // List of radar detection data for APs (Live: IOS-XE 17.12.5)
}

// ApRadar represents radar detection information for specific AP radio.
type ApRadar struct {
	WtpMAC           string `json:"wtp-mac"`             // Access point MAC address (Live: IOS-XE 17.12.5)
	RadioSlotID      int    `json:"radio-slot-id"`       // Radio slot identifier (Live: IOS-XE 17.12.5)
	LastRadarOnRadio string `json:"last-radar-on-radio"` // Last radar detection timestamp (Live: IOS-XE 17.12.5)
}
