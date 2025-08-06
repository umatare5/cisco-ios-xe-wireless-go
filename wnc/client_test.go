package wnc

import (
	"context"
	"os"
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
