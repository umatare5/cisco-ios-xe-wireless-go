// Package rf provides data models for RF operational data.
package rf

// RfOper represents RF operational data response structure
// RF operational data is accessed through RRM (Radio Resource Management) endpoints.
type RfOper struct {
	RfOperData RfOperData `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"` // RF operational data from RRM
}

// RfOperResponse represents RF operational data endpoint response structure.
type RfOperResponse struct {
	RfOperData RfOperData `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"` // RF operational data from RRM
}

// RfOperData represents complete RF operational data container
// RF operational data is provided through RRM subsystem.
type RfOperData struct {
	ApAutoRfDot11Data *ApAutoRfDot11Data `json:"ap-auto-rf-dot11-data,omitempty"` // Auto RF data for 802.11 radios
	ApDot11RadarData  *ApDot11RadarData  `json:"ap-dot11-radar-data,omitempty"`   // Radar detection data for 802.11 radios
}

// ApAutoRfDot11Data represents auto RF data collection for 802.11 radios.
type ApAutoRfDot11Data struct {
	ApAutoRfList []ApAutoRf `json:"ap-auto-rf-dot11-data"` // List of auto RF data for AP radios
}

// ApAutoRf represents auto RF information for specific AP radio.
type ApAutoRf struct {
	WtpMac            string             `json:"wtp-mac"`                       // Access point MAC address
	RadioSlotID       int                `json:"radio-slot-id"`                 // Radio slot identifier
	NeighborRadioInfo *NeighborRadioInfo `json:"neighbor-radio-info,omitempty"` // Neighbor radio information
}

// NeighborRadioInfo represents neighbor radio information collection.
type NeighborRadioInfo struct {
	NeighborRadioList []NeighborRadioList `json:"neighbor-radio-list"`
}

// NeighborRadioList represents neighbor radio entry container.
type NeighborRadioList struct {
	NeighborRadioInfo *NeighborRadioDetail `json:"neighbor-radio-info,omitempty"`
}

// NeighborRadioDetail represents detailed neighbor radio information.
type NeighborRadioDetail struct {
	NeighborRadioMac    string `json:"neighbor-radio-mac"`     // Neighbor radio MAC address
	NeighborRadioSlotID int    `json:"neighbor-radio-slot-id"` // Neighbor radio slot identifier
	Rssi                int    `json:"rssi"`                   // Received Signal Strength Indicator in dBm
	Snr                 int    `json:"snr"`                    // Signal to Noise Ratio in dB
	Channel             int    `json:"channel"`                // Radio channel number
	Power               int    `json:"power"`                  // Transmit power in dBm
	GroupLeaderIP       string `json:"group-leader-ip"`        // Group leader IP address
	ChanWidth           string `json:"chan-width"`             // Channel width setting
	SensorCovered       bool   `json:"sensor-covered"`         // Sensor coverage status flag
}

// ApDot11RadarData represents radar detection data collection for 802.11 radios.
type ApDot11RadarData struct {
	ApRadarList []ApRadar `json:"ap-dot11-radar-data"` // List of radar detection data for AP radios
}

// ApRadar represents radar detection information for specific AP radio.
type ApRadar struct {
	WtpMac           string `json:"wtp-mac"`             // Access point MAC address
	RadioSlotID      int    `json:"radio-slot-id"`       // Radio slot identifier
	LastRadarOnRadio string `json:"last-radar-on-radio"` // Last radar detection timestamp
}
