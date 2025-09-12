package lisp

// LispAgentOper represents LISP agent operational data root structure.
type LispAgentOper struct {
	LispAgentOperData LispAgentOperData `json:"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data"` // LISP agent operational data container
}

// LispAgentOperData represents main container for LISP agent operational data.
type LispAgentOperData struct {
	LispAgentMemoryStats *LispAgentMemoryStats `json:"lisp-agent-memory-stats,omitempty"` // LISP agent memory statistics
	LispWLCCapabilities  *LispWLCCapabilities  `json:"lisp-wlc-capabilities,omitempty"`   // LISP WLC capabilities information
	LispAPCapabilities   []LispAPCapability    `json:"lisp-ap-capabilities,omitempty"`    // LISP AP capabilities per AP type
}

// LispAgentMemoryStats represents LISP agent memory allocation statistics.
type LispAgentMemoryStats struct {
	MallocPskBuf        string `json:"malloc-psk-buf"`         // Allocated PSK buffer count
	FreePskBuf          string `json:"free-psk-buf"`           // Free PSK buffer count
	MallocMapRegMsg     string `json:"malloc-map-reg-msg"`     // Allocated map register message count
	FreeMapRegMsg       string `json:"free-map-reg-msg"`       // Free map register message count
	MallocMapReqMsg     string `json:"malloc-map-req-msg"`     // Allocated map request message count
	FreeMapReqMsg       string `json:"free-map-req-msg"`       // Free map request message count
	MallocLispHANode    string `json:"malloc-lisp-ha-node"`    // Allocated LISP HA node count
	FreeLispHANode      string `json:"free-lisp-ha-node"`      // Free LISP HA node count
	MallocMapServerCtxt string `json:"malloc-map-server-ctxt"` // Allocated map server context count
	FreeMapServerCtxt   string `json:"free-map-server-ctxt"`   // Free map server context count
}

// LispWLCCapabilities represents LISP wireless controller capabilities.
type LispWLCCapabilities struct {
	FabricCapable bool `json:"fabric-capable"` // Fabric mode capability flag
}

// LispAPCapability represents LISP access point capability for specific AP type.
type LispAPCapability struct {
	APType        int  `json:"ap-type"`        // Access point type identifier
	FabricCapable bool `json:"fabric-capable"` // Fabric mode capability flag for this AP type
}
