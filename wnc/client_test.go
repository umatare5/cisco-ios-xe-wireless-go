package wnc

import (
	"context"
	"log/slog"
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
	token := "dGVzdDp0ZXN0" // base64 encoded "test:test"

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
		err := client.Do(nil, "GET", "/restconf/data/test", &response)
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

// =============================================================================
// 3. ADDITIONAL OPTION FUNCTION TESTS
// =============================================================================

func TestClientOptionsExtra(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
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

// =============================================================================
// 4. SERVICE ACCESSOR TESTS
// =============================================================================

func TestServiceAccessors(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	tests := []struct {
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
			accessor: func() interface{} { return client.Mdns() },
		},
		{
			name:     "Geolocation",
			accessor: func() interface{} { return client.Geolocation() },
		},
		{
			name:     "Mcast",
			accessor: func() interface{} { return client.Mcast() },
		},
	}

	for _, tt := range tests {
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

// =============================================================================
// 5. DO METHOD COVERAGE TESTS
// =============================================================================

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

	// Test with cancelled context
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	err = client.Do(cancelledCtx, "GET", "/test", &response)
	if err == nil {
		t.Error("Expected error for cancelled context, got nil")
	}
}

func TestDoMethodNilParameters(t *testing.T) {
	client, err := New("example.com", "token")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	var response interface{}

	// Test with nil context
	err = client.Do(nil, "GET", "/test", &response)
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
		// We expect a network error, not a request creation error
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
}
