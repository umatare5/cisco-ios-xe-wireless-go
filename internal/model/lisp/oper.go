package lisp

// LispAgentOper represents LISP agent operational data root structure.
type LispAgentOper struct {
	CiscoIOSXEWirelessLispAgentOperData struct {
		LispAgentMemoryStats *LispAgentMemoryStats `json:"lisp-agent-memory-stats,omitempty"` // Lisp Agent Memory Statistics (YANG: IOS-XE 17.12.1)
		LispWLCCapabilities  *LispWLCCapabilities  `json:"lisp-wlc-capabilities,omitempty"`   // Wireless Fabric WLC Capabilities (YANG: IOS-XE 17.12.1)
		LispAPCapabilities   []LispAPCapability    `json:"lisp-ap-capabilities,omitempty"`    // Wireless Fabric AP Capabilities (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data"` // Lisp Agent operational data (YANG: IOS-XE 17.12.1)
}

// LispAgentMemoryStats represents LISP agent memory allocation statistics.
type LispAgentMemoryStats struct {
	MallocPskBuf        string `json:"malloc-psk-buf"`         // Malloc count of psk buffer (YANG: IOS-XE 17.12.1)
	FreePskBuf          string `json:"free-psk-buf"`           // Free count of psk buffer (YANG: IOS-XE 17.12.1)
	MallocMapRegMsg     string `json:"malloc-map-reg-msg"`     // Malloc count of map register message (YANG: IOS-XE 17.12.1)
	FreeMapRegMsg       string `json:"free-map-reg-msg"`       // Free count of map register message (YANG: IOS-XE 17.12.1)
	MallocMapReqMsg     string `json:"malloc-map-req-msg"`     // Malloc count of map request message (YANG: IOS-XE 17.12.1)
	FreeMapReqMsg       string `json:"free-map-req-msg"`       // Free count of map request message (YANG: IOS-XE 17.12.1)
	MallocLispHANode    string `json:"malloc-lisp-ha-node"`    // Malloc count of lisp HA node (YANG: IOS-XE 17.12.1)
	FreeLispHANode      string `json:"free-lisp-ha-node"`      // Free count of lisp HA node (YANG: IOS-XE 17.12.1)
	MallocMapServerCtxt string `json:"malloc-map-server-ctxt"` // Malloc count of control plane context (YANG: IOS-XE 17.12.1)
	FreeMapServerCtxt   string `json:"free-map-server-ctxt"`   // Free count of control plane context (YANG: IOS-XE 17.12.1)
}

// LispWLCCapabilities represents LISP wireless controller capabilities.
type LispWLCCapabilities struct {
	FabricCapable bool `json:"fabric-capable"` // WLC Fabric capable (YANG: IOS-XE 17.12.1)
}

// LispAPCapability represents LISP access point capability for specific AP type.
type LispAPCapability struct {
	APType        int  `json:"ap-type"`        // AP Type (YANG: IOS-XE 17.12.1)
	FabricCapable bool `json:"fabric-capable"` // AP Fabric capable (YANG: IOS-XE 17.12.1)
}
