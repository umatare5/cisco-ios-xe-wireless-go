// Package testutil provides testing utilities for the Cisco IOS-XE Wireless Go SDK.
//
// This package offers a unified mock server implementation that simulates RESTCONF API
// responses for comprehensive unit testing. The API uses functional options for flexible
// and composable test server configuration.
//
// # Main Features
//
// - MockServer interface for test server abstraction
// - TestClient interface for test client creation
// - NewMockServer with functional options for flexible server configuration
// - WithSuccessResponses, WithErrorResponses, WithCustomResponse options
// - Support for custom HTTP methods, status codes, and response bodies
// - RESTCONF path normalization and prefix handling
// - Enhanced testing integration with WithTesting option
//
// # Usage Examples
//
//	// Success response testing
//	server := testutil.NewMockServer(
//		testutil.WithSuccessResponses(map[string]string{
//			"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data": `{"status": "ok"}`,
//		}),
//	)
//
//	// Error response testing
//	server := testutil.NewMockServer(
//		testutil.WithErrorResponses([]string{
//			"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data",
//		}, 404),
//	)
//
//	// Complex response testing
//	server := testutil.NewMockServer(
//		testutil.WithCustomResponse("custom-path", testutil.ResponseConfig{
//			StatusCode: 202,
//			Body:       `{"custom": "response"}`,
//			Method:     "POST",
//		}),
//		testutil.WithTesting(t),
//	)
//
//	client := testutil.NewTestClient(server)
//	defer server.Close()
//
// # Known Limitations
//
// - Requires Go 1.21+ for full functionality
// - Mock servers use HTTPS with self-signed certificates
// - Path matching uses simple string contains logic
//
// # Error Handling
//
// All mock server functions return non-nil MockServer instances or panic on configuration errors.
// Use defer server.Close() to properly clean up test servers.
//
// # Requirements
//
// - Go 1.21+
// - testing package for enhanced server capabilities
package testutil
