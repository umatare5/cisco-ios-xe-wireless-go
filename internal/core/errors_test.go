package core

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

// ========================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// ========================================

// TestAPIErrorStructure tests the basic structure of APIError
func TestAPIErrorStructure(t *testing.T) {
	apiErr := &APIError{
		StatusCode: 404,
		Message:    "Resource not found",
		Body:       []byte("detailed error body"),
	}

	expectedError := "API error (HTTP 404): Resource not found"
	if apiErr.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, apiErr.Error())
	}

	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %d", apiErr.StatusCode)
	}

	if apiErr.Message != "Resource not found" {
		t.Errorf("Expected message 'Resource not found', got '%s'", apiErr.Message)
	}
}

// ========================================
// 2. TABLE-DRIVEN TEST PATTERNS
// ========================================

// TestStatusCodeCheckers tests HTTP status code validation functions
func TestStatusCodeCheckers(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		checkFunc  func(int) bool
		expected   bool
	}{
		// Success status codes
		{"Success_200", StatusOK, isSuccessStatusCode, true},
		{"Success_201", StatusCreated, isSuccessStatusCode, true},
		{"Success_202", StatusAccepted, isSuccessStatusCode, true},
		{"Success_204", StatusNoContent, isSuccessStatusCode, true},
		{"NotSuccess_404", StatusNotFound, isSuccessStatusCode, false},
		{"NotSuccess_401", StatusUnauthorized, isSuccessStatusCode, false},

		// Authentication error status codes
		{"Auth_401", StatusUnauthorized, isAuthenticationError, true},
		{"NotAuth_200", StatusOK, isAuthenticationError, false},
		{"NotAuth_403", StatusForbidden, isAuthenticationError, false},

		// Access forbidden status codes
		{"Forbidden_403", StatusForbidden, isAccessForbiddenError, true},
		{"NotForbidden_200", StatusOK, isAccessForbiddenError, false},
		{"NotForbidden_401", StatusUnauthorized, isAccessForbiddenError, false},

		// Not found status codes
		{"NotFound_404", StatusNotFound, isNotFoundError, true},
		{"Found_200", StatusOK, isNotFoundError, false},
		{"Found_403", StatusForbidden, isNotFoundError, false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.checkFunc(tt.statusCode)
			if result != tt.expected {
				t.Errorf("Expected %v for status code %d, got %v", tt.expected, tt.statusCode, result)
			}
		})
	}
}

// TestErrorConstants tests predefined error constants
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
			if tt.err.Error() != tt.expected {
				t.Errorf("Expected error message '%s', got '%s'", tt.expected, tt.err.Error())
			}
		})
	}
}

// ========================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// ========================================

// TestDeadlineExceededError tests context deadline exceeded detection
func TestDeadlineExceededError(t *testing.T) {
	// Test with context deadline exceeded error
	deadlineErr := context.DeadlineExceeded
	if !isDeadlineExceededError(deadlineErr) {
		t.Error("Expected context.DeadlineExceeded to be detected as deadline exceeded")
	}

	// Test with regular error
	regularErr := errors.New("regular error")
	if isDeadlineExceededError(regularErr) {
		t.Error("Expected regular error not to be detected as deadline exceeded")
	}

	// Test with nil error
	if isDeadlineExceededError(nil) {
		t.Error("Expected nil error not to be detected as deadline exceeded")
	}
}

// TestAPIErrorEdgeCases tests edge cases for APIError
func TestAPIErrorEdgeCases(t *testing.T) {
	// Test with empty message
	apiErr := &APIError{
		StatusCode: 500,
		Message:    "",
		Body:       nil,
	}

	expectedError := "API error (HTTP 500): "
	if apiErr.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, apiErr.Error())
	}

	// Test with zero status code
	apiErr = &APIError{
		StatusCode: 0,
		Message:    "Zero status code",
		Body:       []byte{},
	}

	expectedError = "API error (HTTP 0): Zero status code"
	if apiErr.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, apiErr.Error())
	}
}

// TestHTTPStatusConstants tests HTTP status code constants
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
			if tt.constant != tt.expected {
				t.Errorf("Expected %s to be %d, got %d", tt.name, tt.expected, tt.constant)
			}
		})
	}
}
