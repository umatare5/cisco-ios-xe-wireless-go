package rrm

import "time"

// RRMGlobalOper represents RRM global operational response data.
type RRMGlobalOper struct {
	RRMGlobalOperData struct {
		RRMOneShotCounters     []RRMOneShotCounter      `json:"rrm-one-shot-counters,omitempty"`     // Transmit power and channel update count (Live: IOS-XE 17.12.6a)
		RRMChannelParams       []RRMChannelParam        `json:"rrm-channel-params,omitempty"`        // RRM channel parameter data (Live: IOS-XE 17.12.6a)
		RadioOperData24g       []RadioOperData24g       `json:"radio-oper-data-24g,omitempty"`       // 2.4 ghz radio oper data (Live: IOS-XE 17.12.6a)
		RadioOperData5g        []RadioOperData5g        `json:"radio-oper-data-5g,omitempty"`        // 5 ghz radio oper data (Live: IOS-XE 17.12.6a)
		RadioOperData6ghz      []RadioOperData6ghz      `json:"radio-oper-data-6ghz,omitempty"`      // 6 ghz radio oper data (Live: IOS-XE 17.12.6a)
		RadioOperDataDualband  []RadioOperDataDualband  `json:"radio-oper-data-dualband,omitempty"`  // Dual band radio oper data (Live: IOS-XE 17.12.6a)
		SpectrumBandConfigData []SpectrumBandConfigData `json:"spectrum-band-config-data,omitempty"` // AP spectrum config (Live: IOS-XE 17.12.6a)
		RRMClientData          []RRMClientData          `json:"rrm-client-data,omitempty"`           // RRM client data (Live: IOS-XE 17.12.6a)
		RRMFraStats            *RRMFraStats             `json:"rrm-fra-stats,omitempty"`             // RRM flexible radio statistics (Live: IOS-XE 17.12.6a)
		RRMCoverage            []RRMCoverage            `json:"rrm-coverage,omitempty"`              // Coverage information (Live: IOS-XE 17.12.6a)
		SpectrumAqWorstTable   []SpectrumAqWorstTable   `json:"spectrum-aq-worst-table,omitempty"`   // Air quality index data (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data"` // RRM operational data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessRRMGlobalOperRRMOneShotCounters represents the one-shot counter operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRRMOneShotCounters struct {
	RRMOneShotCounters []RRMOneShotCounter `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-one-shot-counters"`
}

// CiscoIOSXEWirelessRRMGlobalOperRRMChannelParams represents the channel parameter operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRRMChannelParams struct {
	RRMChannelParams []RRMChannelParam `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-channel-params"`
}

// CiscoIOSXEWirelessRRMGlobalOperRadioOperData24g represents the 2.4GHz radio operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRadioOperData24g struct {
	RadioOperData24g []RadioOperData24g `json:"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-24g"`
}

// CiscoIOSXEWirelessRRMGlobalOperRadioOperData5g represents the 5GHz radio operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRadioOperData5g struct {
	RadioOperData5g []RadioOperData5g `json:"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-5g"`
}

// CiscoIOSXEWirelessRRMGlobalOperRadioOperData6ghz represents the 6GHz radio operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRadioOperData6ghz struct {
	RadioOperData6ghz []RadioOperData6ghz `json:"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-6ghz"`
}

// CiscoIOSXEWirelessRRMGlobalOperRadioOperDataDualband represents the dual-band radio operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRadioOperDataDualband struct {
	RadioOperDataDualband []RadioOperDataDualband `json:"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-dualband"`
}

// CiscoIOSXEWirelessRRMGlobalOperSpectrumBandConfigData represents the spectrum band configuration data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperSpectrumBandConfigData struct {
	SpectrumBandConfigData []SpectrumBandConfigData `json:"Cisco-IOS-XE-wireless-rrm-global-oper:spectrum-band-config-data"`
}

// CiscoIOSXEWirelessRRMGlobalOperRRMClientData represents the RRM client operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRRMClientData struct {
	RRMClientData []RRMClientData `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-client-data"`
}

// CiscoIOSXEWirelessRRMGlobalOperRRMFraStats represents the RRM flexible radio assignment statistics (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRRMFraStats struct {
	RRMFraStats *RRMFraStats `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-fra-stats"`
}

