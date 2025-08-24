// Package model provides data models for BLE LTX operational data.
package model

// BleLtxOper  represents the structure for BLE LTX operational data.
type BleLtxOper struct {
	BleLtxOperData BleLtxOperData `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data"`
}

// BleLtxOperBleLtxApAntenna  represents the structure for BLE LTX AP antenna data.
type BleLtxOperBleLtxApAntenna struct {
	BleLtxApAntenna []BleLtxApAntenna `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap-antenna"`
}

// BleLtxOperBleLtxAp  represents the structure for BLE LTX AP data.
type BleLtxOperBleLtxAp struct {
	BleLtxAp []BleLtxAp `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap"`
}

type BleLtxOperData struct {
	BleLtxApAntenna []BleLtxApAntenna `json:"ble-ltx-ap-antenna"`
	BleLtxAp        []BleLtxAp        `json:"ble-ltx-ap"`
}

type BleLtxApAntenna struct {
	ApMac               string `json:"ap-mac"`
	BleSlotID           int    `json:"ble-slot-id"`
	BleAntennaID        int    `json:"ble-antenna-id"`
	IsBleAntennaPresent bool   `json:"is-ble-antenna-present"`
	BleAntennaPid       string `json:"ble-antenna-pid"`
	BleAntennaGain      int    `json:"ble-antenna-gain"`
	BleAntennaType      string `json:"ble-antenna-type"`
	BleAntennaMode      string `json:"ble-antenna-mode"`
	BleAntennaDiversity string `json:"ble-antenna-diversity"`
	BleAntennaOptions   string `json:"ble-antenna-options"`
}

type BleLtxAp struct {
	ApMac            string               `json:"ap-mac"`
	Admin            BleLtxApAdmin        `json:"admin"`
	ScanConfig       BleLtxApScanConfig   `json:"scan-config"`
	ProfileIbeacon   BleLtxApProfile      `json:"profile-ibeacon"`
	ProfileEddyURL   BleLtxApProfile      `json:"profile-eddy-url"`
	ProfileEddyUID   BleLtxApProfile      `json:"profile-eddy-uid"`
	ProfileVibeacons BleLtxApProfile      `json:"profile-vibeacons"`
	ScanCounters     BleLtxApProfile      `json:"scan-counters"`
	HostData         BleLtxApProfile      `json:"host-data"`
	FeatureMode      BleLtxApFeatureMode  `json:"feature-mode"`
	DeviceStatus     BleLtxApDeviceStatus `json:"device-status"`
	Capability       BleLtxApCapability   `json:"capability"`
	Streaming        BleLtxApStreaming    `json:"streaming"`
}

type BleLtxApAdmin struct {
	State    string           `json:"state"`
	Feedback BleLtxApFeedback `json:"feedback"`
	Report   BleLtxApReport   `json:"report"`
}

type BleLtxApScanConfig struct {
	IntervalSec int                  `json:"interval-sec"`
	State       string               `json:"state"`
	MaxValue    int                  `json:"max-value"`
	WindowMsec  int                  `json:"window-msec"`
	Filter      string               `json:"filter"`
	Feedback    BleLtxApScanFeedback `json:"feedback"`
	Report      BleLtxApReport       `json:"report"`
}

type BleLtxApScanFeedback struct {
	IntervalSecStatus int `json:"interval-sec-status"`
	StateStatus       int `json:"state-status"`
	MaxValueStatus    int `json:"max-value-status"`
	WindowMsecStatus  int `json:"window-msec-status"`
	FilterStatus      int `json:"filter-status"`
}

type BleLtxApProfile struct {
	Report BleLtxApReport `json:"report"`
}

type BleLtxApFeatureMode struct {
	Feature string         `json:"feature"`
	Mode    string         `json:"mode"`
	Report  BleLtxApReport `json:"report"`
}

type BleLtxApDeviceStatus struct {
	Device string         `json:"device"`
	State  string         `json:"state"`
	Report BleLtxApReport `json:"report"`
}

type BleLtxApCapability struct {
	Ble           bool           `json:"ble"`
	Zigbee        bool           `json:"zigbee"`
	Thread        bool           `json:"thread"`
	Usb           bool           `json:"usb"`
	Iot           bool           `json:"iot"`
	BleHybridMode bool           `json:"ble-hybrid-mode"`
	Report        BleLtxApReport `json:"report"`
}

type BleLtxApStreaming struct {
	State  string         `json:"state"`
	Count  string         `json:"count"`
	Report BleLtxApReport `json:"report"`
}

type BleLtxApFeedback struct {
	StateStatus int `json:"state-status"`
}

type BleLtxApReport struct {
	LastReportTime string `json:"last-report-time,omitempty"`
	Valid          bool   `json:"valid"`
}
