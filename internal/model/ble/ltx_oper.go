package ble

// BleLtxOper represents BLE LTX operational data container.
type BleLtxOper struct {
	BleLtxOperData BleLtxOperData `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data"` // BLE LTX operational data container (YANG: IOS-XE 17.12.1)
}

// BleLtxOperBleLtxApAntenna represents BLE LTX AP antenna collection wrapper.
type BleLtxOperBleLtxApAntenna struct {
	BleLtxApAntenna []BleLtxApAntenna `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap-antenna"` // BLE LTX AP antenna information (YANG: IOS-XE 17.12.1)
}

// BleLtxOperBleLtxAp represents BLE LTX AP collection wrapper.
type BleLtxOperBleLtxAp struct {
	BleLtxAp []BleLtxAp `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap"` // BLE LTX AP data (YANG: IOS-XE 17.12.1)
}

// BleLtxOperData represents BLE LTX operational data container.
type BleLtxOperData struct {
	BleLtxApAntenna []BleLtxApAntenna `json:"ble-ltx-ap-antenna"` // BLE LTX AP antenna information (YANG: IOS-XE 17.12.1)
	BleLtxAp        []BleLtxAp        `json:"ble-ltx-ap"`         // BLE LTX AP data (YANG: IOS-XE 17.12.1)
}

// BleLtxApAntenna represents BLE antenna configuration for specific AP.
type BleLtxApAntenna struct {
	ApMac               string `json:"ap-mac"`                 // BLE LTX AP MAC address (YANG: IOS-XE 17.12.1)
	BleSlotID           int    `json:"ble-slot-id"`            // BLE antenna slot identifier (YANG: IOS-XE 17.12.1)
	BleAntennaID        int    `json:"ble-antenna-id"`         // BLE antenna identifier (YANG: IOS-XE 17.12.1)
	IsBleAntennaPresent bool   `json:"is-ble-antenna-present"` // AP has BLE antenna present (YANG: IOS-XE 17.12.1)
	BleAntennaPid       string `json:"ble-antenna-pid"`        // AP BLE antenna PID (YANG: IOS-XE 17.12.1)
	BleAntennaGain      int    `json:"ble-antenna-gain"`       // AP BLE antenna gain (YANG: IOS-XE 17.12.1)
	BleAntennaType      string `json:"ble-antenna-type"`       // AP BLE antenna type (YANG: IOS-XE 17.12.1)
	BleAntennaMode      string `json:"ble-antenna-mode"`       // AP BLE antenna mode (YANG: IOS-XE 17.12.1)
	BleAntennaDiversity string `json:"ble-antenna-diversity"`  // AP BLE antenna diversity status (YANG: IOS-XE 17.12.1)
	BleAntennaOptions   string `json:"ble-antenna-options"`    // AP BLE antenna options (YANG: IOS-XE 17.12.1)
}

// BleLtxAp represents BLE LTX configuration and status for specific AP.
type BleLtxAp struct {
	ApMac            string               `json:"ap-mac"`            // AP MAC address (Live: IOS-XE 17.12.5)
	Admin            BleLtxApAdmin        `json:"admin"`             // AP administrative state (Live: IOS-XE 17.12.5)
	ScanConfig       BleLtxApScanConfig   `json:"scan-config"`       // AP BLE scan configuration (Live: IOS-XE 17.12.5)
	ProfileIbeacon   BleLtxApProfile      `json:"profile-ibeacon"`   // AP BLE iBeacon chirping profile (Live: IOS-XE 17.12.5)
	ProfileEddyURL   BleLtxApProfile      `json:"profile-eddy-url"`  // AP BLE Eddystone URL chirping profile (Live: IOS-XE 17.12.5)
	ProfileEddyUID   BleLtxApProfile      `json:"profile-eddy-uid"`  // AP BLE Eddystone UID chirping profile (Live: IOS-XE 17.12.5)
	ProfileVibeacons BleLtxApProfile      `json:"profile-vibeacons"` // AP BLE viBeacons chirping profiles (Live: IOS-XE 17.12.5)
	ScanCounters     BleLtxApScanCounters `json:"scan-counters"`     // AP BLE scan counters (Live: IOS-XE 17.12.5)
	HostData         BleLtxApHostData     `json:"host-data"`         // AP BLE host data (Live: IOS-XE 17.12.5)
	FeatureMode      BleLtxApFeatureMode  `json:"feature-mode"`      // AP BLE LTX feature mode (Live: IOS-XE 17.12.5)
	DeviceStatus     BleLtxApDeviceStatus `json:"device-status"`     // AP BLE interface status (Live: IOS-XE 17.12.5)
	Capability       BleLtxApCapability   `json:"capability"`        // AP BLE capability (Live: IOS-XE 17.12.5)
	Streaming        BleLtxApStreaming    `json:"streaming"`         // AP BLE telemetry streaming (Live: IOS-XE 17.12.5)
}

