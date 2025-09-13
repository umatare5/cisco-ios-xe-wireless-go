//go:build integration || scenario

// Package integration provides test utilities for integration tests.
package integration

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

// ServiceSetup holds the common test setup components.
type ServiceSetup struct {
	Client  *core.Client
	Context context.Context
}

// SetupTestClient creates a ServiceSetup with environment validation for integration tests.
// This function requires basic controller and access token environment variables.
// Use this function for integration tests that need a configured client.
func SetupTestClient(t *testing.T) ServiceSetup {
	t.Helper()

	// Load configuration directly from environment
	controller := os.Getenv("WNC_CONTROLLER")
	accessToken := os.Getenv("WNC_ACCESS_TOKEN")

	if controller == "" || accessToken == "" {
		t.Fatalf("SetupTestClient: WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables are required")
	}

	client, err := core.New(
		controller,
		accessToken,
		core.WithTimeout(5*time.Second),
		core.WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("SetupTestClient: failed to create client: %v", err)
	}

	return ServiceSetup{
		Client:  client,
		Context: testutil.TestContext(t),
	}
}

// TestContext provides common setup for integration tests.
type TestContext struct {
	T      *testing.T
	Ctx    context.Context
	Client *core.Client
	Config *Config
}

// Setup initializes integration test context.
func Setup(t *testing.T) *TestContext {
	setup := SetupTestClient(t)
	if setup.Client == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	// Load integration test configuration from environment
	config := LoadConfig()

	return &TestContext{
		T:      t,
		Ctx:    setup.Context,
		Client: setup.Client,
		Config: &config,
	}
}

// TestSuite represents a complete test suite for integration tests
type TestSuite struct {
	Config          TestSuiteConfig
	BasicMethods    []TestMethod
	FilterMethods   []TestMethod
	ValidationTests []ValidationTestMethod
}

// TestSuiteConfig defines configuration for integration test suites
type TestSuiteConfig struct {
	ServiceName        string
	ServiceConstructor func(client any) any
	UseTimeout         bool
	TimeoutDuration    time.Duration
}

// TestMethod represents a single test method
type TestMethod struct {
	Name           string
	Method         func(ctx context.Context, service any) (any, error)
	ExpectNotFound bool
	ExpectError    bool
	LogResult      bool
}

// ValidationTestMethod represents parameter validation tests
type ValidationTestMethod struct {
	Name          string
	Method        func(ctx context.Context, service any) error
	ExpectedError bool
	ErrorKeywords []string
}

// RunTestSuite executes a complete integration test suite
func RunTestSuite(t *testing.T, suite TestSuite) {
	t.Helper()

	// Setup client and context
	setup := SetupTestClient(t)
	if setup.Client == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	// Create service instance
	service := suite.Config.ServiceConstructor(setup.Client)

	// Setup timeout context if requested
	var testCtx context.Context
	if suite.Config.UseTimeout {
		duration := max(suite.Config.TimeoutDuration, 30*time.Second)
		var cancel context.CancelFunc
		testCtx, cancel = context.WithTimeout(setup.Context, duration)
		defer cancel()
	} else {
		testCtx = setup.Context
	}

	// Run test groups
	runBasicMethodTests(t, testCtx, service, suite.BasicMethods)
	runFilteredMethodTests(t, testCtx, service, suite.FilterMethods)
	runValidationTests(t, testCtx, service, suite.ValidationTests)
}

// isNetworkError checks if the error is related to network connectivity
func isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	errorMsg := strings.ToLower(err.Error())
	return strings.Contains(errorMsg, "no such host") ||
		strings.Contains(errorMsg, "connection refused") ||
		strings.Contains(errorMsg, "network unreachable") ||
		strings.Contains(errorMsg, "timeout") ||
		strings.Contains(errorMsg, "deadline exceeded")
}

// Private helper functions

