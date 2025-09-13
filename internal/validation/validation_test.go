package validation

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// ========================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// ========================================

// TestValidationUnit_Constants_Success tests validation constants.
func TestValidationUnit_Constants_Success(t *testing.T) {
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
			testutil.AssertIntEquals(t, tt.value, tt.expected, tt.name+" constant")
		})
	}
}

// ========================================
// 2. TABLE-DRIVEN TEST PATTERNS
// ========================================

// TestValidationUnit_ControllerValidation_Success tests controller validation.
func TestValidationUnit_ControllerValidation_Success(t *testing.T) {
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
			testutil.AssertBoolEquals(t, result, tt.expected, "IsValidController for "+tt.controller)
		})
	}
}

// TestValidationUnit_TokenValidation_Success tests access token validation.
func TestValidationUnit_TokenValidation_Success(t *testing.T) {
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
			testutil.AssertBoolEquals(t, result, tt.expected, "IsValidAccessToken for "+tt.token)
		})
	}
}

// TestValidationUnit_TimeoutValidation_Success tests timeout validation.
func TestValidationUnit_TimeoutValidation_Success(t *testing.T) {
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
			testutil.AssertBoolEquals(t, result, tt.expected, "IsPositiveTimeout for "+tt.timeout.String())
		})
	}
}

// ========================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// ========================================

// TestValidationUnit_ErrorTemplates_Success tests error message templates.
func TestValidationUnit_ErrorTemplates_Success(t *testing.T) {
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
			testutil.AssertStringEquals(t, tt.template, tt.expected, tt.name+" template")
		})
	}
}

// TestValidationUnit_BoundaryConditions_Success tests boundary conditions.
func TestValidationUnit_BoundaryConditions_Success(t *testing.T) {
	t.Run("ControllerValidation", func(t *testing.T) {
		// Test empty string
		testutil.AssertBoolEquals(t, IsValidController(""), false, "empty controller should be invalid")

		// Test single character
		testutil.AssertBoolEquals(t, IsValidController("a"), true, "single character controller should be valid")

		// Test with spaces
		testutil.AssertBoolEquals(t, IsValidController("test host"), true, "controller with spaces should be valid")
	})

	t.Run("AccessTokenValidation", func(t *testing.T) {
		// Test empty string
		testutil.AssertBoolEquals(t, IsValidAccessToken(""), false, "empty token should be invalid")

		// Test single character
		testutil.AssertBoolEquals(t, IsValidAccessToken("a"), true, "single character token should be valid")
	})

	t.Run("TimeoutValidation", func(t *testing.T) {
		// Test exactly at threshold
		threshold := time.Duration(ValidationTimeoutThreshold) * time.Second
		testutil.AssertBoolEquals(t, IsPositiveTimeout(threshold), false, "timeout at threshold should be invalid")

		// Test just above threshold
		justAbove := threshold + time.Nanosecond
		testutil.AssertBoolEquals(t, IsPositiveTimeout(justAbove), true, "timeout just above threshold should be valid")

		// Test very large timeout
		largeTimeout := 24 * time.Hour
		testutil.AssertBoolEquals(t, IsPositiveTimeout(largeTimeout), true, "very large timeout should be valid")
	})
}

// TestValidationUnit_MACValidation_Success tests MAC address validation.
func TestValidationUnit_MACValidation_Success(t *testing.T) {
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
			err := ValidateMACAddress(tt.macAddr)
			if tt.valid {
				testutil.AssertNoError(t, err, "MAC address "+tt.macAddr+" should be valid")
			} else {
				testutil.AssertError(t, err, "MAC address "+tt.macAddr+" should be invalid")
			}
		})
	}
}

// TestValidationUnit_MACNormalization_Success tests MAC address normalization.
func TestValidationUnit_MACNormalization_Success(t *testing.T) {
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
			result, err := NormalizeMACAddress(tt.input)
			testutil.AssertNoError(t, err, "NormalizeMACAddress should not return error for "+tt.input)
			testutil.AssertStringEquals(t, result, tt.expected, "NormalizeMACAddress for "+tt.input)
		})
	}

	// Test error cases
	errorCases := []struct {
		name  string
		input string
	}{
		{"EmptyString", ""},
		{"TooShort", "00:11:22:33:44"},
		{"TooLong", "00:11:22:33:44:55:66"},
		{"InvalidHex", "00:11:22:33:44:5z"},
	}

	for _, tt := range errorCases {
		t.Run("Error_"+tt.name, func(t *testing.T) {
			_, err := NormalizeMACAddress(tt.input)
			testutil.AssertError(t, err, "NormalizeMACAddress should return error for "+tt.input)
		})
	}
}

