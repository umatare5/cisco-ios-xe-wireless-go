package framework

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

// RestconfOperationTestCase defines a test case for RESTCONF operations using builder pattern
type RestconfOperationTestCase struct {
	Name          string
	SetupContext  func() context.Context
	TagName       string
	ExpectedError string
}

// RestconfSetOperationPattern defines TagSetOper pattern for testing set operations with RestconfBuilder
type RestconfSetOperationPattern struct {
	ServiceName     string
	TestCases       []RestconfOperationTestCase
	BaseEndpoint    string // Base endpoint constant from routes package
	Operation       func(t *testing.T, ctx context.Context, server *httptest.Server, tagName string) error
	BuildEndpointFn func(baseEndpoint, tagName string) string // Function to build parameterized endpoint
}

// RestconfGetOperationPattern defines TagSetOper pattern for testing get operations with RestconfBuilder
type RestconfGetOperationPattern struct {
	ServiceName     string
	TestCases       []RestconfOperationTestCase
	BaseEndpoint    string
	DefaultResponse string
	Operation       func(t *testing.T, ctx context.Context, server *httptest.Server, tagName string) (interface{}, error)
	BuildEndpointFn func(baseEndpoint, tagName string) string
}

// RunTagSetOperTests executes set operation tests using RestconfBuilder pattern
func RunTagSetOperTests(t *testing.T, pattern RestconfSetOperationPattern) {
	t.Helper()

	t.Run(pattern.ServiceName+"_TagSetOperations", func(t *testing.T) {
		for _, tc := range pattern.TestCases {
			t.Run(tc.Name, func(t *testing.T) {
				// Build endpoint using builder function or default pattern
				var endpoint string
				if pattern.BuildEndpointFn != nil {
					endpoint = pattern.BuildEndpointFn(pattern.BaseEndpoint, tc.TagName)
				} else {
					// Default pattern: baseEndpoint + "=" + tagName
					endpoint = pattern.BaseEndpoint + "=" + tc.TagName
				}

				var server *httptest.Server

				// Setup server based on test case
				switch {
				case isSuccessTest(tc.Name):
					// Success case - empty response for set operations
					endpoints := map[string]string{endpoint: ``}
					server = mock.NewRESTCONFSuccessServer(endpoints)
				case isContextCanceledTest(tc.Name):
					// For context cancellation test
					endpoints := map[string]string{endpoint: `{"delayed": "response"}`}
					server = mock.NewRESTCONFSuccessServer(endpoints)
				default:
					// For error status tests - extract status code from test name
					statusCode := getStatusCodeFromTestName(tc.Name)
					paths := []string{endpoint}
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

// RunTagGetOperTests executes get operation tests using RestconfBuilder pattern
func RunTagGetOperTests(t *testing.T, pattern RestconfGetOperationPattern) {
	t.Helper()

	t.Run(pattern.ServiceName+"_TagGetOperations", func(t *testing.T) {
		for _, tc := range pattern.TestCases {
			t.Run(tc.Name, func(t *testing.T) {
				// Build endpoint using builder function or default pattern
				var endpoint string
				if pattern.BuildEndpointFn != nil {
					endpoint = pattern.BuildEndpointFn(pattern.BaseEndpoint, tc.TagName)
				} else {
					// Default pattern: baseEndpoint + "=" + tagName
					endpoint = pattern.BaseEndpoint + "=" + tc.TagName
				}

				var server *httptest.Server

				// Setup server based on test case
				switch {
				case isSuccessTest(tc.Name):
					// Success case - return default response or empty
					responseData := pattern.DefaultResponse
					if responseData == "" {
						responseData = `{"test": "data"}`
					}
					endpoints := map[string]string{endpoint: responseData}
					server = mock.NewRESTCONFSuccessServer(endpoints)
				case isContextCanceledTest(tc.Name):
					// For context cancellation test
					endpoints := map[string]string{endpoint: `{"delayed": "response"}`}
					server = mock.NewRESTCONFSuccessServer(endpoints)
				default:
					// For error status tests - extract status code from test name
					statusCode := getStatusCodeFromTestName(tc.Name)
					paths := []string{endpoint}
					server = mock.NewRESTCONFErrorServer(paths, statusCode)
				}
				defer server.Close()

				// Setup context
				ctx := tc.SetupContext()

				// Execute operation
				result, err := pattern.Operation(t, ctx, server, tc.TagName)

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

				// For success cases, result can be nil or non-nil depending on endpoint
				if result != nil {
					t.Logf("Successfully retrieved result for %s", tc.TagName)
				}
			})
		}
	})
}

// GenerateBaseEndpoint creates a base endpoint from service name and tag type
