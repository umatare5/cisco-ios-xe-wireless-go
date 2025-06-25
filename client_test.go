// Package wnc provides test functionality for the Cisco Wireless Network Controller API client.
package wnc

import (
	"context"
	"os"
	"testing"
	"time"
)

// Test constants
const (
	defaultTestTimeout  = 20 * time.Second
	quickTestTimeout    = 5 * time.Second
	extendedTestTimeout = 30 * time.Second
	microTestTimeout    = 2 * time.Microsecond
)

// createTestConfig creates a test configuration from environment variables
func createTestConfig(t *testing.T) *testConfig {
	t.Helper()

	controller := os.Getenv(EnvVarController)
	accessToken := os.Getenv(EnvVarAccessToken)

	if controller == "" || accessToken == "" {
		t.Skip("WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set for integration tests")
	}

	return &testConfig{
		Controller:  controller,
		AccessToken: accessToken,
		Timeout:     defaultTestTimeout,
	}
}

// testConfig represents configuration for test operations
type testConfig struct {
	Controller  string
	AccessToken string
	Timeout     time.Duration
}

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestClientStructure tests the basic structure of the client
func TestClientStructure(t *testing.T) {
	controller := ExampleTestHostname
	token := TestAccessTokenValue

	config := Config{
		Controller:  controller,
		AccessToken: token,
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create basic client: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil client")
	}

	// Test client fields are properly set
	if client.controller != controller {
		t.Errorf("Expected controller %s, got %s", controller, client.controller)
	}
}

