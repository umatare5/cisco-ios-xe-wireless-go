package ap

// APReloadRPCPayload represents complete payload for AP reload RPC calls.
type APReloadRPCPayload struct {
	Input APReloadRPCInput `json:"input"`
}

// APTagPayload represents complete payload for AP tag assignment.
type APTagPayload struct {
	ApTag APCfgApTagData `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
}

// APConfigRPCPayload represents complete payload for AP configuration RPC calls.
type APConfigRPCPayload struct {
	Input APConfigRPCInput `json:"Cisco-IOS-XE-wireless-access-point-cfg-rpc:input"`
}

// APSlotConfigRPCPayload represents complete payload for AP slot configuration RPC calls.
type APSlotConfigRPCPayload struct {
	Input APSlotConfigRPCInput `json:"Cisco-IOS-XE-wireless-access-point-cfg-rpc:input"`
}

// APReloadRPCInput represents input structure for AP reload RPC calls.
type APReloadRPCInput struct {
	APName  string `json:"ap-name,omitempty"`  // AP name identifier
	MACAddr string `json:"mac-addr,omitempty"` // AP MAC address identifier
}

// APCfgApTagData represents tag data structure for AP tag assignment.
type APCfgApTagData struct {
	APMac     string `json:"ap-mac"`     // AP MAC address
	SiteTag   string `json:"site-tag"`   // Site tag assigned to the AP
	PolicyTag string `json:"policy-tag"` // Policy tag assigned to the AP
	RFTag     string `json:"rf-tag"`     // RF tag assigned to the AP
}

// APConfigRPCInput represents input structure for AP configuration RPC calls.
type APConfigRPCInput struct {
	Mode    string `json:"mode"`               // Configuration mode
	MACAddr string `json:"mac-addr,omitempty"` // AP MAC address identifier
	APName  string `json:"ap-name,omitempty"`  // AP name identifier
}

// APSlotConfigRPCInput represents input structure for AP slot configuration RPC calls.
type APSlotConfigRPCInput struct {
	Mode    string `json:"mode"`     // Configuration mode
	SlotID  int    `json:"slot-id"`  // Radio slot identifier (0=2.4GHz, 1=5GHz, 2=6GHz)
	Band    string `json:"band"`     // Radio band specification
	MACAddr string `json:"mac-addr"` // AP MAC address
}
