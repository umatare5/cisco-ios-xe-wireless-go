package wnc

import (
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"
	"time"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestHTTPTimeoutConstants tests HTTP timeout constants
func TestHTTPTimeoutConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Duration
		expected time.Duration
	}{
		{"DefaultTLSHandshakeTimeout", DefaultTLSHandshakeTimeout, 10 * time.Second},
		{"DefaultResponseHeaderTimeout", DefaultResponseHeaderTimeout, 10 * time.Second},
		{"DefaultIdleConnTimeout", DefaultIdleConnTimeout, 90 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be %v, got %v", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// TestHTTPHeaderConstants tests HTTP header constants
func TestHTTPHeaderConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"HTTPHeaderKeyAuthorization", HTTPHeaderKeyAuthorization, "Authorization"},
		{"HTTPHeaderKeyAccept", HTTPHeaderKeyAccept, "Accept"},
		{"HTTPHeaderKeyUserAgent", HTTPHeaderKeyUserAgent, "User-Agent"},
		{"HTTPHeaderKeyContentType", HTTPHeaderKeyContentType, "Content-Type"},
		{"HTTPHeaderValueBasicPrefix", HTTPHeaderValueBasicPrefix, "Basic "},
		{"HTTPHeaderValueYANGData", HTTPHeaderValueYANGData, "application/yang-data+json"},
		{"HTTPHeaderUserAgent", HTTPHeaderUserAgent, "wnc-go-client/1.0"},
		{"HTTPHeaderAccept", HTTPHeaderAccept, "application/yang-data+json"},
		{"HTTPHeaderContentType", HTTPHeaderContentType, "application/yang-data+json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be '%s', got '%s'", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// =============================================================================
// 2. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// TestCreateHTTPTransport tests HTTP transport creation
func TestCreateHTTPTransport(t *testing.T) {
	tests := []struct {
		name               string
		insecureSkipVerify bool
	}{
		{"SecureTransport", false},
		{"InsecureTransport", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				insecureSkipVerify: tt.insecureSkipVerify,
			}

			transport := client.createHTTPTransport()

			if transport == nil {
				t.Fatal("Expected transport to be created")
			}

			if transport.TLSClientConfig == nil {
				t.Fatal("Expected TLS config to be set")
			}

			if transport.TLSClientConfig.InsecureSkipVerify != tt.insecureSkipVerify {
				t.Errorf("Expected InsecureSkipVerify to be %v, got %v",
					tt.insecureSkipVerify, transport.TLSClientConfig.InsecureSkipVerify)
			}

			// Test timeout settings
			if transport.TLSHandshakeTimeout != DefaultTLSHandshakeTimeout {
				t.Errorf("Expected TLSHandshakeTimeout to be %v, got %v",
					DefaultTLSHandshakeTimeout, transport.TLSHandshakeTimeout)
			}

			if transport.ResponseHeaderTimeout != DefaultResponseHeaderTimeout {
				t.Errorf("Expected ResponseHeaderTimeout to be %v, got %v",
					DefaultResponseHeaderTimeout, transport.ResponseHeaderTimeout)
			}

			if transport.IdleConnTimeout != DefaultIdleConnTimeout {
				t.Errorf("Expected IdleConnTimeout to be %v, got %v",
					DefaultIdleConnTimeout, transport.IdleConnTimeout)
			}

			// Test HTTP/2 and other settings
			if transport.ForceAttemptHTTP2 {
				t.Error("Expected ForceAttemptHTTP2 to be false")
			}

			if transport.DisableKeepAlives {
				t.Error("Expected DisableKeepAlives to be false")
			}

			if transport.DisableCompression {
				t.Error("Expected DisableCompression to be false")
			}
		})
	}
}

// TestBuildRequestURL tests request URL construction
func TestBuildRequestURL(t *testing.T) {
	tests := []struct {
		name       string
		controller string
		endpoint   string
		expected   string
	}{
		{"SimpleEndpoint", "example.com", "/api/test", "https://example.com/api/test"},
		{"EndpointWithQuery", "192.168.1.100", "/restconf/data?fields=all", "https://192.168.1.100/restconf/data?fields=all"},
		{"ControllerWithPort", "wnc.local:8443", "/status", "https://wnc.local:8443/status"},
		{"EmptyEndpoint", "test.com", "", "https://test.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				controller: tt.controller,
			}

			result := client.buildRequestURL(tt.endpoint)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// TestBuildHTTPHeaders tests HTTP header construction
func TestBuildHTTPHeaders(t *testing.T) {
	tests := []struct {
		name         string
		accessToken  string
		expectedAuth string
	}{
		{"ValidToken", "dGVzdDp0ZXN0", "Basic dGVzdDp0ZXN0"},
		{"EmptyToken", "", "Basic "},
		{"LongToken", "YWRtaW46cGFzc3dvcmQxMjM0NTY3ODkw", "Basic YWRtaW46cGFzc3dvcmQxMjM0NTY3ODkw"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				accessToken: tt.accessToken,
			}

			headers := client.buildHTTPHeaders()

			// Check required headers
			if headers[HTTPHeaderKeyAuthorization] != tt.expectedAuth {
				t.Errorf("Expected Authorization header '%s', got '%s'",
					tt.expectedAuth, headers[HTTPHeaderKeyAuthorization])
			}

			if headers[HTTPHeaderKeyAccept] != HTTPHeaderValueYANGData {
				t.Errorf("Expected Accept header '%s', got '%s'",
					HTTPHeaderValueYANGData, headers[HTTPHeaderKeyAccept])
			}

			if headers[HTTPHeaderKeyUserAgent] != HTTPHeaderUserAgent {
				t.Errorf("Expected User-Agent header '%s', got '%s'",
					HTTPHeaderUserAgent, headers[HTTPHeaderKeyUserAgent])
			}

			// Should have exactly 3 headers
			if len(headers) != 3 {
				t.Errorf("Expected 3 headers, got %d", len(headers))
			}
		})
	}
}

