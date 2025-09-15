//go:build integration || scenario

// Package integration provides test utilities for integration tests.
package integration

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
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
	LogResult      bool

	// Optional callback for custom error handling. Returns true if handled.
	WhenError func(t *testing.T, methodName string, err error) bool
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

	// Create service instance and test context
	service := suite.Config.ServiceConstructor(setup.Client)
	testCtx := createTestContext(setup.Context, suite.Config)

	// Run test groups
	runBasicMethodTests(t, testCtx, service, suite.BasicMethods, suite.Config.ServiceName)
	runFilteredMethodTests(t, testCtx, service, suite.FilterMethods, suite.Config.ServiceName)
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

// createTestContext creates an appropriate context based on suite configuration
func createTestContext(baseCtx context.Context, config TestSuiteConfig) context.Context {
	if !config.UseTimeout {
		return baseCtx
	}

	duration := max(config.TimeoutDuration, 30*time.Second)
	testCtx, cancel := context.WithTimeout(baseCtx, duration)

	// Note: We can't defer cancel() here as the context needs to live beyond this function
	// The caller should handle cancellation or use a parent context that gets cancelled
	_ = cancel // Mark as intentionally unused to avoid context leak warning
	return testCtx
}

// runBasicMethodTests executes basic method test group
func runBasicMethodTests(t *testing.T, testCtx context.Context, service any, methods []TestMethod, serviceName string) {
	if methods == nil {
		return
	}

	t.Run("BasicMethods", func(t *testing.T) {
		for _, method := range methods {
			t.Run(method.Name, func(t *testing.T) {
				result, err := method.Method(testCtx, service)
				handleMethodResult(t, method, result, err, "successfully retrieved data", serviceName)
			})
		}
	})
}