// BleLtxApAdmin represents BLE AP administrative configuration.
type BleLtxApAdmin struct {
	State    string           `json:"state"`    // State of the overall BLE hardware module (Live: IOS-XE 17.12.5)
	Feedback BleLtxApFeedback `json:"feedback"` // Feedback of the last config command (Live: IOS-XE 17.12.5)
	Report   BleLtxApReport   `json:"report"`   // Status of the last report from AP BLE (Live: IOS-XE 17.12.5)
}

// BleLtxApScanConfig represents BLE AP scanning configuration.
type BleLtxApScanConfig struct {
	IntervalSec int                  `json:"interval-sec"` // AP BLE scan cycle time in seconds (Live: IOS-XE 17.12.5)
	State       string               `json:"state"`        // AP BLE scan enable flag (Live: IOS-XE 17.12.5)
	MaxValue    int                  `json:"max-value"`    // AP BLE max number of scan performed in cycle (Live: IOS-XE 17.12.5)
	WindowMsec  int                  `json:"window-msec"`  // AP BLE scan time during each cycle in milliseconds (Live: IOS-XE 17.12.5)
	Filter      string               `json:"filter"`       // AP BLE flag to enable/disable MAC based scan filtering (Live: IOS-XE 17.12.5)
	Feedback    BleLtxApScanFeedback `json:"feedback"`     // Feedback of the last config command (Live: IOS-XE 17.12.5)
	Report      BleLtxApReport       `json:"report"`       // Status of the last report from AP BLE (Live: IOS-XE 17.12.5)
}

// BleLtxApScanFeedback represents BLE AP scan configuration feedback.
type BleLtxApScanFeedback struct {
	IntervalSecStatus int `json:"interval-sec-status"` // Reported status of AP BLE scan cycle time (Live: IOS-XE 17.12.5)
	StateStatus       int `json:"state-status"`        // Reported status of AP BLE scan enable flag (Live: IOS-XE 17.12.5)
	MaxValueStatus    int `json:"max-value-status"`    // Reported status of AP BLE max number of scans (Live: IOS-XE 17.12.5)
	WindowMsecStatus  int `json:"window-msec-status"`  // Reported status of AP BLE scan time during each cycle (Live: IOS-XE 17.12.5)
	FilterStatus      int `json:"filter-status"`       // Reported status of AP BLE MAC based scan filtering flag (Live: IOS-XE 17.12.5)
}

// BleLtxApProfile represents BLE AP profile configuration.
type BleLtxApProfile struct {
	Report BleLtxApReport `json:"report"` // Status of the last report from AP BLE (Live: IOS-XE 17.12.5)
}

// BleLtxApScanCounters represents BLE AP scan statistics counters.
type BleLtxApScanCounters struct {
	Total          *uint32        `json:"total,omitempty"`           // Total scan records (YANG: IOS-XE 17.12.1)
	DNALtx         *uint32        `json:"dna-ltx,omitempty"`         // Total DNA LTX records (YANG: IOS-XE 17.12.1)
	SystemTlm      *uint32        `json:"system-tlm,omitempty"`      // Total system telemetry records (YANG: IOS-XE 17.12.1)
	EventTlm       *uint32        `json:"event-tlm,omitempty"`       // Total event telemetry records (YANG: IOS-XE 17.12.1)
	RegularTlm     *uint32        `json:"regular-tlm,omitempty"`     // Total regular telemetry records (YANG: IOS-XE 17.12.1)
	Emergency      *uint32        `json:"emergency,omitempty"`       // Total emergency records (YANG: IOS-XE 17.12.1)
	EventEmergency *uint32        `json:"event-emergency,omitempty"` // Total event emergency records (YANG: IOS-XE 17.12.1)
	Other          *uint32        `json:"other,omitempty"`           // Other records received by AP (YANG: IOS-XE 17.12.1)
	Report         BleLtxApReport `json:"report"`                    // Status of the last report from AP BLE (Live: IOS-XE 17.12.5)
}

