// Package ble provides Bluetooth Low Energy operational data functionality for the Cisco Wireless Network Controller API.
package ble

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// BleLtxOperBasePath defines the base path for BLE LTX operational data endpoints
	BleLtxOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data"
	// BleLtxOperEndpoint retrieves complete BLE LTX operational data
	BleLtxOperEndpoint = BleLtxOperBasePath
	// BleLtxApAntennaEndpoint retrieves BLE LTX AP antenna information
	BleLtxApAntennaEndpoint = BleLtxOperBasePath + "/ble-ltx-ap-antenna"
	// BleLtxApEndpoint retrieves BLE LTX AP information
	BleLtxApEndpoint = BleLtxOperBasePath + "/ble-ltx-ap"
)

// BleLtxOperResponse represents the complete BLE LTX operational data response
type BleLtxOperResponse struct {
	CiscoIOSXEWirelessBleLtxOperBleLtxOperData struct {
		BleLtxApAntenna []BleLtxApAntenna `json:"ble-ltx-ap-antenna"`
		BleLtxAp        []BleLtxAp        `json:"ble-ltx-ap"`
	} `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data"`
}

// BleLtxApAntennaResponse represents the BLE LTX AP antenna information response
type BleLtxApAntennaResponse struct {
	BleLtxApAntenna []BleLtxApAntenna `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap-antenna"`
}

// BleLtxApResponse represents the BLE LTX AP information response
type BleLtxApResponse struct {
	BleLtxAp []BleLtxAp `json:"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap"`
}

// BleLtxApAntenna contains BLE LTX antenna information for an access point
type BleLtxApAntenna struct {
	ApMac               string `json:"ap-mac"`                 // Access point MAC address
	BleSlotID           int    `json:"ble-slot-id"`            // BLE slot identifier
	BleAntennaID        int    `json:"ble-antenna-id"`         // BLE antenna identifier
	IsBleAntennaPresent bool   `json:"is-ble-antenna-present"` // Whether BLE antenna is present
	BleAntennaPid       string `json:"ble-antenna-pid"`        // BLE antenna product ID
	BleAntennaGain      int    `json:"ble-antenna-gain"`       // BLE antenna gain in dBi
	BleAntennaType      string `json:"ble-antenna-type"`       // BLE antenna type
	BleAntennaMode      string `json:"ble-antenna-mode"`       // BLE antenna mode
	BleAntennaDiversity string `json:"ble-antenna-diversity"`  // BLE antenna diversity setting
	BleAntennaOptions   string `json:"ble-antenna-options"`    // BLE antenna options
}

// BleLtxAp contains comprehensive BLE LTX information for an access point
type BleLtxAp struct {
	ApMac string `json:"ap-mac"` // Access point MAC address
	Admin struct {
		State    string `json:"state"` // Administrative state
		Feedback struct {
			StateStatus int `json:"state-status"` // State status code
		} `json:"feedback"`
		Report struct {
			LastReportTime string `json:"last-report-time"` // Last report timestamp
			Valid          bool   `json:"valid"`            // Report validity
		} `json:"report"`
	} `json:"admin"`
	ScanConfig struct {
		IntervalSec int    `json:"interval-sec"` // Scan interval in seconds
		State       string `json:"state"`        // Scan state
		MaxValue    int    `json:"max-value"`    // Maximum scan value
		WindowMsec  int    `json:"window-msec"`  // Scan window in milliseconds
		Filter      string `json:"filter"`       // Scan filter settings
		Feedback    struct {
			IntervalSecStatus int `json:"interval-sec-status"` // Interval status code
			StateStatus       int `json:"state-status"`        // State status code
			MaxValueStatus    int `json:"max-value-status"`    // Max value status code
			WindowMsecStatus  int `json:"window-msec-status"`  // Window status code
			FilterStatus      int `json:"filter-status"`       // Filter status code
		} `json:"feedback"`
		Report struct {
			LastReportTime string `json:"last-report-time"` // Last report timestamp
			Valid          bool   `json:"valid"`            // Report validity
		} `json:"report"`
	} `json:"scan-config"`
	ProfileIbeacon struct {
		Report struct {
			Valid bool `json:"valid"` // iBeacon profile validity
		} `json:"report"`
	} `json:"profile-ibeacon"`
	ProfileEddyUrl struct {
		Report struct {
			Valid bool `json:"valid"` // Eddystone URL profile validity
		} `json:"report"`
	} `json:"profile-eddy-url"`
	ProfileEddyUid struct {
		Report struct {
			Valid bool `json:"valid"` // Eddystone UID profile validity
		} `json:"report"`
	} `json:"profile-eddy-uid"`
	ProfileVibeacons struct {
		Report struct {
			Valid bool `json:"valid"` // Vibeacons profile validity
		} `json:"report"`
	} `json:"profile-vibeacons"`
	ScanCounters struct {
		Report struct {
			Valid bool `json:"valid"` // Scan counters validity
		} `json:"report"`
	} `json:"scan-counters"`
	HostData struct {
		Report struct {
			Valid bool `json:"valid"` // Host data validity
		} `json:"report"`
	} `json:"host-data"`
	FeatureMode struct {
		Feature string `json:"feature"` // BLE feature type
		Mode    string `json:"mode"`    // Feature mode
		Report  struct {
			LastReportTime string `json:"last-report-time"` // Last report timestamp
			Valid          bool   `json:"valid"`            // Report validity
		} `json:"report"`
	} `json:"feature-mode"`
	DeviceStatus struct {
		Device string `json:"device"` // Device identifier
		State  string `json:"state"`  // Device state
		Report struct {
			LastReportTime string `json:"last-report-time"` // Last report timestamp
			Valid          bool   `json:"valid"`            // Report validity
		} `json:"report"`
	} `json:"device-status"`
	Capability struct {
		Ble           bool `json:"ble"`             // BLE capability
		Zigbee        bool `json:"zigbee"`          // Zigbee capability
		Thread        bool `json:"thread"`          // Thread capability
		Usb           bool `json:"usb"`             // USB capability
		Iot           bool `json:"iot"`             // IoT capability
		BleHybridMode bool `json:"ble-hybrid-mode"` // BLE hybrid mode capability
		Report        struct {
			LastReportTime string `json:"last-report-time"` // Last report timestamp
			Valid          bool   `json:"valid"`            // Report validity
		} `json:"report"`
	} `json:"capability"`
	Streaming struct {
		State  string `json:"state"` // Streaming state
		Report struct {
			LastReportTime string `json:"last-report-time"` // Last report timestamp
			Valid          bool   `json:"valid"`            // Report validity
		} `json:"report"`
		Count string `json:"count"` // Streaming count
	} `json:"streaming"`
}

func GetBleLtxOper(client *wnc.Client, ctx context.Context) (*BleLtxOperResponse, error) {
	var data BleLtxOperResponse
	err := client.SendAPIRequest(ctx, BleLtxOperEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetBleLtxApAntenna(client *wnc.Client, ctx context.Context) (*BleLtxApAntennaResponse, error) {
	var data BleLtxApAntennaResponse
	err := client.SendAPIRequest(ctx, BleLtxApAntennaEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetBleLtxAp(client *wnc.Client, ctx context.Context) (*BleLtxApResponse, error) {
	var data BleLtxApResponse
	err := client.SendAPIRequest(ctx, BleLtxApEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
