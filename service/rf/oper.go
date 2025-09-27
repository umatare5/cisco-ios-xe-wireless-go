package rf

// CiscoIOSXEWirelessRFOper represents RF operational data response structure.
type CiscoIOSXEWirelessRFOper struct {
	CiscoIOSXEWirelessRFOperData struct {
		ApAutoRFDot11Data *ApAutoRFDot11Data `json:"ap-auto-rf-dot11-data,omitempty"` // Auto RF data for 802.11 radios (Live: IOS-XE 17.12.6a)
		ApDot11RadarData  *ApDot11RadarData  `json:"ap-dot11-radar-data,omitempty"`   // Radar detection data for 802.11 radios (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"` // RRM operational data container (Live: IOS-XE 17.12.6a)
}

// ApAutoRFDot11Data represents auto RF data collection for 802.11 radios.
type ApAutoRFDot11Data struct {
	ApAutoRFList []ApAutoRF `json:"ap-auto-rf-dot11-data"` // List of auto RF data for AP radios (Live: IOS-XE 17.12.6a)
}

// ApAutoRF represents auto RF information for specific AP radio.
type ApAutoRF struct {
	WtpMAC            string             `json:"wtp-mac"`                       // Access point MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID       int                `json:"radio-slot-id"`                 // Radio slot identifier (Live: IOS-XE 17.12.6a)
	NeighborRadioInfo *NeighborRadioInfo `json:"neighbor-radio-info,omitempty"` // Neighbor radio information data (Live: IOS-XE 17.12.6a)
}

// NeighborRadioInfo represents neighbor radio information collection.
type NeighborRadioInfo struct {
	NeighborRadioList []NeighborRadioList `json:"neighbor-radio-list"` // List of neighbor radio entries (Live: IOS-XE 17.12.6a)
}

// NeighborRadioList represents neighbor radio entry container.
type NeighborRadioList struct {
	NeighborRadioInfo *NeighborRadioDetail `json:"neighbor-radio-info,omitempty"` // Neighbor radio detail data (Live: IOS-XE 17.12.6a)
}

// NeighborRadioDetail represents detailed neighbor radio information.
type NeighborRadioDetail struct {
	NeighborRadioMAC    string `json:"neighbor-radio-mac"`     // Neighbor radio MAC address (Live: IOS-XE 17.12.6a)
	NeighborRadioSlotID int    `json:"neighbor-radio-slot-id"` // Neighbor radio slot identifier (Live: IOS-XE 17.12.6a)
	RSSI                int    `json:"rssi"`                   // Received Signal Strength in dBm (Live: IOS-XE 17.12.6a)
	SNR                 int    `json:"snr"`                    // Signal to Noise Ratio in dB (Live: IOS-XE 17.12.6a)
	Channel             int    `json:"channel"`                // Operating channel number (Live: IOS-XE 17.12.6a)
	Power               int    `json:"power"`                  // Transmit power in dBm (Live: IOS-XE 17.12.6a)
	GroupLeaderIP       string `json:"group-leader-ip"`        // RRM group leader IP address (Live: IOS-XE 17.12.6a)
	ChanWidth           string `json:"chan-width"`             // Channel width configuration (Live: IOS-XE 17.12.6a)
	SensorCovered       bool   `json:"sensor-covered"`         // Sensor coverage status flag (Live: IOS-XE 17.12.6a)
}

// ApDot11RadarData represents radar detection data collection for 802.11 radios.
type ApDot11RadarData struct {
	ApRadarList []ApRadar `json:"ap-dot11-radar-data"` // List of radar detection data for APs (Live: IOS-XE 17.12.6a)
}

// ApRadar represents radar detection information for specific AP radio.
type ApRadar struct {
	WtpMAC           string `json:"wtp-mac"`             // Access point MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID      int    `json:"radio-slot-id"`       // Radio slot identifier (Live: IOS-XE 17.12.6a)
	LastRadarOnRadio string `json:"last-radar-on-radio"` // Last radar detection timestamp (Live: IOS-XE 17.12.6a)
}
