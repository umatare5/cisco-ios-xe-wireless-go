package client

// ClientDeauthRPCPayload represents complete payload for client deauthentication RPC calls.
type ClientDeauthRPCPayload struct {
	Input ClientDeauthRPCInput `json:"input"`
}

// ClientDeauthRPCInput represents input structure for client deauthentication RPC calls.
//
// The three identifiers are the arms of the RPC's mandatory choice, so exactly one is filled and
// the other two are absent rather than empty. zone-id is not declared: it defaults to 0 and the
// controller resolves every arm within that zone.
type ClientDeauthRPCInput struct {
	MACAddr  string `json:"mac-addr,omitempty"`  // Client MAC address identifier
	IPAddr   string `json:"ip-addr,omitempty"`   // Client IP address identifier
	Username string `json:"user-name,omitempty"` // Client username identifier
}
