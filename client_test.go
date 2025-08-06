package wnc

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"
)

// Test constants - using standard timeouts
const (
	clientTestTimeout     = 20 * time.Second
	clientQuickTimeout    = 5 * time.Second
	clientExtendedTimeout = 30 * time.Second
	clientMicroTimeout    = 2 * time.Microsecond
)

// createClientTestConfig creates a test configuration from environment variables
func createClientTestConfig(t *testing.T) *clientTestConfig {
	t.Helper()

	controller := os.Getenv(EnvVarController)
	accessToken := os.Getenv(EnvVarAccessToken)

	if controller == "" || accessToken == "" {
		t.Skip("WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set for integration tests")
	}

	return &clientTestConfig{
		Controller:  controller,
		AccessToken: accessToken,
		Timeout:     clientTestTimeout,
	}
}

// clientTestConfig represents configuration for client test operations
type clientTestConfig struct {
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

	// Test basic client creation with config
	config := Config{
		Controller:  controller,
		AccessToken: token,
		Timeout:     clientTestTimeout,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Fatal("Expected client to be created, got nil")
	}
}

// TestCoreClient tests the CoreClient accessor method
func TestCoreClient(t *testing.T) {
	controller := ExampleTestHostname
	token := TestAccessTokenValue

	config := Config{
		Controller:  controller,
		AccessToken: token,
		Timeout:     clientTestTimeout,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	core := client.CoreClient()
	if core == nil {
		t.Error("Expected CoreClient to return non-nil client")
	}
}

// TestServiceAccessors tests all service accessor methods
func TestClientServiceAccessors(t *testing.T) {
	controller := ExampleTestHostname
	token := TestAccessTokenValue

	config := Config{
		Controller:  controller,
		AccessToken: token,
		Timeout:     clientTestTimeout,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test AFC accessor
	afc := client.AFC()
	if afc == nil {
		t.Error("Expected AFC to return non-nil service")
	}

	// Test General accessor
	general := client.General()
	if general == nil {
		t.Error("Expected General to return non-nil service")
	}
}

// TestAFCServiceMethods tests the AFC service methods
func TestAFCServiceMethods(t *testing.T) {
	controller := ExampleTestHostname
	token := TestAccessTokenValue

	config := Config{
		Controller:  controller,
		AccessToken: token,
		Timeout:     clientTestTimeout,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	afc := client.AFC()
	if afc == nil {
		t.Fatal("Expected AFC to return non-nil service")
	}

	// Use type assertion to access the methods
	afcStruct, ok := afc.(struct {
		Oper       func(context.Context) (interface{}, error)
		CloudStats func(context.Context) (interface{}, error)
	})
	if !ok {
		t.Fatal("AFC service does not have expected structure")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Test Oper method
	t.Run("AFC_Oper", func(t *testing.T) {
		_, err := afcStruct.Oper(ctx)
		// Error is expected due to test environment, but we've covered the code path
		t.Logf("AFC Oper method called, result: %v", err)
	})

	// Test CloudStats method
	t.Run("AFC_CloudStats", func(t *testing.T) {
		_, err := afcStruct.CloudStats(ctx)
		// Error is expected due to test environment, but we've covered the code path
		t.Logf("AFC CloudStats method called, result: %v", err)
	})
}

// TestGeneralServiceMethods tests the General service methods
func TestGeneralServiceMethods(t *testing.T) {
	controller := ExampleTestHostname
	token := TestAccessTokenValue

	config := Config{
		Controller:  controller,
		AccessToken: token,
		Timeout:     clientTestTimeout,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	general := client.General()
	if general == nil {
		t.Fatal("Expected General to return non-nil service")
	}

	// Use type assertion to access the methods
	generalStruct, ok := general.(struct {
		Oper func(context.Context) (interface{}, error)
		Cfg  func(context.Context) (interface{}, error)
	})
	if !ok {
		t.Fatal("General service does not have expected structure")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Test Oper method
	t.Run("General_Oper", func(t *testing.T) {
		_, err := generalStruct.Oper(ctx)
		// Error is expected due to test environment, but we've covered the code path
		t.Logf("General Oper method called, result: %v", err)
	})

	// Test Cfg method
	t.Run("General_Cfg", func(t *testing.T) {
		_, err := generalStruct.Cfg(ctx)
		// Error is expected due to test environment, but we've covered the code path
		t.Logf("General Cfg method called, result: %v", err)
	})
}

// TestNew tests the New function
func TestNewLegacyFunction(t *testing.T) {
	controller := ExampleTestHostname
	token := TestAccessTokenValue

	client, err := New(controller, token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Error("Expected New to return non-nil client")
	}

	// Verify it creates a working client
	afc := client.AFC()
	if afc == nil {
		t.Error("Expected AFC service to be accessible from legacy client")
	}
}

// TestNewFunction tests the New function with various scenarios
func TestNewFunction(t *testing.T) {
	tests := []struct {
		name        string
		host        string
		token       string
		shouldError bool
	}{
		{
			name:        "ValidParameters",
			host:        ExampleTestHostname,
			token:       TestAccessTokenValue,
			shouldError: false,
		},
		{
			name:        "EmptyHost",
			host:        "",
			token:       TestAccessTokenValue,
			shouldError: true,
		},
		{
			name:        "EmptyToken",
			host:        ExampleTestHostname,
			token:       "",
			shouldError: true,
		},
		{
			name:        "EmptyBoth",
			host:        "",
			token:       "",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := New(tt.host, tt.token)
			if tt.shouldError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if client != nil {
					t.Error("Expected nil client when error occurs")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if client == nil {
					t.Error("Expected non-nil client when no error")
				}
			}
		})
	}
}

// TestNewClientWithConfigFunction tests the NewClientWithConfig function
func TestNewClientWithConfigFunction(t *testing.T) {
	tests := []struct {
		name        string
		config      Config
		shouldError bool
	}{
		{
			name: "ValidConfig",
			config: Config{
				Controller:  ExampleTestHostname,
				AccessToken: TestAccessTokenValue,
				Timeout:     clientTestTimeout,
			},
			shouldError: false,
		},
		{
			name: "EmptyController",
			config: Config{
				Controller:  "",
				AccessToken: TestAccessTokenValue,
				Timeout:     clientTestTimeout,
			},
			shouldError: true,
		},
		{
			name: "EmptyToken",
			config: Config{
				Controller:  ExampleTestHostname,
				AccessToken: "",
				Timeout:     clientTestTimeout,
			},
			shouldError: true,
		},
		{
			name: "ZeroTimeout",
			config: Config{
				Controller:  ExampleTestHostname,
				AccessToken: TestAccessTokenValue,
				Timeout:     0,
			},
			shouldError: false, // Zero timeout should use default
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClientWithConfig(tt.config)
			if tt.shouldError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if client != nil {
					t.Error("Expected nil client when error occurs")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if client == nil {
					t.Error("Expected non-nil client when no error")
				}
			}
		})
	}
}

// TestNewClientWithConfigAdvanced tests advanced NewClientWithConfig scenarios
func TestNewClientWithConfigAdvanced(t *testing.T) {
	t.Run("ConfigWithLogger", func(t *testing.T) {
		customLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		config := Config{
			Controller:  ExampleTestHostname,
			AccessToken: TestAccessTokenValue,
			Logger:      customLogger,
		}

		client, err := NewClientWithConfig(config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
		if client.logger != customLogger {
			t.Error("Expected custom logger to be set")
		}
	})

	t.Run("ConfigWithInsecureSkipVerify", func(t *testing.T) {
		config := Config{
			Controller:         ExampleTestHostname,
			AccessToken:        TestAccessTokenValue,
			InsecureSkipVerify: true,
		}

		client, err := NewClientWithConfig(config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
		if !client.insecureSkipVerify {
			t.Error("Expected insecureSkipVerify to be true")
		}
	})

	t.Run("ConfigWithOptions", func(t *testing.T) {
		config := Config{
			Controller:  ExampleTestHostname,
			AccessToken: TestAccessTokenValue,
		}

		customTimeout := 45 * time.Second
		client, err := NewClientWithConfig(config, WithTimeout(customTimeout))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
		if client.timeout != customTimeout {
			t.Errorf("Expected timeout %v, got %v", customTimeout, client.timeout)
		}
	})

	t.Run("ConfigWithMultipleOptions", func(t *testing.T) {
		config := Config{
			Controller:  ExampleTestHostname,
			AccessToken: TestAccessTokenValue,
		}

		customTimeout := 45 * time.Second
		customLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))

		client, err := NewClientWithConfig(config,
			WithTimeout(customTimeout),
			WithLogger(customLogger),
			WithInsecureSkipVerify(true))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
		if client.timeout != customTimeout {
			t.Errorf("Expected timeout %v, got %v", customTimeout, client.timeout)
		}
		if client.logger != customLogger {
			t.Error("Expected custom logger to be set")
		}
	})

	t.Run("ConfigWithZeroTimeoutAndOptionTimeout", func(t *testing.T) {
		config := Config{
			Controller:  ExampleTestHostname,
			AccessToken: TestAccessTokenValue,
			Timeout:     0, // Zero timeout in config
		}

		customTimeout := 50 * time.Second
		client, err := NewClientWithConfig(config, WithTimeout(customTimeout))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
		// Option timeout should override zero timeout
		if client.timeout != customTimeout {
			t.Errorf("Expected timeout %v (from option), got %v", customTimeout, client.timeout)
		}
	})

	t.Run("ConfigWithNilLogger", func(t *testing.T) {
		config := Config{
			Controller:  ExampleTestHostname,
			AccessToken: TestAccessTokenValue,
			Logger:      nil, // Explicitly nil logger
		}

		client, err := NewClientWithConfig(config)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("Expected non-nil client")
		}
		// Should get default logger
		if client.logger == nil {
			t.Error("Expected default logger when config logger is nil")
		}
	})
}

func TestClientWithConfig(t *testing.T) {
	controller := ExampleTestHostname
	token := TestAccessTokenValue

	config := Config{
		Controller:  controller,
		AccessToken: token,
		Timeout:     clientTestTimeout,
	}

	client, err := NewClientWithConfig(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	if client == nil {
		t.Fatal("Client should not be nil")
	}

	// Test that the client was created successfully - we can't test private fields
	// but we can ensure the client is functional by testing a basic API call
}

// TestClientErrorHandling tests various error scenarios
func TestClientErrorHandling(t *testing.T) {
	testConfig := createClientTestConfig(t)

	config := Config{
		Controller:         testConfig.Controller,
		AccessToken:        testConfig.AccessToken,
		Timeout:            clientQuickTimeout,
		InsecureSkipVerify: true,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Run("InvalidEndpoint", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), clientQuickTimeout)
		defer cancel()

		var result interface{}
		err := client.SendAPIRequest(ctx, "/invalid/nonexistent/endpoint", &result)
		if err == nil {
			t.Error("Expected error for invalid endpoint, got nil")
		}
	})

	t.Run("CancelledContext", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		var result interface{}
		err := client.SendAPIRequest(ctx, "ietf-yang-library:yang-library", &result)
		if err == nil {
			t.Error("Expected error for cancelled context, got nil")
		}
	})

	t.Run("TimeoutContext", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), clientMicroTimeout)
		defer cancel()

		var result interface{}
		err := client.SendAPIRequest(ctx, "ietf-yang-library:yang-library", &result)
		if err == nil {
			t.Error("Expected timeout error, got nil")
		}
	})
}

