package core

import (
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// ========================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// ========================================

// TestConfigStructure tests the basic structure of Config
func TestConfigStructure(t *testing.T) {
	config := Config{
		Controller:         "core.example.com",
		AccessToken:        "test-token",
		Timeout:            30 * time.Second,
		InsecureSkipVerify: true,
		Logger:             slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	if config.Controller != "core.example.com" {
		t.Errorf("Expected Controller to be 'core.example.com', got '%s'", config.Controller)
	}

	if config.AccessToken != "test-token" {
		t.Errorf("Expected AccessToken to be 'test-token', got '%s'", config.AccessToken)
	}

	if config.Timeout != 30*time.Second {
		t.Errorf("Expected Timeout to be 30s, got %v", config.Timeout)
	}

	if !config.InsecureSkipVerify {
		t.Error("Expected InsecureSkipVerify to be true")
	}

	if config.Logger == nil {
		t.Error("Expected Logger to be set")
	}
}

// ========================================
// 2. TABLE-DRIVEN TEST PATTERNS
// ========================================

// TestConfigDefaults tests Config with default values
func TestConfigDefaults(t *testing.T) {
	testCases := []struct {
		name   string
		config Config
	}{
		{"EmptyConfig", Config{}},
		{"PartialConfig", Config{Controller: "test.com"}},
		{"MinimalConfig", Config{Controller: "test.com", AccessToken: "token"}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Test that config can be created with various combinations
			config := tt.config

			// Verify the config structure is valid and accessible
			if config.Controller != tt.config.Controller {
				t.Errorf("Expected Controller to match, got '%s'", config.Controller)
			}

			if config.AccessToken != tt.config.AccessToken {
				t.Errorf("Expected AccessToken to match, got '%s'", config.AccessToken)
			}

			// Test that zero values are handled correctly
			if config.Timeout < 0 {
				t.Error("Timeout should not be negative")
			}

			// Test that boolean field is properly set
			if config.InsecureSkipVerify != tt.config.InsecureSkipVerify {
				t.Error("InsecureSkipVerify should match original")
			}

			// Test logger field accessibility
			if config.Logger != tt.config.Logger {
				t.Error("Logger should match original")
			}
		})
	}
}

// TestConfigValidation tests various configuration scenarios
func TestConfigValidation(t *testing.T) {
	testCases := []struct {
		name               string
		controller         string
		accessToken        string
		timeout            time.Duration
		insecureSkipVerify bool
		logger             *slog.Logger
		expectedValid      bool
	}{
		{
			name:               "ValidConfig",
			controller:         "core.example.com",
			accessToken:        "dGVzdDp0ZXN0",
			timeout:            30 * time.Second,
			insecureSkipVerify: false,
			logger:             slog.New(slog.NewTextHandler(io.Discard, nil)),
			expectedValid:      true,
		},
		{
			name:               "ValidInsecureConfig",
			controller:         "192.168.1.100",
			accessToken:        "test-token",
			timeout:            15 * time.Second,
			insecureSkipVerify: true,
			logger:             nil,
			expectedValid:      true,
		},
		{
			name:               "EmptyController",
			controller:         "",
			accessToken:        "test-token",
			timeout:            30 * time.Second,
			insecureSkipVerify: false,
			logger:             slog.New(slog.NewTextHandler(io.Discard, nil)),
			expectedValid:      false,
		},
		{
			name:               "EmptyAccessToken",
			controller:         "core.example.com",
			accessToken:        "",
			timeout:            30 * time.Second,
			insecureSkipVerify: false,
			logger:             slog.New(slog.NewTextHandler(io.Discard, nil)),
			expectedValid:      false,
		},
		{
			name:               "ZeroTimeout",
			controller:         "core.example.com",
			accessToken:        "test-token",
			timeout:            0,
			insecureSkipVerify: false,
			logger:             slog.New(slog.NewTextHandler(io.Discard, nil)),
			expectedValid:      true, // Zero timeout is valid, will use default
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			config := Config{
				Controller:         tt.controller,
				AccessToken:        tt.accessToken,
				Timeout:            tt.timeout,
				InsecureSkipVerify: tt.insecureSkipVerify,
				Logger:             tt.logger,
			}

			// Basic validation using our validation functions
			controllerValid := validation.IsValidController(config.Controller)
			tokenValid := validation.IsValidAccessToken(config.AccessToken)

			isValid := controllerValid && tokenValid

			if isValid != tt.expectedValid {
				t.Errorf("Expected config validity to be %v, got %v", tt.expectedValid, isValid)
			}

			// Test all fields are properly set
			if config.Timeout != tt.timeout {
				t.Errorf("Expected Timeout to be %v, got %v", tt.timeout, config.Timeout)
			}

			if config.InsecureSkipVerify != tt.insecureSkipVerify {
				t.Errorf(
					"Expected InsecureSkipVerify to be %v, got %v",
					tt.insecureSkipVerify,
					config.InsecureSkipVerify,
				)
			}

			if config.Logger != tt.logger {
				t.Errorf("Expected Logger to match test case logger")
			}
		})
	}
}

