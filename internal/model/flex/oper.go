package flex

// FlexOper represents the FlexConnect operational data from WNC 17.12.5.
type FlexOper struct {
	FlexOperData FlexOperData `json:"Cisco-IOS-XE-wireless-flex-oper:flex-oper-data"`
}

// FlexOperData represents the root operational data structure.
type FlexOperData struct {
	FlexAPModeData []FlexAPModeInfo `json:"flex-ap-mode-data,omitempty"`
}

// FlexAPModeInfo represents FlexConnect AP mode information.
type FlexAPModeInfo struct {
	WTPMac        string            `json:"wtp-mac"`
	WTPMode       string            `json:"wtp-mode"`
	HomeAPEnabled *bool             `json:"home-ap-enabled,omitempty"`
	ClearMode     *bool             `json:"clear-mode,omitempty"`
	APSubMode     string            `json:"ap-sub-mode,omitempty"`
	APFabricData  *FlexAPFabricData `json:"ap-fabric-data,omitempty"`
}

// FlexAPFabricData represents fabric-related operational data.
type FlexAPFabricData struct {
	IsFabricAP *bool `json:"is-fabric-ap,omitempty"`
}
