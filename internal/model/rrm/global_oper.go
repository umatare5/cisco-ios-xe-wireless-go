package rrm

import "time"

// RrmGlobalOper represents RRM global operational response data.
type RrmGlobalOper struct {
	RrmGlobalOperData RrmGlobalOperData `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data"` // RRM global operational data container
}

// RrmGlobalOperData represents RRM global operational data container.
type RrmGlobalOperData struct {
	RrmOneShotCounters     []RrmOneShotCounter      `json:"rrm-one-shot-counters,omitempty"`     // RRM one-shot counters
	RrmChannelParams       []RrmChannelParam        `json:"rrm-channel-params,omitempty"`        // RRM channel parameters
	RadioOperData24g       []RadioOperData24g       `json:"radio-oper-data-24g,omitempty"`       // 2.4GHz radio operational data
	RadioOperData5g        []RadioOperData5g        `json:"radio-oper-data-5g,omitempty"`        // 5GHz radio operational data
	RadioOperData6ghz      []RadioOperData6ghz      `json:"radio-oper-data-6ghz,omitempty"`      // 6GHz radio operational data
	RadioOperDataDualband  []RadioOperDataDualband  `json:"radio-oper-data-dualband,omitempty"`  // Dual band radio operational data
	SpectrumBandConfigData []SpectrumBandConfigData `json:"spectrum-band-config-data,omitempty"` // Spectrum band configuration data
	RrmClientData          []RrmClientData          `json:"rrm-client-data,omitempty"`           // RRM client data
	RrmFraStats            *RrmFraStats             `json:"rrm-fra-stats,omitempty"`             // Flexible Radio Assignment statistics
	RrmCoverage            []RrmCoverage            `json:"rrm-coverage,omitempty"`              // RRM coverage data
	SpectrumAqWorstTable   []SpectrumAqWorstTable   `json:"spectrum-aq-worst-table,omitempty"`   // Spectrum Air Quality worst interference table
}

// RrmOneShotCounter represents one-shot counter data.
type RrmOneShotCounter struct {
	PhyType      string `json:"phy-type"`      // PHY type identifier
	PowerCounter int    `json:"power-counter"` // Power counter value
}

// RrmChannelParam represents channel parameter data.
type RrmChannelParam struct {
	PhyType        string `json:"phy-type"`        // PHY type identifier
	MinDwell       int    `json:"min-dwell"`       // Minimum dwell time
	AvgDwell       int    `json:"avg-dwell"`       // Average dwell time
	MaxDwell       int    `json:"max-dwell"`       // Maximum dwell time
	MinRssi        int    `json:"min-rssi"`        // Minimum RSSI value
	MaxRssi        int    `json:"max-rssi"`        // Maximum RSSI value
	AvgRssi        int    `json:"avg-rssi"`        // Average RSSI value
	ChannelCounter int    `json:"channel-counter"` // Channel counter
}

// RadioOperData24g represents 2.4GHz radio operational data.
type RadioOperData24g struct {
	WtpMac          string    `json:"wtp-mac"`                    // Access point MAC address
	RadioSlotID     int       `json:"radio-slot-id"`              // Radio slot identifier
	ApMac           string    `json:"ap-mac"`                     // Access point MAC address
	SlotID          int       `json:"slot-id"`                    // Slot identifier
	Name            string    `json:"name"`                       // Radio name
	SpectrumCapable []any     `json:"spectrum-capable,omitempty"` // Spectrum capability list
	NumSlots        int       `json:"num-slots"`                  // Number of slots
	MeshRadioRole   string    `json:"mesh-radio-role"`            // Mesh radio role
	ApUpTime        time.Time `json:"ap-up-time"`                 // Access point uptime
	CapwapUpTime    time.Time `json:"capwap-up-time"`             // CAPWAP uptime
}

// RadioOperData5g represents 5GHz radio operational data.
type RadioOperData5g struct {
	WtpMac          string    `json:"wtp-mac"`                    // Access point MAC address
	RadioSlotID     int       `json:"radio-slot-id"`              // Radio slot identifier
	ApMac           string    `json:"ap-mac"`                     // Access point MAC address
	SlotID          int       `json:"slot-id"`                    // Slot identifier
	Name            string    `json:"name"`                       // Radio name
	SpectrumCapable []any     `json:"spectrum-capable,omitempty"` // Spectrum capability list
	NumSlots        int       `json:"num-slots"`                  // Number of slots
	MeshRadioRole   string    `json:"mesh-radio-role"`            // Mesh radio role
	ApUpTime        time.Time `json:"ap-up-time"`                 // Access point uptime
	CapwapUpTime    time.Time `json:"capwap-up-time"`             // CAPWAP uptime
}

// RadioOperData6ghz represents 6GHz radio operational data.
type RadioOperData6ghz struct {
	WtpMac          string    `json:"wtp-mac"`                    // Access point MAC address
	RadioSlotID     int       `json:"radio-slot-id"`              // Radio slot identifier
	ApMac           string    `json:"ap-mac"`                     // Access point MAC address
	SlotID          int       `json:"slot-id"`                    // Slot identifier
	Name            string    `json:"name"`                       // Radio name
	SpectrumCapable []any     `json:"spectrum-capable,omitempty"` // Spectrum capability list
	NumSlots        int       `json:"num-slots"`                  // Number of slots
	MeshRadioRole   string    `json:"mesh-radio-role"`            // Mesh radio role
	ApUpTime        time.Time `json:"ap-up-time"`                 // Access point uptime
	CapwapUpTime    time.Time `json:"capwap-up-time"`             // CAPWAP uptime
}

