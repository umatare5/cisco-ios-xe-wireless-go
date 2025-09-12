// Package testutil provides public testing utilities for the Cisco IOS-XE Wireless Go client library.
//
// This package serves as the public testing API that hides internal implementation details
// while providing essential testing functionality for service layers.
//
// # Design Principles
//
// Following enterprise Go project patterns (Kubernetes, Docker, Prometheus):
//
//   - Interface-based abstraction to hide internal types
//   - Minimal public surface area
//   - Clear separation between public and internal APIs
//   - Enterprise-grade testing patterns
//
// # Basic Usage
//
// Create a mock server and test client:
//
//	responses := map[string]string{
//		"some-endpoint": `{"result": "success"}`,
//	}
//	mockServer := testutil.NewMockServer(responses)
//	defer mockServer.Close()
//
//	testClient := testutil.NewTestClient(mockServer)
//	service := someservice.NewService(testClient.Core().(*core.Client))
//
// # Error Testing
//
// Create mock servers that return errors:
//
//	errorPaths := []string{"some-endpoint"}
//	errorServer := testutil.NewMockErrorServer(errorPaths, 404)
//	defer errorServer.Close()
//
// # Design Goals
//
//   - Enable service layer testing without internal package dependencies
//   - Provide consistent testing patterns across all services
//   - Maintain compatibility with CI/CD environments
//   - Support both success and error testing scenarios
//
// This package is part of Phase 3 of the Go conventions compliance initiative,
// resolving internal package boundary violations while maintaining test coverage.
package testutil
