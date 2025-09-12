package ble

// BleMgmtOper represents BLE management operational data container.
type BleMgmtOper struct {
	BleMgmtOperData BleMgmtOperData `json:"Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data"` // BLE Management operational data
}

// BleMgmtOperData represents BLE management operational data container.
type BleMgmtOperData struct {
	BleMgmtAp  []BleMgmtAp  `json:"ble-mgmt-ap,omitempty"`  // BLE Management AP data (YANG: IOS-XE 17.12.1+)
	BleMgmtCmx []BleMgmtCmx `json:"ble-mgmt-cmx,omitempty"` // BLE Management CMX data (YANG: IOS-XE 17.12.1+)
}

// BleMgmtAp represents BLE management data for each AP.
type BleMgmtAp struct {
	ApMac        string  `json:"ap-mac"`                  // AP MAC address (YANG: IOS-XE 17.12.1+)
	IsNew        *bool   `json:"is-new,omitempty"`        // The AP just joined the controller (YANG: IOS-XE 17.12.1+)
	CmxID        *uint64 `json:"cmx-id,omitempty"`        // ID of CMX controlling the AP (YANG: IOS-XE 17.12.1+)
	BleInterface *string `json:"ble-interface,omitempty"` // AP BLE interface (YANG: IOS-XE 17.12.1+)
	RadioState   *string `json:"radio-state,omitempty"`   // AP BLE radio state (YANG: IOS-XE 17.12.1+)
	OperState    *bool   `json:"oper-state,omitempty"`    // AP BLE Operational state (YANG: IOS-XE 17.12.1+)
}

// BleMgmtCmx represents BLE management data for each CMX.
type BleMgmtCmx struct {
	CmxID        uint64  `json:"cmx-id"`                   // CMX Identifier (YANG: IOS-XE 17.12.1+)
	OperState    *bool   `json:"oper-state,omitempty"`     // Operational state (YANG: IOS-XE 17.12.1+)
	ReasonDown   *string `json:"reason-down,omitempty"`    // Reason for BLE operational state down (YANG: IOS-XE 17.12.1+)
	AdminState   *bool   `json:"admin-state,omitempty"`    // Administrative state (YANG: IOS-XE 17.12.1+)
	CmxResolved  *bool   `json:"cmx-resolved,omitempty"`   // Whether CMX is reachable (YANG: IOS-XE 17.12.1+)
	ScanInterval *uint32 `json:"scan-interval,omitempty"`  // BLE scan interval (YANG: IOS-XE 17.12.1+)
	BleSourceMac *string `json:"ble-source-mac,omitempty"` // Source MAC used for BLE traffic from AP (YANG: IOS-XE 17.12.1+)
}