// runFilteredMethodTests executes filtered method test group
func runFilteredMethodTests(t *testing.T, testCtx context.Context, service any, methods []TestMethod, serviceName string) {
	if methods == nil {
		return
	}

	t.Run("FilteredMethods", func(t *testing.T) {
		for _, method := range methods {
			t.Run(method.Name, func(t *testing.T) {
				result, err := method.Method(testCtx, service)
				handleMethodResult(t, method, result, err, "successfully retrieved filtered data", serviceName)
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
func handleMethodResult(t *testing.T, method TestMethod, result any, err error, successMsg string, serviceName string) {
	// Handle expected not-found scenarios first
	if method.ExpectNotFound && err != nil {
		if isNetworkError(err) {
			t.Skipf("%s: skipping due to network error (controller unreachable): %v", method.Name, err)
			return
		}
		if core.IsNotFoundError(err) {
			t.Logf("%s: endpoint not supported (404): %v", method.Name, err)
			return
		}
		t.Errorf("%s: expected 404 error but got: %v", method.Name, err)
		return
	}

	// Handle errors
	if err != nil {
		if isNetworkError(err) {
			t.Skipf("%s: skipping due to network error (controller unreachable): %v", method.Name, err)
			return
		}
		if method.WhenError != nil && method.WhenError(t, method.Name, err) {
			return // Custom error handler processed the error
		}
		t.Errorf("%s: unexpected error: %v", method.Name, err)
		return
	}

	// Handle nil result case
	if result == nil {
		t.Errorf("%s: returned nil data", method.Name)
		return
	}

	// Success path: save response and log result
	saveTestResponse(t, method.Name, result)
	if method.LogResult {
		t.Logf("%s: %s", method.Name, successMsg)
	}
}

// handleValidationResult processes parameter validation test results
func handleValidationResult(t *testing.T, validationTest ValidationTestMethod, err error) {
	if !validationTest.ExpectedError {
		if err != nil {
			t.Errorf("%s: unexpected error: %v", validationTest.Name, err)
			return
		}
		t.Logf("%s: parameter validation passed", validationTest.Name)
		return
	}

	if err == nil {
		t.Errorf("%s: expected error but got none", validationTest.Name)
		return
	}

	// Check for expected keywords in error message
	if len(validationTest.ErrorKeywords) > 0 {
		errorMsg := strings.ToLower(err.Error())
		found := false
		for _, keyword := range validationTest.ErrorKeywords {
			if strings.Contains(errorMsg, strings.ToLower(keyword)) {
				found = true
				break
			}
		}
		if !found {
			t.Logf("%s: error did not contain expected keywords %v: %v",
				validationTest.Name, validationTest.ErrorKeywords, err)
		}
	}

	t.Logf("%s: correctly rejected invalid parameter: %v", validationTest.Name, err)
}

// saveTestResponse saves the test response to testdata directory for future reference
func saveTestResponse(t *testing.T, methodName string, response any) {
	t.Helper()

	if response == nil {
		return
	}

	// Get project root directory
	cwd, err := os.Getwd()
	if err != nil {
		t.Logf("%s: warning: failed to get working directory: %v", methodName, err)
		return
	}
	rootDir := findProjectRoot(cwd)

	// Create testdata directory
	serviceDir := extractServiceDirFromTestName(t)
	testdataDir := filepath.Join(rootDir, "testdata", "integration", serviceDir)
	if err := os.MkdirAll(testdataDir, 0o755); err != nil {
		t.Logf("%s: warning: failed to create testdata directory: %v", methodName, err)
		return
	}

	// Marshal response to JSON
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		t.Logf("%s: warning: failed to marshal response: %v", methodName, err)
		return
	}

	// Save to file
	filename := methodName + ".json"
	fullPath := filepath.Join(testdataDir, filename)
	if err := os.WriteFile(fullPath, jsonData, 0o644); err != nil {
		t.Logf("%s: warning: failed to save test response: %v", methodName, err)
		return
	}

	t.Logf("%s: saved response to testdata/integration/%s/%s", methodName, serviceDir, filename)
}

// findProjectRoot walks up the directory tree to find the project root containing go.mod
func findProjectRoot(startDir string) string {
	currentDir := startDir

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			return startDir
		}
		currentDir = parent
	}
}

// extractServiceDirFromTestName extracts the service directory name from the test name.
// It looks for the pattern "Test{Service}ServiceIntegration_*" and extracts the service part.
// For example: "TestAFCServiceIntegration_GetOperationalOperations_Success" -> "afc"
func extractServiceDirFromTestName(t *testing.T) string {
	testName := t.Name()

	// Handle test hierarchy by taking the top-level test name
	if idx := strings.Index(testName, "/"); idx != -1 {
		testName = testName[:idx]
	}

	// Extract service name from pattern "Test{Service}ServiceIntegration_*"
	if strings.HasPrefix(testName, "Test") && strings.Contains(testName, "ServiceIntegration") {
		// Remove "Test" prefix
		withoutTest := strings.TrimPrefix(testName, "Test")

		// Find "ServiceIntegration" and extract the part before it
		if idx := strings.Index(withoutTest, "ServiceIntegration"); idx != -1 {
			serviceName := withoutTest[:idx]
			return strings.ToLower(serviceName)
		}
	}

	// Fallback: try to extract from current working directory or test file pattern
	cwd, err := os.Getwd()
	if err != nil {
		return strings.ToLower(testName)
	}

	if !strings.Contains(cwd, "tests/integration") {
		return strings.ToLower(testName)
	}

	// Look for pattern like "*_service_test.go"
	files, err := filepath.Glob(filepath.Join(cwd, "*_service_test.go"))
	if err != nil || len(files) == 0 {
		return strings.ToLower(testName)
	}

	for _, file := range files {
		basename := filepath.Base(file)
		if strings.HasSuffix(basename, "_service_test.go") {
			serviceName := strings.TrimSuffix(basename, "_service_test.go")
			return serviceName
		}
	}

	return strings.ToLower(testName)
}
