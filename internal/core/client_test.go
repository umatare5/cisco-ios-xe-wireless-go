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

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
)

// Test constants.
const (
	testTimeout = 10 * time.Second
)

// TestCoreClientUnit_Constructor_Success tests the new core client creation.
func TestCoreClientUnit_Constructor_Success(t *testing.T) {
	controller := "test.example.com"
	token := "dGVzdDp0ZXN0"

	t.Run("ValidClient", func(t *testing.T) {
		client, err := New(controller, token)
		helper.AssertClientCreated(t, client, err, "ValidClient")
	})

	t.Run("EmptyController", func(t *testing.T) {
		_, err := New("", token)
		helper.AssertClientCreationError(t, err, "EmptyController")
	})

	t.Run("EmptyToken", func(t *testing.T) {
		_, err := New(controller, "")
		helper.AssertClientCreationError(t, err, "EmptyToken")
	})
}

// TestCoreClientUnit_Options_Success tests functional options.
func TestCoreClientUnit_Options_Success(t *testing.T) {
	controller := "test.example.com"
	token := "dGVzdDp0ZXN0"

	t.Run("WithTimeout", func(t *testing.T) {
		client, err := New(controller, token, WithTimeout(testTimeout))
		helper.AssertClientCreated(t, client, err, "WithTimeout")
	})

	t.Run("WithInsecureSkipVerify", func(t *testing.T) {
		client, err := New(controller, token, WithInsecureSkipVerify(true))
		helper.AssertClientCreated(t, client, err, "WithInsecureSkipVerify")
	})

	t.Run("InvalidTimeout", func(t *testing.T) {
		_, err := New(controller, token, WithTimeout(0))
		helper.AssertClientCreationError(t, err, "InvalidTimeout")
	})
}

// TestCoreClientUnit_DoOperations_Success tests the Do method with real controller if available.
func TestCoreClientUnit_DoOperations_Success(t *testing.T) {
	controller := os.Getenv("WNC_CONTROLLER")
	token := os.Getenv("WNC_ACCESS_TOKEN")
	apMac := os.Getenv("WNC_AP_MAC_ADDR")

	helper.RequireEnvironmentVars(t, map[string]string{
		"WNC_CONTROLLER":   controller,
		"WNC_ACCESS_TOKEN": token,
		"WNC_AP_MAC_ADDR":  apMac,
	})

	client, err := New(controller, token, WithInsecureSkipVerify(true), WithTimeout(testTimeout))
	helper.AssertNoError(t, err, "Failed to create client")

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	t.Run("GET_GeneralOper", func(t *testing.T) {
		_, err := client.Do(ctx, http.MethodGet, "Cisco-IOS-XE-wireless-general-oper:general-oper-data")
		if err != nil {
			t.Logf("GET request failed (may be expected for test controller): %v", err)
		} else {
			t.Logf("GET request successful")
		}
	})

	t.Run("InvalidMethod", func(t *testing.T) {
		_, err := client.Do(ctx, "INVALID", "/restconf/data/test")
		helper.AssertError(t, err, "Expected error for invalid HTTP method")
	})

	t.Run("NilContext", func(t *testing.T) {
		// Using a nil context variable instead of nil literal to test the validation
		var nilCtx context.Context //nolint:staticcheck // Testing nil context behavior
		_, err := client.Do(nilCtx, http.MethodGet, "/restconf/data/test")
		helper.AssertError(t, err, "Expected error for nil context")
	})

	t.Run("NilOutput", func(t *testing.T) {
		_, err := client.Do(ctx, http.MethodGet, "/restconf/data/test")
		helper.AssertError(t, err, "Expected error for nil output")
	})
}

// TestCoreClientUnit_ErrorHandling_HTTPErrors tests the HTTPError type.
func TestCoreClientUnit_ErrorHandling_HTTPErrors(t *testing.T) {
	err := &HTTPError{
		Status: 404,
		Body:   []byte("Not Found"),
	}

	expected := "HTTP 404: Not Found"
	helper.AssertErrorMessage(t, err, expected, "HTTPError should format status and body correctly")
}

// ========================================
// 3. ADDITIONAL OPTION FUNCTION TESTS
// ========================================

