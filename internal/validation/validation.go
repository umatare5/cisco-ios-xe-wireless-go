// Package validation provides common validation functions for the Cisco Wireless Network Controller client.
package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Default tag values for AP operations
const (
	// DefaultSiteTag represents the default site tag value
	DefaultSiteTag = "default-site-tag"

	// DefaultPolicyTag represents the default policy tag value
	DefaultPolicyTag = "default-policy-tag"

	// DefaultRFTag represents the default RF tag value
	DefaultRFTag = "default-rf-tag"
)

// Validation constants
const (
	// MinEndpointLength is the minimum length for API endpoints
	MinEndpointLength = 10

	// MinTokenLength is the minimum length for authentication tokens
	MinTokenLength = 8

	// ValidationTimeoutThreshold is the minimum timeout for validation
	ValidationTimeoutThreshold = 1
)

// Error message templates for validation errors
const (
	// EndpointMismatchErrorTemplate is used for endpoint validation errors
	EndpointMismatchErrorTemplate = "Expected %s = %s, got %s"

	// EmptyEndpointErrorTemplate is used when an endpoint is empty
	EmptyEndpointErrorTemplate = "%s endpoint is empty"

	// ShortEndpointErrorTemplate is used when an endpoint is too short
	ShortEndpointErrorTemplate = "%s endpoint is too short: %s"

	// InvalidEndpointErrorTemplate is used for invalid endpoint formats
	InvalidEndpointErrorTemplate = "%s endpoint has invalid format: %s"
)

// Core validation functions

// IsValidController checks if controller address is valid (non-empty)
func IsValidController(controller string) bool {
	return IsNonEmptyString(controller)
}

// IsValidAccessToken checks if access token is valid (non-empty)
func IsValidAccessToken(accessToken string) bool {
	return IsNonEmptyString(accessToken)
}

// IsPositiveTimeout checks if timeout is greater than validation threshold
func IsPositiveTimeout(timeout time.Duration) bool {
	return timeout > ValidationTimeoutThreshold*time.Second
}

// IsValidTimeout returns true if timeout is greater than zero
func IsValidTimeout(timeout time.Duration) bool {
	return timeout > 0
}

// NOTE: IsInsecureSkipVerify function was removed as it provided no additional
// value over direct boolean usage

// ValidateNonEmptyString validates that a string is not empty after trimming whitespace
func ValidateNonEmptyString(s, fieldName string) error {
	if !IsNonEmptyString(s) {
		return fmt.Errorf("%s cannot be empty", fieldName)
	}
	return nil
}

// String validation predicates

// IsNonEmptyString checks if a string is not empty after trimming whitespace
func IsNonEmptyString(s string) bool {
	return strings.TrimSpace(s) != ""
}

// IsStringEmpty returns true if the string is literally empty (not trimmed)
// This is different from !IsNonEmptyString as it doesn't trim whitespace
func IsStringEmpty(str string) bool {
	return str == ""
}

// MAC address validation functions

// ValidateAPMac validates MAC address format
// MAC address can be formatted as aa:bb:cc:dd:ee:ff, aa-bb-cc-dd-ee-ff, or aabbccddeeff
func ValidateAPMac(mac string) error {
	normalized := normalizeMACAddress(mac)

	if len(normalized) != 12 {
		return errors.New("MAC address must be 12 hex characters")
	}

	for _, c := range normalized {
		if !IsHexChar(c) {
			return fmt.Errorf("MAC address contains invalid character: %c", c)
		}
	}

	return nil
}

// NormalizeAPMac normalizes MAC address to colon-separated format (aa:bb:cc:dd:ee:ff)
func NormalizeAPMac(mac string) string {
	normalized := strings.ToLower(normalizeMACAddress(mac))

	var result strings.Builder
	for i := 0; i < len(normalized); i += 2 {
		if i > 0 {
			result.WriteString(":")
		}
		result.WriteString(normalized[i : i+2])
	}

	return result.String()
}

// normalizeMACAddress removes all separators from MAC address
func normalizeMACAddress(mac string) string {
	normalized := strings.ReplaceAll(mac, ":", "")
	normalized = strings.ReplaceAll(normalized, "-", "")
	normalized = strings.ReplaceAll(normalized, ".", "")
	return normalized
}

