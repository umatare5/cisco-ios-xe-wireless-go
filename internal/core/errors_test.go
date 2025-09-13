package core

import (
	"errors"
	"net/http"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// ========================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// ========================================

// TestAPIErrorStructure tests the basic structure of APIError.
func TestAPIErrorStructure(t *testing.T) {
	apiErr := &APIError{
		StatusCode: 404,
		Message:    "Resource not found",
		Body:       []byte("detailed error body"),
	}

	expectedError := "API error (HTTP 404): Resource not found"
	testutil.AssertStringEquals(t, apiErr.Error(), expectedError, "APIError.Error() should return expected format")

	testutil.AssertIntEquals(t, apiErr.StatusCode, http.StatusNotFound, "APIError.StatusCode should be 404")

	testutil.AssertStringEquals(t, apiErr.Message, "Resource not found", "APIError.Message should match expected value")
}

// TestErrorConstants tests predefined error constants.
func TestErrorConstants(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected string
	}{
		{"AuthenticationFailed", ErrAuthenticationFailed, "authentication failed: invalid credentials"},
		{"AccessForbidden", ErrAccessForbidden, "access forbidden: insufficient permissions"},
		{"ResourceNotFound", ErrResourceNotFound, "resource not found"},
		{"InvalidConfiguration", ErrInvalidConfiguration, "invalid client configuration"},
		{"RequestTimeout", ErrRequestTimeout, "request timeout"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AssertStringEquals(t, tt.err.Error(), tt.expected, "Error constant should return expected message")
		})
	}
}

// TestAPIErrorEdgeCases tests edge cases for APIError.
func TestAPIErrorEdgeCases(t *testing.T) {
	// Test with empty message
	apiErr := &APIError{
		StatusCode: 500,
		Message:    "",
		Body:       nil,
	}

	expectedError := "API error (HTTP 500): "
	testutil.AssertStringEquals(t, apiErr.Error(), expectedError,
		"APIError with empty message should return expected format")

	// Test with zero status code
	apiErr = &APIError{
		StatusCode: 0,
		Message:    "Zero status code",
		Body:       []byte{},
	}

	expectedError = "API error (HTTP 0): Zero status code"
	testutil.AssertStringEquals(t, apiErr.Error(), expectedError,
		"APIError with zero status code should return expected format")
}

// TestHTTPStatusConstants tests HTTP status code constants.
func TestHTTPStatusConstants(t *testing.T) {
	testCases := []struct {
		name     string
		constant int
		expected int
	}{
		{"StatusOK", StatusOK, 200},
		{"StatusBadRequest", StatusBadRequest, 400},
		{"StatusUnauthorized", StatusUnauthorized, 401},
		{"StatusForbidden", StatusForbidden, 403},
		{"StatusNotFound", StatusNotFound, 404},
		{"StatusMethodNotAllowed", StatusMethodNotAllowed, 405},
		{"StatusConflict", StatusConflict, 409},
		{"StatusUnprocessableEntity", StatusUnprocessableEntity, 422},
		{"StatusInternalServerError", StatusInternalServerError, 500},
		{"StatusBadGateway", StatusBadGateway, 502},
		{"StatusServiceUnavailable", StatusServiceUnavailable, 503},
		{"StatusGatewayTimeout", StatusGatewayTimeout, 504},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AssertIntEquals(t, tt.constant, tt.expected, "HTTP status constant should match expected value")
		})
	}
}

// TestIsNotFoundError tests the IsNotFoundError function.
func TestIsNotFoundError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil_error",
			err:      nil,
			expected: false,
		},
		{
			name: "http_error_404",
			err: &HTTPError{
				Status: http.StatusNotFound,
				Body:   []byte("Not Found"),
			},
			expected: true,
		},
		{
			name: "http_error_500",
			err: &HTTPError{
				Status: http.StatusInternalServerError,
				Body:   []byte("Internal Server Error"),
			},
			expected: false,
		},
		{
			name: "api_error_404",
			err: &APIError{
				StatusCode: http.StatusNotFound,
				Message:    "Resource not found",
			},
			expected: true,
		},
		{
			name: "api_error_403",
			err: &APIError{
				StatusCode: http.StatusForbidden,
				Message:    "Access forbidden",
			},
			expected: false,
		},
		{
			name:     "string_error_with_404",
			err:      errors.New("API error (HTTP 404): resource not found"),
			expected: true,
		},
		{
			name:     "string_error_with_not_found",
			err:      errors.New("resource not found"),
			expected: true,
		},
		{
			name:     "string_error_with_not_found_uppercase",
			err:      errors.New("API error: Not Found"),
			expected: true,
		},
		{
			name:     "string_error_without_404",
			err:      errors.New("internal server error"),
			expected: false,
		},
		{
			name:     "generic_error",
			err:      errors.New("some other error"),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsNotFoundError(tc.err)
			testutil.AssertBoolEquals(t, result, tc.expected, "IsNotFoundError should return expected boolean value")
		})
	}
}
