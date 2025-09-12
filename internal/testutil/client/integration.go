//go:build integration

package client

import (
	"context"
	"strings"
	"testing"
	"time"
)

// TestConstants for standard test parameters
type TestConstants struct {
	TestAPMac     string
	TestEthMAC    string
	TestWtpMAC    string
	TestWlanID    int
	TestSlot      int
	TestRequestID string
	TestLocation  string
}

// TestSuiteConfig defines configuration for integration test suites
type TestSuiteConfig struct {
	ServiceName        string
	ServiceConstructor func(client any) any
	UseTimeout         bool
	TimeoutDuration    time.Duration
}

// IntegrationTestMethod represents a single test method
type IntegrationTestMethod struct {
	Name           string
	Method         func(ctx context.Context, service any) (any, error)
	ExpectNotFound bool
	ExpectError    bool
	LogResult      bool
}

// IntegrationTestSuite represents a complete test suite for integration tests
type IntegrationTestSuite struct {
	Config          TestSuiteConfig
	BasicMethods    []IntegrationTestMethod
	FilterMethods   []IntegrationTestMethod
	ValidationTests []ValidationTestMethod
}

// ValidationTestMethod represents parameter validation tests
type ValidationTestMethod struct {
	Name          string
	Method        func(ctx context.Context, service any) error
	ExpectedError bool
	ErrorKeywords []string
}

// RunIntegrationTestSuite executes a complete integration test suite
func RunIntegrationTestSuite(t *testing.T, suite IntegrationTestSuite) {
	t.Helper()

	// Setup client and context using existing client package
	setup := SetupRequiredClient(t)

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

// isNotFoundError checks if an error is a 404 not found error
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	lowerMsg := strings.ToLower(err.Error())
	return strings.Contains(lowerMsg, "404") ||
		strings.Contains(lowerMsg, "not found") ||
		strings.Contains(lowerMsg, "no data found")
}

// Private helper functions

// runBasicMethodTests executes basic method test group
func runBasicMethodTests(t *testing.T, testCtx context.Context, service any, methods []IntegrationTestMethod) {
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
func runFilteredMethodTests(t *testing.T, testCtx context.Context, service any, methods []IntegrationTestMethod) {
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
func handleMethodResult(t *testing.T, method IntegrationTestMethod, result any, err error, successMsg string) {
	if method.ExpectError {
		handleExpectedError(t, method.Name, err)
		return
	}

	if method.ExpectNotFound && err != nil {
		handleExpectedNotFound(t, method.Name, err)
		return
	}

	if err != nil {
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
	if isNotFoundError(err) {
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