// RadioOperDataDualband represents dual band radio operational data.
type RadioOperDataDualband struct {
	WtpMac          string    `json:"wtp-mac"`                    // Access point MAC address
	RadioSlotID     int       `json:"radio-slot-id"`              // Radio slot identifier
	ApMac           string    `json:"ap-mac"`                     // Access point MAC address
	SlotID          int       `json:"slot-id"`                    // Slot identifier
	Name            string    `json:"name"`                       // Radio name
	SpectrumCapable []any     `json:"spectrum-capable,omitempty"` // Spectrum capability list
	NumSlots        int       `json:"num-slots"`                  // Number of slots
	MeshRadioRole   string    `json:"mesh-radio-role"`            // Mesh radio role
	ApUpTime        time.Time `json:"ap-up-time"`                 // Access point uptime
	CapwapUpTime    time.Time `json:"capwap-up-time"`             // CAPWAP uptime
}

// SpectrumBandConfigData represents spectrum band configuration data.
type SpectrumBandConfigData struct {
	ApMac              string               `json:"ap-mac"`                         // Access point MAC address
	SpectrumBandConfig []SpectrumBandConfig `json:"spectrum-band-config,omitempty"` // Spectrum band configurations
}

// SpectrumBandConfig represents spectrum band configuration for a specific band.
type SpectrumBandConfig struct {
	BandID             string `json:"band-id"`              // Band identifier
	SpectrumAdminState bool   `json:"spectrum-admin-state"` // Spectrum administrative state
}

// RrmClientData represents RRM client data.
type RrmClientData struct {
	PhyType         string    `json:"phy-type"`        // PHY type identifier
	LastChdRun      time.Time `json:"last-chd-run"`    // Last CHD run timestamp
	Disassociations int       `json:"disassociations"` // Number of disassociations
	Rejections      int       `json:"rejections"`      // Number of rejections
}

// RrmFraStats represents Flexible Radio Assignment statistics.
type RrmFraStats struct {
	DualBandMonitorTo24ghz int `json:"dual-band-monitor-to-24ghz"` // Dual band monitor to 2.4GHz transitions
	DualBandMonitorTo5ghz  int `json:"dual-band-monitor-to-5ghz"`  // Dual band monitor to 5GHz transitions
	DualBand24ghzTo5ghz    int `json:"dual-band-24ghz-to-5ghz"`    // Dual band 2.4GHz to 5GHz transitions
	DualBand24ghzToMonitor int `json:"dual-band-24ghz-to-monitor"` // Dual band 2.4GHz to monitor transitions
	DualBand5ghzTo24ghz    int `json:"dual-band-5ghz-to-24ghz"`    // Dual band 5GHz to 2.4GHz transitions
	DualBand5ghzToMonitor  int `json:"dual-band-5ghz-to-monitor"`  // Dual band 5GHz to monitor transitions
	SecRadioMonitorTo5ghz  int `json:"sec-radio-monitor-to-5ghz"`  // Secondary radio monitor to 5GHz transitions
	SecRadio5ghzToMonitor  int `json:"sec-radio-5ghz-to-monitor"`  // Secondary radio 5GHz to monitor transitions
	DualBand6ghzTo5ghz     int `json:"dual-band-6ghz-to-5ghz"`     // Dual band 6GHz to 5GHz transitions
	DualBand5ghzTo6ghz     int `json:"dual-band-5ghz-to-6ghz"`     // Dual band 5GHz to 6GHz transitions
}

// RrmCoverage represents RRM coverage data.
type RrmCoverage struct {
	WtpMac            string     `json:"wtp-mac"`             // Access point MAC address
	RadioSlotID       int        `json:"radio-slot-id"`       // Radio slot identifier
	FailedClientCount int        `json:"failed-client-count"` // Number of failed clients
	SNRInfo           []SNRInfo  `json:"snr-info,omitempty"`  // Signal-to-Noise Ratio information
	RSSIInfo          []RSSIInfo `json:"rssi-info,omitempty"` // Received Signal Strength Indicator information
}

// SNRInfo represents Signal-to-Noise Ratio information.
type SNRInfo struct {
	SNR        int `json:"snr"`         // Signal-to-Noise Ratio in dB
	NumClients int `json:"num-clients"` // Number of clients at this SNR level
}

// RSSIInfo represents Received Signal Strength Indicator information.
type RSSIInfo struct {
	RSSI       int `json:"rssi"`        // Received Signal Strength Indicator in dBm
	NumClients int `json:"num-clients"` // Number of clients at this RSSI level
}

// SpectrumAqWorstTable represents spectrum Air Quality worst interference table entry.
type SpectrumAqWorstTable struct {
	BandID               int    `json:"band-id"`                 // Band identifier
	DetectingApName      string `json:"detecting-ap-name"`       // Name of AP detecting interference
	ChannelNum           int    `json:"channel-num"`             // Channel number
	MinAqi               int    `json:"min-aqi"`                 // Minimum Air Quality Index
	Aqi                  int    `json:"aqi"`                     // Current Air Quality Index
	TotalIntfDeviceCount int    `json:"total-intf-device-count"` // Total interference device count
	WtpCaSiCapable       string `json:"wtp-ca-si-capable"`       // WTP spectrum intelligence capability
	ScanRadioType        string `json:"scan-radio-type"`         // Scanning radio type
}
