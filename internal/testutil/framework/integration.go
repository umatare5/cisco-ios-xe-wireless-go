package framework

import (
	"context"
	"strings"
	"testing"
	"time"
)

// =============================================================================
// INTEGRATION TESTS FRAMEWORK
// =============================================================================

// IntegrationTestCase defines a test case for integration tests
type IntegrationTestCase struct {
	Name            string
	Description     string
	SkipOnShortMode bool
	ExpectError     bool
	ExpectNotFound  bool
	LogResultLength bool
	UseTimeout      bool
}

// IntegrationTestPattern defines the pattern for testing integration operations
type IntegrationTestPattern struct {
	ServiceName string
	TestCases   []IntegrationTestCase
	Operation   func(t *testing.T, setup ServiceSetup, tc IntegrationTestCase) (interface{}, error)
}

// ServiceSetup represents test setup with client and context
type ServiceSetup struct {
	Client  interface{}
	Context context.Context
}

// TestOperationTimeout represents the default timeout for test operations
const TestOperationTimeout = 30 * time.Second

// SetupRequiredClient sets up a client for integration tests
func SetupRequiredClient(t *testing.T) ServiceSetup {
	// This function should be implemented to return a proper ServiceSetup
	// For now, return a basic setup
	return ServiceSetup{
		Client:  nil,
		Context: context.Background(),
	}
}

// RunStandardIntegrationTests executes standardized integration operation tests
func RunStandardIntegrationTests(t *testing.T, pattern IntegrationTestPattern) {
	t.Helper()

	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	setup := SetupRequiredClient(t)

	for _, tc := range pattern.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.SkipOnShortMode && testing.Short() {
				t.Skip("Skipping test in short mode")
			}

			// Setup context with timeout if requested
			ctx := setup.Context
			if tc.UseTimeout {
				timeoutCtx, cancel := context.WithTimeout(setup.Context, TestOperationTimeout)
				defer cancel()
				ctx = timeoutCtx
			}

			// Execute operation with proper context
			tempSetup := ServiceSetup{
				Client:  setup.Client,
				Context: ctx,
			}

			result, err := pattern.Operation(t, tempSetup, tc)

			// Validate results using early returns
			if tc.ExpectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				t.Logf("Got expected error: %v", err)
				return
			}

			if tc.ExpectNotFound && err != nil {
				// Import core for IsNotFoundError check if needed
				if strings.Contains(err.Error(), "404") ||
					strings.Contains(err.Error(), "not found") {
					t.Logf("Endpoint not supported (404): %v", err)
					return
				}
				t.Errorf("Expected 404 error but got: %v", err)
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Errorf("Expected non-nil result")
				return
			}

			// Log result length if requested (for collections)
			if tc.LogResultLength {
				switch resultSlice := result.(type) {
				case []interface{}:
					t.Logf("Retrieved %d items", len(resultSlice))
				case interface{ Len() int }:
					t.Logf("Retrieved %d items", resultSlice.Len())
				default:
					t.Logf("Retrieved result: %T", result)
				}
			}
		})
	}
}
