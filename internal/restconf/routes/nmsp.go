package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// NMSP (Network Mobility Services Protocol) Operational Endpoints
//
// These constants define the RESTCONF API endpoints for NMSP operational
// data based on Cisco-IOS-XE-wireless-nmsp-oper YANG model.

const (
	// NMSPOperBasePath defines the base path for NMSP operational data endpoints
	NMSPOperBasePath = restconf.YANGModelPrefix + "nmsp-oper:nmsp-oper-data"
)

// NMSP Operational Endpoints
const (
	// EndpointGetOper provides the endpoint for retrieving all NMSP operational data
	EndpointGetOper = NMSPOperBasePath

	// EndpointGetOperClientRegistration provides the endpoint for retrieving client registration data
	EndpointGetOperClientRegistration = NMSPOperBasePath + "/client-registration"

	// EndpointGetOperCmxConnection provides the endpoint for retrieving CMX connection data
	EndpointGetOperCmxConnection = NMSPOperBasePath + "/cmx-connection"

	// EndpointGetOperCmxCloudInfo provides the endpoint for retrieving CMX cloud information
	EndpointGetOperCmxCloudInfo = NMSPOperBasePath + "/cmx-cloud-info"
)
