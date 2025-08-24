package validation

import (
	"fmt"
	"strings"
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
		{"MinEndpointLength", MinEndpointLength, 10},
		{"MinTokenLength", MinTokenLength, 8},
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
		{"SpaceOnlyToken", " ", false}, // Whitespace-only string is invalid
		{"TabToken", "\t", false},      // Whitespace-only string is invalid
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

// TestValidateAPMac tests MAC address validation
func TestValidateAPMac(t *testing.T) {
	testCases := []struct {
		name    string
		macAddr string
		valid   bool
	}{
		{"ValidColonFormat", "00:11:22:33:44:55", true},
		{"ValidHyphenFormat", "00-11-22-33-44-55", true},
		{"ValidNoSeparator", "001122334455", true},
		{"ValidUppercase", "AA:BB:CC:DD:EE:FF", true},
		{"ValidMixed", "12:34:56:78:9a:bc", true},
		{"EmptyString", "", false},
		{"TooShort", "00:11:22:33:44", false},
		{"TooLong", "00:11:22:33:44:55:66", false},
		{"InvalidHex", "00:11:22:33:44:gg", false},
		{"InvalidChar", "00:11:22:33:44:5z", false},
		{"NonHexString", "xyz", false},
		{"OnlyNumbers", "123456789012", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAPMac(tt.macAddr)
			if tt.valid && err != nil {
				t.Errorf("Expected MAC address %s to be valid, got error: %v", tt.macAddr, err)
			}
			if !tt.valid && err == nil {
				t.Errorf("Expected MAC address %s to be invalid, but validation passed", tt.macAddr)
			}
		})
	}
}

// TestNormalizeAPMac tests MAC address normalization
func TestNormalizeAPMac(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"ColonFormat", "00:11:22:33:44:55", "00:11:22:33:44:55"},
		{"HyphenFormat", "00-11-22-33-44-55", "00:11:22:33:44:55"},
		{"NoSeparator", "001122334455", "00:11:22:33:44:55"},
		{"UppercaseToLowercase", "AA:BB:CC:DD:EE:FF", "aa:bb:cc:dd:ee:ff"},
		{"MixedCase", "Aa:Bb:Cc:Dd:Ee:Ff", "aa:bb:cc:dd:ee:ff"},
		{"DotFormat", "00.11.22.33.44.55", "00:11:22:33:44:55"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeAPMac(tt.input)
			if result != tt.expected {
				t.Errorf("Expected NormalizeAPMac(%s) to be %s, got %s", tt.input, tt.expected, result)
			}
		})
	}
}

// TestValidateNonEmptyString tests string validation
func TestValidateNonEmptyString(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		fieldName string
		valid     bool
	}{
		{"ValidString", "test", "field", true},
		{"EmptyString", "", "field", false},
		{"WhitespaceOnly", "   ", "field", false},
		{"TabOnly", "\t", "field", false},
		{"NewlineOnly", "\n", "field", false},
		{"ValidWithWhitespace", " test ", "field", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNonEmptyString(tt.input, tt.fieldName)
			if tt.valid && err != nil {
				t.Errorf("Expected string %q to be valid, got error: %v", tt.input, err)
			}
			if !tt.valid && err == nil {
				t.Errorf("Expected string %q to be invalid, but validation passed", tt.input)
			}
			if !tt.valid && err != nil && !strings.Contains(err.Error(), tt.fieldName) {
				t.Errorf("Expected error to contain field name %q, got: %v", tt.fieldName, err)
			}
		})
	}
}

// TestIsStringEmpty tests string empty checking
func TestIsStringEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"EmptyString", "", true},
		{"NonEmptyString", "test", false},
		{"SpaceString", " ", false},
		{"TabString", "\t", false},
		{"NewlineString", "\n", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsStringEmpty(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestIsHexChar tests the IsHexChar function
func TestIsHexChar(t *testing.T) {
	validHexChars := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a',
		'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F',
	}
	for _, char := range validHexChars {
		t.Run(fmt.Sprintf("valid_%c", char), func(t *testing.T) {
			if !IsHexChar(char) {
				t.Errorf("Expected %c to be a valid hex character", char)
			}
		})
	}

	invalidHexChars := []rune{'g', 'G', 'z', 'Z', '!', '@', ' ', '\t', '\n', ':', '-', '.'}
	for _, char := range invalidHexChars {
		t.Run(fmt.Sprintf("invalid_%c", char), func(t *testing.T) {
			if IsHexChar(char) {
				t.Errorf("Expected %c to be an invalid hex character", char)
			}
		})
	}
}
