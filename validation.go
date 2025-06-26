package wnc

import "time"

// Validation constants
const (
	// MinEndpointLengthChars is the minimum character length for API endpoints
	MinEndpointLengthChars = 10

	// MinTokenLengthChars is the minimum character length for authentication tokens
	MinTokenLengthChars = 8

	// MinEndpointLength is the minimum length for API endpoints
	MinEndpointLength = MinEndpointLengthChars

	// MinTokenLength is the minimum length for authentication tokens
	MinTokenLength = MinTokenLengthChars

	// ZeroTimeoutSeconds represents zero timeout for validation tests
	ZeroTimeoutSeconds = 0

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

// isValidController checks if controller address is valid
func isValidController(controller string) bool {
	return controller != ""
}

// isValidAccessToken checks if access token is valid
func isValidAccessToken(accessToken string) bool {
	return accessToken != ""
}

// isPositiveTimeout checks if timeout is greater than validation threshold
func isPositiveTimeout(timeout time.Duration) bool {
	return timeout > ValidationTimeoutThreshold*time.Second
}
