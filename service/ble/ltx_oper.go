package ble

// CiscoIOSXEWirelessBLELtxOper represents BLE LTX operational data container.
type CiscoIOSXEWirelessBLELtxOper struct {
	CiscoIOSXEWirelessBLELtxOperData struct {
		BLELtxApAntenna []BLELtxApAntenna `json:"ble-ltx-ap-antenna"` // BLE LTX AP antenna information (YANG: IOS-XE 17.12.1)
		BLELtxAp        []BLELtxAp        `json:"ble-ltx-ap"`         // BLE LTX AP data (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data"` // BLE LTX operational data (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessBLELtxOperBLELtxApAntenna represents BLE LTX AP antenna collection wrapper.
type CiscoIOSXEWirelessBLELtxOperBLELtxApAntenna struct {
	BLELtxApAntenna []BLELtxApAntenna `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap-antenna"`
}

// CiscoIOSXEWirelessBLELtxOperBLELtxAp represents BLE LTX AP collection wrapper.
type CiscoIOSXEWirelessBLELtxOperBLELtxAp struct {
	BLELtxAp []BLELtxAp `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap"`
}

// BLELtxApAntenna represents BLE antenna configuration for specific AP.
type BLELtxApAntenna struct {
	ApMAC               string `json:"ap-mac"`                 // BLE LTX AP MAC address (YANG: IOS-XE 17.12.1)
	BLESlotID           int    `json:"ble-slot-id"`            // BLE antenna slot identifier (YANG: IOS-XE 17.12.1)
	BLEAntennaID        int    `json:"ble-antenna-id"`         // BLE antenna identifier (YANG: IOS-XE 17.12.1)
	IsBLEAntennaPresent bool   `json:"is-ble-antenna-present"` // AP has BLE antenna present (YANG: IOS-XE 17.12.1)
	BLEAntennaPid       string `json:"ble-antenna-pid"`        // AP BLE antenna PID (YANG: IOS-XE 17.12.1)
	BLEAntennaGain      int    `json:"ble-antenna-gain"`       // AP BLE antenna gain (YANG: IOS-XE 17.12.1)
	BLEAntennaType      string `json:"ble-antenna-type"`       // AP BLE antenna type (YANG: IOS-XE 17.12.1)
	BLEAntennaMode      string `json:"ble-antenna-mode"`       // AP BLE antenna mode (YANG: IOS-XE 17.12.1)
	BLEAntennaDiversity string `json:"ble-antenna-diversity"`  // AP BLE antenna diversity status (YANG: IOS-XE 17.12.1)
	BLEAntennaOptions   string `json:"ble-antenna-options"`    // AP BLE antenna options (YANG: IOS-XE 17.12.1)
}

// BLELtxAp represents BLE LTX configuration and status for specific AP.
type BLELtxAp struct {
	ApMAC            string               `json:"ap-mac"`            // AP MAC address (Live: IOS-XE 17.12.6a)
	Admin            BLELtxApAdmin        `json:"admin"`             // AP administrative state (Live: IOS-XE 17.12.6a)
	ScanConfig       BLELtxApScanConfig   `json:"scan-config"`       // AP BLE scan configuration (Live: IOS-XE 17.12.6a)
	ProfileIbeacon   BLELtxApProfile      `json:"profile-ibeacon"`   // AP BLE iBeacon chirping profile (Live: IOS-XE 17.12.6a)
	ProfileEddyURL   BLELtxApProfile      `json:"profile-eddy-url"`  // AP BLE Eddystone URL chirping profile (Live: IOS-XE 17.12.6a)
	ProfileEddyUID   BLELtxApProfile      `json:"profile-eddy-uid"`  // AP BLE Eddystone UID chirping profile (Live: IOS-XE 17.12.6a)
	ProfileVibeacons BLELtxApProfile      `json:"profile-vibeacons"` // AP BLE viBeacons chirping profiles (Live: IOS-XE 17.12.6a)
	ScanCounters     BLELtxApScanCounters `json:"scan-counters"`     // AP BLE scan counters (Live: IOS-XE 17.12.6a)
	HostData         BLELtxApHostData     `json:"host-data"`         // AP BLE host data (Live: IOS-XE 17.12.6a)
	FeatureMode      BLELtxApFeatureMode  `json:"feature-mode"`      // AP BLE LTX feature mode (Live: IOS-XE 17.12.6a)
	DeviceStatus     BLELtxApDeviceStatus `json:"device-status"`     // AP BLE interface status (Live: IOS-XE 17.12.6a)
	Capability       BLELtxApCapability   `json:"capability"`        // AP BLE capability (Live: IOS-XE 17.12.6a)
	Streaming        BLELtxApStreaming    `json:"streaming"`         // AP BLE telemetry streaming (Live: IOS-XE 17.12.6a)
}

// BLELtxApAdmin represents BLE AP administrative configuration.
type BLELtxApAdmin struct {
	State    string           `json:"state"`    // State of the overall BLE hardware module (Live: IOS-XE 17.12.6a)
	Feedback BLELtxApFeedback `json:"feedback"` // Feedback of the last config command (Live: IOS-XE 17.12.6a)
	Report   BLELtxApReport   `json:"report"`   // Status of the last report from AP BLE (Live: IOS-XE 17.12.6a)
}

// BLELtxApScanConfig represents BLE AP scanning configuration.
type BLELtxApScanConfig struct {
	IntervalSec int                  `json:"interval-sec"` // AP BLE scan cycle time in seconds (Live: IOS-XE 17.12.6a)
	State       string               `json:"state"`        // AP BLE scan enable flag (Live: IOS-XE 17.12.6a)
	MaxValue    int                  `json:"max-value"`    // AP BLE max number of scan performed in cycle (Live: IOS-XE 17.12.6a)
	WindowMsec  int                  `json:"window-msec"`  // AP BLE scan time during each cycle in milliseconds (Live: IOS-XE 17.12.6a)
	Filter      string               `json:"filter"`       // AP BLE flag to enable/disable MAC based scan filtering (Live: IOS-XE 17.12.6a)
	Feedback    BLELtxApScanFeedback `json:"feedback"`     // Feedback of the last config command (Live: IOS-XE 17.12.6a)
	Report      BLELtxApReport       `json:"report"`       // Status of the last report from AP BLE (Live: IOS-XE 17.12.6a)
}

// BLELtxApScanFeedback represents BLE AP scan configuration feedback.
type BLELtxApScanFeedback struct {
	IntervalSecStatus int `json:"interval-sec-status"` // Reported status of AP BLE scan cycle time (Live: IOS-XE 17.12.6a)
	StateStatus       int `json:"state-status"`        // Reported status of AP BLE scan enable flag (Live: IOS-XE 17.12.6a)
	MaxValueStatus    int `json:"max-value-status"`    // Reported status of AP BLE max number of scans (Live: IOS-XE 17.12.6a)
	WindowMsecStatus  int `json:"window-msec-status"`  // Reported status of AP BLE scan time during each cycle (Live: IOS-XE 17.12.6a)
	FilterStatus      int `json:"filter-status"`       // Reported status of AP BLE MAC based scan filtering flag (Live: IOS-XE 17.12.6a)
}

// BLELtxApProfile represents BLE AP profile configuration.
type BLELtxApProfile struct {
	Report BLELtxApReport `json:"report"` // Status of the last report from AP BLE (Live: IOS-XE 17.12.6a)
}

// BLELtxApScanCounters represents BLE AP scan statistics counters.
type BLELtxApScanCounters struct {
	Total          *uint32        `json:"total,omitempty"`           // Total scan records (YANG: IOS-XE 17.12.1)
	DNALtx         *uint32        `json:"dna-ltx,omitempty"`         // Total DNA LTX records (YANG: IOS-XE 17.12.1)
	SystemTlm      *uint32        `json:"system-tlm,omitempty"`      // Total system telemetry records (YANG: IOS-XE 17.12.1)
	EventTlm       *uint32        `json:"event-tlm,omitempty"`       // Total event telemetry records (YANG: IOS-XE 17.12.1)
	RegularTlm     *uint32        `json:"regular-tlm,omitempty"`     // Total regular telemetry records (YANG: IOS-XE 17.12.1)
	Emergency      *uint32        `json:"emergency,omitempty"`       // Total emergency records (YANG: IOS-XE 17.12.1)
	EventEmergency *uint32        `json:"event-emergency,omitempty"` // Total event emergency records (YANG: IOS-XE 17.12.1)
	Other          *uint32        `json:"other,omitempty"`           // Other records received by AP (YANG: IOS-XE 17.12.1)
	Report         BLELtxApReport `json:"report"`                    // Status of the last report from AP BLE (Live: IOS-XE 17.12.6a)
}

// BLELtxApHostData represents BLE AP host device information.
type BLELtxApHostData struct {
	DeviceName     *string        `json:"device-name,omitempty"`     // AP BLE device name (YANG: IOS-XE 17.12.1)
	BLEMac         *string        `json:"ble-mac,omitempty"`         // BLE over the Air MAC address (YANG: IOS-XE 17.12.1)
	APIVersion     *uint16        `json:"api-version,omitempty"`     // AP BLE API version (YANG: IOS-XE 17.12.1)
	FWVersion      *string        `json:"fw-version,omitempty"`      // AP BLE Device/major/minor/revision (YANG: IOS-XE 17.12.1)
	AdvertiseCount *uint32        `json:"advertise-count,omitempty"` // Total number of broadcasts since powered on (YANG: IOS-XE 17.12.1)
	UptimeDsec     *uint32        `json:"uptime-dsec,omitempty"`     // Time since chip was last powered on in deciseconds (YANG: IOS-XE 17.12.1)
	ActiveProfile  *string        `json:"active-profile,omitempty"`  // AP BLE active chirping profile (YANG: IOS-XE 17.12.1)
	Report         BLELtxApReport `json:"report"`                    // Status of the last report from AP BLE (Live: IOS-XE 17.12.6a)
}

// BLELtxApFeatureMode represents BLE AP feature mode configuration.
type BLELtxApFeatureMode struct {
	Feature string         `json:"feature"` // Type of BLE feature (Live: IOS-XE 17.12.6a)
	Mode    string         `json:"mode"`    // Mode of the BLE device (Live: IOS-XE 17.12.6a)
	Report  BLELtxApReport `json:"report"`  // Status of the last report from AP BLE (Live: IOS-XE 17.12.6a)
}

// BLELtxApDeviceStatus represents BLE AP device status information.
type BLELtxApDeviceStatus struct {
	Device string         `json:"device"` // Type of BLE device (Live: IOS-XE 17.12.6a)
	State  string         `json:"state"`  // State of BLE device (Live: IOS-XE 17.12.6a)
	Report BLELtxApReport `json:"report"` // Status of the last report from AP BLE (Live: IOS-XE 17.12.6a)
}

// BLELtxApCapability represents BLE AP capabilities information.
type BLELtxApCapability struct {
	BLE           bool           `json:"ble"`             // BLE capability (Live: IOS-XE 17.12.6a)
	Zigbee        bool           `json:"zigbee"`          // Zigbee capability (Live: IOS-XE 17.12.6a)
	Thread        bool           `json:"thread"`          // Thread capability (Live: IOS-XE 17.12.6a)
	Usb           bool           `json:"usb"`             // USB capability (Live: IOS-XE 17.12.6a)
	Iot           bool           `json:"iot"`             // IOT capability (Live: IOS-XE 17.12.6a)
	BLEHybridMode bool           `json:"ble-hybrid-mode"` // BLE hybrid mode capability (Live: IOS-XE 17.12.6a)
	Report        BLELtxApReport `json:"report"`          // Status of the last report from AP BLE (Live: IOS-XE 17.12.6a)
}

// BLELtxApStreaming represents BLE AP streaming configuration.
type BLELtxApStreaming struct {
	State  string         `json:"state"`  // State of the BLE telemetry streaming (Live: IOS-XE 17.12.6a)
	Count  string         `json:"count"`  // BLE telemetry streaming record counter (Live: IOS-XE 17.12.6a)
	Report BLELtxApReport `json:"report"` // BLE telemetry streaming report (Live: IOS-XE 17.12.6a)
}

// BLELtxApFeedback represents BLE AP feedback status.
type BLELtxApFeedback struct {
	StateStatus int `json:"state-status"` // Reported status of overall BLE hardware module (Live: IOS-XE 17.12.6a)
}

// BLELtxApReport represents BLE AP report information.
type BLELtxApReport struct {
	LastReportTime string `json:"last-report-time,omitempty"` // Timestamp of the last AP report received (Live: IOS-XE 17.12.6a)
	Valid          bool   `json:"valid"`                      // Indicates whether data is valid (Live: IOS-XE 17.12.6a)
}
