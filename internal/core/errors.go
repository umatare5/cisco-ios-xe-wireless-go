package core

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// Custom error types for better error handling and debugging
var (
	// ErrAuthenticationFailed indicates that authentication with the WNC failed due to invalid credentials
	ErrAuthenticationFailed = errors.New("authentication failed: invalid credentials")
	// ErrAccessForbidden indicates that the client lacks sufficient permissions for the requested operation
	ErrAccessForbidden = errors.New("access forbidden: insufficient permissions")
	// ErrResourceNotFound indicates that the requested resource or endpoint was not found
	ErrResourceNotFound = errors.New("resource not found")
	// ErrInvalidConfiguration indicates that the client configuration is invalid or incomplete
	ErrInvalidConfiguration = errors.New("invalid client configuration")
	// ErrRequestTimeout indicates that the request exceeded the configured timeout period
	ErrRequestTimeout = errors.New("request timeout")
)

// HTTPError represents an HTTP error response from the API
type HTTPError struct {
	Status int    // HTTP status code
	Body   []byte // Response body
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Status, string(e.Body))
}

// HTTP status code constants
const (
	// Success status codes
	StatusOK        = http.StatusOK
	StatusCreated   = http.StatusCreated
	StatusAccepted  = http.StatusAccepted
	StatusNoContent = http.StatusNoContent

	// Client error status codes
	StatusBadRequest          = http.StatusBadRequest
	StatusUnauthorized        = http.StatusUnauthorized
	StatusForbidden           = http.StatusForbidden
	StatusNotFound            = http.StatusNotFound
	StatusMethodNotAllowed    = http.StatusMethodNotAllowed
	StatusConflict            = http.StatusConflict
	StatusUnprocessableEntity = http.StatusUnprocessableEntity

	// Server error status codes
	StatusInternalServerError = http.StatusInternalServerError
	StatusBadGateway          = http.StatusBadGateway
	StatusServiceUnavailable  = http.StatusServiceUnavailable
	StatusGatewayTimeout      = http.StatusGatewayTimeout
)

// APIError represents an API-specific error with HTTP status code and message
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Body       []byte `json:"-"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (HTTP %d): %s", e.StatusCode, e.Message)
}

// isSuccessStatusCode checks if HTTP status code indicates success
func isSuccessStatusCode(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

// isAuthenticationError checks if HTTP status code indicates authentication failure
func isAuthenticationError(statusCode int) bool {
	return statusCode == StatusUnauthorized
}

// isAccessForbiddenError checks if HTTP status code indicates access forbidden
func isAccessForbiddenError(statusCode int) bool {
	return statusCode == StatusForbidden
}

// isNotFoundError checks if HTTP status code indicates resource not found
func isNotFoundError(statusCode int) bool {
	return statusCode == StatusNotFound
}

// isDeadlineExceededError checks if error is due to context deadline exceeded
func isDeadlineExceededError(err error) bool {
	return errors.Is(err, context.DeadlineExceeded)
}