// TestValidationUnit_StringValidation_Success tests string validation.
func TestValidationUnit_StringValidation_Success(t *testing.T) {
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
			if tt.valid {
				testutil.AssertNoError(t, err, "string \""+tt.input+"\" should be valid")
			} else {
				testutil.AssertError(t, err, "string \""+tt.input+"\" should be invalid")
				testutil.AssertStringContains(t, err.Error(), tt.fieldName, "error should contain field name")
			}
		})
	}
}

// TestValidationUnit_StringEmptyCheck_Success tests string empty checking.
// TestIsValidTimeout tests the IsValidTimeout function
// TestValidationUnit_TimeoutCheck_Success tests timeout validation.
func TestValidationUnit_TimeoutCheck_Success(t *testing.T) {
	testCases := []struct {
		name     string
		timeout  time.Duration
		expected bool
	}{
		{"PositiveTimeout", 30 * time.Second, true},
		{"MinimumTimeout", 1 * time.Millisecond, true},
		{"ZeroTimeout", 0, false},
		{"NegativeTimeout", -1 * time.Second, false},
		{"LargeTimeout", 24 * time.Hour, true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidTimeout(tt.timeout)
			testutil.AssertBoolEquals(t, result, tt.expected, fmt.Sprintf("IsValidTimeout(%v)", tt.timeout))
		})
	}
}

// TestIsNonEmptyString tests the IsNonEmptyString function.
func TestValidationUnit_NonEmptyStringCheck_Success(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ValidString", "test", true},
		{"EmptyString", "", false},
		{"WhitespaceOnly", "   ", false},
		{"TabOnly", "\t", false},
		{"NewlineOnly", "\n", false},
		{"ValidWithWhitespace", " test ", true},
		{"SingleChar", "a", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNonEmptyString(tt.input)
			testutil.AssertBoolEquals(t, result, tt.expected, fmt.Sprintf("IsNonEmptyString(%q)", tt.input))
		})
	}
}

// TestIsValidMACAddr tests the IsValidMACAddr function.
func TestValidationUnit_MACFormatCheck_Success(t *testing.T) {
	testCases := []struct {
		name     string
		mac      string
		expected bool
	}{
		{"ValidColonFormat", "00:11:22:33:44:55", true},
		{"ValidHyphenFormat", "00-11-22-33-44-55", true},
		{"ValidNoSeparator", "001122334455", true},
		{"InvalidFormat", "invalid", false},
		{"EmptyString", "", false},
		{"TooShort", "00:11:22", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidMACAddr(tt.mac)
			testutil.AssertBoolEquals(t, result, tt.expected, fmt.Sprintf("IsValidMACAddr(%q)", tt.mac))
		})
	}
}

// TestValidationUnit_TagValidation_Success tests tag validation functions.
func TestValidationUnit_TagValidation_Success(t *testing.T) {
	t.Run("IsNonEmptyString", func(t *testing.T) {
		testCases := []struct {
			name     string
			tag      string
			expected bool
		}{
			{"ValidTag", "site-tag-1", true},
			{"EmptyTag", "", false},
			{"WhitespaceTag", "   ", false},
			{"DefaultTag", DefaultSiteTag, true},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				result := IsNonEmptyString(tt.tag)
				testutil.AssertBoolEquals(t, result, tt.expected, fmt.Sprintf("IsNonEmptyString(%q)", tt.tag))
			})
		}
	})

	t.Run("IsNonEmptyString", func(t *testing.T) {
		testCases := []struct {
			name     string
			tag      string
			expected bool
		}{
			{"ValidTag", "policy-tag-1", true},
			{"EmptyTag", "", false},
			{"WhitespaceTag", "   ", false},
			{"DefaultTag", DefaultPolicyTag, true},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				result := IsNonEmptyString(tt.tag)
				testutil.AssertBoolEquals(t, result, tt.expected, fmt.Sprintf("IsNonEmptyString(%q)", tt.tag))
			})
		}
	})

	t.Run("IsNonEmptyString", func(t *testing.T) {
		testCases := []struct {
			name     string
			tag      string
			expected bool
		}{
			{"ValidTag", "rf-tag-1", true},
			{"EmptyTag", "", false},
			{"WhitespaceTag", "   ", false},
			{"DefaultTag", DefaultRFTag, true},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				result := IsNonEmptyString(tt.tag)
				testutil.AssertBoolEquals(t, result, tt.expected, fmt.Sprintf("IsNonEmptyString(%q)", tt.tag))
			})
		}
	})
}

// TestValidationUnit_SlotIDValidation_Success tests slot ID validation.
func TestValidationUnit_SlotIDValidation_Success(t *testing.T) {
	testCases := []struct {
		name    string
		slotID  int
		wantErr bool
	}{
		{"ValidSlot0", 0, false},
		{"ValidSlot1", 1, false},
		{"ValidSlot100", 100, false},
		{"InvalidNegative", -1, true},
		{"InvalidNegativeLarge", -100, true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSlotID(tt.slotID)
			if tt.wantErr {
				testutil.AssertError(t, err, fmt.Sprintf("ValidateSlotID(%d)", tt.slotID))
			} else {
				testutil.AssertNoError(t, err, fmt.Sprintf("ValidateSlotID(%d)", tt.slotID))
			}
		})
	}
}

