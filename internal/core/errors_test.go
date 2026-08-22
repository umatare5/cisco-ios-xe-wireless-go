package core

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// ========================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// ========================================

// TestCoreErrorsUnit_APIErrorUnwrap_Success pins the status-to-sentinel map that
// wnc.go:46-53 documents.
func TestCoreErrorsUnit_APIErrorUnwrap_Success(t *testing.T) {
	mapped := []struct {
		status   int
		sentinel error
	}{
		{http.StatusUnauthorized, ErrAuthenticationFailed},
		{http.StatusForbidden, ErrAccessForbidden},
		{http.StatusNotFound, ErrResourceNotFound},
	}
	for _, tc := range mapped {
		var err error = &APIError{StatusCode: tc.status}
		testutil.AssertTrue(t, errors.Is(err, tc.sentinel), http.StatusText(tc.status))
	}

	var unmapped error = &APIError{StatusCode: http.StatusInternalServerError, Message: "boom"}
	testutil.AssertFalse(t, errors.Is(unmapped, ErrResourceNotFound), "500 matches no sentinel")
	testutil.AssertFalse(t, errors.Is(unmapped, ErrAuthenticationFailed), "500 matches no sentinel")

	var apiErr *APIError
	testutil.AssertTrue(t, errors.As(unmapped, &apiErr), "errors.As still reaches *APIError")
	testutil.AssertIntEquals(t, apiErr.StatusCode, http.StatusInternalServerError, "StatusCode")

	testutil.AssertFalse(t, IsNotFoundError(errors.New("resource not found")),
		"a message is no longer evidence of a 404")

	// Both paths that report absence answer true. A guard rejecting an empty list key
	// returns the sentinel before a request exists, so a caller asking "was this absent"
	// must not have to know which path produced the error.
	testutil.AssertTrue(t, IsNotFoundError(ErrResourceNotFound), "the sentinel itself")
	testutil.AssertTrue(t, IsNotFoundError(fmt.Errorf("read failed: %w", ErrResourceNotFound)),
		"the sentinel wrapped")
	testutil.AssertTrue(t, IsNotFoundError(&APIError{StatusCode: http.StatusNotFound}),
		"a 404 from the controller")
	testutil.AssertFalse(t, IsNotFoundError(&APIError{StatusCode: http.StatusBadRequest}),
		"a 400 is not absence")
}

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
			expected: false,
		},
		{
			name:     "string_error_with_not_found",
			err:      errors.New("resource not found"),
			expected: false,
		},
		{
			name:     "string_error_with_not_found_uppercase",
			err:      errors.New("API error: Not Found"),
			expected: false,
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
