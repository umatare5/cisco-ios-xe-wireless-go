package routes

// NMSP (Network Mobility Services Protocol) Operational Paths
//
// These constants define the RESTCONF API paths for NMSP operational
// data based on Cisco-IOS-XE-wireless-nmsp-oper YANG model.

// NMSP Operational Paths.
const (
	// NMSPOperPath provides the path for retrieving all NMSP operational data.
	NMSPOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"

	// NMSPClientRegistrationPath provides the path for retrieving client registration data.
	NMSPClientRegistrationPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data/client-registration"

	// NMSPCmxConnectionPath provides the path for retrieving CMX connection data.
	NMSPCmxConnectionPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data/cmx-connection"

	// NMSPCmxCloudInfoPath provides the path for retrieving CMX cloud information.
	NMSPCmxCloudInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data/cmx-cloud-info"
)
