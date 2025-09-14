package validation

import (
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// TestValidationUnit_Constructor_Success tests validation function availability.
func TestValidationUnit_Constructor_Success(t *testing.T) {
	// Test constants are accessible
	testutil.AssertIntEquals(t, MinEndpointLength, 10, "MinEndpointLength constant")
	testutil.AssertIntEquals(t, MinTokenLength, 8, "MinTokenLength constant")
	testutil.AssertIntEquals(t, ValidationTimeoutThreshold, 1, "ValidationTimeoutThreshold constant")

	// Test default tags are accessible
	testutil.AssertStringEquals(t, DefaultSiteTag, "default-site-tag", "DefaultSiteTag constant")
	testutil.AssertStringEquals(t, DefaultPolicyTag, "default-policy-tag", "DefaultPolicyTag constant")
	testutil.AssertStringEquals(t, DefaultRFTag, "default-rf-tag", "DefaultRFTag constant")
}

// TestValidationUnit_GetOperations_Success tests basic validation functions.
func TestValidationUnit_GetOperations_Success(t *testing.T) {
	// Controller validation
	testutil.AssertBoolEquals(t, IsValidController("core.example.com"), true, "valid controller")
	testutil.AssertBoolEquals(t, IsValidController(""), false, "empty controller")
	testutil.AssertBoolEquals(t, IsValidController("192.168.1.100"), true, "IP controller")

	// Token validation
	testutil.AssertBoolEquals(t, IsValidAccessToken("valid-token"), true, "valid token")
	testutil.AssertBoolEquals(t, IsValidAccessToken(""), false, "empty token")
	testutil.AssertBoolEquals(t, IsValidAccessToken("short"), true, "short token is valid (only checks non-empty)")

	// MAC validation
	testutil.AssertBoolEquals(t, IsValidMACAddr("00:11:22:33:44:55"), true, "valid colon MAC")
	testutil.AssertBoolEquals(t, IsValidMACAddr("00-11-22-33-44-55"), true, "valid hyphen MAC")
	testutil.AssertBoolEquals(t, IsValidMACAddr("001122334455"), true, "valid no-separator MAC")
	testutil.AssertBoolEquals(t, IsValidMACAddr("invalid"), false, "invalid MAC")
	testutil.AssertBoolEquals(t, IsValidMACAddr(""), false, "empty MAC")

	// Timeout validation
	testutil.AssertBoolEquals(t, IsValidTimeout(30*time.Second), true, "valid timeout")
	testutil.AssertBoolEquals(t, IsValidTimeout(0), false, "zero timeout")
	testutil.AssertBoolEquals(t, IsValidTimeout(-1*time.Second), false, "negative timeout")
	testutil.AssertBoolEquals(t, IsPositiveTimeout(30*time.Second), true, "positive timeout")
	testutil.AssertBoolEquals(t, IsPositiveTimeout(0), false, "zero positive timeout")

	// String validation
	testutil.AssertBoolEquals(t, IsNonEmptyString("test"), true, "valid string")
	testutil.AssertBoolEquals(t, IsNonEmptyString(""), false, "empty string")
	testutil.AssertBoolEquals(t, IsNonEmptyString("   "), false, "whitespace string")
}

// TestValidationUnit_SetOperations_Success tests MAC normalization and validation functions.
func TestValidationUnit_SetOperations_Success(t *testing.T) {
	// MAC normalization
	normalized, err := NormalizeMACAddress("00:11:22:33:44:55")
	testutil.AssertNoError(t, err, "colon format normalization should not error")
	testutil.AssertStringEquals(t, normalized, "00:11:22:33:44:55", "colon format normalization")

	normalized, err = NormalizeMACAddress("00-11-22-33-44-55")
	testutil.AssertNoError(t, err, "hyphen format normalization should not error")
	testutil.AssertStringEquals(t, normalized, "00:11:22:33:44:55", "hyphen format normalization")

	normalized, err = NormalizeMACAddress("001122334455")
	testutil.AssertNoError(t, err, "no-separator format normalization should not error")
	testutil.AssertStringEquals(t, normalized, "00:11:22:33:44:55", "no-separator format normalization")

	// MAC validation with error
	err = ValidateMACAddress("00:11:22:33:44:55")
	testutil.AssertNoError(t, err, "valid MAC validation")

	err = ValidateMACAddress("")
	testutil.AssertError(t, err, "empty MAC validation should error")

	err = ValidateMACAddress("invalid")
	testutil.AssertError(t, err, "invalid MAC validation should error")
}

// TestValidationUnit_ErrorHandling_Success tests error template functions.
func TestValidationUnit_ErrorHandling_Success(t *testing.T) {
	// Test error templates are available
	testutil.AssertStringEquals(
		t,
		EndpointMismatchErrorTemplate,
		"Expected %s = %s, got %s",
		"endpoint mismatch template",
	)
	testutil.AssertStringEquals(
		t,
		EmptyEndpointErrorTemplate,
		"%s endpoint is empty",
		"empty endpoint template",
	)
	testutil.AssertStringEquals(
		t,
		ShortEndpointErrorTemplate,
		"%s endpoint is too short: %s",
		"short endpoint template",
	)
	testutil.AssertStringEquals(
		t,
		InvalidEndpointErrorTemplate,
		"%s endpoint has invalid format: %s",
		"invalid endpoint template",
	)
}

// TestValidationUnit_ValidationErrors_Success tests boundary conditions and edge cases.
func TestValidationUnit_ValidationErrors_Success(t *testing.T) {
	// Boundary timeout testing
	threshold := time.Duration(ValidationTimeoutThreshold) * time.Second
	testutil.AssertBoolEquals(
		t,
		IsPositiveTimeout(threshold),
		false,
		"timeout at threshold should be invalid",
	)

	justAbove := threshold + time.Nanosecond
	testutil.AssertBoolEquals(
		t,
		IsPositiveTimeout(justAbove),
		true,
		"timeout just above threshold should be valid",
	)

	// Single character validation
	testutil.AssertBoolEquals(t, IsValidController("a"), true, "single character controller should be valid")
	testutil.AssertBoolEquals(
		t,
		IsValidAccessToken("a"),
		true,
		"single character token is valid (only checks non-empty)",
	)

	// Composite validation scenarios
	controller := "test.cisco.com"
	token := "test-token"
	testutil.AssertBoolEquals(
		t,
		IsValidController(controller) && IsValidAccessToken(token),
		true,
		"valid controller and token",
	)

	emptyController := ""
	testutil.AssertBoolEquals(
		t,
		IsValidController(emptyController) && IsValidAccessToken(token),
		false,
		"empty controller fails composite",
	)

	// Additional function coverage
	testutil.AssertBoolEquals(t, IsStringEmpty(""), true, "empty string is empty")
	testutil.AssertBoolEquals(t, IsStringEmpty("test"), false, "non-empty string is not empty")

	// Tag validation (HasValidTags checks if at least one tag is provided)
	testutil.AssertBoolEquals(t, HasValidTags("site", "policy", "rf"), true, "all valid tags")
	testutil.AssertBoolEquals(t, HasValidTags("", "policy", "rf"), true, "empty site tag but others valid")
	testutil.AssertBoolEquals(t, HasValidTags("", "", ""), false, "all empty tags")

	// MAC or name validation (HasValidMACOrName requires exactly one, not both)
	testutil.AssertBoolEquals(t, HasValidMACOrName("00:11:22:33:44:55", ""), true, "valid MAC only")
	testutil.AssertBoolEquals(t, HasValidMACOrName("", "ap1"), true, "valid name only")
	testutil.AssertBoolEquals(t, HasValidMACOrName("00:11:22:33:44:55", "ap1"), false, "both MAC and name not allowed")
	testutil.AssertBoolEquals(t, HasValidMACOrName("", ""), false, "empty MAC and name")
	testutil.AssertBoolEquals(t, HasEitherMACOrName("00:11:22:33:44:55", ""), true, "has MAC")
	testutil.AssertBoolEquals(t, HasEitherMACOrName("", "ap1"), true, "has name")

	// Value selection
	testutil.AssertStringEquals(t, SelectNonEmptyValue("primary", "default"), "primary", "primary value selected")
	testutil.AssertStringEquals(t, SelectNonEmptyValue("", "default"), "default", "default value selected")

	// Tag assignment validation
	testutil.AssertBoolEquals(t, IsValidTagAssignment("tag1", "site"), true, "valid site tag assignment")
	testutil.AssertBoolEquals(t, IsValidTagAssignment("tag1", "policy"), true, "valid policy tag assignment")
	testutil.AssertBoolEquals(t, IsValidTagAssignment("tag1", "rf"), true, "valid rf tag assignment")
	testutil.AssertBoolEquals(t, IsValidTagAssignment("", "site"), false, "empty tag assignment")
	testutil.AssertBoolEquals(t, IsValidTagAssignment("tag1", "invalid"), false, "invalid tag type")

	// String validation with field name
	err := ValidateNonEmptyString("test", "testField")
	testutil.AssertNoError(t, err, "valid string validation with field name")

	err = ValidateNonEmptyString("", "testField")
	testutil.AssertError(t, err, "empty string validation with field name should error")

	// Slot ID validation
	err = ValidateSlotID(1)
	testutil.AssertNoError(t, err, "valid slot ID")

	err = ValidateSlotID(-1)
	testutil.AssertError(t, err, "invalid slot ID should error")

	// Spatial stream validation
	err = ValidateSpatialStream(2)
	testutil.AssertNoError(t, err, "valid spatial stream")

	err = ValidateSpatialStream(-1)
	testutil.AssertError(t, err, "invalid spatial stream should error")

	// WLAN ID validation
	err = ValidateWlanID("1")
	testutil.AssertNoError(t, err, "valid WLAN ID")

	err = ValidateWlanID("")
	testutil.AssertError(t, err, "empty WLAN ID should error")
}