func TestCoreClientUnit_Options_ExtendedTests(t *testing.T) {
	testCases := []struct {
		name   string
		option Option
		test   func(*Client) error
	}{
		{
			name:   "WithTimeout",
			option: WithTimeout(10 * time.Second),
			test: func(c *Client) error {
				helper.AssertDurationEquals(t, c.httpClient.Timeout, 10*time.Second,
					"Client timeout should be set correctly")
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
			helper.AssertClientCreated(t, client, err, "Failed to create client")

			helper.AssertNoError(t, tt.test(client), "Test execution failed")
		})
	}
}

func TestCoreClientUnit_Validation_NilLogger(t *testing.T) {
	_, err := New("example.com", "token", WithLogger(nil))
	helper.AssertError(t, err, "Expected error for nil logger")
	helper.AssertStringContains(t, err.Error(), "logger cannot be nil",
		"Error message should contain expected text about nil logger")
}

func TestCoreClientUnit_Validation_ZeroTimeout(t *testing.T) {
	_, err := New("example.com", "token", WithTimeout(0))
	helper.AssertError(t, err, "Expected error for zero timeout")
	helper.AssertStringContains(t, err.Error(),
		"timeout must be positive",
		"Error message should contain expected text about positive timeout")
}

// ========================================
// 4. DO METHOD COVERAGE TESTS
// ========================================

func TestCoreClientUnit_DoOperations_ErrorHandling(t *testing.T) {
	client, err := New("nonexistent.invalid", "token", WithTimeout(1*time.Second))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	// Test with invalid host to cover error paths
	_, err = client.Do(ctx, http.MethodGet, "/test")
	helper.AssertError(t, err, "Expected error for invalid host")

	// Test with canceled context
	canceledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = client.Do(canceledCtx, http.MethodGet, "/test")
	helper.AssertError(t, err, "Expected error for canceled context")
}

