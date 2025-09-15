package testutil

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

// TestClient represents a test client with core functionality hidden.
type TestClient interface {
	// Core returns the underlying core client for service initialization.
	// The concrete type is *core.Client but returned as interface{} to hide internals.
	Core() interface{}
}

// MockServer represents a mock RESTCONF server for testing.
type MockServer interface {
	URL() string
	Close()
}

// ResponseConfig defines configuration for custom responses.
type ResponseConfig struct {
	StatusCode int
	Body       string
	Method     string // HTTP method, defaults to "GET" if empty
}

// MockServerOption configures a MockServer.
type MockServerOption func(*mockServerConfig)

// mockServerConfig holds configuration for MockServer.
type mockServerConfig struct {
	testing         *testing.T
	successPaths    map[string]string         // path -> response body (200 OK)
	errorPaths      map[string]int            // path -> status code
	customResponses map[string]ResponseConfig // path -> full response config
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

// WithSuccessResponse adds a successful response for a specific path.
func WithSuccessResponse(path, body string) MockServerOption {
	return func(cfg *mockServerConfig) {
		if cfg.successPaths == nil {
			cfg.successPaths = make(map[string]string)
		}
		cfg.successPaths[path] = body
	}
}

// WithSuccessResponses adds multiple successful responses.
func WithSuccessResponses(responses map[string]string) MockServerOption {
	return func(cfg *mockServerConfig) {
		if cfg.successPaths == nil {
			cfg.successPaths = make(map[string]string)
		}
		for path, body := range responses {
			cfg.successPaths[path] = body
		}
	}
}

// WithErrorResponse adds an error response for a specific path.
func WithErrorResponse(path string, statusCode int) MockServerOption {
	return func(cfg *mockServerConfig) {
		if cfg.errorPaths == nil {
			cfg.errorPaths = make(map[string]int)
		}
		cfg.errorPaths[path] = statusCode
	}
}

// WithErrorResponses adds error responses for multiple paths.
func WithErrorResponses(paths []string, statusCode int) MockServerOption {
	return func(cfg *mockServerConfig) {
		if cfg.errorPaths == nil {
			cfg.errorPaths = make(map[string]int)
		}
		for _, path := range paths {
			cfg.errorPaths[path] = statusCode
		}
	}
}

// WithCustomResponse adds a fully customizable response.
func WithCustomResponse(path string, config ResponseConfig) MockServerOption {
	return func(cfg *mockServerConfig) {
		if cfg.customResponses == nil {
			cfg.customResponses = make(map[string]ResponseConfig)
		}
		cfg.customResponses[path] = config
	}
}

// WithTesting provides a testing.T instance for enhanced server capabilities.
func WithTesting(t *testing.T) MockServerOption {
	return func(cfg *mockServerConfig) {
		cfg.testing = t
	}
}

// NewTestClient creates a test client for the given mock server.
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

// NewMockServer creates a mock RESTCONF server with functional options.
// This is the unified constructor that supports all testing scenarios.
//
// Examples:
//
//	// Simple success responses
//	server := NewMockServer(WithSuccessResponses(map[string]string{"path": "response"}))
//
//	// Error responses
//	server := NewMockServer(WithErrorResponses([]string{"path"}, 404))
//
//	// Mixed responses
//	server := NewMockServer(
//	  WithSuccessResponse("success-path", `{"status":"ok"}`),
//	  WithErrorResponse("error-path", 500),
//	  WithTesting(t),
//	)
func NewMockServer(opts ...MockServerOption) MockServer {
	cfg := &mockServerConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	// If we have custom responses or testing context, use the flexible server
	if len(cfg.customResponses) > 0 || cfg.testing != nil {
		return newAdvancedMockServer(cfg)
	}

	// Handle simple success-only case
	if len(cfg.successPaths) > 0 && len(cfg.errorPaths) == 0 {
		server := NewRESTCONFSuccessServer(cfg.successPaths)
		return &mockServerImpl{server: server}
	}

	// Handle simple error-only case
	if len(cfg.errorPaths) > 0 && len(cfg.successPaths) == 0 {
		// Convert error paths map to slice and uniform status code
		paths := make([]string, 0, len(cfg.errorPaths))
		var statusCode int
		for path, status := range cfg.errorPaths {
			paths = append(paths, path)
			if statusCode == 0 {
				statusCode = status
			}
		}
		server := NewRESTCONFErrorServer(paths, statusCode)
		return &mockServerImpl{server: server}
	}

	// Mixed case - use advanced server
	return newAdvancedMockServer(cfg)
}

// newAdvancedMockServer creates a server using the flexible RESTCONFServer.
func newAdvancedMockServer(cfg *mockServerConfig) MockServer {
	t := cfg.testing
	if t == nil {
		// Create a minimal testing.T substitute for RESTCONFServer requirement
		t = &testing.T{}
	}

	server := NewRESTCONFServer(t)

	// Add success responses
	for path, body := range cfg.successPaths {
		server.AddHandler("GET", path, func() (int, string) {
			return 200, body
		})
	}

	// Add error responses
	for path, statusCode := range cfg.errorPaths {
		server.AddHandler("GET", path, func() (int, string) {
			return statusCode, ""
		})
	}

	// Add custom responses
	for path, config := range cfg.customResponses {
		method := config.Method
		if method == "" {
			method = "GET"
		}
		server.AddHandler(method, path, func() (int, string) {
			return config.StatusCode, config.Body
		})
	}

	return &mockServerImpl{server: server.Server}
}
