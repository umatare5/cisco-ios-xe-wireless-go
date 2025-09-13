package validation

import (
	"fmt"
	"strings"
	"time"
)

const (
	// MACAddressHexCharLength is the required length for MAC address hex characters.
	MACAddressHexCharLength = 12
)

// Default tag values for AP operations.
const (
	// DefaultSiteTag represents the default site tag value.
	DefaultSiteTag = "default-site-tag"

	// DefaultPolicyTag represents the default policy tag value.
	DefaultPolicyTag = "default-policy-tag"

	// DefaultRFTag represents the default RF tag value.
	DefaultRFTag = "default-rf-tag"
)

// Validation constants.
const (
	// MinEndpointLength is the minimum length for API endpoints.
	MinEndpointLength = 10

	// MinTokenLength is the minimum length for authentication tokens.
	MinTokenLength = 8

	// ValidationTimeoutThreshold is the minimum timeout for validation.
	ValidationTimeoutThreshold = 1
)

// Error message templates for validation errors.
const (
	// EndpointMismatchErrorTemplate is used for endpoint validation errors.
	EndpointMismatchErrorTemplate = "Expected %s = %s, got %s"

	// EmptyEndpointErrorTemplate is used when an endpoint is empty.
	EmptyEndpointErrorTemplate = "%s endpoint is empty"

	// ShortEndpointErrorTemplate is used when an endpoint is too short.
	ShortEndpointErrorTemplate = "%s endpoint is too short: %s"

	// InvalidEndpointErrorTemplate is used for invalid endpoint formats.
	InvalidEndpointErrorTemplate = "%s endpoint has invalid format: %s"
)

// Core validation functions

// IsValidController checks if controller address is valid (non-empty).
func IsValidController(controller string) bool {
	return IsNonEmptyString(controller)
}

// IsValidAccessToken checks if access token is valid (non-empty).
func IsValidAccessToken(accessToken string) bool {
	return IsNonEmptyString(accessToken)
}

// IsPositiveTimeout checks if timeout is greater than validation threshold.
func IsPositiveTimeout(timeout time.Duration) bool {
	return timeout > ValidationTimeoutThreshold*time.Second
}

// IsValidTimeout returns true if timeout is greater than zero.
func IsValidTimeout(timeout time.Duration) bool {
	return timeout > 0
}

// ValidateNonEmptyString validates that a string is not empty after trimming whitespace.
func ValidateNonEmptyString(s, fieldName string) error {
	if !IsNonEmptyString(s) {
		return fmt.Errorf("validation failed: %w",
			fmt.Errorf("field %s cannot be empty or contain only whitespace", fieldName))
	}
	return nil
}

// String validation predicates

// IsNonEmptyString checks if a string is not empty after trimming whitespace.
func IsNonEmptyString(s string) bool {
	return strings.TrimSpace(s) != ""
}

// IsStringEmpty returns true if the string is literally empty (not trimmed)
// This is different from !IsNonEmptyString as it doesn't trim whitespace.
func IsStringEmpty(str string) bool {
	return str == ""
}

// MAC address validation functions

// ValidateMACAddress validates MAC address format.
// MAC address can be formatted as aa:bb:cc:dd:ee:ff, aa-bb-cc-dd-ee-ff, or aabbccddeeff.
func ValidateMACAddress(mac string) error {
	normalized := normalizeMACAddress(mac)

	if len(normalized) != MACAddressHexCharLength {
		return fmt.Errorf("MAC address must be %d hex characters, got %d", MACAddressHexCharLength, len(normalized))
	}

	for _, c := range normalized {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			return fmt.Errorf("invalid hexadecimal character '%c' in MAC address %s", c, mac)
		}
	}

	return nil
}

// NormalizeMACAddress validates and normalizes MAC address to colon-separated format (aa:bb:cc:dd:ee:ff).
// Returns normalized MAC address and validation error if invalid.
func NormalizeMACAddress(mac string) (string, error) {
	if err := ValidateMACAddress(mac); err != nil {
		return "", err
	}

	normalized := strings.ToLower(normalizeMACAddress(mac))
	var result strings.Builder
	for i := 0; i < len(normalized); i += 2 {
		if i > 0 {
			result.WriteString(":")
		}
		result.WriteString(normalized[i : i+2])
	}

	return result.String(), nil
}

// normalizeMACAddress removes all separators from MAC address.
func normalizeMACAddress(mac string) string {
	normalized := strings.ReplaceAll(mac, ":", "")
	normalized = strings.ReplaceAll(normalized, "-", "")
	normalized = strings.ReplaceAll(normalized, ".", "")
	return normalized
}

// IsValidMACAddr performs MAC address validation and returns boolean.
func IsValidMACAddr(mac string) bool {
	return ValidateMACAddress(mac) == nil
}

// ValidateSlotID validates slot ID (radio slot, antenna slot, etc.)
func ValidateSlotID(slotID int) error {
	if slotID < 0 {
		return fmt.Errorf("slot ID must be non-negative, got %d", slotID)
	}
	return nil
}

// ValidateSpatialStream validates the spatial stream parameter (must be 1-8).
func ValidateSpatialStream(spatialStream int) error {
	if spatialStream < 1 || spatialStream > 8 {
		return fmt.Errorf("spatial stream must be between 1 and 8, got %d", spatialStream)
	}
	return nil
}

// ValidateWlanID validates WLAN ID format.
func ValidateWlanID(wlanID string) error {
	return ValidateNonEmptyString(wlanID, "WLAN ID")
}

// HasValidTags checks if at least one tag is provided.
func HasValidTags(siteTag, policyTag, rfTag string) bool {
	return IsNonEmptyString(siteTag) || IsNonEmptyString(policyTag) || IsNonEmptyString(rfTag)
}

// HasValidMACOrName checks if either MAC address or AP name is provided (but not both).
func HasValidMACOrName(apMac, apName string) bool {
	return (IsNonEmptyString(apMac) && IsStringEmpty(apName)) ||
		(IsStringEmpty(apMac) && IsNonEmptyString(apName))
}

// HasEitherMACOrName checks if at least one of MAC address or AP name is provided.
func HasEitherMACOrName(apMac, apName string) bool {
	return IsNonEmptyString(apMac) || IsNonEmptyString(apName)
}

// General utility predicates

// SelectNonEmptyValue returns the primary value if not empty, otherwise returns the default value.
func SelectNonEmptyValue(primary, defaultValue string) string {
	if IsNonEmptyString(primary) {
		return primary
	}
	return defaultValue
}

// IsValidTagAssignment validates tag value based on tag type.
func IsValidTagAssignment(tagValue, tagType string) bool {
	switch tagType {
	case "site":
		return IsNonEmptyString(tagValue)
	case "policy":
		return IsNonEmptyString(tagValue)
	case "rf":
		return IsNonEmptyString(tagValue)
	default:
		return false
	}
}