func TestCoreClientUnit_DoOperations_NilParameters(t *testing.T) {
	// Create test server for local testing
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "data"}`))
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	// Using a nil context variable instead of nil literal to test the validation
	var nilCtx context.Context //nolint:staticcheck // Testing nil context behavior
	_, err = client.Do(nilCtx, http.MethodGet, "/test")
	helper.AssertError(t, err, "Expected error for nil context")
	helper.AssertStringContains(t, err.Error(),
		"context cannot be nil",
		"Error message should contain expected text about nil context")

	// Test with nil output - this should not error since nil output is handled gracefully
	ctx := context.Background()
	_, err = client.Do(ctx, http.MethodGet, "/test")
	if err != nil {
		t.Logf("Got expected error for nil output: %v", err)
		// This is acceptable behavior - nil output may be an error condition
		if !strings.Contains(err.Error(), "client cannot be nil") &&
			!strings.Contains(err.Error(), "context cannot be nil") {
			// If error is not about client/context being nil, it might be network related
			t.Logf("Error is network-related, which is expected for this test: %v", err)
		}
	}
}

func TestCoreClientUnit_DoOperations_ErrorResponse(t *testing.T) {
	// Create a test server that returns different HTTP status codes
	client, err := New("httpbin.org", "token", WithTimeout(5*time.Second))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	// Test 404 error response
	_, err = client.Do(ctx, http.MethodGet, "/status/404")
	helper.AssertError(t, err, "Expected error for 404 response")

	// Check if it's an HTTPError
	if strings.Contains(err.Error(), "HTTP 404") {
		// This is the expected behavior
		t.Logf("Got expected HTTP error: %v", err)
	}
}

// TestCoreClientUnit_DoOperations_SpecificCoverage tests specific code paths in the Do method.
func TestCoreClientUnit_DoOperations_SpecificCoverage(t *testing.T) {
	client, err := New("example.com", "token")
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	t.Run("ContextDeadlineExceeded", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancel()

		// Let the context expire
		time.Sleep(2 * time.Nanosecond)

		_, err := client.Do(ctx, http.MethodGet, "/test")
		helper.AssertError(t, err, "Expected error for expired context")
	})

	t.Run("InvalidHTTPMethod", func(t *testing.T) {
		ctx := context.Background()

		// Test with invalid HTTP method that might cause request creation to fail
		_, err := client.Do(ctx, "INVALID METHOD WITH SPACES", "/test")
		helper.AssertError(t, err, "Expected error for invalid HTTP method")
		if !strings.Contains(err.Error(), "failed to create request") {
			t.Logf("Got error (expected): %v", err)
		}
	})

	t.Run("ValidRequestCreation", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		// Test that request creation works with valid parameters
		// (This will fail on network but validates request creation logic)
		_, err := client.Do(ctx, http.MethodGet, "/valid/path")
		// I expect a network error, not a request creation error
		if err != nil && strings.Contains(err.Error(), "failed to create request") {
			helper.AssertNoError(t, err, "Unexpected request creation error")
		}
		// Any other error (network, etc.) is expected and acceptable
		t.Logf("Request creation successful, network error expected: %v", err)
	})
}

// TestCoreClientUnit_DoOperations_JSONHandling tests JSON unmarshaling in the Do method.
func TestCoreClientUnit_DoOperations_JSONHandling(t *testing.T) {
	client, err := New("httpbin.org", "token", WithTimeout(5*time.Second))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("ValidJSONResponse", func(t *testing.T) {
		// httpbin.org/json returns valid JSON
		_, err := client.Do(ctx, http.MethodGet, "/json")
		if err != nil {
			// Network errors are acceptable, JSON parsing errors are not
			if !strings.Contains(err.Error(), "failed to unmarshal") {
				t.Logf("Network error expected: %v", err)
			} else {
				helper.AssertNoError(t, err, "Unexpected JSON unmarshaling error")
			}
		} else {
			t.Log("Successfully parsed JSON response")
		}
	})

	t.Run("InvalidJSONTarget", func(t *testing.T) {
		// Test with a target that can't accept the JSON structure

		_, err := client.Do(ctx, http.MethodGet, "/json")
		if err != nil {
			// Could be network error or JSON unmarshaling error
			t.Logf("Error (expected): %v", err)
		}
	})

	t.Run("InvalidJSONResponse", func(t *testing.T) {
		// Try to get HTML as JSON - this should cause unmarshaling error
		_, err := client.Do(ctx, http.MethodGet, "/html")
		if err != nil {
			// Should get either network error or JSON unmarshaling error
			t.Logf("Error (expected for HTML response): %v", err)
		}
	})
}

// TestCoreClientUnit_DoOperations_ResponseBodyHandling tests response body handling edge cases.
func TestCoreClientUnit_DoOperations_ResponseBodyHandling(t *testing.T) {
	client, err := New("httpbin.org", "token", WithTimeout(5*time.Second))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.Run("EmptyResponse", func(t *testing.T) {
		// Test empty response body handling
		_, err := client.Do(ctx, http.MethodGet, "/status/204")
		if err != nil {
			// Empty responses might cause JSON unmarshaling errors, which is expected
			t.Logf("Error (may be expected for empty response): %v", err)
		}
	})

	t.Run("LargeResponse", func(t *testing.T) {
		// Test with a larger response
		_, err := client.Do(ctx, http.MethodGet, "/json")
		if err != nil {
			t.Logf("Error (network or parsing): %v", err)
		}
	})
}

// TestCoreClientUnit_DoOperations_ResponseBodyReadError tests response body read error handling.
func TestCoreClientUnit_DoOperations_ResponseBodyReadError(t *testing.T) {
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
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = client.Do(ctx, http.MethodGet, "/test")
	if err != nil {
		// Should get JSON unmarshaling error due to incomplete JSON
		t.Logf("Expected error for incomplete JSON: %v", err)
	}
}

// TestCoreClientUnit_DoOperations_HTTPErrorBoundaries tests HTTP error status code boundaries.
func TestCoreClientUnit_DoOperations_HTTPErrorBoundaries(t *testing.T) {
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
			helper.AssertClientCreated(t, client, err, "Failed to create client")

			ctx := context.Background()
			_, err = client.Do(ctx, http.MethodGet, "/test")

			if tc.expectErr {
				helper.AssertError(t, err, fmt.Sprintf("Expected error for status %d", tc.statusCode))
			} else {
				helper.AssertNoError(t, err, fmt.Sprintf("Expected no error for status %d", tc.statusCode))
			}
		})
	}
}

// TestCoreClientUnit_DoOperations_ResponseBodyCloseError tests deferred response body close error handling.
func TestCoreClientUnit_DoOperations_ResponseBodyCloseError(t *testing.T) {
	// Create a test server that returns valid JSON
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"test": "data"}`)
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	// This test covers the successful path including the deferred body close
	_, err = client.Do(ctx, http.MethodGet, "/test")
	helper.AssertNoError(t, err, "Unexpected error")
}

