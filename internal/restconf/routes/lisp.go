package routes

// LISP (Locator/ID Separation Protocol) Operational Paths
//
// These constants define the RESTCONF API paths for LISP agent operational
// data based on Cisco-IOS-XE-wireless-lisp-agent-oper YANG model.

// LISP Operational Paths
//
// Note: LISP configuration endpoints are not available on this controller.
// Only operational data is provided through the lisp-agent-oper-data endpoint.
const (
	// LISPOperPath provides the path for retrieving all LISP agent operational data.
	LISPOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data"

	// LISPMemoryStatsPath provides the path for retrieving LISP memory statistics.
	LISPMemoryStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data/lisp-agent-memory-stats"

	// LISPCapabilitiesPath provides the path for retrieving LISP WLC capabilities.
	LISPCapabilitiesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data/lisp-wlc-capabilities"
)
