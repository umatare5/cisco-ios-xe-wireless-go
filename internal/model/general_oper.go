// Package model contains generated response structures for the Cisco WNC API.
package model

// GeneralOperResponse represents the complete general operational data response
type GeneralOperResponse struct {
	CiscoIOSXEWirelessGeneralOperData struct {
		MgmtIntfData MgmtIntfData `json:"mgmt-intf-data"`
	} `json:"Cisco-IOS-XE-wireless-general-oper:general-oper-data"`
}

// GeneralOperMgmtIntfDataResponse represents the management interface data response
type GeneralOperMgmtIntfDataResponse struct {
	MgmtIntfData MgmtIntfData `json:"Cisco-IOS-XE-wireless-general-oper:mgmt-intf-data"`
}

// MgmtIntfData contains management interface configuration and status information
type MgmtIntfData struct {
	IntfName string `json:"intf-name"` // Interface name
	IntfType string `json:"intf-type"` // Interface type
	IntfID   int    `json:"intf-id"`   // Interface ID
	MgmtIP   string `json:"mgmt-ip"`   // Management IP address
	NetMask  string `json:"net-mask"`  // Network mask
	MgmtMAC  string `json:"mgmt-mac"`  // Management MAC address
}
