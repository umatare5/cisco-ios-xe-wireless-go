// Package mock provides mock server implementations for testing RESTCONF services.
//
// This package contains mock RESTCONF servers and test utilities for simulating
// Cisco Catalyst 9800 controller responses. It supports both success and error
// scenarios for comprehensive testing of service operations.
//
// Key components:
//   - RESTCONF Server: NewRESTCONFSuccessServer, NewRESTCONFErrorServer
//   - Client Creation: NewTLSClientForServer for connecting to test servers
//   - Flexible Server: RESTCONFServer with customizable handlers
//   - Test Hooks: Support for error path testing coverage
package mock