// CiscoIOSXEWirelessRRMGlobalOperRRMCoverage represents the RRM coverage operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperRRMCoverage struct {
	RRMCoverage []RRMCoverage `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-coverage"`
}

// CiscoIOSXEWirelessRRMGlobalOperSpectrumAqWorstTable represents the spectrum air quality worst table data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMGlobalOperSpectrumAqWorstTable struct {
	SpectrumAqWorstTable []SpectrumAqWorstTable `json:"Cisco-IOS-XE-wireless-rrm-global-oper:spectrum-aq-worst-table"`
}

// RRMOneShotCounter represents one-shot counter data.
type RRMOneShotCounter struct {
	PhyType      string `json:"phy-type"`      // Radio type (Live: IOS-XE 17.12.6a)
	PowerCounter int    `json:"power-counter"` // Transmit power update count (Live: IOS-XE 17.12.6a)
}

// RRMChannelParam represents channel parameter data.
type RRMChannelParam struct {
	PhyType        string `json:"phy-type"`        // Radio type (Live: IOS-XE 17.12.6a)
	MinDwell       int    `json:"min-dwell"`       // Minimum channel dwell time (Live: IOS-XE 17.12.6a)
	AvgDwell       int    `json:"avg-dwell"`       // Average channel dwell time (Live: IOS-XE 17.12.6a)
	MaxDwell       int    `json:"max-dwell"`       // Maximum channel dwell time (Live: IOS-XE 17.12.6a)
	MinRSSI        int    `json:"min-rssi"`        // Minimum channel energy level (Live: IOS-XE 17.12.6a)
	MaxRSSI        int    `json:"max-rssi"`        // Maximum channel energy level (Live: IOS-XE 17.12.6a)
	AvgRSSI        int    `json:"avg-rssi"`        // Average channel energy level (Live: IOS-XE 17.12.6a)
	ChannelCounter int    `json:"channel-counter"` // Channel Update Count (Live: IOS-XE 17.12.6a)
}

// RadioOperData24g represents 2.4GHz radio operational data.
type RadioOperData24g struct {
	WtpMAC          string    `json:"wtp-mac"`                    // MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID     int       `json:"radio-slot-id"`              // Radio slot identifier (Live: IOS-XE 17.12.6a)
	ApMAC           string    `json:"ap-mac"`                     // MAC address (Live: IOS-XE 17.12.6a)
	SlotID          int       `json:"slot-id"`                    // Slot identifier (Live: IOS-XE 17.12.6a)
	Name            string    `json:"name"`                       // WTP name (Live: IOS-XE 17.12.6a)
	SpectrumCapable []any     `json:"spectrum-capable,omitempty"` // AP is cleanair capable or not (Live: IOS-XE 17.12.6a)
	NumSlots        int       `json:"num-slots"`                  // Number of slots (Live: IOS-XE 17.12.6a)
	MeshRadioRole   string    `json:"mesh-radio-role"`            // Radio Role (Live: IOS-XE 17.12.6a)
	ApUpTime        time.Time `json:"ap-up-time"`                 // AP up time (Live: IOS-XE 17.12.6a)
	CAPWAPUpTime    time.Time `json:"capwap-up-time"`             // CAPWAP uptime (Live: IOS-XE 17.12.6a)
}

// RadioOperData5g represents 5GHz radio operational data.
type RadioOperData5g struct {
	WtpMAC          string    `json:"wtp-mac"`                    // MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID     int       `json:"radio-slot-id"`              // Radio slot identifier (Live: IOS-XE 17.12.6a)
	ApMAC           string    `json:"ap-mac"`                     // MAC address (Live: IOS-XE 17.12.6a)
	SlotID          int       `json:"slot-id"`                    // Slot identifier (Live: IOS-XE 17.12.6a)
	Name            string    `json:"name"`                       // WTP name (Live: IOS-XE 17.12.6a)
	SpectrumCapable []any     `json:"spectrum-capable,omitempty"` // AP is cleanair capable or not (Live: IOS-XE 17.12.6a)
	NumSlots        int       `json:"num-slots"`                  // Number of slots (Live: IOS-XE 17.12.6a)
	MeshRadioRole   string    `json:"mesh-radio-role"`            // Radio Role (Live: IOS-XE 17.12.6a)
	ApUpTime        time.Time `json:"ap-up-time"`                 // AP up time (Live: IOS-XE 17.12.6a)
	CAPWAPUpTime    time.Time `json:"capwap-up-time"`             // CAPWAP uptime (Live: IOS-XE 17.12.6a)
}