// TestClientValidation tests client structure validation
func TestClientValidation(t *testing.T) {
	tests := []struct {
		name       string
		controller string
		token      string
		shouldFail bool
	}{
		{"ValidConfig", ExampleTestHostname, TestAccessTokenValue, false},
		{"EmptyController", "", TestAccessTokenValue, true},
		{"EmptyToken", ExampleTestHostname, "", true},
		{"EmptyBoth", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := Config{
				Controller:  tt.controller,
				AccessToken: tt.token,
			}
			_, err := NewClient(config)
			if tt.shouldFail && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.shouldFail && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

// =============================================================================
// 2. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// TestClientConfig tests client configuration options
func TestClientConfig(t *testing.T) {
	tests := []struct {
		name       string
		controller string
		token      string
		options    []ClientOption
		valid      bool
	}{
		{
			name:       "ValidBasicConfig",
			controller: ExampleTestHostname,
			token:      TestAccessTokenValue,
			options:    []ClientOption{},
			valid:      true,
		},
		{
			name:       "EmptyController",
			controller: "",
			token:      TestAccessTokenValue,
			options:    []ClientOption{},
			valid:      false,
		},
		{
			name:       "EmptyToken",
			controller: ExampleTestHostname,
			token:      "",
			options:    []ClientOption{},
			valid:      false,
		},
		{
			name:       "WithTimeout",
			controller: ExampleTestHostname,
			token:      TestAccessTokenValue,
			options:    []ClientOption{WithTimeout(30 * time.Second)},
			valid:      true,
		},
		{
			name:       "WithInsecureSkipVerify",
			controller: ExampleTestHostname,
			token:      TestAccessTokenValue,
			options:    []ClientOption{WithInsecureSkipVerify(true)},
			valid:      true,
		},
		{
			name:       "InvalidTimeout",
			controller: ExampleTestHostname,
			token:      TestAccessTokenValue,
			options:    []ClientOption{WithTimeout(0)},
			valid:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := Config{
				Controller:  tt.controller,
				AccessToken: tt.token,
			}

			client, err := NewClientWithConfig(config, tt.options...)
			t.Logf("Test %s: client=%v, err=%v", tt.name, client != nil, err)
			if tt.valid {
				if err != nil {
					t.Errorf("Expected valid config to succeed, got error: %v", err)
				}
				if client == nil {
					t.Error("Expected client to be created for valid config")
				}
			} else {
				if err == nil {
					t.Error("Expected invalid config to fail")
				}
				if client != nil {
					t.Error("Expected client to be nil for invalid config")
				}
			}
		})
	}
}

// TestClientOptsTable tests various client options
func TestClientOptsTable(t *testing.T) {
	tests := []struct {
		name   string
		option ClientOption
		valid  bool
	}{
		{
			name:   "ValidTimeout",
			option: WithTimeout(30 * time.Second),
			valid:  true,
		},
		{
			name:   "ZeroTimeout",
			option: WithTimeout(0),
			valid:  false,
		},
		{
			name:   "NegativeTimeout",
			option: WithTimeout(-1 * time.Second),
			valid:  false,
		},
		{
			name:   "InsecureSkipVerifyTrue",
			option: WithInsecureSkipVerify(true),
			valid:  true,
		},
		{
			name:   "InsecureSkipVerifyFalse",
			option: WithInsecureSkipVerify(false),
			valid:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := Config{
				Controller:  ExampleTestHostname,
				AccessToken: TestAccessTokenValue,
			}
			client, err := NewClientWithConfig(config, tt.option)
			if tt.valid {
				if err != nil {
					t.Errorf("Expected valid option to succeed, got error: %v", err)
				}
				if client == nil {
					t.Error("Expected client to be created for valid option")
				}
			} else {
				if err == nil {
					t.Error("Expected invalid option to fail")
				}
				if client != nil {
					t.Error("Expected client to be nil for invalid option")
				}
			}
		})
	}
}

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// TestClientFailures tests client initialization failure scenarios
func TestClientFailures(t *testing.T) {
	// Test with empty controller
	config := Config{
		Controller:  "",
		AccessToken: TestAccessTokenValue,
	}
	client, err := NewClient(config)
	if err == nil {
		t.Fatal("Expected error for empty controller, got nil")
	}
	if client != nil {
		t.Fatal("Expected nil client for empty controller")
	}

	// Test with empty token
	config = Config{
		Controller:  ExampleTestHostname,
		AccessToken: "",
	}
	client, err = NewClient(config)
	if err == nil {
		t.Fatal("Expected error for empty token, got nil")
	}
	if client != nil {
		t.Fatal("Expected nil client for empty token")
	}

	// Test with zero timeout option (should get default)
	config = Config{
		Controller:  ExampleTestHostname,
		AccessToken: TestAccessTokenValue,
		Timeout:     0,
	}
	client, err = NewClient(config)
	if err != nil {
		t.Fatalf("Expected zero timeout to get default, but got error: %v", err)
	}
	if client == nil {
		t.Fatal("Expected client to be created with default timeout")
	}
	if client.timeout != DefaultTimeout {
		t.Errorf("Expected default timeout %v, got %v", DefaultTimeout, client.timeout)
	}
}

// TestContextCancellation tests context cancellation for client operations
func TestContextCancellation(t *testing.T) {
	testConfig := createTestConfig(t)
	config := Config{
		Controller:  testConfig.Controller,
		AccessToken: testConfig.AccessToken,
		Timeout:     extendedTestTimeout,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test immediate context cancellation
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Create a simple struct to receive the response
	var response interface{}

	// This should fail due to cancelled context
	err = client.SendAPIRequest(ctx, "/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
	if err == nil {
		t.Fatal("Expected error for cancelled context, got nil")
	}

	// Test context timeout
	ctx, cancel = context.WithTimeout(context.Background(), microTestTimeout)
	defer cancel()

	time.Sleep(2 * microTestTimeout) // Ensure timeout

	err = client.SendAPIRequest(ctx, "/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}
}

// TestInvalidEndpoint tests requests to an invalid controller
func TestInvalidEndpoint(t *testing.T) {
	config := Config{
		Controller:  "invalid.controller.local",
		AccessToken: TestAccessTokenValue,
		Timeout:     quickTestTimeout,
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), quickTestTimeout)
	defer cancel()

	var response interface{}
	err = client.SendAPIRequest(ctx, "/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
	if err == nil {
		t.Fatal("Expected error for invalid controller, got nil")
	}
}

// =============================================================================
// 4. INTEGRATION TESTS (API Communication & Full Workflow Tests)
// =============================================================================

// TestClientFunctions tests the core client functions
func TestClientFunctions(t *testing.T) {
	testConfig := createTestConfig(t)
	config := Config{
		Controller:         testConfig.Controller,
		AccessToken:        testConfig.AccessToken,
		Timeout:            defaultTestTimeout,
		InsecureSkipVerify: true,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()

	t.Run("GET_GeneralOper", func(t *testing.T) {
		var response interface{}
		err := client.SendAPIRequest(ctx, "/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
		if err != nil {
			t.Logf("GET request failed (may be expected): %v", err)
		} else {
			t.Logf("GET request successful")
			if response == nil {
				t.Error("Expected non-nil response")
			}
		}
	})

	t.Run("GET_APOper", func(t *testing.T) {
		var response interface{}
		err := client.SendAPIRequest(ctx, "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data", &response)
		if err != nil {
			t.Logf("GET AP oper request failed (may be expected): %v", err)
		} else {
			t.Logf("GET AP oper request successful")
			if response == nil {
				t.Error("Expected non-nil response")
			}
		}
	})
}

// TestRealController tests client with real controller
func TestRealController(t *testing.T) {
	testConfig := createTestConfig(t)
	config := Config{
		Controller:         testConfig.Controller,
		AccessToken:        testConfig.AccessToken,
		Timeout:            defaultTestTimeout,
		InsecureSkipVerify: true,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()

	// Test basic connectivity
	var response interface{}
	err = client.SendAPIRequest(ctx, "/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
	if err != nil {
		t.Logf("Controller connection test failed: %v", err)
	} else {
		t.Logf("Controller connection successful")
	}
}

// =============================================================================
// 5. OTHER TESTS
// =============================================================================

// TestClientDefaults tests client default values
func TestClientDefaults(t *testing.T) {
	config := Config{
		Controller:  ExampleTestHostname,
		AccessToken: TestAccessTokenValue,
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test default timeout is set
	if client.timeout == 0 {
		t.Error("Expected default timeout to be set")
	}
}