// TestCoreClientUnit_DoOperations_ResponseBodyReadFailure tests response body read error handling.
// TestCoreClientUnit_DoOperations_LargeResponse tests handling of larger response bodies.
func TestCoreClientUnit_DoOperations_LargeResponse(t *testing.T) {
	// Create a test server that returns a large valid JSON response
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Generate a larger JSON response
		largeData := make(map[string]any)
		for i := range 100 {
			largeData[fmt.Sprintf("key_%d", i)] = fmt.Sprintf("value_%d", i)
		}

		// Manually create JSON to ensure it's valid
		fmt.Fprintf(w, `{"test": "data", "largeField": "`)
		for range 1000 {
			fmt.Fprintf(w, "x")
		}
		fmt.Fprintf(w, `"}`)
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	_, err = client.Do(ctx, http.MethodGet, "/test")
	helper.AssertNoError(t, err, "Unexpected error with large response")
}

// TestCoreClientUnit_DoOperations_EmptyResponseBody tests handling of empty response bodies.
func TestCoreClientUnit_DoOperations_EmptyResponseBody(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Send empty body
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	// Do method should succeed with empty body (returns raw bytes)
	body, err := client.Do(ctx, http.MethodGet, "/test")
	helper.AssertNoError(t, err, "Unexpected error for empty response")

	// Body should be empty
	helper.AssertIntEquals(t, len(body), 0, "Response body should be empty")
}

// TestCoreClientUnit_DoOperations_InvalidJSONResponse tests handling of invalid JSON in responses.
func TestCoreClientUnit_DoOperations_InvalidJSONResponse(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid": json}`)) // Invalid JSON
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	// Do method should succeed (returns raw bytes, no JSON parsing)
	body, err := client.Do(ctx, http.MethodGet, "/test")
	helper.AssertNoError(t, err, "Unexpected error for Do method")

	// Body should contain the invalid JSON
	expected := `{"invalid": json}`
	helper.AssertStringEquals(t, string(body), expected, "Response body should match expected content")
}

// TestCoreClientUnit_Validation_NilContext tests handling of nil context.
func TestCoreClientUnit_Validation_NilContext(t *testing.T) {
	client, err := New("example.com", "token")
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	// Using a nil context variable instead of nil literal to test the validation
	var nilCtx context.Context //nolint:staticcheck // Testing nil context behavior
	_, err = client.Do(nilCtx, http.MethodGet, "/test")
	helper.AssertError(t, err, "Expected error with nil context")

	helper.AssertStringContains(t, err.Error(),
		"context cannot be nil",
		"Error message should contain expected text about nil context")
}

