package ble

// BleLtxOper represents BLE LTX operational data container.
type BleLtxOper struct {
	BleLtxOperData BleLtxOperData `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data"` // BLE LTX operational data
}

// BleLtxOperBleLtxApAntenna represents BLE LTX AP antenna collection wrapper.
type BleLtxOperBleLtxApAntenna struct {
	BleLtxApAntenna []BleLtxApAntenna `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap-antenna"` // BLE AP antenna configurations
}

// BleLtxOperBleLtxAp represents BLE LTX AP collection wrapper.
type BleLtxOperBleLtxAp struct {
	BleLtxAp []BleLtxAp `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap"` // BLE LTX AP configurations
}

// BleLtxOperData represents BLE LTX operational data container.
type BleLtxOperData struct {
	BleLtxApAntenna []BleLtxApAntenna `json:"ble-ltx-ap-antenna"` // BLE AP antenna configurations
	BleLtxAp        []BleLtxAp        `json:"ble-ltx-ap"`         // BLE AP operational data
}

// BleLtxApAntenna represents BLE antenna configuration for specific AP.
type BleLtxApAntenna struct {
	ApMac               string `json:"ap-mac"`                 // Access point MAC address
	BleSlotID           int    `json:"ble-slot-id"`            // BLE slot identifier
	BleAntennaID        int    `json:"ble-antenna-id"`         // BLE antenna identifier
	IsBleAntennaPresent bool   `json:"is-ble-antenna-present"` // BLE antenna presence status
	BleAntennaPid       string `json:"ble-antenna-pid"`        // BLE antenna product identifier
	BleAntennaGain      int    `json:"ble-antenna-gain"`       // BLE antenna gain value
	BleAntennaType      string `json:"ble-antenna-type"`       // BLE antenna type
	BleAntennaMode      string `json:"ble-antenna-mode"`       // BLE antenna mode setting
	BleAntennaDiversity string `json:"ble-antenna-diversity"`  // BLE antenna diversity configuration
	BleAntennaOptions   string `json:"ble-antenna-options"`    // BLE antenna options
}

// BleLtxAp represents BLE LTX configuration and status for specific AP.
type BleLtxAp struct {
	ApMac            string               `json:"ap-mac"`            // Access point MAC address
	Admin            BleLtxApAdmin        `json:"admin"`             // Administrative configuration
	ScanConfig       BleLtxApScanConfig   `json:"scan-config"`       // Scanning configuration
	ProfileIbeacon   BleLtxApProfile      `json:"profile-ibeacon"`   // iBeacon profile configuration
	ProfileEddyURL   BleLtxApProfile      `json:"profile-eddy-url"`  // Eddystone URL profile
	ProfileEddyUID   BleLtxApProfile      `json:"profile-eddy-uid"`  // Eddystone UID profile
	ProfileVibeacons BleLtxApProfile      `json:"profile-vibeacons"` // viBeacons profile configuration
	ScanCounters     BleLtxApScanCounters `json:"scan-counters"`     // Scan statistics counters
	HostData         BleLtxApHostData     `json:"host-data"`         // Host device information
	FeatureMode      BleLtxApFeatureMode  `json:"feature-mode"`      // Feature mode configuration
	DeviceStatus     BleLtxApDeviceStatus `json:"device-status"`     // Device status information
	Capability       BleLtxApCapability   `json:"capability"`        // Device capabilities
	Streaming        BleLtxApStreaming    `json:"streaming"`         // Streaming configuration
}

// BleLtxApAdmin represents BLE AP administrative configuration.
type BleLtxApAdmin struct {
	State    string           `json:"state"`    // Administrative state
	Feedback BleLtxApFeedback `json:"feedback"` // Configuration feedback
	Report   BleLtxApReport   `json:"report"`   // Report information
}

// BleLtxApScanConfig represents BLE AP scanning configuration.
type BleLtxApScanConfig struct {
	IntervalSec int                  `json:"interval-sec"` // Scan interval in seconds
	State       string               `json:"state"`        // Scan state
	MaxValue    int                  `json:"max-value"`    // Maximum scan value
	WindowMsec  int                  `json:"window-msec"`  // Scan window in milliseconds
	Filter      string               `json:"filter"`       // Scan filter setting
	Feedback    BleLtxApScanFeedback `json:"feedback"`     // Scan configuration feedback
	Report      BleLtxApReport       `json:"report"`       // Report information
}

