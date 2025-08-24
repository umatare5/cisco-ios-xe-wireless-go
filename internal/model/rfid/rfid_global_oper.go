// Package model provides data models for RFID global operational data.
package model

// RfidGlobalOper  represents the RFID global operational data.
type RfidGlobalOper struct {
	RfidGlobalOperData RfidGlobalOperData `json:"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data"`
}

// RfidGlobalOperRfidDataDetail  represents the RFID data detail.
type RfidGlobalOperRfidDataDetail struct {
	RfidDataDetail []RfidDataDetail `json:"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-data-detail"`
}

// RfidGlobalOperRfidRadioData  represents the RFID radio data.
type RfidGlobalOperRfidRadioData struct {
	RfidRadioData []RfidRadioData `json:"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-radio-data"`
}

type RfidGlobalOperData struct {
	RfidDataDetail []RfidDataDetail `json:"rfid-data-detail"`
	RfidRadioData  []RfidRadioData  `json:"rfid-radio-data"`
}

type RfidDataDetail struct {
	RfidMacAddr string `json:"rfid-mac-addr"`
	// Add other RFID detail fields as needed
}

type RfidRadioData struct {
	RfidMacAddr string `json:"rfid-mac-addr"`
	ApMacAddr   string `json:"ap-mac-addr"`
	Slot        int    `json:"slot"`
	// Add other RFID radio data fields as needed
}
