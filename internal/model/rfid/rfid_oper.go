package model

// RfidOper  represents the RFID operational data.
type RfidOper struct {
	RfidOperData RfidOperData `json:"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data"`
}

// RfidOperRfidData  represents the RFID data.
type RfidOperRfidData struct {
	RfidData []RfidDataItem `json:"Cisco-IOS-XE-wireless-rfid-oper:rfid-data"`
}

type RfidOperData struct {
	RfidData []RfidDataItem `json:"rfid-data"`
}

type RfidDataItem struct {
	RfidMacAddr string `json:"rfid-mac-addr"`
	// Add other RFID operational data fields as needed
}