// BleLtxApScanFeedback represents BLE AP scan configuration feedback.
type BleLtxApScanFeedback struct {
	IntervalSecStatus int `json:"interval-sec-status"` // Interval configuration status
	StateStatus       int `json:"state-status"`        // State configuration status
	MaxValueStatus    int `json:"max-value-status"`    // Max value configuration status
	WindowMsecStatus  int `json:"window-msec-status"`  // Window configuration status
	FilterStatus      int `json:"filter-status"`       // Filter configuration status
}

// BleLtxApProfile represents BLE AP profile configuration.
type BleLtxApProfile struct {
	Report BleLtxApReport `json:"report"` // Profile report information
}

// BleLtxApScanCounters represents BLE AP scan statistics counters.
type BleLtxApScanCounters struct {
	Total          *uint32        `json:"total,omitempty"`           // Total scan count
	DNALtx         *uint32        `json:"dna-ltx,omitempty"`         // DNA LTX scan count
	SystemTlm      *uint32        `json:"system-tlm,omitempty"`      // System telemetry count
	EventTlm       *uint32        `json:"event-tlm,omitempty"`       // Event telemetry count
	RegularTlm     *uint32        `json:"regular-tlm,omitempty"`     // Regular telemetry count
	Emergency      *uint32        `json:"emergency,omitempty"`       // Emergency event count
	EventEmergency *uint32        `json:"event-emergency,omitempty"` // Event emergency count
	Other          *uint32        `json:"other,omitempty"`           // Other scan count
	Report         BleLtxApReport `json:"report"`                    // Report information
}

// BleLtxApHostData represents BLE AP host device information.
type BleLtxApHostData struct {
	DeviceName     *string        `json:"device-name,omitempty"`     // Host device name
	BLEMac         *string        `json:"ble-mac,omitempty"`         // BLE MAC address
	APIVersion     *uint16        `json:"api-version,omitempty"`     // API version
	FWVersion      *string        `json:"fw-version,omitempty"`      // Firmware version
	AdvertiseCount *uint32        `json:"advertise-count,omitempty"` // Advertisement count
	UptimeDsec     *uint32        `json:"uptime-dsec,omitempty"`     // Uptime in deciseconds
	ActiveProfile  *string        `json:"active-profile,omitempty"`  // Active profile name
	Report         BleLtxApReport `json:"report"`                    // Report information
}

// BleLtxApFeatureMode represents BLE AP feature mode configuration.
type BleLtxApFeatureMode struct {
	Feature string         `json:"feature"` // Feature type
	Mode    string         `json:"mode"`    // Mode setting
	Report  BleLtxApReport `json:"report"`  // Report information
}

// BleLtxApDeviceStatus represents BLE AP device status information.
type BleLtxApDeviceStatus struct {
	Device string         `json:"device"` // Device identifier
	State  string         `json:"state"`  // Device state
	Report BleLtxApReport `json:"report"` // Report information
}

// BleLtxApCapability represents BLE AP capabilities information.
type BleLtxApCapability struct {
	Ble           bool           `json:"ble"`             // BLE capability
	Zigbee        bool           `json:"zigbee"`          // Zigbee capability
	Thread        bool           `json:"thread"`          // Thread capability
	Usb           bool           `json:"usb"`             // USB capability
	Iot           bool           `json:"iot"`             // IoT capability
	BleHybridMode bool           `json:"ble-hybrid-mode"` // BLE hybrid mode capability
	Report        BleLtxApReport `json:"report"`          // Report information
}

// BleLtxApStreaming represents BLE AP streaming configuration.
type BleLtxApStreaming struct {
	State  string         `json:"state"`  // Streaming state
	Count  string         `json:"count"`  // Streaming count
	Report BleLtxApReport `json:"report"` // Report information
}

// BleLtxApFeedback represents BLE AP feedback status.
type BleLtxApFeedback struct {
	StateStatus int `json:"state-status"` // State configuration status
}

// BleLtxApReport represents BLE AP report information.
type BleLtxApReport struct {
	LastReportTime string `json:"last-report-time,omitempty"` // Last report timestamp
	Valid          bool   `json:"valid"`                      // Report validity status
}
