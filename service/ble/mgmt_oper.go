package ble

// BLEMgmtOper represents BLE management operational data container.
type BLEMgmtOper struct {
	BLEMgmtOperData struct {
		BLEMgmtAp  []BLEMgmtAp  `json:"ble-mgmt-ap,omitempty"`  // BLE Management AP data (YANG: IOS-XE 17.12.1)
		BLEMgmtCmx []BLEMgmtCmx `json:"ble-mgmt-cmx,omitempty"` // BLE Management CMX data (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data"` // BLE management operational data (YANG: IOS-XE 17.12.1)
}

// BLEMgmtAp represents BLE management data for each AP.
type BLEMgmtAp struct {
	ApMAC        string  `json:"ap-mac"`                  // AP MAC address (YANG: IOS-XE 17.12.1)
	IsNew        *bool   `json:"is-new,omitempty"`        // The AP just joined the controller (YANG: IOS-XE 17.12.1)
	CmxID        *uint64 `json:"cmx-id,omitempty"`        // ID of CMX controlling the AP (YANG: IOS-XE 17.12.1)
	BLEInterface *string `json:"ble-interface,omitempty"` // AP BLE interface (YANG: IOS-XE 17.12.1)
	RadioState   *string `json:"radio-state,omitempty"`   // AP BLE radio state (YANG: IOS-XE 17.12.1)
	OperState    *bool   `json:"oper-state,omitempty"`    // AP BLE Operational state (YANG: IOS-XE 17.12.1)
}

// BLEMgmtCmx represents BLE management data for each CMX.
type BLEMgmtCmx struct {
	CmxID        uint64  `json:"cmx-id"`                   // CMX Identifier (YANG: IOS-XE 17.12.1)
	OperState    *bool   `json:"oper-state,omitempty"`     // Operational state (YANG: IOS-XE 17.12.1)
	ReasonDown   *string `json:"reason-down,omitempty"`    // Reason for BLE operational state down (YANG: IOS-XE 17.12.1)
	AdminState   *bool   `json:"admin-state,omitempty"`    // Administrative state (YANG: IOS-XE 17.12.1)
	CmxResolved  *bool   `json:"cmx-resolved,omitempty"`   // Whether CMX is reachable (YANG: IOS-XE 17.12.1)
	ScanInterval *uint32 `json:"scan-interval,omitempty"`  // BLE scan interval (YANG: IOS-XE 17.12.1)
	BLESourceMAC *string `json:"ble-source-mac,omitempty"` // Source MAC used for BLE traffic from AP (YANG: IOS-XE 17.12.1)
}