// RadioOperData6ghz represents 6GHz radio operational data.
type RadioOperData6ghz struct {
	WtpMAC          string    `json:"wtp-mac"`                    // MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID     int       `json:"radio-slot-id"`              // Radio slot identifier (Live: IOS-XE 17.12.6a)
	ApMAC           string    `json:"ap-mac"`                     // MAC address (Live: IOS-XE 17.12.6a)
	SlotID          int       `json:"slot-id"`                    // Slot identifier (Live: IOS-XE 17.12.6a)
	Name            string    `json:"name"`                       // WTP name (Live: IOS-XE 17.12.6a)
	SpectrumCapable []any     `json:"spectrum-capable,omitempty"` // AP is cleanair capable or not (Live: IOS-XE 17.12.6a)
	NumSlots        int       `json:"num-slots"`                  // Number of slots (Live: IOS-XE 17.12.6a)
	MeshRadioRole   string    `json:"mesh-radio-role"`            // Radio Role (Live: IOS-XE 17.12.6a)
	ApUpTime        time.Time `json:"ap-up-time"`                 // AP up time (Live: IOS-XE 17.12.6a)
	CAPWAPUpTime    time.Time `json:"capwap-up-time"`             // CAPWAP uptime (Live: IOS-XE 17.12.6a)
}

// RadioOperDataDualband represents dual band radio operational data.
type RadioOperDataDualband struct {
	WtpMAC          string    `json:"wtp-mac"`                    // MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID     int       `json:"radio-slot-id"`              // Radio slot identifier (Live: IOS-XE 17.12.6a)
	ApMAC           string    `json:"ap-mac"`                     // MAC address (Live: IOS-XE 17.12.6a)
	SlotID          int       `json:"slot-id"`                    // Slot identifier (Live: IOS-XE 17.12.6a)
	Name            string    `json:"name"`                       // WTP name (Live: IOS-XE 17.12.6a)
	SpectrumCapable []any     `json:"spectrum-capable,omitempty"` // AP is cleanair capable or not (Live: IOS-XE 17.12.6a)
	NumSlots        int       `json:"num-slots"`                  // Number of slots (Live: IOS-XE 17.12.6a)
	MeshRadioRole   string    `json:"mesh-radio-role"`            // Radio Role (Live: IOS-XE 17.12.6a)
	ApUpTime        time.Time `json:"ap-up-time"`                 // AP up time (Live: IOS-XE 17.12.6a)
	CAPWAPUpTime    time.Time `json:"capwap-up-time"`             // CAPWAP uptime (Live: IOS-XE 17.12.6a)
}

// SpectrumBandConfigData represents spectrum band configuration data.
type SpectrumBandConfigData struct {
	ApMAC              string               `json:"ap-mac"`                         // MAC address of the AP (Live: IOS-XE 17.12.6a)
	SpectrumBandConfig []SpectrumBandConfig `json:"spectrum-band-config,omitempty"` // Spectrum band config (Live: IOS-XE 17.12.6a)
}

// SpectrumBandConfig represents spectrum band configuration for a specific band.
type SpectrumBandConfig struct {
	BandID             string `json:"band-id"`              // Band of AP for configuration (Live: IOS-XE 17.12.6a)
	SpectrumAdminState bool   `json:"spectrum-admin-state"` // Spectrum admin state (Live: IOS-XE 17.12.6a)
}

// RRMClientData represents RRM client data.
type RRMClientData struct {
	PhyType         string    `json:"phy-type"`        // Radio type (Live: IOS-XE 17.12.6a)
	LastChdRun      time.Time `json:"last-chd-run"`    // Timestamp at which CHD algorithm was last run (Live: IOS-XE 17.12.6a)
	Disassociations int       `json:"disassociations"` // Number of dissociations by client (Live: IOS-XE 17.12.6a)
	Rejections      int       `json:"rejections"`      // Number of rejections by client (Live: IOS-XE 17.12.6a)
}

