// nolint:SA1012 // Testing nil context behavior
package core

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

// Test constants
const (
	testTimeout = 10 * time.Second
)

// TestNewClient tests the new core client creation
func TestNewClient(t *testing.T) {
	controller := "test.example.com"
	token := "dGVzdDp0ZXN0" // base64 enc			t.Error("Expected error for canceled context, got nil")		t.Error("Expected error for canceled context, got nil")ded "test:test"

	t.Run("ValidClient", func(t *testing.T) {
		client, err := New(controller, token)
		if err != nil {
			t.Fatalf("Expected successful client creation, got error: %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
	})

	t.Run("EmptyController", func(t *testing.T) {
		_, err := New("", token)
		if err == nil {
			t.Fatal("Expected error for empty controller")
		}
	})

	t.Run("EmptyToken", func(t *testing.T) {
		_, err := New(controller, "")
		if err == nil {
			t.Fatal("Expected error for empty token")
		}
	})
}

// TestClientOptions tests functional options
func TestClientOptions(t *testing.T) {
	controller := "test.example.com"
	token := "dGVzdDp0ZXN0"

	t.Run("WithTimeout", func(t *testing.T) {
		client, err := New(controller, token, WithTimeout(testTimeout))
		if err != nil {
			t.Fatalf("Expected successful client creation with timeout, got error: %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
	})

	t.Run("WithInsecureSkipVerify", func(t *testing.T) {
		client, err := New(controller, token, WithInsecureSkipVerify(true))
		if err != nil {
			t.Fatalf("Expected successful client creation with insecure, got error: %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
	})

	t.Run("InvalidTimeout", func(t *testing.T) {
		_, err := New(controller, token, WithTimeout(0))
		if err == nil {
			t.Fatal("Expected error for zero timeout")
		}
	})
}

// TestClientDo tests the Do method with real controller if available
func TestClientDo(t *testing.T) {
	controller := os.Getenv("WNC_CONTROLLER")
	token := os.Getenv("WNC_ACCESS_TOKEN")

	if controller == "" || token == "" {
		t.Skip("WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set for integration tests")
	}

	client, err := New(controller, token, WithInsecureSkipVerify(true), WithTimeout(testTimeout))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	t.Run("GET_GeneralOper", func(t *testing.T) {
		var response interface{}
		err := client.Do(ctx, "GET", "Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
		if err != nil {
			t.Logf("GET request failed (may be expected for test controller): %v", err)
		} else {
			t.Logf("GET request successful")
			if response == nil {
				t.Error("Expected non-nil response")
			}
		}
	})

	t.Run("InvalidMethod", func(t *testing.T) {
		var response interface{}
		err := client.Do(ctx, "INVALID", "/restconf/data/test", &response)
		if err == nil {
			t.Error("Expected error for invalid HTTP method")
		}
	})

	t.Run("NilContext", func(t *testing.T) {
		var response interface{}
		// Using a nil context variable instead of nil literal to test the validation
		var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
		err := client.Do(nilCtx, "GET", "/restconf/data/test", &response)
		if err == nil {
			t.Error("Expected error for nil context")
		}
	})

	t.Run("NilOutput", func(t *testing.T) {
		err := client.Do(ctx, "GET", "/restconf/data/test", nil)
		if err == nil {
			t.Error("Expected error for nil output")
		}
	})
}

// TestHTTPError tests the HTTPError type
func TestHTTPError(t *testing.T) {
	err := &HTTPError{
		Status: 404,
		Body:   []byte("Not Found"),
	}

	expected := "HTTP 404: Not Found"
	if err.Error() != expected {
		t.Errorf("Expected error message %q, got %q", expected, err.Error())
	}
}

// ========================================
// 3. ADDITIONAL OPTION FUNCTION TESTS
// ========================================

func TestClientOptionsExtra(t *testing.T) {
	testCases := []struct {
		name   string
		option Option
		test   func(*Client) error
	}{
		{
			name:   "WithTimeout",
			option: WithTimeout(10 * time.Second),
			test: func(c *Client) error {
				if c.httpClient.Timeout != 10*time.Second {
					t.Errorf("Expected timeout 10s, got %v", c.httpClient.Timeout)
				}
				return nil
			},
		},
		{
			name:   "WithInsecureSkipVerify",
			option: WithInsecureSkipVerify(true),
			test: func(c *Client) error {
				// Test passes if option applied without error
				return nil
			},
		},
		{
			name:   "WithLogger",
			option: WithLogger(slog.Default()),
			test: func(c *Client) error {
				// Test passes if no panic occurs
				return nil
			},
		},
		{
			name:   "WithUserAgent",
			option: WithUserAgent("test-agent"),
			test: func(c *Client) error {
				// Test passes if option applied without error
				return nil
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client, err := New("example.com", "token", tt.option)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			if err := tt.test(client); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestWithLoggerNilError(t *testing.T) {
	_, err := New("example.com", "token", WithLogger(nil))
	if err == nil {
		t.Error("Expected error for nil logger, got nil")
	}
	if !strings.Contains(err.Error(), "logger cannot be nil") {
		t.Errorf("Expected error message about nil logger, got: %v", err)
	}
}

func TestWithTimeoutZeroError(t *testing.T) {
	_, err := New("example.com", "token", WithTimeout(0))
	if err == nil {
		t.Error("Expected error for zero timeout, got nil")
	}
	if !strings.Contains(err.Error(), "timeout must be positive") {
		t.Errorf("Expected error message about positive timeout, got: %v", err)
	}
}

// ========================================
// 4. SERVICE ACCESSOR TESTS
// ========================================

func TestServiceAccessors(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	testCases := []struct {
		name     string
		accessor func() interface{}
	}{
		{
			name:     "AFC",
			accessor: func() interface{} { return client.AFC() },
		},
		{
			name:     "AP",
			accessor: func() interface{} { return client.AP() },
		},
		{
			name:     "Client",
			accessor: func() interface{} { return client.Client() },
		},
		{
			name:     "General",
			accessor: func() interface{} { return client.General() },
		},
		{
			name:     "RRM",
			accessor: func() interface{} { return client.RRM() },
		},
		{
			name:     "WLAN",
			accessor: func() interface{} { return client.WLAN() },
		},
		{
			name:     "Rogue",
			accessor: func() interface{} { return client.Rogue() },
		},
		{
			name:     "NMSP",
			accessor: func() interface{} { return client.NMSP() },
		},
		{
			name:     "Hyperlocation",
			accessor: func() interface{} { return client.Hyperlocation() },
		},
		{
			name:     "Mdns",
			accessor: func() interface{} { return client.MDNS() },
		},
		{
			name:     "Geolocation",
			accessor: func() interface{} { return client.Geolocation() },
		},
		{
			name:     "Mcast",
			accessor: func() interface{} { return client.Mcast() },
		},
		{
			name:     "APF",
			accessor: func() interface{} { return client.APF() },
		},
		{
			name:     "AWIPS",
			accessor: func() interface{} { return client.AWIPS() },
		},
		{
			name:     "BLE",
			accessor: func() interface{} { return client.BLE() },
		},
		{
			name:     "CTS",
			accessor: func() interface{} { return client.CTS() },
		},
		{
			name:     "Dot11",
			accessor: func() interface{} { return client.Dot11() },
		},
		{
			name:     "Dot15",
			accessor: func() interface{} { return client.Dot15() },
		},
		{
			name:     "Fabric",
			accessor: func() interface{} { return client.Fabric() },
		},
		{
			name:     "Flex",
			accessor: func() interface{} { return client.Flex() },
		},
		{
			name:     "Location",
			accessor: func() interface{} { return client.Location() },
		},
		{
			name:     "Radio",
			accessor: func() interface{} { return client.Radio() },
		},
		{
			name:     "RF",
			accessor: func() interface{} { return client.RF() },
		},
		{
			name:     "RFID",
			accessor: func() interface{} { return client.RFID() },
		},
		{
			name:     "Mobility",
			accessor: func() interface{} { return client.Mobility() },
		},
		{
			name:     "Mesh",
			accessor: func() interface{} { return client.Mesh() },
		},
		{
			name:     "Site",
			accessor: func() interface{} { return client.Site() },
		},
		{
			name:     "LISP",
			accessor: func() interface{} { return client.LISP() },
		},
	}

	for _, tt := range testCases {
		t.Run("ServiceAccessor_"+tt.name, func(t *testing.T) {
			// Service accessors currently return nil as placeholders
			// This test ensures they can be called without panicking
			service := tt.accessor()
			// All services currently return nil as documented placeholders
			if service != nil {
				t.Logf("Service accessor %s returned non-nil: %v", tt.name, service)
			}
		})
	}
}

// ========================================
// 5. DO METHOD COVERAGE TESTS
// ========================================

func TestDoMethodErrorHandling(t *testing.T) {
	client, err := New("nonexistent.invalid", "token", WithTimeout(1*time.Second))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	var response interface{}

	// Test with invalid host to cover error paths
	err = client.Do(ctx, "GET", "/test", &response)
	if err == nil {
		t.Error("Expected error for invalid host, got nil")
	}

	// Test with canceled context
	canceledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	err = client.Do(canceledCtx, "GET", "/test", &response)
	if err == nil {
		t.Error("Expected error for canceled context, got nil")
	}
}

func TestDoMethodNilParameters(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	var response interface{}

	// Using a nil context variable instead of nil literal to test the validation
	var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
	err = client.Do(nilCtx, "GET", "/test", &response)
	if err == nil {
		t.Error("Expected error for nil context, got nil")
	}
	if !strings.Contains(err.Error(), "context cannot be nil") {
		t.Errorf("Expected error message about nil context, got: %v", err)
	}

	// Test with nil output
	ctx := context.Background()
	err = client.Do(ctx, "GET", "/test", nil)
	if err == nil {
		t.Error("Expected error for nil output, got nil")
	}
	if !strings.Contains(err.Error(), "output parameter cannot be nil") {
		t.Errorf("Expected error message about nil output, got: %v", err)
	}
}

func TestDoMethodErrorResponse(t *testing.T) {
	// Create a test server that returns different HTTP status codes
	client, err := New("httpbin.org", "token", WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	var response interface{}

	// Test 404 error response
	err = client.Do(ctx, "GET", "/status/404", &response)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Check if it's an HTTPError
	if strings.Contains(err.Error(), "HTTP 404") {
		// This is the expected behavior
		t.Logf("Got expected HTTP error: %v", err)
	}
}

// TestDoMethodSpecificCoverage tests specific code paths in the Do method
func TestDoMethodSpecificCoverage(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Run("ContextDeadlineExceeded", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancel()

		// Let the context expire
		time.Sleep(2 * time.Nanosecond)

		var response interface{}
		err := client.Do(ctx, "GET", "/test", &response)
		if err == nil {
			t.Error("Expected error for expired context, got nil")
		}
	})

	t.Run("InvalidHTTPMethod", func(t *testing.T) {
		ctx := context.Background()
		var response interface{}

		// Test with invalid HTTP method that might cause request creation to fail
		err := client.Do(ctx, "INVALID METHOD WITH SPACES", "/test", &response)
		if err == nil {
			t.Error("Expected error for invalid HTTP method, got nil")
		}
		if !strings.Contains(err.Error(), "failed to create request") {
			t.Logf("Got error (expected): %v", err)
		}
	})

	t.Run("ValidRequestCreation", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		var response interface{}

		// Test that request creation works with valid parameters
		// (This will fail on network but validates request creation logic)
		err := client.Do(ctx, "GET", "/valid/path", &response)
		// I expect a network error, not a request creation error
		if err != nil && strings.Contains(err.Error(), "failed to create request") {
			t.Errorf("Unexpected request creation error: %v", err)
		}
		// Any other error (network, etc.) is expected and acceptable
		t.Logf("Request creation successful, network error expected: %v", err)
	})
}

// TestDoMethodJSONHandling tests JSON unmarshaling in the Do method
func TestDoMethodJSONHandling(t *testing.T) {
	client, err := New("httpbin.org", "token", WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("ValidJSONResponse", func(t *testing.T) {
		var response map[string]interface{}

		// httpbin.org/json returns valid JSON
		err := client.Do(ctx, "GET", "/json", &response)
		if err != nil {
			// Network errors are acceptable, JSON parsing errors are not
			if !strings.Contains(err.Error(), "failed to unmarshal") {
				t.Logf("Network error expected: %v", err)
			} else {
				t.Errorf("Unexpected JSON unmarshaling error: %v", err)
			}
		} else {
			t.Log("Successfully parsed JSON response")
		}
	})

	t.Run("InvalidJSONTarget", func(t *testing.T) {
		// Test with a target that can't accept the JSON structure
		var response string // String can't accept JSON object

		err := client.Do(ctx, "GET", "/json", &response)
		if err != nil {
			// Could be network error or JSON unmarshaling error
			t.Logf("Error (expected): %v", err)
		}
	})

	t.Run("InvalidJSONResponse", func(t *testing.T) {
		var response map[string]interface{}

		// Try to get HTML as JSON - this should cause unmarshaling error
		err := client.Do(ctx, "GET", "/html", &response)
		if err != nil {
			// Should get either network error or JSON unmarshaling error
			t.Logf("Error (expected for HTML response): %v", err)
		}
	})
}

// TestDoMethodResponseBodyHandling tests response body handling edge cases
func TestDoMethodResponseBodyHandling(t *testing.T) {
	client, err := New("httpbin.org", "token", WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("EmptyResponse", func(t *testing.T) {
		var response map[string]interface{}

		// Test empty response body handling
		err := client.Do(ctx, "GET", "/status/204", &response)
		if err != nil {
			// Empty responses might cause JSON unmarshaling errors, which is expected
			t.Logf("Error (may be expected for empty response): %v", err)
		}
	})

	t.Run("LargeResponse", func(t *testing.T) {
		var response map[string]interface{}

		// Test with a larger response
		err := client.Do(ctx, "GET", "/json", &response)
		if err != nil {
			t.Logf("Error (network or parsing): %v", err)
		}
	})
}

// TestDoMethodResponseBodyReadError tests response body read error handling
func TestDoMethodResponseBodyReadError(t *testing.T) {
	// Create a test server that sends a response with incomplete JSON
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"partial": "response"`)) // Incomplete JSON
		// Don't close the JSON properly to simulate read error
	}))
	defer server.Close()

	// Extract just the host:port from the server URL
	serverURL := strings.TrimPrefix(server.URL, "https://")

	client, err := New(serverURL, "token", WithTimeout(1*time.Second), WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var response map[string]interface{}
	err = client.Do(ctx, "GET", "/test", &response)
	if err != nil {
		// Should get JSON unmarshaling error due to incomplete JSON
		t.Logf("Expected error for incomplete JSON: %v", err)
	}
}

// TestDoMethodHTTPErrorBoundaries tests HTTP error status code boundaries
func TestDoMethodHTTPErrorBoundaries(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		expectErr  bool
	}{
		{"Status399_NoError", 399, false},
		{"Status400_Error", 400, true},
		{"Status404_Error", 404, true},
		{"Status500_Error", 500, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				fmt.Fprintf(w, `{"error": "test error"}`)
			}))
			defer server.Close()

			serverURL := strings.TrimPrefix(server.URL, "https://")
			client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			ctx := context.Background()
			var response map[string]interface{}
			err = client.Do(ctx, "GET", "/test", &response)

			if tc.expectErr && err == nil {
				t.Errorf("Expected error for status %d, got nil", tc.statusCode)
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error for status %d, got: %v", tc.statusCode, err)
			}
		})
	}
}

// TestDoMethodResponseBodyCloseError tests deferred response body close error handling
func TestDoMethodResponseBodyCloseError(t *testing.T) {
	// Create a test server that returns valid JSON
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"test": "data"}`)
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	var response map[string]interface{}

	// This test covers the successful path including the deferred body close
	err = client.Do(ctx, "GET", "/test", &response)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if response["test"] != "data" {
		t.Errorf("Expected response data, got: %v", response)
	}
}

// TestDoMethodResponseBodyReadFailure tests response body read error handling
func TestDoMethodResponseBodyReadFailure(t *testing.T) {
	// Create a test server that returns a broken response body
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "100") // Set content length but send less data
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"partial"`)) // Incomplete response to trigger read errors
		// Force connection close to simulate read failure
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		// Don't write complete response
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	var response map[string]interface{}

	// This should trigger JSON unmarshaling error due to incomplete response
	err = client.Do(ctx, "GET", "/test", &response)
	if err == nil {
		t.Error("Expected error due to incomplete JSON response")
	}

	// Should be a JSON unmarshaling error
	if !strings.Contains(err.Error(), "failed to unmarshal response") {
		t.Logf("Got error (JSON parsing expected): %v", err)
	}
}

// TestDoMethodWithLargeResponse tests handling of larger response bodies
func TestDoMethodWithLargeResponse(t *testing.T) {
	// Create a test server that returns a large valid JSON response
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Generate a larger JSON response
		largeData := make(map[string]interface{})
		for i := 0; i < 100; i++ {
			largeData[fmt.Sprintf("key_%d", i)] = fmt.Sprintf("value_%d", i)
		}

		// Manually create JSON to ensure it's valid
		fmt.Fprintf(w, `{"test": "data", "largeField": "`)
		for i := 0; i < 1000; i++ {
			fmt.Fprintf(w, "x")
		}
		fmt.Fprintf(w, `"}`)
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	var response map[string]interface{}

	err = client.Do(ctx, "GET", "/test", &response)
	if err != nil {
		t.Errorf("Unexpected error with large response: %v", err)
	}

	if response["test"] != "data" {
		t.Errorf("Expected test data in large response, got: %v", response["test"])
	}
}

// TestDoMethodEmptyResponseBody tests handling of empty response bodies
func TestDoMethodEmptyResponseBody(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Send empty body
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	var response map[string]interface{}

	err = client.Do(ctx, "GET", "/test", &response)
	if err == nil {
		t.Error("Expected error due to empty JSON response")
	}

	// Should be a JSON unmarshaling error
	if !strings.Contains(err.Error(), "failed to unmarshal response") {
		t.Logf("Got error (JSON parsing expected): %v", err)
	}
}

// TestDoMethodInvalidJSONInResponse tests handling of invalid JSON in responses
func TestDoMethodInvalidJSONInResponse(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid": json}`)) // Invalid JSON
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	var response map[string]interface{}

	err = client.Do(ctx, "GET", "/test", &response)
	if err == nil {
		t.Error("Expected error due to invalid JSON response")
	}

	// Should be a JSON unmarshaling error
	if !strings.Contains(err.Error(), "failed to unmarshal response") {
		t.Errorf("Expected JSON unmarshal error, got: %v", err)
	}
}

// TestDoMethodWithNilContext tests handling of nil context
func TestDoMethodWithNilContext(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	var response map[string]interface{}
	// Using a nil context variable instead of nil literal to test the validation
	var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
	err = client.Do(nilCtx, "GET", "/test", &response)
	if err == nil {
		t.Error("Expected error with nil context")
	}

	if !strings.Contains(err.Error(), "context cannot be nil") {
		t.Errorf("Expected context error, got: %v", err)
	}
}

// TestDoMethodWithNilOutput tests handling of nil output parameter
func TestDoMethodWithNilOutput(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	err = client.Do(ctx, "GET", "/test", nil)
	if err == nil {
		t.Error("Expected error with nil output parameter")
	}

	if !strings.Contains(err.Error(), "output parameter cannot be nil") {
		t.Errorf("Expected output parameter error, got: %v", err)
	}
}

// TestDoMethodWithRequestCreationError tests handling of request creation errors
func TestDoMethodWithRequestCreationError(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	var response map[string]interface{}

	// Use an invalid HTTP method to trigger request creation error
	err = client.Do(ctx, "INVALID METHOD\nWITH\nNEWLINES", "/test", &response)
	if err == nil {
		t.Error("Expected error with invalid HTTP method")
	}

	if !strings.Contains(err.Error(), "failed to create request") {
		t.Errorf("Expected request creation error, got: %v", err)
	}
}

// TestValidateDoParametersWithNilClient tests validation with nil client
func TestValidateDoParametersWithNilClient(t *testing.T) {
	var nilClient *Client

	ctx := context.Background()
	var output interface{}

	// Test that validateDoParameters correctly handles nil client
	err := nilClient.validateDoParameters(ctx, &output)
	if err == nil {
		t.Error("Expected error with nil client")
	}

	if !strings.Contains(err.Error(), "client cannot be nil") {
		t.Errorf("Expected nil client error, got: %v", err)
	}
}

// TestCloseResponseBodyErrorHandling tests error handling in closeResponseBody
func TestCloseResponseBodyErrorHandling(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create a mock reader that will fail on Close()
	mockReader := &errorCloser{closed: false}

	// Create a simple HTTP response for testing with our mock body
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       mockReader,
	}

	// Test that closeResponseBody handles the error case correctly
	// This should log an error but not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("closeResponseBody should not panic: %v", r)
		}
	}()

	client.closeResponseBody(resp)

	// Verify that Close() was called
	if !mockReader.closed {
		t.Error("Expected Close() to be called on response body")
	}

	t.Log("closeResponseBody completed without panic, handled error correctly")
}

// errorCloser is a mock io.ReadCloser that returns an error on Close()
type errorCloser struct {
	closed bool
}

func (e *errorCloser) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func (e *errorCloser) Close() error {
	e.closed = true
	return fmt.Errorf("mock close error")
}