// TestValidateSpatialStream tests spatial stream validation.
func TestValidationUnit_SpatialStreamValidation_Success(t *testing.T) {
	testCases := []struct {
		name    string
		stream  int
		wantErr bool
	}{
		{"Valid1", 1, false},
		{"Valid4", 4, false},
		{"Valid8", 8, false},
		{"Invalid0", 0, true},
		{"Invalid9", 9, true},
		{"InvalidNegative", -1, true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSpatialStream(tt.stream)
			if tt.wantErr {
				testutil.AssertError(t, err, fmt.Sprintf("ValidateSpatialStream(%d)", tt.stream))
			} else {
				testutil.AssertNoError(t, err, fmt.Sprintf("ValidateSpatialStream(%d)", tt.stream))
			}
		})
	}
}

// TestValidateWlanID tests WLAN ID validation.
func TestValidationUnit_WlanIDValidation_Success(t *testing.T) {
	testCases := []struct {
		name    string
		wlanID  string
		wantErr bool
	}{
		{"ValidID", "1", false},
		{"ValidName", "test-wlan", false},
		{"EmptyID", "", true},
		{"WhitespaceID", "   ", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateWlanID(tt.wlanID)
			if tt.wantErr {
				testutil.AssertError(t, err, fmt.Sprintf("ValidateWlanID(%q)", tt.wlanID))
			} else {
				testutil.AssertNoError(t, err, fmt.Sprintf("ValidateWlanID(%q)", tt.wlanID))
			}
		})
	}
}

// TestValidateMACAddress tests rogue address validation.
func TestValidationUnit_RogueAddressValidation_Success(t *testing.T) {
	testCases := []struct {
		name    string
		address string
		wantErr bool
	}{
		{"ValidAddress", "00:11:22:33:44:55", false},
		{"EmptyAddress", "", true},
		{"WhitespaceAddress", "   ", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMACAddress(tt.address)
			if tt.wantErr {
				testutil.AssertError(t, err, fmt.Sprintf("ValidateMACAddress(%q)", tt.address))
			} else {
				testutil.AssertNoError(t, err, fmt.Sprintf("ValidateMACAddress(%q)", tt.address))
			}
		})
	}
}

// TestValidateMACAddress tests rogue client address validation.
func TestValidationUnit_RogueClientAddressValidation_Success(t *testing.T) {
	testCases := []struct {
		name    string
		address string
		wantErr bool
	}{
		{"ValidAddress", "00:11:22:33:44:55", false},
		{"EmptyAddress", "", true},
		{"WhitespaceAddress", "   ", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMACAddress(tt.address)
			if tt.wantErr {
				testutil.AssertError(t, err, fmt.Sprintf("ValidateMACAddress(%q)", tt.address))
			} else {
				testutil.AssertNoError(t, err, fmt.Sprintf("ValidateMACAddress(%q)", tt.address))
			}
		})
	}
}

// TestHelperFunctions tests direct Go idioms instead of helper functions.
func TestValidationUnit_DirectGoIdioms_Success(t *testing.T) {
	t.Run("DirectNilCheck", func(t *testing.T) {
		var nilInterface interface{} = nil
		result := nilInterface == nil
		testutil.AssertBoolEquals(t, result, true, "nilInterface == nil")

		notNilString := "not nil"
		result = notNilString == ""
		testutil.AssertBoolEquals(t, result, false, "notNilString is not empty")

		// Note: typed nil pointers are not nil in Go interface comparison
		var nilPointer *int
		result = nilPointer == nil
		testutil.AssertBoolEquals(t, result, true, "nilPointer == nil")
	})

	t.Run("DirectErrorCheck", func(t *testing.T) {
		var err error = nil
		result := err != nil
		testutil.AssertBoolEquals(t, result, false, "err != nil with nil error")

		err = errors.New("test error")
		result = err != nil
		testutil.AssertBoolEquals(t, result, true, "err != nil with actual error")
	})

	t.Run("DirectValidityCheck", func(t *testing.T) {
		response := "response"
		var err error = nil
		result := response != "" && err == nil
		testutil.AssertBoolEquals(t, result, true, "response != \"\" && err == nil")

		response = ""
		err = nil
		result = response != "" && err == nil
		testutil.AssertBoolEquals(t, result, false, "empty response fails validity check")

		response = "response"
		err = errors.New("error")
		result = response != "" && err == nil
		testutil.AssertBoolEquals(t, result, false, "response with error fails validity check")
	})

	t.Run("DirectFieldComparison", func(t *testing.T) {
		// Test literal equality comparison
		testString := "test"
		result := testString == "test"
		testutil.AssertBoolEquals(t, result, true, "\"test\" == \"test\"")

		result = "test" == "different"
		testutil.AssertBoolEquals(t, result, false, "\"test\" == \"different\"")
	})
}