// RRMFraStats represents Flexible Radio Assignment statistics.
type RRMFraStats struct {
	DualBandMonitorTo24ghz int `json:"dual-band-monitor-to-24ghz"` // Dual-band radio transition from monitor to 2.4GHz (Live: IOS-XE 17.12.6a)
	DualBandMonitorTo5ghz  int `json:"dual-band-monitor-to-5ghz"`  // Dual-band radio transition from monitor to 5GHz (Live: IOS-XE 17.12.6a)
	DualBand24ghzTo5ghz    int `json:"dual-band-24ghz-to-5ghz"`    // Dual-band radio transition from 2.4GHz to 5GHz (Live: IOS-XE 17.12.6a)
	DualBand24ghzToMonitor int `json:"dual-band-24ghz-to-monitor"` // Dual-band radio transition from 2.4GHz to monitor (Live: IOS-XE 17.12.6a)
	DualBand5ghzTo24ghz    int `json:"dual-band-5ghz-to-24ghz"`    // Dual-band radio transition from 5GHz to 2.4GHz (Live: IOS-XE 17.12.6a)
	DualBand5ghzToMonitor  int `json:"dual-band-5ghz-to-monitor"`  // Dual-band radio transition from 5GHz to monitor (Live: IOS-XE 17.12.6a)
	SecRadioMonitorTo5ghz  int `json:"sec-radio-monitor-to-5ghz"`  // Secondary radio transition from monitor to 5GHz (Live: IOS-XE 17.12.6a)
	SecRadio5ghzToMonitor  int `json:"sec-radio-5ghz-to-monitor"`  // Secondary radio transition from 5GHz to monitor (Live: IOS-XE 17.12.6a)
	DualBand6ghzTo5ghz     int `json:"dual-band-6ghz-to-5ghz"`     // Dual-band radio transition from 6GHz to 5GHz (Live: IOS-XE 17.12.6a)
	DualBand5ghzTo6ghz     int `json:"dual-band-5ghz-to-6ghz"`     // Dual-band radio transition from 5GHz to 6GHz (Live: IOS-XE 17.12.6a)
}

// RRMCoverage represents RRM coverage data.
type RRMCoverage struct {
	WtpMAC            string     `json:"wtp-mac"`             // MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID       int        `json:"radio-slot-id"`       // Radio slot identifier (Live: IOS-XE 17.12.6a)
	FailedClientCount int        `json:"failed-client-count"` // Number of failed clients (Live: IOS-XE 17.12.6a)
	SNRInfo           []SNRInfo  `json:"snr-info,omitempty"`  // Client signal to noise ratio (Live: IOS-XE 17.12.6a)
	RSSIInfo          []RSSIInfo `json:"rssi-info,omitempty"` // Received signal strength from the client (Live: IOS-XE 17.12.6a)
}

// SNRInfo represents Signal-to-Noise Ratio information.
type SNRInfo struct {
	SNR        int `json:"snr"`         // Client signal to noise ratio (Live: IOS-XE 17.12.6a)
	NumClients int `json:"num-clients"` // Number of clients per SNR (Live: IOS-XE 17.12.6a)
}

// RSSIInfo represents Received Signal Strength Indicator information.
type RSSIInfo struct {
	RSSI       int `json:"rssi"`        // Received signal strength from the client (Live: IOS-XE 17.12.6a)
	NumClients int `json:"num-clients"` // Number of clients per RSSI (Live: IOS-XE 17.12.6a)
}

// SpectrumAqWorstTable represents spectrum Air Quality worst interference table entry.
type SpectrumAqWorstTable struct {
	BandID               int    `json:"band-id"`                 // Band ID (Live: IOS-XE 17.12.6a)
	DetectingApName      string `json:"detecting-ap-name"`       // AP name (Live: IOS-XE 17.12.6a)
	ChannelNum           int    `json:"channel-num"`             // Channel number (Live: IOS-XE 17.12.6a)
	MinAqi               int    `json:"min-aqi"`                 // Min air quality index (Live: IOS-XE 17.12.6a)
	Aqi                  int    `json:"aqi"`                     // Air quality index (Live: IOS-XE 17.12.6a)
	TotalIntfDeviceCount int    `json:"total-intf-device-count"` // Interference device count (Live: IOS-XE 17.12.6a)
	WtpCaSiCapable       string `json:"wtp-ca-si-capable"`       // AP SI capable or not (Live: IOS-XE 17.12.6a)
	ScanRadioType        string `json:"scan-radio-type"`         // Scan radio type (Live: IOS-XE 17.12.6a)
}
