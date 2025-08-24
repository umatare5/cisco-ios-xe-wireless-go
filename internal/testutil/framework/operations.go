package framework

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

// =============================================================================
// SET OPERATIONS (Create/Update/Delete) FRAMEWORK
// =============================================================================

// SetOperationTestCase defines a test case for set operations (Create/Update/Delete)
type SetOperationTestCase struct {
	Name          string
	SetupContext  func() context.Context
	TagName       string
	ExpectedError string
}

// SetOperationTestPattern defines the pattern for testing set operations
type SetOperationTestPattern struct {
	ServiceName    string
	TestCases      []SetOperationTestCase
	EndpointFormat string // e.g., "Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=%s"
	Operation      func(t *testing.T, ctx context.Context, server *httptest.Server, tagName string) error
}

// RunStandardSetOperTests executes standardized set operation tests
func RunStandardSetOperTests(t *testing.T, pattern SetOperationTestPattern) {
	t.Helper()

	t.Run(pattern.ServiceName+"_SetOperations", func(t *testing.T) {
		for _, tc := range pattern.TestCases {
			t.Run(tc.Name, func(t *testing.T) {
				var server *httptest.Server

				// Setup server based on test case
				switch {
				case isSuccessTest(tc.Name):
					// Success case - empty response for set operations
					var endpoints map[string]string
					if strings.Contains(pattern.EndpointFormat, "%s") {
						endpoints = map[string]string{
							fmt.Sprintf(pattern.EndpointFormat, tc.TagName): ``,
						}
					} else {
						endpoints = map[string]string{
							pattern.EndpointFormat: ``,
						}
					}
					server = mock.NewRESTCONFSuccessServer(endpoints)
				case isContextCanceledTest(tc.Name):
					// For context cancellation test
					var endpoints map[string]string
					if strings.Contains(pattern.EndpointFormat, "%s") {
						endpoints = map[string]string{
							fmt.Sprintf(pattern.EndpointFormat, tc.TagName): `{"delayed": "response"}`,
						}
					} else {
						endpoints = map[string]string{
							pattern.EndpointFormat: `{"delayed": "response"}`,
						}
					}
					server = mock.NewRESTCONFSuccessServer(endpoints)
				default:
					// For error status tests - extract status code from test name
					statusCode := getStatusCodeFromTestName(tc.Name)
					var paths []string
					if strings.Contains(pattern.EndpointFormat, "%s") {
						paths = []string{fmt.Sprintf(pattern.EndpointFormat, tc.TagName)}
					} else {
						paths = []string{pattern.EndpointFormat}
					}
					server = mock.NewRESTCONFErrorServer(paths, statusCode)
				}
				defer server.Close()

				// Setup context
				ctx := tc.SetupContext()

				// Execute operation
				err := pattern.Operation(t, ctx, server, tc.TagName)

				// Validate results using early returns
				if !isSuccessTest(tc.Name) {
					if isNilError(err) {
						t.Errorf("Expected error containing '%s' but got none", tc.ExpectedError)
						return
					}
					if !containsError(err, tc.ExpectedError) {
						t.Errorf(
							"Expected error containing '%s' but got: %v",
							tc.ExpectedError,
							err,
						)
					}
					return
				}

				if !isNilError(err) {
					t.Errorf("Unexpected error: %v", err)
				}
			})
		}
	})
}

// =============================================================================
// LIST OPERATIONS (Collection Retrieval) FRAMEWORK
// =============================================================================

// ListOperationTestCase defines a test case for LIST operations (collection retrieval)
type ListOperationTestCase struct {
	Name          string
	SetupContext  func() context.Context
	ExpectedError string // Expected error substring, empty for success
	ExpectResult  bool   // Whether to expect non-nil result
	ResponseData  string // JSON response data for success cases
}

// ListOperationTestPattern defines the pattern for testing LIST operations
type ListOperationTestPattern struct {
	ServiceName     string
	TestCases       []ListOperationTestCase
	Endpoint        string // e.g., "path/to/collection"
	Operation       func(t *testing.T, ctx context.Context, server *httptest.Server) (interface{}, error)
	DefaultResponse string // Default success response if ResponseData not provided
}