// ========================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// ========================================

// TestConfigFieldTypes tests that all fields have correct types
func TestConfigFieldTypes(t *testing.T) {
	config := Config{}

	// Test field types and assignments
	config.Controller = "string"
	if config.Controller != "string" {
		t.Error("Controller field assignment failed")
	}

	config.AccessToken = "string"
	if config.AccessToken != "string" {
		t.Error("AccessToken field assignment failed")
	}

	config.Timeout = time.Duration(0)
	if config.Timeout != 0 {
		t.Error("Timeout field assignment failed")
	}

	config.InsecureSkipVerify = false
	if config.InsecureSkipVerify != false {
		t.Error("InsecureSkipVerify field assignment failed")
	}

	config.Logger = (*slog.Logger)(nil)
	if config.Logger != nil {
		t.Error("Logger field assignment failed")
	}

	// Verify conditional assignments work without type errors
	if config.Controller == "string" {
		config.Controller = "default"
		if config.Controller != "default" {
			t.Error("Conditional Controller assignment failed")
		}
	}

	if config.AccessToken == "string" {
		config.AccessToken = "default-token"
		if config.AccessToken != "default-token" {
			t.Error("Conditional AccessToken assignment failed")
		}
	}

	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
		if config.Timeout != DefaultTimeout {
			t.Error("Conditional Timeout assignment failed")
		}
	}

	// Test boolean toggle
	originalSkipVerify := config.InsecureSkipVerify
	config.InsecureSkipVerify = !config.InsecureSkipVerify
	if config.InsecureSkipVerify == originalSkipVerify {
		t.Error("Boolean toggle assignment failed")
	}

	// Test logger assignment
	if config.Logger == nil {
		config.Logger = slog.Default()
		if config.Logger == nil {
			t.Error("Logger assignment to default failed")
		}
	}
}

// TestConfigCopy tests config copying/assignment
func TestConfigCopy(t *testing.T) {
	original := Config{
		Controller:         "original.com",
		AccessToken:        "original-token",
		Timeout:            45 * time.Second,
		InsecureSkipVerify: true,
		Logger:             slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	// Test struct copy
	cloned := original

	// Verify copy has same values
	if cloned.Controller != original.Controller {
		t.Errorf("Expected Controller to be copied, got '%s' vs '%s'", cloned.Controller, original.Controller)
	}

	if cloned.AccessToken != original.AccessToken {
		t.Errorf("Expected AccessToken to be copied, got '%s' vs '%s'", cloned.AccessToken, original.AccessToken)
	}

	if cloned.Timeout != original.Timeout {
		t.Errorf("Expected Timeout to be copied, got %v vs %v", cloned.Timeout, original.Timeout)
	}

	if cloned.InsecureSkipVerify != original.InsecureSkipVerify {
		t.Errorf(
			"Expected InsecureSkipVerify to be copied, got %v vs %v",
			cloned.InsecureSkipVerify,
			original.InsecureSkipVerify,
		)
	}

	if cloned.Logger != original.Logger {
		t.Error("Expected Logger to be copied (same reference)")
	}

	// Test that modifying copy doesn't affect original
	cloned.Controller = "modified.com"
	if original.Controller == "modified.com" {
		t.Error("Expected original Controller to remain unchanged")
	}
}

// TestConfigZeroValue tests zero value behavior
func TestConfigZeroValue(t *testing.T) {
	var config Config

	// Test zero values
	if config.Controller != "" {
		t.Errorf("Expected zero Controller to be empty, got '%s'", config.Controller)
	}

	if config.AccessToken != "" {
		t.Errorf("Expected zero AccessToken to be empty, got '%s'", config.AccessToken)
	}

	if config.Timeout != 0 {
		t.Errorf("Expected zero Timeout to be 0, got %v", config.Timeout)
	}

	if config.InsecureSkipVerify {
		t.Error("Expected zero InsecureSkipVerify to be false")
	}

	if config.Logger != nil {
		t.Error("Expected zero Logger to be nil")
	}
}

// TestConfigComparison tests config equality
func TestConfigComparison(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	config1 := Config{
		Controller:         "test.com",
		AccessToken:        "token",
		Timeout:            30 * time.Second,
		InsecureSkipVerify: false,
		Logger:             logger,
	}

	config2 := Config{
		Controller:         "test.com",
		AccessToken:        "token",
		Timeout:            30 * time.Second,
		InsecureSkipVerify: false,
		Logger:             logger,
	}

	// Test individual field equality
	if config1.Controller != config2.Controller {
		t.Error("Expected Controllers to be equal")
	}

	if config1.AccessToken != config2.AccessToken {
		t.Error("Expected AccessTokens to be equal")
	}

	if config1.Timeout != config2.Timeout {
		t.Error("Expected Timeouts to be equal")
	}

	if config1.InsecureSkipVerify != config2.InsecureSkipVerify {
		t.Error("Expected InsecureSkipVerify to be equal")
	}

	if config1.Logger != config2.Logger {
		t.Error("Expected Loggers to be equal (same reference)")
	}
}
