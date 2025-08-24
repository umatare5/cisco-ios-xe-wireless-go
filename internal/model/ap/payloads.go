// Package model provides AP-specific payloads for the Cisco IOS-XE Wireless Network Controller API.
package model

// APReloadRPCPayload  represents the complete payload for AP reload RPC calls
type APReloadRPCPayload struct {
	Input APReloadRPCInput `json:"input"`
}

// APTagPayload  represents the complete payload for AP tag assignment
type APTagPayload struct {
	ApTag APCfgApTagData `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
}

// APConfigRPCPayload  represents the complete payload for AP configuration RPC calls
type APConfigRPCPayload struct {
	Input APConfigRPCInput `json:"Cisco-IOS-XE-wireless-access-point-cfg-rpc:input"`
}

// APSlotConfigRPCPayload  represents the complete payload for AP slot configuration RPC calls
type APSlotConfigRPCPayload struct {
	Input APSlotConfigRPCInput `json:"Cisco-IOS-XE-wireless-access-point-cfg-rpc:input"`
}

// APReloadRPCInput  represents the input structure for AP reload RPC calls
type APReloadRPCInput struct {
	APName  string `json:"ap-name,omitempty"`
	MACAddr string `json:"mac-addr,omitempty"`
}

// APCfgApTagData  represents the tag data structure for AP tag assignment
type APCfgApTagData struct {
	APMac     string `json:"ap-mac"`
	SiteTag   string `json:"site-tag"`
	PolicyTag string `json:"policy-tag"`
	RFTag     string `json:"rf-tag"`
}

// APConfigRPCInput  represents the input structure for AP configuration RPC calls
type APConfigRPCInput struct {
	Mode    string `json:"mode"`
	MACAddr string `json:"mac-addr,omitempty"`
	APName  string `json:"ap-name,omitempty"`
}

// APSlotConfigRPCInput  represents the input structure for AP slot configuration RPC calls
type APSlotConfigRPCInput struct {
	Mode    string `json:"mode"`
	SlotID  int    `json:"slot-id"`
	Band    string `json:"band"`
	MACAddr string `json:"mac-addr"`
}