// RunStandardListOperTests executes standardized LIST operation tests
func RunStandardListOperTests(t *testing.T, pattern ListOperationTestPattern) {
	t.Helper()

	t.Run(pattern.ServiceName+"_ListOperations", func(t *testing.T) {
		for _, tc := range pattern.TestCases {
			t.Run(tc.Name, func(t *testing.T) {
				var server *httptest.Server

				// Setup server based on test case
				switch {
				case isSuccessTest(tc.Name):
					// Success case - return collection data
					responseData := tc.ResponseData
					if responseData == "" {
						responseData = pattern.DefaultResponse
					}
					endpoints := map[string]string{
						pattern.Endpoint: responseData,
					}
					server = mock.NewRESTCONFSuccessServer(endpoints)
				case isContextCanceledTest(tc.Name):
					// For context cancellation test
					endpoints := map[string]string{
						pattern.Endpoint: `{"delayed": "response"}`,
					}
					server = mock.NewRESTCONFSuccessServer(endpoints)
				default:
					// For error status tests - extract status code from test name
					statusCode := getStatusCodeFromTestName(tc.Name)
					paths := []string{pattern.Endpoint}
					server = mock.NewRESTCONFErrorServer(paths, statusCode)
				}
				defer server.Close()

				// Setup context
				ctx := tc.SetupContext()

				// Execute operation
				result, err := pattern.Operation(t, ctx, server)

				// Validate results using early returns
				if tc.ExpectedError != "" {
					if isNilError(err) {
						t.Errorf("Expected error containing '%s' but got none", tc.ExpectedError)
						return
					}
					if !containsError(err, tc.ExpectedError) {
						t.Errorf(
							"Expected error containing '%s' but got: %v",
							tc.ExpectedError,
							err,
						)
						return
					}
					return
				}

				if !isNilError(err) {
					t.Errorf("Unexpected error: %v", err)
					return
				}

				// Check result for success cases
				if tc.ExpectResult && isNilResult(result) {
					t.Errorf("Expected non-nil result for successful case")
				}
			})
		}
	})
}

// =============================================================================
// HELPER FUNCTIONS
// =============================================================================

// CreateTimeoutContext creates a context that will timeout immediately for testing context cancellation
func CreateTimeoutContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	time.Sleep(5 * time.Millisecond) // Force timeout
	return ctx
}

// Predicate functions for test classification
func isSuccessTest(testName string) bool {
	return strings.Contains(testName, "Success")
}

func isContextCanceledTest(testName string) bool {
	return strings.Contains(testName, "ContextCanceled")
}

func isUnauthorizedTest(testName string) bool {
	return strings.Contains(testName, "401") || strings.Contains(testName, "Unauthorized")
}

func isForbiddenTest(testName string) bool {
	return strings.Contains(testName, "403") || strings.Contains(testName, "Forbidden")
}

func isNotFoundTest(testName string) bool {
	return strings.Contains(testName, "404") || strings.Contains(testName, "NotFound")
}

func isInternalServerErrorTest(testName string) bool {
	return strings.Contains(testName, "500") || strings.Contains(testName, "InternalServerError")
}

// Error checking helpers
func containsError(err error, expectedError string) bool {
	if err == nil {
		return false
	}
	errorMessage := err.Error()

	// Handle context cancellation variations
	if expectedError == "context deadline exceeded" {
		return strings.Contains(errorMessage, "context deadline exceeded") ||
			strings.Contains(errorMessage, "context canceled")
	}

	return strings.Contains(errorMessage, expectedError)
}

func isNilError(err error) bool {
	return err == nil
}

func isNilResult(result interface{}) bool {
	return result == nil
}

// HTTP status code extraction from test names
func getStatusCodeFromTestName(testName string) int {
	switch {
	case strings.Contains(strings.ToLower(testName), "400") || strings.Contains(strings.ToLower(testName), "client error") || strings.Contains(strings.ToLower(testName), "bad request"):
		return http.StatusBadRequest
	case isUnauthorizedTest(testName):
		return http.StatusUnauthorized
	case isForbiddenTest(testName):
		return http.StatusForbidden
	case isNotFoundTest(testName):
		return http.StatusNotFound
	case strings.Contains(strings.ToLower(testName), "409") || strings.Contains(strings.ToLower(testName), "conflict"):
		return http.StatusConflict
	case isInternalServerErrorTest(testName):
		return http.StatusInternalServerError
	case strings.Contains(strings.ToLower(testName), "502") || strings.Contains(strings.ToLower(testName), "bad gateway"):
		return http.StatusBadGateway
	case strings.Contains(strings.ToLower(testName), "503") || strings.Contains(strings.ToLower(testName), "service unavailable"):
		return http.StatusServiceUnavailable
	default:
		return http.StatusNotFound
	}
}
