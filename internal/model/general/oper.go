package general

// GeneralOper represents the complete general operational data response.
type GeneralOper struct {
	CiscoIOSXEWirelessGeneralOperData struct {
		MgmtIntfData MgmtIntfData `json:"mgmt-intf-data"` // Controller wireless interface data (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-general-oper:general-oper-data"` // General operational data (Live: IOS-XE 17.12.5)
}

// GeneralOperMgmtIntfData represents the management interface data response.
type GeneralOperMgmtIntfData struct {
	MgmtIntfData MgmtIntfData `json:"Cisco-IOS-XE-wireless-general-oper:mgmt-intf-data"` // Controller wireless interface data (Live: IOS-XE 17.12.5)
}

// MgmtIntfData represents management interface configuration and status.
type MgmtIntfData struct {
	IntfName string `json:"intf-name"` // Controller management interface name (Live: IOS-XE 17.12.5)
	IntfType string `json:"intf-type"` // Controller management interface type (Live: IOS-XE 17.12.5)
	IntfID   int    `json:"intf-id"`   // Controller management interface id (Live: IOS-XE 17.12.5)
	MgmtIP   string `json:"mgmt-ip"`   // Controller management IPv4 address (Live: IOS-XE 17.12.5)
	NetMask  string `json:"net-mask"`  // Controller management interface netmask (Live: IOS-XE 17.12.5)
	MgmtMAC  string `json:"mgmt-mac"`  // Controller management interface MAC (Live: IOS-XE 17.12.5)
}
