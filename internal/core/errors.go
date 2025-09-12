package core

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// HTTP status code constants.
const (
	// Success status codes.
	StatusOK = http.StatusOK

	// Client error status codes.
	StatusBadRequest          = http.StatusBadRequest
	StatusUnauthorized        = http.StatusUnauthorized
	StatusForbidden           = http.StatusForbidden
	StatusNotFound            = http.StatusNotFound
	StatusMethodNotAllowed    = http.StatusMethodNotAllowed
	StatusConflict            = http.StatusConflict
	StatusUnprocessableEntity = http.StatusUnprocessableEntity

	// Server error status codes.
	StatusInternalServerError = http.StatusInternalServerError
	StatusBadGateway          = http.StatusBadGateway
	StatusServiceUnavailable  = http.StatusServiceUnavailable
	StatusGatewayTimeout      = http.StatusGatewayTimeout
)

// Custom error types for better error handling and debugging.
var (
	// ErrAuthenticationFailed indicates that authentication with the WNC failed due to invalid credentials.
	ErrAuthenticationFailed = errors.New("authentication failed: invalid credentials")
	// ErrAccessForbidden indicates that the client lacks sufficient permissions for the requested operation.
	ErrAccessForbidden = errors.New("access forbidden: insufficient permissions")
	// ErrResourceNotFound indicates that the requested resource or endpoint was not found.
	ErrResourceNotFound = errors.New("resource not found")
	// ErrInvalidConfiguration indicates that the client configuration is invalid or incomplete.
	ErrInvalidConfiguration = errors.New("invalid client configuration")
	// ErrRequestTimeout indicates that the request exceeded the configured timeout period.
	ErrRequestTimeout = errors.New("request timeout")
)

// HTTPError represents an HTTP error response from the API.
type HTTPError struct {
	Status int
	Body   []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Status, string(e.Body))
}

// APIError represents an API-specific error with HTTP status code and message.
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Body       []byte `json:"-"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (HTTP %d): %s", e.StatusCode, e.Message)
}

// IsNotFoundError checks if the error is a 404 not found error.
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return httpErr.Status == http.StatusNotFound
	}

	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == http.StatusNotFound
	}

	errStr := err.Error()
	return strings.Contains(errStr, "404") ||
		strings.Contains(errStr, "not found") ||
		strings.Contains(errStr, "Not Found")
}
