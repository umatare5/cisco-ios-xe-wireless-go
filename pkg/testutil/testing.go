// Package testutil provides public testing utilities for Cisco IOS-XE Wireless Go client.
// Following Kubernetes patterns for public testing APIs that hide internal implementation details.
package testutil

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

// TestClient represents a test client with core functionality hidden.
// Following Docker patterns for interface-based testing abstraction.
type TestClient interface {
	// Core returns the underlying core client for service initialization.
	// The concrete type is *core.Client but returned as interface{} to hide internals.
	Core() interface{}
}

// MockServer represents a mock RESTCONF server for testing.
// Following Prometheus patterns for test server abstraction.
type MockServer interface {
	URL() string
	Close()
	NewTestClient(t *testing.T) TestClient
}

// testClientImpl implements TestClient interface hiding internal details.
type testClientImpl struct {
	client *core.Client
}

// Core returns the core client as interface{} to hide the concrete type.
func (tc *testClientImpl) Core() interface{} {
	return tc.client
}

// mockServerImpl implements MockServer interface wrapping httptest.Server.
type mockServerImpl struct {
	server *httptest.Server
}

// URL returns the mock server URL.
func (ms *mockServerImpl) URL() string {
	return ms.server.URL
}

// Close shuts down the mock server.
func (ms *mockServerImpl) Close() {
	ms.server.Close()
}

// NewTestClient creates a test client connected to this mock server.
// Following Kubernetes patterns for test client creation.
func (ms *mockServerImpl) NewTestClient(t *testing.T) TestClient {
	// Parse the server URL to get the host
	u, err := url.Parse(ms.server.URL)
	if err != nil {
		t.Fatalf("failed to parse server URL: %v", err)
	}

	client, err := core.New(u.Host, "test-token", core.WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	return &testClientImpl{client: client}
}

// NewTestClient creates a test client for the given mock server.
// Following Terraform patterns for test client abstraction.
func NewTestClient(server MockServer) TestClient {
	// Create a minimal test client without requiring *testing.T
	serverImpl, ok := server.(*mockServerImpl)
	if !ok {
		panic("testutil: server must be created with NewMockServer()")
	}

	// Parse the server URL to get the host
	u, err := url.Parse(serverImpl.server.URL)
	if err != nil {
		panic("testutil: failed to parse server URL: " + err.Error())
	}

	// Create core client using the mock server host
	client, err := core.New(u.Host, "test-token", core.WithInsecureSkipVerify(true))
	if err != nil {
		panic("testutil: failed to create test client: " + err.Error())
	}

	return &testClientImpl{client: client}
}

// NewMockServer creates a mock RESTCONF server with the specified responses.
// Following Kubernetes controller-runtime patterns for mock server creation.
func NewMockServer(responses map[string]string) MockServer {
	server := mock.NewRESTCONFSuccessServer(responses)
	return &mockServerImpl{server: server}
}

// NewMockErrorServer creates a mock server that returns errors for specified paths.
// Following Prometheus patterns for error simulation in tests.
func NewMockErrorServer(errorPaths []string, statusCode int) MockServer {
	server := mock.NewRESTCONFErrorServer(errorPaths, statusCode)
	return &mockServerImpl{server: server}
}

// NewMockServerWithCustomErrors creates a mock server with custom error messages.
// Following Kubernetes patterns for detailed error simulation in tests.
func NewMockServerWithCustomErrors(t *testing.T, errorConfig map[string]ErrorConfig) MockServer {
	server := mock.NewRESTCONFServer(t)
	for pathPattern, config := range errorConfig {
		server.AddHandler("GET", pathPattern, func() (int, string) {
			return config.StatusCode, config.ErrorMessage
		})
	}
	return &mockServerImpl{server: server.Server}
}

// ErrorConfig defines configuration for custom error responses.
type ErrorConfig struct {
	StatusCode   int
	ErrorMessage string
}
