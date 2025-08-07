package validation

import (
	"testing"
	"time"
)

// ========================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// ========================================

// TestValidationConstants tests validation constants
func TestValidationConstants(t *testing.T) {
	testCases := []struct {
		name     string
		value    int
		expected int
	}{
		{"MinEndpointLengthChars", MinEndpointLengthChars, 10},
		{"MinTokenLengthChars", MinTokenLengthChars, 8},
		{"MinEndpointLength", MinEndpointLength, 10},
		{"MinTokenLength", MinTokenLength, 8},
		{"ZeroTimeoutSeconds", ZeroTimeoutSeconds, 0},
		{"ValidationTimeoutThreshold", ValidationTimeoutThreshold, 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %s to be %d, got %d", tt.name, tt.expected, tt.value)
			}
		})
	}
}

// ========================================
// 2. TABLE-DRIVEN TEST PATTERNS
// ========================================

// TestIsValidController tests controller validation
func TestIsValidController(t *testing.T) {
	testCases := []struct {
		name       string
		controller string
		expected   bool
	}{
		{"ValidController", "core.example.com", true},
		{"ValidIP", "192.168.1.100", true},
		{"ValidLocalhost", "localhost", true},
		{"EmptyController", "", false},
		{"ValidWithPort", "core.example.com:443", true},
		{"ValidHostname", "test.local", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidController(tt.controller)
			if result != tt.expected {
				t.Errorf("Expected IsValidController('%s') to be %v, got %v", tt.controller, tt.expected, result)
			}
		})
	}
}

// TestIsValidAccessToken tests access token validation
func TestIsValidAccessToken(t *testing.T) {
	testCases := []struct {
		name     string
		token    string
		expected bool
	}{
		{"ValidToken", "dGVzdDp0ZXN0", true},
		{"ValidLongToken", "YWRtaW46cGFzc3dvcmQxMjM0NTY3ODkw", true},
		{"ValidShortToken", "dGVzdA==", true},
		{"EmptyToken", "", false},
		{"SpaceOnlyToken", " ", true}, // Non-empty string is valid
		{"TabToken", "\t", true},      // Non-empty string is valid
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidAccessToken(tt.token)
			if result != tt.expected {
				t.Errorf("Expected IsValidAccessToken('%s') to be %v, got %v", tt.token, tt.expected, result)
			}
		})
	}
}

// TestIsPositiveTimeout tests timeout validation
func TestIsPositiveTimeout(t *testing.T) {
	testCases := []struct {
		name     string
		timeout  time.Duration
		expected bool
	}{
		{"ValidTimeout", 30 * time.Second, true},
		{"MinimumValidTimeout", 2 * time.Second, true},
		{"ZeroTimeout", 0, false},
		{"NegativeTimeout", -1 * time.Second, false},
		{"VeryShortTimeout", 500 * time.Millisecond, false},  // Less than 1 second
		{"ExactlyOneSecond", 1 * time.Second, false},         // Equal to threshold
		{"JustOverOneSecond", 1001 * time.Millisecond, true}, // Just over threshold
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPositiveTimeout(tt.timeout)
			if result != tt.expected {
				t.Errorf("Expected IsPositiveTimeout(%v) to be %v, got %v", tt.timeout, tt.expected, result)
			}
		})
	}
}

// ========================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// ========================================

// TestValidationErrorTemplates tests error message templates
func TestValidationErrorTemplates(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		expected string
	}{
		{"EndpointMismatchErrorTemplate", EndpointMismatchErrorTemplate, "Expected %s = %s, got %s"},
		{"EmptyEndpointErrorTemplate", EmptyEndpointErrorTemplate, "%s endpoint is empty"},
		{"ShortEndpointErrorTemplate", ShortEndpointErrorTemplate, "%s endpoint is too short: %s"},
		{"InvalidEndpointErrorTemplate", InvalidEndpointErrorTemplate, "%s endpoint has invalid format: %s"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.template != tt.expected {
				t.Errorf("Expected %s to be '%s', got '%s'", tt.name, tt.expected, tt.template)
			}
		})
	}
}

// TestValidationFunctionsBoundaryConditions tests boundary conditions
func TestValidationFunctionsBoundaryConditions(t *testing.T) {
	t.Run("ControllerValidation", func(t *testing.T) {
		// Test empty string
		if IsValidController("") {
			t.Error("Expected empty controller to be invalid")
		}

		// Test single character
		if !IsValidController("a") {
			t.Error("Expected single character controller to be valid")
		}

		// Test with spaces
		if !IsValidController("test host") {
			t.Error("Expected controller with spaces to be valid")
		}
	})

	t.Run("AccessTokenValidation", func(t *testing.T) {
		// Test empty string
		if IsValidAccessToken("") {
			t.Error("Expected empty token to be invalid")
		}

		// Test single character
		if !IsValidAccessToken("a") {
			t.Error("Expected single character token to be valid")
		}

		// Test unicode
		if !IsValidAccessToken("ユーザー") {
			t.Error("Expected unicode token to be valid")
		}
	})

	t.Run("TimeoutValidation", func(t *testing.T) {
		// Test exactly at threshold
		threshold := time.Duration(ValidationTimeoutThreshold) * time.Second
		if IsPositiveTimeout(threshold) {
			t.Error("Expected timeout at threshold to be invalid")
		}

		// Test just above threshold
		justAbove := threshold + time.Nanosecond
		if !IsPositiveTimeout(justAbove) {
			t.Error("Expected timeout just above threshold to be valid")
		}

		// Test very large timeout
		largeTimeout := 24 * time.Hour
		if !IsPositiveTimeout(largeTimeout) {
			t.Error("Expected very large timeout to be valid")
		}
	})
}