// TestCompositeValidationFunctions tests composite validation functions.
func TestValidationUnit_CompositeValidation_Success(t *testing.T) {
	t.Run("HasValidTags", func(t *testing.T) {
		testCases := []struct {
			name      string
			siteTag   string
			policyTag string
			rfTag     string
			expected  bool
		}{
			{"AllValid", "site", "policy", "rf", true},
			{"OnlySite", "site", "", "", true},
			{"OnlyPolicy", "", "policy", "", true},
			{"OnlyRF", "", "", "rf", true},
			{"NoneValid", "", "", "", false},
			{"AllWhitespace", "   ", "   ", "   ", false},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				result := HasValidTags(tt.siteTag, tt.policyTag, tt.rfTag)
				testutil.AssertBoolEquals(t, result, tt.expected,
					fmt.Sprintf("HasValidTags(%q, %q, %q)", tt.siteTag, tt.policyTag, tt.rfTag))
			})
		}
	})

	t.Run("HasValidMACOrName", func(t *testing.T) {
		testCases := []struct {
			name     string
			apMac    string
			apName   string
			expected bool
		}{
			{"OnlyMAC", "00:11:22:33:44:55", "", true},
			{"OnlyName", "", "ap-name", true},
			{"Both", "00:11:22:33:44:55", "ap-name", false},
			{"Neither", "", "", false},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				result := HasValidMACOrName(tt.apMac, tt.apName)
				testutil.AssertBoolEquals(t, result, tt.expected,
					fmt.Sprintf("HasValidMACOrName(%q, %q)", tt.apMac, tt.apName))
			})
		}
	})

	t.Run("HasEitherMACOrName", func(t *testing.T) {
		testCases := []struct {
			name     string
			apMac    string
			apName   string
			expected bool
		}{
			{"OnlyMAC", "00:11:22:33:44:55", "", true},
			{"OnlyName", "", "ap-name", true},
			{"Both", "00:11:22:33:44:55", "ap-name", true},
			{"Neither", "", "", false},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				result := HasEitherMACOrName(tt.apMac, tt.apName)
				testutil.AssertBoolEquals(t, result, tt.expected,
					fmt.Sprintf("HasEitherMACOrName(%q, %q)", tt.apMac, tt.apName))
			})
		}
	})
}

// TestSelectNonEmptyValue tests value selection.
func TestValidationUnit_SelectNonEmptyValue_Success(t *testing.T) {
	testCases := []struct {
		name         string
		primary      string
		defaultValue string
		expected     string
	}{
		{"PrimaryNotEmpty", "primary", "default", "primary"},
		{"PrimaryEmpty", "", "default", "default"},
		{"PrimaryWhitespace", "   ", "default", "default"},
		{"BothEmpty", "", "", ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := SelectNonEmptyValue(tt.primary, tt.defaultValue)
			testutil.AssertStringEquals(t, result, tt.expected,
				fmt.Sprintf("SelectNonEmptyValue(%q, %q)", tt.primary, tt.defaultValue))
		})
	}
}

// TestIsValidTagAssignment tests tag assignment validation.
func TestValidationUnit_TagAssignmentValidation_Success(t *testing.T) {
	testCases := []struct {
		name     string
		tagValue string
		tagType  string
		expected bool
	}{
		{"ValidSiteTag", "site-tag", "site", true},
		{"ValidPolicyTag", "policy-tag", "policy", true},
		{"ValidRFTag", "rf-tag", "rf", true},
		{"InvalidType", "tag-value", "invalid", false},
		{"EmptyValue", "", "site", false},
		{"EmptyType", "tag-value", "", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidTagAssignment(tt.tagValue, tt.tagType)
			testutil.AssertBoolEquals(t, result, tt.expected,
				fmt.Sprintf("IsValidTagAssignment(%q, %q)", tt.tagValue, tt.tagType))
		})
	}
}

// TestDefaultTagConstants tests default tag constants.
func TestValidationUnit_DefaultTagConstants_Success(t *testing.T) {
	testCases := []struct {
		name     string
		constant string
		expected string
	}{
		{"DefaultSiteTag", DefaultSiteTag, "default-site-tag"},
		{"DefaultPolicyTag", DefaultPolicyTag, "default-policy-tag"},
		{"DefaultRFTag", DefaultRFTag, "default-rf-tag"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AssertStringEquals(t, tt.constant, tt.expected, tt.name)
		})
	}
}