// IsHexChar checks if character is hexadecimal
func IsHexChar(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// IsValidAPMacFormat performs MAC address validation and returns boolean
func IsValidAPMacFormat(mac string) bool {
	return ValidateAPMac(mac) == nil
}

// Specific entity validation predicates

// IsValidSiteTag checks if site tag is valid
func IsValidSiteTag(siteTag string) bool {
	return IsNonEmptyString(siteTag)
}

// IsValidPolicyTag checks if policy tag is valid
func IsValidPolicyTag(policyTag string) bool {
	return IsNonEmptyString(policyTag)
}

// IsValidRFTag checks if RF tag is valid
func IsValidRFTag(rfTag string) bool {
	return IsNonEmptyString(rfTag)
}

// ValidateRFIDMAC validates RFID MAC address format and emptiness
func ValidateRFIDMAC(macAddr string) error {
	return ValidateAPMac(macAddr)
}

// ValidateBothMACs validates both RFID and AP MAC addresses
func ValidateBothMACs(rfidMAC, apMAC string) error {
	if err := ValidateRFIDMAC(rfidMAC); err != nil {
		return err
	}
	return ValidateAPMac(apMAC)
}

// ValidateSlotID validates slot ID (radio slot, antenna slot, etc.)
func ValidateSlotID(slotID int) error {
	if slotID < 0 {
		return fmt.Errorf("slot ID must be non-negative, got %d", slotID)
	}
	return nil
}

// ValidateSpatialStream validates the spatial stream parameter (must be 1-8)
func ValidateSpatialStream(spatialStream int) error {
	if spatialStream < 1 || spatialStream > 8 {
		return fmt.Errorf("spatial stream must be between 1 and 8, got %d", spatialStream)
	}
	return nil
}

// ValidateWlanID validates WLAN ID format
func ValidateWlanID(wlanID string) error {
	return ValidateNonEmptyString(wlanID, "WLAN ID")
}

// ValidateRogueAddress validates rogue device address
func ValidateRogueAddress(address string) error {
	return ValidateNonEmptyString(address, "rogue address")
}

// ValidateRogueClientAddress validates rogue client address
func ValidateRogueClientAddress(address string) error {
	return ValidateNonEmptyString(address, "rogue client address")
}

// ValidateClassType validates class type parameter
func ValidateClassType(classType string) error {
	return ValidateNonEmptyString(classType, "class type")
}

// ValidateContainmentLevel validates containment level parameter
func ValidateContainmentLevel(level int) error {
	if level < 0 {
		return fmt.Errorf("containment level must be non-negative, got %d", level)
	}
	return nil
}

// ValidateProfileName validates a profile name (RF profile, radio profile, etc.)
func ValidateProfileName(name, profileType string) error {
	return ValidateNonEmptyString(name, profileType+" profile name")
}

// =============================================================================
// Helper Functions for Value Checking
// =============================================================================

// IsNil returns true if the value is nil
func IsNil(value interface{}) bool { // interface{} needed for generic nil checking of any type
	return value == nil
}

// HasError returns true if an error occurred
func HasError(err error) bool {
	return err != nil
}

// IsValid returns true if the response is not nil and contains valid data
func IsValid(resp interface{}, err error) bool { // interface{} needed for generic validation of any response type
	return !HasError(err) && !IsNil(resp)
}

// FieldMatches returns true if the field value matches the expected value
func FieldMatches(fieldValue, expectedValue string) bool {
	return fieldValue == expectedValue
}

// SelectNonEmptyValue returns the primary value if not empty, otherwise returns the default value

// Composite validation functions

// HasValidTags checks if at least one tag is provided
func HasValidTags(siteTag, policyTag, rfTag string) bool {
	return IsValidSiteTag(siteTag) || IsValidPolicyTag(policyTag) || IsValidRFTag(rfTag)
}

// HasValidMACOrName checks if either MAC address or AP name is provided (but not both)
func HasValidMACOrName(apMac, apName string) bool {
	return (IsNonEmptyString(apMac) && IsStringEmpty(apName)) ||
		(IsStringEmpty(apMac) && IsNonEmptyString(apName))
}

// HasEitherMACOrName checks if at least one of MAC address or AP name is provided
func HasEitherMACOrName(apMac, apName string) bool {
	return IsNonEmptyString(apMac) || IsNonEmptyString(apName)
}

// General utility predicates

// SelectNonEmptyValue returns the primary value if not empty, otherwise returns the default value
func SelectNonEmptyValue(primary, defaultValue string) string {
	if IsNonEmptyString(primary) {
		return primary
	}
	return defaultValue
}

// IsValidTagAssignment validates tag value based on tag type
func IsValidTagAssignment(tagValue, tagType string) bool {
	switch tagType {
	case "site":
		return IsValidSiteTag(tagValue)
	case "policy":
		return IsValidPolicyTag(tagValue)
	case "rf":
		return IsValidRFTag(tagValue)
	default:
		return false
	}
}

// =============================================================================
// Helper Functions for Value Checking
// =============================================================================

// IsNilOrEmpty checks if a value is nil or zero/empty using reflection
func IsNilOrEmpty[T any](v T) bool {
	val := reflect.ValueOf(v)

	// Handle nil interface case
	if !val.IsValid() {
		return true
	}

	// Handle different kinds
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return val.IsNil()
	case reflect.Array:
		return val.Len() == 0
	case reflect.String:
		return val.String() == ""
	default:
		// For basic types, use zero value comparison
		return val.IsZero()
	}
}