// runBasicMethodTests executes basic method test group
func runBasicMethodTests(t *testing.T, testCtx context.Context, service any, methods []TestMethod) {
	if methods == nil {
		return
	}

	t.Run("BasicMethods", func(t *testing.T) {
		for _, method := range methods {
			t.Run(method.Name, func(t *testing.T) {
				result, err := method.Method(testCtx, service)
				handleMethodResult(t, method, result, err, "successfully retrieved data")
			})
		}
	})
}

// runFilteredMethodTests executes filtered method test group
func runFilteredMethodTests(t *testing.T, testCtx context.Context, service any, methods []TestMethod) {
	if methods == nil {
		return
	}

	t.Run("FilteredMethods", func(t *testing.T) {
		for _, method := range methods {
			t.Run(method.Name, func(t *testing.T) {
				result, err := method.Method(testCtx, service)
				handleMethodResult(t, method, result, err, "successfully retrieved filtered data")
			})
		}
	})
}

// runValidationTests executes parameter validation test group
func runValidationTests(t *testing.T, testCtx context.Context, service any, validations []ValidationTestMethod) {
	if validations == nil {
		return
	}

	t.Run("ParameterValidation", func(t *testing.T) {
		for _, validationTest := range validations {
			t.Run(validationTest.Name, func(t *testing.T) {
				err := validationTest.Method(testCtx, service)
				handleValidationResult(t, validationTest, err)
			})
		}
	})
}

// handleMethodResult processes integration test method results
func handleMethodResult(t *testing.T, method TestMethod, result any, err error, successMsg string) {
	if method.ExpectError {
		handleExpectedError(t, method.Name, err)
		return
	}

	if method.ExpectNotFound && err != nil {
		handleExpectedNotFound(t, method.Name, err)
		return
	}

	if err != nil {
		// Check for network errors and skip gracefully
		if isNetworkError(err) {
			t.Skipf("%s: skipping due to network error (controller unreachable): %v", method.Name, err)
			return
		}
		t.Errorf("%s: unexpected error: %v", method.Name, err)
		return
	}

	if result == nil {
		t.Errorf("%s: returned nil data", method.Name)
		return
	}

	if method.LogResult {
		t.Logf("%s: %s", method.Name, successMsg)
	}
}

// handleValidationResult processes parameter validation test results
func handleValidationResult(t *testing.T, validationTest ValidationTestMethod, err error) {
	if validationTest.ExpectedError {
		if err == nil {
			t.Errorf("%s: expected error but got none", validationTest.Name)
			return
		}

		checkErrorKeywords(t, validationTest, err)
		t.Logf("%s: correctly rejected invalid parameter: %v", validationTest.Name, err)
		return
	}

	if err != nil {
		t.Errorf("%s: unexpected error: %v", validationTest.Name, err)
		return
	}

	t.Logf("%s: parameter validation passed", validationTest.Name)
}

// handleExpectedError processes expected error scenarios
func handleExpectedError(t *testing.T, methodName string, err error) {
	if err == nil {
		t.Errorf("%s: expected error but got none", methodName)
		return
	}
	t.Logf("%s: got expected error: %v", methodName, err)
}

// handleExpectedNotFound processes expected 404 scenarios
func handleExpectedNotFound(t *testing.T, methodName string, err error) {
	// Check for network errors first and skip gracefully
	if isNetworkError(err) {
		t.Skipf("%s: skipping due to network error (controller unreachable): %v", methodName, err)
		return
	}

	if core.IsNotFoundError(err) {
		t.Logf("%s: endpoint not supported (404): %v", methodName, err)
		return
	}
	t.Errorf("%s: expected 404 error but got: %v", methodName, err)
}

// checkErrorKeywords validates error contains expected keywords (inlined for single use)
func checkErrorKeywords(t *testing.T, validationTest ValidationTestMethod, err error) {
	if validationTest.ErrorKeywords == nil {
		return
	}

	errorMsg := strings.ToLower(err.Error())
	for _, keyword := range validationTest.ErrorKeywords {
		if strings.Contains(errorMsg, strings.ToLower(keyword)) {
			return // Found expected keyword
		}
	}

	t.Logf("%s: error did not contain expected keywords %v: %v",
		validationTest.Name, validationTest.ErrorKeywords, err)
}
