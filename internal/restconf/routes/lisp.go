package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// LISP (Locator/ID Separation Protocol) Operational Endpoints
//
// These constants define the RESTCONF API endpoints for LISP agent operational
// data based on Cisco-IOS-XE-wireless-lisp-agent-oper YANG model.

const (
	// LispOperBasePath defines the base path for LISP agent operational data endpoints
	LispOperBasePath = restconf.YANGModelPrefix + "lisp-agent-oper:lisp-agent-oper-data"
)

// LISP Operational Endpoints
//
// Note: LISP configuration endpoints are not available on this controller.
// Only operational data is provided through the lisp-agent-oper-data endpoint.
const (
	// LispOperEndpoint provides the endpoint for retrieving all LISP agent operational data
	LispOperEndpoint = LispOperBasePath

	// LispOperMemoryStatsEndpoint provides the endpoint for retrieving LISP memory statistics
	LispOperMemoryStatsEndpoint = LispOperBasePath + "/lisp-agent-memory-stats"

	// LispOperCapabilitiesEndpoint provides the endpoint for retrieving LISP WLC capabilities
	LispOperCapabilitiesEndpoint = LispOperBasePath + "/lisp-wlc-capabilities"
)