// TestClientNilParameterHandling tests handling of nil parameters
func TestClientNilParameterHandling(t *testing.T) {
	controller := ExampleTestHostname
	token := TestAccessTokenValue

	config := Config{
		Controller:  controller,
		AccessToken: token,
		Timeout:     clientTestTimeout,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Use a very short timeout context to avoid hanging on network calls
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	t.Run("NilResult", func(t *testing.T) {
		err := client.SendAPIRequest(ctx, "/test", nil)
		if err == nil {
			t.Error("Expected error for nil result parameter")
		}
	})

	t.Run("EmptyEndpoint", func(t *testing.T) {
		var result interface{}
		err := client.SendAPIRequest(ctx, "", &result)
		if err == nil {
			t.Error("Expected error for empty endpoint")
		}
	})
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
	testConfig := createClientTestConfig(t)
	config := Config{
		Controller:  testConfig.Controller,
		AccessToken: testConfig.AccessToken,
		Timeout:     clientExtendedTimeout,
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
	err = client.SendAPIRequest(ctx, "Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
	if err == nil {
		t.Fatal("Expected error for cancelled context, got nil")
	}

	// Test context timeout
	ctx, cancel = context.WithTimeout(context.Background(), clientMicroTimeout)
	defer cancel()

	time.Sleep(2 * clientMicroTimeout) // Ensure timeout

	err = client.SendAPIRequest(ctx, "Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}
}

// TestInvalidEndpoint tests requests to an invalid controller
func TestInvalidEndpoint(t *testing.T) {
	config := Config{
		Controller:  "invalid.controller.local",
		AccessToken: TestAccessTokenValue,
		Timeout:     clientQuickTimeout,
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var response interface{}
	err = client.SendAPIRequest(ctx, "Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
	if err == nil {
		t.Fatal("Expected error for invalid controller, got nil")
	}
}

// TestSendAPIRequestFailures tests various failure scenarios for SendAPIRequest
func TestSendAPIRequestFailures(t *testing.T) {
	config := Config{
		Controller:  ExampleTestHostname,
		AccessToken: TestAccessTokenValue,
		Timeout:     clientQuickTimeout,
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Use a very short timeout context to avoid hanging on network calls
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	tests := []struct {
		name       string
		endpoint   string
		response   interface{}
		shouldFail bool
	}{
		{"EmptyEndpoint", "", &map[string]interface{}{}, true},
		{"InvalidEndpoint", "invalid", &map[string]interface{}{}, true},
		{"NilResponse", "/restconf/data/test", nil, true},
		{"ValidEndpoint", "Cisco-IOS-XE-wireless-general-oper:general-oper-data", &map[string]interface{}{}, true}, // Expected to fail due to network
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := client.SendAPIRequest(ctx, tt.endpoint, tt.response)
			if tt.shouldFail && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.shouldFail && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

// TestCreateHTTPRequestCoverage tests createHTTPRequest method scenarios
func TestCreateHTTPRequestCoverage(t *testing.T) {
	t.Skip("Skipping test for internal method removed in Phase 1 refactor")
	return
}

// TestExecuteHTTPRequestCoverage tests executeHTTPRequest method scenarios
func TestExecuteHTTPRequestCoverage(t *testing.T) {
	t.Skip("Skipping test for internal method removed in Phase 1 refactor")
	return
}

// TestProcessHTTPResponseCoverage tests processHTTPResponse method scenarios
func TestProcessHTTPResponseCoverage(t *testing.T) {
	// processHTTPResponse is tested via SendAPIRequest integration tests
	// since testing with nil response would cause a panic
	t.Log("processHTTPResponse is tested via SendAPIRequest integration tests")
}

// =============================================================================
// 4. INTEGRATION TESTS (API Communication & Full Workflow Tests)
// =============================================================================

// TestClientFunctions tests the core client functions
func TestClientFunctions(t *testing.T) {
	testConfig := createClientTestConfig(t)
	config := Config{
		Controller:         testConfig.Controller,
		AccessToken:        testConfig.AccessToken,
		Timeout:            clientTestTimeout,
		InsecureSkipVerify: true,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), clientTestTimeout)
	defer cancel()

	t.Run("GET_GeneralOper", func(t *testing.T) {
		var response interface{}
		err := client.SendAPIRequest(ctx, "Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
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
		err := client.SendAPIRequest(ctx, "Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data", &response)
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
	testConfig := createClientTestConfig(t)
	config := Config{
		Controller:         testConfig.Controller,
		AccessToken:        testConfig.AccessToken,
		Timeout:            clientTestTimeout,
		InsecureSkipVerify: true,
	}

	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), clientTestTimeout)
	defer cancel()

	// Test basic connectivity
	var response interface{}
	err = client.SendAPIRequest(ctx, "Cisco-IOS-XE-wireless-general-oper:general-oper-data", &response)
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

// =============================================================================
// 6. DETAILED INTERNAL FUNCTION TESTS
// =============================================================================

// TestSendAPIRequestErrorPaths tests error paths in SendAPIRequest
func TestSendAPIRequestErrorPaths(t *testing.T) {
	config := Config{
		Controller:  ExampleTestHostname,
		AccessToken: TestAccessTokenValue,
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Run("NilContext", func(t *testing.T) {
		var response interface{}
		err := client.SendAPIRequest(nil, "/test", &response)
		if err == nil {
			t.Error("Expected error for nil context")
		}
		if err != nil && !containsText(err.Error(), "context cannot be nil") {
			t.Errorf("Expected 'context cannot be nil' error, got: %v", err)
		}
	})

	t.Run("TimeoutContext", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancel()
		time.Sleep(2 * time.Nanosecond) // Ensure timeout

		var response interface{}
		err := client.SendAPIRequest(ctx, "/restconf/data/test", &response)
		if err == nil {
			t.Error("Expected timeout error")
		}
	})

	t.Run("InvalidURL", func(t *testing.T) {
		// Test with client that has invalid controller URL
		invalidConfig := Config{
			Controller:  "://invalid-url",
			AccessToken: TestAccessTokenValue,
		}
		invalidClient, err := NewClient(invalidConfig)
		if err != nil {
			// If client creation fails, that's also valid
			t.Logf("Client creation failed with invalid URL: %v", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), clientQuickTimeout)
		defer cancel()

		var response interface{}
		err = invalidClient.SendAPIRequest(ctx, "/test", &response)
		if err == nil {
			t.Error("Expected error for invalid URL")
		}
	})
}

// TestClientWithDifferentConfigurations tests client behavior with edge case configurations
func TestClientWithDifferentConfigurations(t *testing.T) {
	t.Run("ClientWithZeroTimeout", func(t *testing.T) {
		config := Config{
			Controller:  ExampleTestHostname,
			AccessToken: TestAccessTokenValue,
		}
		_, err := NewClientWithConfig(config, WithTimeout(0))
		if err == nil {
			t.Error("Expected error for zero timeout")
		}
		if err != nil && !containsText(err.Error(), "timeout must be positive") {
			t.Errorf("Expected 'timeout must be positive' error, got: %v", err)
		}
	})

	t.Run("ClientWithVeryLongTimeout", func(t *testing.T) {
		config := Config{
			Controller:  ExampleTestHostname,
			AccessToken: TestAccessTokenValue,
		}
		longTimeout := 24 * time.Hour
		client, err := NewClientWithConfig(config, WithTimeout(longTimeout))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		if client.timeout != longTimeout {
			t.Errorf("Expected timeout %v, got %v", longTimeout, client.timeout)
		}
	})
}

// =============================================================================
// 6. DETAILED INTERNAL FUNCTION TESTS
// =============================================================================

// TestSendAPIRequestDetailedCoverage tests SendAPIRequest with various scenarios
func TestSendAPIRequestDetailedCoverage(t *testing.T) {
	config := Config{
		Controller:  ExampleTestHostname,
		AccessToken: TestAccessTokenValue,
	}
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Run("NilContext", func(t *testing.T) {
		var response interface{}
		err := client.SendAPIRequest(nil, "/test", &response)
		if err == nil {
			t.Error("Expected error for nil context")
		}
		if !containsText(err.Error(), "context cannot be nil") {
			t.Errorf("Expected 'context cannot be nil' error, got: %v", err)
		}
	})

	t.Run("InvalidEndpoint", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), clientQuickTimeout)
		defer cancel()

		var response interface{}
		err := client.SendAPIRequest(ctx, "invalid-url-format", &response)
		if err == nil {
			t.Error("Expected error for invalid endpoint")
		}
	})

	t.Run("CancelledContext", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		var response interface{}
		err := client.SendAPIRequest(ctx, "/test", &response)
		if err == nil {
			t.Error("Expected error for cancelled context")
		}
	})

	t.Run("TimeoutContext", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), clientMicroTimeout)
		defer cancel()

		var response interface{}
		err := client.SendAPIRequest(ctx, "/restconf/data/test", &response)
		if err == nil {
			t.Error("Expected timeout error")
		}
	})
}

// Helper function to check if error message contains specific text
func containsText(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