// BleLtxApHostData represents BLE AP host device information.
type BleLtxApHostData struct {
	DeviceName     *string        `json:"device-name,omitempty"`     // AP BLE device name (YANG: IOS-XE 17.12.1)
	BLEMac         *string        `json:"ble-mac,omitempty"`         // BLE over the Air MAC address (YANG: IOS-XE 17.12.1)
	APIVersion     *uint16        `json:"api-version,omitempty"`     // AP BLE API version (YANG: IOS-XE 17.12.1)
	FWVersion      *string        `json:"fw-version,omitempty"`      // AP BLE Device/major/minor/revision (YANG: IOS-XE 17.12.1)
	AdvertiseCount *uint32        `json:"advertise-count,omitempty"` // Total number of broadcasts since powered on (YANG: IOS-XE 17.12.1)
	UptimeDsec     *uint32        `json:"uptime-dsec,omitempty"`     // Time since chip was last powered on in deciseconds (YANG: IOS-XE 17.12.1)
	ActiveProfile  *string        `json:"active-profile,omitempty"`  // AP BLE active chirping profile (YANG: IOS-XE 17.12.1)
	Report         BleLtxApReport `json:"report"`                    // Status of the last report from AP BLE (Live: IOS-XE 17.12.5)
}

// BleLtxApFeatureMode represents BLE AP feature mode configuration.
type BleLtxApFeatureMode struct {
	Feature string         `json:"feature"` // Type of BLE feature (Live: IOS-XE 17.12.5)
	Mode    string         `json:"mode"`    // Mode of the BLE device (Live: IOS-XE 17.12.5)
	Report  BleLtxApReport `json:"report"`  // Status of the last report from AP BLE (Live: IOS-XE 17.12.5)
}

// BleLtxApDeviceStatus represents BLE AP device status information.
type BleLtxApDeviceStatus struct {
	Device string         `json:"device"` // Type of BLE device (Live: IOS-XE 17.12.5)
	State  string         `json:"state"`  // State of BLE device (Live: IOS-XE 17.12.5)
	Report BleLtxApReport `json:"report"` // Status of the last report from AP BLE (Live: IOS-XE 17.12.5)
}

// BleLtxApCapability represents BLE AP capabilities information.
type BleLtxApCapability struct {
	Ble           bool           `json:"ble"`             // BLE capability (Live: IOS-XE 17.12.5)
	Zigbee        bool           `json:"zigbee"`          // Zigbee capability (Live: IOS-XE 17.12.5)
	Thread        bool           `json:"thread"`          // Thread capability (Live: IOS-XE 17.12.5)
	Usb           bool           `json:"usb"`             // USB capability (Live: IOS-XE 17.12.5)
	Iot           bool           `json:"iot"`             // IOT capability (Live: IOS-XE 17.12.5)
	BleHybridMode bool           `json:"ble-hybrid-mode"` // BLE hybrid mode capability (Live: IOS-XE 17.12.5)
	Report        BleLtxApReport `json:"report"`          // Status of the last report from AP BLE (Live: IOS-XE 17.12.5)
}

// BleLtxApStreaming represents BLE AP streaming configuration.
type BleLtxApStreaming struct {
	State  string         `json:"state"`  // State of the BLE telemetry streaming (Live: IOS-XE 17.12.5)
	Count  string         `json:"count"`  // BLE telemetry streaming record counter (Live: IOS-XE 17.12.5)
	Report BleLtxApReport `json:"report"` // BLE telemetry streaming report (Live: IOS-XE 17.12.5)
}

// BleLtxApFeedback represents BLE AP feedback status.
type BleLtxApFeedback struct {
	StateStatus int `json:"state-status"` // Reported status of overall BLE hardware module (Live: IOS-XE 17.12.5)
}

// BleLtxApReport represents BLE AP report information.
type BleLtxApReport struct {
	LastReportTime string `json:"last-report-time,omitempty"` // Timestamp of the last AP report received (Live: IOS-XE 17.12.5)
	Valid          bool   `json:"valid"`                      // Indicates whether data is valid (Live: IOS-XE 17.12.5)
}