// TestBuildHTTPHeadersWithAcceptType tests HTTP headers with custom Accept type
func TestBuildHTTPHeadersWithAcceptType(t *testing.T) {
	tests := []struct {
		name       string
		acceptType string
	}{
		{"JSONAccept", "application/json"},
		{"XMLAccept", "application/xml"},
		{"YANGAccept", HTTPHeaderValueYANGData},
		{"EmptyAccept", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				accessToken: "test-token",
			}

			headers := client.buildHTTPHeadersWithAcceptType(tt.acceptType)

			if headers[HTTPHeaderKeyAccept] != tt.acceptType {
				t.Errorf("Expected Accept header '%s', got '%s'",
					tt.acceptType, headers[HTTPHeaderKeyAccept])
			}

			// Other headers should remain the same
			if headers[HTTPHeaderKeyAuthorization] != "Basic test-token" {
				t.Errorf("Expected Authorization header 'Basic test-token', got '%s'",
					headers[HTTPHeaderKeyAuthorization])
			}

			if headers[HTTPHeaderKeyUserAgent] != HTTPHeaderUserAgent {
				t.Errorf("Expected User-Agent header '%s', got '%s'",
					HTTPHeaderUserAgent, headers[HTTPHeaderKeyUserAgent])
			}
		})
	}
}

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// TestSetRequestHeaders tests request header setting
func TestSetRequestHeaders(t *testing.T) {
	client := &Client{}
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	headers := map[string]string{
		"Authorization": "Basic test-token",
		"Accept":        "application/json",
		"Custom-Header": "custom-value",
	}

	client.setRequestHeaders(req, headers)

	// Check that all headers were set
	for key, expectedValue := range headers {
		actualValue := req.Header.Get(key)
		if actualValue != expectedValue {
			t.Errorf("Expected header '%s' to be '%s', got '%s'",
				key, expectedValue, actualValue)
		}
	}
}

// TestHandleHTTPError tests HTTP error handling
func TestHandleHTTPError(t *testing.T) {
	tests := []struct {
		name         string
		statusCode   int
		responseBody []byte
		requestURL   string
		expectedErr  error
	}{
		{"AuthenticationError", 401, []byte("Unauthorized"), "https://test.com/api", ErrAuthenticationFailed},
		{"AccessForbiddenError", 403, []byte("Forbidden"), "https://test.com/api", ErrAccessForbidden},
		{"NotFoundError", 404, []byte("Not Found"), "https://test.com/api", ErrResourceNotFound},
		{"SuccessStatus", 200, []byte("OK"), "https://test.com/api", nil},
		{"SuccessStatusCreated", 201, []byte("Created"), "https://test.com/api", nil},
		{"SuccessStatusAccepted", 202, []byte("Accepted"), "https://test.com/api", nil},
		{"SuccessStatusNoContent", 204, []byte(""), "https://test.com/api", nil},
		{"ServerError", 500, []byte("Internal Server Error"), "https://test.com/api", nil},      // Should return error
		{"BadGateway", 502, []byte("Bad Gateway"), "https://test.com/api", nil},                 // Should return error
		{"ServiceUnavailable", 503, []byte("Service Unavailable"), "https://test.com/api", nil}, // Should return error
		{"BadRequest", 400, []byte("Bad Request"), "https://test.com/api", nil},                 // Should return error
		{"Conflict", 409, []byte("Conflict"), "https://test.com/api", nil},                      // Should return error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a no-op logger using slog
			logger := slog.New(slog.NewTextHandler(io.Discard, nil))

			client := &Client{
				logger: logger,
			}

			err := client.handleHTTPError(tt.statusCode, tt.responseBody, tt.requestURL)

			if tt.expectedErr == nil {
				// For success status codes (200-299), expect no error
				if tt.statusCode >= 200 && tt.statusCode < 300 {
					if err != nil {
						t.Errorf("Expected no error for success status %d, got %v", tt.statusCode, err)
					}
				} else {
					// For non-success status codes, expect an error
					if err == nil {
						t.Errorf("Expected an error for status %d, got nil", tt.statusCode)
					} else if !strings.Contains(err.Error(), "HTTP error") {
						t.Errorf("Expected HTTP error message for status %d, got %v", tt.statusCode, err)
					}
				}
			} else {
				if err == nil {
					t.Errorf("Expected error %v, got nil", tt.expectedErr)
				} else if err != tt.expectedErr {
					t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
				}
			}
		})
	}
}

// TestHeaderValueConsistency tests that header value constants are consistent
func TestHeaderValueConsistency(t *testing.T) {
	// Test that HTTPHeaderAccept equals HTTPHeaderValueYANGData
	if HTTPHeaderAccept != HTTPHeaderValueYANGData {
		t.Errorf("Expected HTTPHeaderAccept (%s) to equal HTTPHeaderValueYANGData (%s)",
			HTTPHeaderAccept, HTTPHeaderValueYANGData)
	}

	// Test that HTTPHeaderContentType equals HTTPHeaderValueYANGData
	if HTTPHeaderContentType != HTTPHeaderValueYANGData {
		t.Errorf("Expected HTTPHeaderContentType (%s) to equal HTTPHeaderValueYANGData (%s)",
			HTTPHeaderContentType, HTTPHeaderValueYANGData)
	}

	// Test that Basic prefix ends with space
	if !strings.HasSuffix(HTTPHeaderValueBasicPrefix, " ") {
		t.Errorf("Expected HTTPHeaderValueBasicPrefix (%s) to end with space", HTTPHeaderValueBasicPrefix)
	}
}
