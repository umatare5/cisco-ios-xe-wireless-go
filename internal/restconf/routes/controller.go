package routes

// Controller Management RPC Operations
//
// These constants define the RESTCONF API paths for controller management
// operations such as system reload, restart, and administrative operations.
// Based on the Cisco-IOS-XE-rpc and cisco-ia YANG model specifications.

// Controller RPC Operations.
const (
	// ControllerReloadRPC defines the RPC for WNC controller reload operations
	// This operation follows the Cisco-IOS-XE-rpc:reload YANG model specification.
	ControllerReloadRPC = RESTCONFOperationsPath + "/Cisco-IOS-XE-rpc:reload"

	// ControllerSaveConfigRPC defines the RPC that copies the running configuration to the startup
	// configuration. cisco-ia is the one module name here that is neither capitalised nor
	// Cisco-IOS-XE-* prefixed, and the controller publishes it in that form.
	ControllerSaveConfigRPC = RESTCONFOperationsPath + "/cisco-ia:save-config"
)

// Controller Device Operational Data
//
// The subject of this service is the controller itself and not only its wireless namespace, so
// these paths read the platform's own operational data under a module that is not a
// Cisco-IOS-XE-wireless-* one. ControllerReloadRPC above is the same case for an RPC.

// Controller Operational Paths.
const (
	// ControllerBootTimePath retrieves the instant at which the controller last booted.
	ControllerBootTimePath = RESTCONFDataPath + "/Cisco-IOS-XE-device-hardware-oper:device-hardware-data/device-hardware/device-system-data/boot-time"
)
