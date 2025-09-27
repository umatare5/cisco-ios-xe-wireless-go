package general

// CiscoIOSXEWirelessGeneralOper represents the complete general operational data response.
type CiscoIOSXEWirelessGeneralOper struct {
	CiscoIOSXEWirelessGeneralOperData struct {
		MgmtIntfData MgmtIntfData `json:"mgmt-intf-data"` // Controller wireless interface data (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-general-oper:general-oper-data"` // General operational data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessGeneralOperMgmtIntfData represents the management interface data response.
type CiscoIOSXEWirelessGeneralOperMgmtIntfData struct {
	MgmtIntfData MgmtIntfData `json:"Cisco-IOS-XE-wireless-general-oper:mgmt-intf-data"` // Controller wireless interface data (Live: IOS-XE 17.12.6a)
}

// MgmtIntfData represents management interface configuration and status.
type MgmtIntfData struct {
	IntfName string `json:"intf-name"` // Controller management interface name (Live: IOS-XE 17.12.6a)
	IntfType string `json:"intf-type"` // Controller management interface type (Live: IOS-XE 17.12.6a)
	IntfID   int    `json:"intf-id"`   // Controller management interface id (Live: IOS-XE 17.12.6a)
	MgmtIP   string `json:"mgmt-ip"`   // Controller management IPv4 address (Live: IOS-XE 17.12.6a)
	NetMask  string `json:"net-mask"`  // Controller management interface netmask (Live: IOS-XE 17.12.6a)
	MgmtMAC  string `json:"mgmt-mac"`  // Controller management interface MAC (Live: IOS-XE 17.12.6a)
}