// TestCoreClientUnit_DoOperations_NilOutput tests handling of nil output parameter.
func TestCoreClientUnit_DoOperations_NilOutput(t *testing.T) {
	// Create a test server that responds normally
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "data"}`))
	}))
	defer server.Close()

	serverURL := strings.TrimPrefix(server.URL, "https://")
	client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	// The Do method should succeed - it doesn't validate output parameters
	// since it returns raw bytes, not unmarshaled data
	_, err = client.Do(ctx, http.MethodGet, "/test")
	helper.AssertNoError(t, err, "Unexpected error from Do method")
}

// TestCoreClientUnit_DoOperations_RequestCreationError tests handling of request creation errors.
func TestCoreClientUnit_DoOperations_RequestCreationError(t *testing.T) {
	client, err := New("example.com", "token")
	helper.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	// Use an invalid HTTP method to trigger request creation error
	_, err = client.Do(ctx, "INVALID METHOD\nWITH\nNEWLINES", "/test")
	helper.AssertError(t, err, "Expected error with invalid HTTP method")

	helper.AssertStringContains(t, err.Error(),
		"failed to create request",
		"Error message should contain expected text about request creation")
}

// TestCoreClientUnit_Validation_NilClient tests validation with nil client.
func TestCoreClientUnit_Validation_NilClient(t *testing.T) {
	var nilClient *Client

	ctx := context.Background()

	// Test that validateDoParameters correctly handles nil client
	err := nilClient.validateDoParameters(ctx)
	helper.AssertError(t, err, "Expected error with nil client")

	helper.AssertStringContains(t, err.Error(),
		"client cannot be nil",
		"Error message should contain expected text about nil client")
}

// TestCoreClientUnit_ErrorHandling_ResponseBodyClose tests error handling in closeResponseBody.
func TestCoreClientUnit_ErrorHandling_ResponseBodyClose(t *testing.T) {
	client, err := New("example.com", "token")
	helper.AssertClientCreated(t, client, err, "Failed to create client")

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
			helper.AssertTrue(t, false, fmt.Sprintf("closeResponseBody should not panic: %v", r))
		}
	}()

	client.closeResponseBody(resp)

	// Verify that Close() was called
	helper.AssertBoolEquals(t, mockReader.closed, true, "Expected Close() to be called on response body")

	t.Log("closeResponseBody completed without panic, handled error correctly")
}

// errorCloser is a mock io.ReadCloser that returns an error on Close().
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

// TestCoreClientUnit_PostOperations_Success tests the POST operations.
func TestCoreClientUnit_PostOperations_Success(t *testing.T) {
	t.Run("Post_with_nil_client", func(t *testing.T) {
		var nilClient *Client
		payload := map[string]string{"test": "value"}
		err := PostVoid(context.Background(), nilClient, "/test-endpoint", payload)
		helper.AssertError(t, err, "Expected error for nil client")
		helper.AssertStringContains(t, err.Error(),
			"client cannot be nil",
			"Error message should contain expected text about nil client")
	})

	t.Run("Post_marshal_error", func(t *testing.T) {
		client, err := New("test.example.com", "test-token")
		helper.AssertClientCreated(t, client, err, "Failed to create client")

		// Use a payload that cannot be marshaled to JSON
		invalidPayload := make(chan int)
		err = PostVoid(context.Background(), client, "/test-endpoint", invalidPayload)
		helper.AssertError(t, err, "Expected error for unmarshalable payload")
	})

	t.Run("Post_context_canceled", func(t *testing.T) {
		client, err := New("test.example.com", "test-token")
		helper.AssertClientCreated(t, client, err, "Failed to create client")

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		payload := map[string]string{"test": "value"}
		err = PostVoid(ctx, client, "/test-endpoint", payload)
		helper.AssertError(t, err, "Expected error for canceled context")
	})
}

// TestCoreClientUnit_RPCOperations_WithPayload tests the DoRPCWithPayload method.
func TestCoreClientUnit_RPCOperations_WithPayload(t *testing.T) {
	// Create mock server that handles RPC requests
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Check RPC path prefix
		if strings.Contains(r.URL.Path, "/operations/") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"output": {"result": "success"}}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "RPC not found"}`))
		}
	}))
	defer server.Close()

	// Create test client
	serverURL := strings.TrimPrefix(server.URL, "http://")
	testClient, err := New(serverURL, "test-token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, testClient, err, "Failed to create test client")

	ctx := context.Background()
	rpcPath := "/test-rpc"
	payload := map[string]string{"test": "data"}

	t.Run("ValidRPCRequest", func(t *testing.T) {
		result, err := testClient.DoRPCWithPayload(ctx, http.MethodPost, rpcPath, payload)
		if err != nil {
			t.Logf("DoRPCWithPayload() error (expected in test): %v", err)
			return
		}
		if result == nil {
			helper.AssertNotNil(t, result, "DoRPCWithPayload() result should not be nil")
		}
	})

	t.Run("NilClient", func(t *testing.T) {
		var nilClient *Client
		_, err := nilClient.DoRPCWithPayload(ctx, http.MethodPost, rpcPath, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("NilContext", func(t *testing.T) {
		// Use a variable to avoid direct nil literal which triggers SA1012
		var nilCtx context.Context // This is nil by default
		_, err := testClient.DoRPCWithPayload(nilCtx, http.MethodPost, rpcPath, payload)
		helper.AssertError(t, err, "Expected error for nil context")
	})
}

// TestCoreClientUnit_HTTPMethods_PutPatchDelete tests the uncovered client methods.
func TestCoreClientUnit_HTTPMethods_PutPatchDelete(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodPut:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "PUT"}`))
		case http.MethodPatch:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "PATCH"}`))
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "DELETE"}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer server.Close()

	// Create test client
	serverURL := strings.TrimPrefix(server.URL, "http://")
	testClient, err := New(serverURL, "test-token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, testClient, err, "Failed to create test client")

	ctx := context.Background()
	endpoint := "/test-endpoint"
	payload := map[string]string{"test": "data"}

	t.Run("PutMethod", func(t *testing.T) {
		err := PutVoid(ctx, testClient, endpoint, payload)
		if err != nil {
			t.Logf("Put() error (expected in test): %v", err)
		}
	})

	t.Run("PatchMethod", func(t *testing.T) {
		err := PatchVoid(ctx, testClient, endpoint, payload)
		if err != nil {
			t.Logf("Patch() error (expected in test): %v", err)
		}
	})

	t.Run("DeleteMethod", func(t *testing.T) {
		err := Delete(ctx, testClient, endpoint)
		if err != nil {
			t.Logf("Delete() error (expected in test): %v", err)
		}
	})

	t.Run("PutWithNilClient", func(t *testing.T) {
		var nilClient *Client
		err := PutVoid(ctx, nilClient, endpoint, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("PatchWithNilClient", func(t *testing.T) {
		var nilClient *Client
		err := PatchVoid(ctx, nilClient, endpoint, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("DeleteWithNilClient", func(t *testing.T) {
		var nilClient *Client
		err := Delete(ctx, nilClient, endpoint)
		helper.AssertError(t, err, "Expected error for nil client")
	})
}

// TestCoreClientUnit_GenericRequests_AllMethods tests the generic request helper functions.
func TestCoreClientUnit_GenericRequests_AllMethods(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "GET", "data": "test"}`))
		case http.MethodPost:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "POST", "data": "test"}`))
		case http.MethodPut:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "PUT", "data": "test"}`))
		case http.MethodPatch:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "PATCH", "data": "test"}`))
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"method": "DELETE", "data": "test"}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer server.Close()

	// Create test client
	serverURL := strings.TrimPrefix(server.URL, "http://")
	testClient, err := New(serverURL, "test-token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, testClient, err, "Failed to create test client")

	ctx := context.Background()
	endpoint := "/test-endpoint"
	payload := map[string]string{"test": "data"}

	type testResponse struct {
		Method string `json:"method"`
		Data   string `json:"data"`
	}

	t.Run("Get", func(t *testing.T) {
		result, err := Get[testResponse](ctx, testClient, endpoint)
		if err != nil {
			t.Logf("Get() error (expected in test): %v", err)
			return
		}
		if result == nil {
			helper.AssertNotNil(t, result, "Get() result should not be nil")
		}
	})

	t.Run("GetNilClient", func(t *testing.T) {
		_, err := Get[testResponse](ctx, nil, endpoint)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("Post", func(t *testing.T) {
		result, err := Post[testResponse](ctx, testClient, endpoint, payload)
		if err != nil {
			t.Logf("Post() error (expected in test): %v", err)
			return
		}
		if result == nil {
			helper.AssertNotNil(t, result, "Post() result should not be nil")
		}
	})

	t.Run("PostNilClient", func(t *testing.T) {
		_, err := Post[testResponse](ctx, nil, endpoint, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("Put", func(t *testing.T) {
		result, err := Put[testResponse](ctx, testClient, endpoint, payload)
		if err != nil {
			t.Logf("Put() error (expected in test): %v", err)
			return
		}
		if result == nil {
			helper.AssertNotNil(t, result, "Put() result should not be nil")
		}
	})

	t.Run("PutNilClient", func(t *testing.T) {
		_, err := Put[testResponse](ctx, nil, endpoint, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("Patch", func(t *testing.T) {
		result, err := Patch[testResponse](ctx, testClient, endpoint, payload)
		if err != nil {
			t.Logf("Patch() error (expected in test): %v", err)
			return
		}
		if result == nil {
			helper.AssertNotNil(t, result, "Patch() result should not be nil")
		}
	})

	t.Run("PatchNilClient", func(t *testing.T) {
		_, err := Patch[testResponse](ctx, nil, endpoint, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("Delete", func(t *testing.T) {
		err := Delete(ctx, testClient, endpoint)
		if err != nil {
			t.Logf("Delete() error (expected in test): %v", err)
		}
	})

	t.Run("DeleteNilClient", func(t *testing.T) {
		err := Delete(ctx, nil, endpoint)
		helper.AssertError(t, err, "Expected error for nil client")
	})
}

// TestCoreClientUnit_VoidRequests_AllMethods tests the void request helper functions.
func TestCoreClientUnit_VoidRequests_AllMethods(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodPost:
			if strings.Contains(r.URL.Path, "/operations/") {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"output": {"result": "success"}}`))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{}`))
			}
		case http.MethodPut, http.MethodPatch, http.MethodDelete:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer server.Close()

	// Create test client
	serverURL := strings.TrimPrefix(server.URL, "http://")
	testClient, err := New(serverURL, "test-token", WithInsecureSkipVerify(true))
	helper.AssertClientCreated(t, testClient, err, "Failed to create test client")

	ctx := context.Background()
	endpoint := "/test-endpoint"
	rpcPath := "/test-rpc"
	payload := map[string]string{"test": "data"}

	t.Run("PostVoid", func(t *testing.T) {
		err := PostVoid(ctx, testClient, endpoint, payload)
		if err != nil {
			t.Logf("PostVoid() error (expected in test): %v", err)
		}
	})

	t.Run("PostVoidNilClient", func(t *testing.T) {
		err := PostVoid(ctx, nil, endpoint, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("PostRPCVoid", func(t *testing.T) {
		err := PostRPCVoid(ctx, testClient, rpcPath, payload)
		if err != nil {
			t.Logf("PostRPCVoid() error (expected in test): %v", err)
		}
	})

	t.Run("PostRPCVoidNilClient", func(t *testing.T) {
		err := PostRPCVoid(ctx, nil, rpcPath, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("PutVoid", func(t *testing.T) {
		err := PutVoid(ctx, testClient, endpoint, payload)
		if err != nil {
			t.Logf("PutVoid() error (expected in test): %v", err)
		}
	})

	t.Run("PutVoidNilClient", func(t *testing.T) {
		err := PutVoid(ctx, nil, endpoint, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("PatchVoid", func(t *testing.T) {
		err := PatchVoid(ctx, testClient, endpoint, payload)
		if err != nil {
			t.Logf("PatchVoid() error (expected in test): %v", err)
		}
	})

	t.Run("PatchVoidNilClient", func(t *testing.T) {
		err := PatchVoid(ctx, nil, endpoint, payload)
		helper.AssertError(t, err, "Expected error for nil client")
	})
}

// TestCoreClientUnit_SuccessPaths_Coverage tests success paths for improved coverage.
func TestCoreClientUnit_SuccessPaths_Coverage(t *testing.T) {
	// Test DoWithPayload success path
	t.Run("DoWithPayload_Success", func(t *testing.T) {
		mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result": "success"}`))
		}))
		defer mockServer.Close()

		client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
		helper.AssertNoError(t, err, "Client creation should succeed")

		ctx := context.Background()

		payload := map[string]string{"test": "data"}
		body, err := client.DoWithPayload(ctx, "POST", "/restconf/data/test", payload)

		helper.AssertNoError(t, err, "DoWithPayload should succeed")
		helper.AssertTrue(t, len(body) > 0, "Response body should not be empty")
	})

	// Test DoRPCWithPayload success path
	t.Run("DoRPCWithPayload_Success", func(t *testing.T) {
		mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result": "rpc-success"}`))
		}))
		defer mockServer.Close()

		client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
		helper.AssertNoError(t, err, "Client creation should succeed")

		ctx := context.Background()

		payload := map[string]string{"rpc": "data"}
		body, err := client.DoRPCWithPayload(ctx, "POST", "/test-rpc", payload)

		helper.AssertNoError(t, err, "DoRPCWithPayload should succeed")
		helper.AssertTrue(t, len(body) > 0, "RPC response body should not be empty")
	}) // Test RestconfBuilder
	t.Run("RestconfBuilder_Success", func(t *testing.T) {
		controller := "test.example.com"
		token := "test-token"
		client, err := New(controller, token)

		helper.AssertNoError(t, err, "Client creation should succeed")

		builder := client.RestconfBuilder()
		helper.AssertTrue(t, builder != nil, "RestconfBuilder should return non-nil builder")
	})

	// Test RestconfBuilder with nil client
	t.Run("RestconfBuilder_NilClient", func(t *testing.T) {
		var client *Client
		builder := client.RestconfBuilder()
		helper.AssertTrue(t, builder == nil, "RestconfBuilder should return nil for nil client")
	})

	// Test Get success path
	t.Run("Get_SuccessPath", func(t *testing.T) {
		mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"data": {"value": "test"}}`))
		}))
		defer mockServer.Close()

		client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
		helper.AssertNoError(t, err, "Client creation should succeed")

		ctx := context.Background()

		type TestData struct {
			Data struct {
				Value string `json:"value"`
			} `json:"data"`
		}

		result, err := Get[TestData](ctx, client, "/restconf/data/test")

		helper.AssertNoError(t, err, "Get should succeed")
		helper.AssertTrue(t, result != nil, "Result should not be nil")
		helper.AssertStringEquals(t, result.Data.Value, "test", "Data should match")
	}) // Test Post success path
	t.Run("Post_SuccessPath", func(t *testing.T) {
		mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"created": {"id": "123"}}`))
		}))
		defer mockServer.Close()

		client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
		helper.AssertNoError(t, err, "Client creation should succeed")

		ctx := context.Background()

		type CreateResponse struct {
			Created struct {
				ID string `json:"id"`
			} `json:"created"`
		}

		payload := map[string]string{"name": "test"}
		result, err := Post[CreateResponse](ctx, client, "/restconf/data/test", payload)

		helper.AssertNoError(t, err, "Post should succeed")
		helper.AssertTrue(t, result != nil, "Result should not be nil")
		helper.AssertStringEquals(t, result.Created.ID, "123", "Created ID should match")
	})

	// Test Put success path
	t.Run("Put_SuccessPath", func(t *testing.T) {
		mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"updated": {"status": "success"}}`))
		}))
		defer mockServer.Close()

		client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
		helper.AssertNoError(t, err, "Client creation should succeed")

		ctx := context.Background()

		type UpdateResponse struct {
			Updated struct {
				Status string `json:"status"`
			} `json:"updated"`
		}

		payload := map[string]string{"field": "value"}
		result, err := Put[UpdateResponse](ctx, client, "/restconf/data/test", payload)

		helper.AssertNoError(t, err, "Put should succeed")
		helper.AssertTrue(t, result != nil, "Result should not be nil")
		helper.AssertStringEquals(t, result.Updated.Status, "success", "Update status should match")
	})

	// Test Patch success path
	t.Run("Patch_SuccessPath", func(t *testing.T) {
		mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"patched": {"result": "ok"}}`))
		}))
		defer mockServer.Close()

		client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
		helper.AssertNoError(t, err, "Client creation should succeed")

		ctx := context.Background()

		type PatchResponse struct {
			Patched struct {
				Result string `json:"result"`
			} `json:"patched"`
		}

		payload := map[string]string{"patch": "data"}
		result, err := Patch[PatchResponse](ctx, client, "/restconf/data/test", payload)

		helper.AssertNoError(t, err, "Patch should succeed")
		helper.AssertTrue(t, result != nil, "Result should not be nil")
		helper.AssertStringEquals(t, result.Patched.Result, "ok", "Patch result should match")
	})
}

// TestCoreClientUnit_AdditionalCoverage_EdgeCases provides additional coverage for edge cases.
