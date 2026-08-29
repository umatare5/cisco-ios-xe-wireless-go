package core

import (
	"errors"
	"fmt"
	"net/http"
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
	// Every error New returns wraps it: a malformed authority, an empty token, or an option that
	// refused its argument.
	ErrInvalidConfiguration = errors.New("invalid client configuration")
	// ErrRequestTimeout indicates that the request exceeded the configured timeout period.
	ErrRequestTimeout = errors.New("request timeout")
)

// APIError represents an API-specific error with HTTP status code and message.
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Body       []byte `json:"-"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (HTTP %d): %s", e.StatusCode, e.Message)
}

// Unwrap returns the sentinel that describes the status, so a caller can match with
// errors.Is instead of reading StatusCode. A status with no sentinel unwraps to nil,
// which ends the chain.
func (e *APIError) Unwrap() error {
	switch e.StatusCode {
	case http.StatusUnauthorized:
		return ErrAuthenticationFailed
	case http.StatusForbidden:
		return ErrAccessForbidden
	case http.StatusNotFound:
		return ErrResourceNotFound
	default:
		return nil
	}
}

// IsNotFoundError reports whether err carries an HTTP 404 from the controller.
//
// ErrResourceNotFound answers true as well: a guard that rejects an empty list key returns
// it before a request is built, and a caller asking "was this absent" must not have to know
// which of the two paths produced the error.
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrResourceNotFound)
}
