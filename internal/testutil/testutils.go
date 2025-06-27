// Package testutil provides testing utilities and helper functions for the Cisco Wireless Network Controller API client.
package testutil

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// Test timeout constants in seconds for better readability
const (
	// DefaultTestTimeoutSeconds defines the standard timeout for most tests
	DefaultTestTimeoutSeconds = 40

	// ExtendedTestTimeoutSeconds defines timeout for longer operations
	ExtendedTestTimeoutSeconds = 60

	// ComprehensiveTestTimeoutSeconds defines timeout for comprehensive test suites
	ComprehensiveTestTimeoutSeconds = 150

	// QuickTestTimeoutSeconds defines timeout for fast operations
	QuickTestTimeoutSeconds = 5

	// MicroTestTimeoutMicroseconds defines timeout for immediate cancellation tests
	MicroTestTimeoutMicroseconds = 1

	// StandardNetworkTimeoutSeconds defines timeout for network operations
	StandardNetworkTimeoutSeconds = 60
)

// Test-related constants - using constants from main package where possible
const (
	// TestDataDirName defines the default test data directory name
	TestDataDirName = "test_data"

	// JSONFileExtension defines the JSON file extension
	JSONFileExtension = ".json"

	// FilePermissionMode defines the default file permission mode
	FilePermissionMode = 0755

	// JSONIndentSpaces defines the number of spaces for JSON indentation
	JSONIndentSpaces = 2

	// JSONIndentPrefix defines the prefix for JSON indentation
	JSONIndentPrefix = ""

	// JSONIndentString defines the indentation string for JSON
	JSONIndentString = "  "

	// TestDataDir is the directory where test data files are saved
	TestDataDir = TestDataDirName

	// DefaultTestTimeout is the standard timeout for most tests
	DefaultTestTimeout = DefaultTestTimeoutSeconds * time.Second

	// ExtendedTestTimeout is used for longer operations
	ExtendedTestTimeout = ExtendedTestTimeoutSeconds * time.Second

	// ComprehensiveTestTimeout is used for comprehensive test suites
	ComprehensiveTestTimeout = ComprehensiveTestTimeoutSeconds * time.Second

	// QuickTestTimeout is used for fast operations
	QuickTestTimeout = QuickTestTimeoutSeconds * time.Second

	// MicroTestTimeout is used for immediate cancellation tests
	MicroTestTimeout = MicroTestTimeoutMicroseconds * time.Microsecond
)

// Common test endpoint validation constants - using main package constants
const (
	MinEndpointLength = wnc.MinEndpointLength
)

// TestConfig represents configuration for test operations
type TestConfig struct {
	Controller  string
	AccessToken string
	Timeout     time.Duration
}

// NewTestConfig creates a new test configuration
func NewTestConfig(controller, accessToken string, timeout time.Duration) *TestConfig {
	return &TestConfig{
		Controller:  controller,
		AccessToken: accessToken,
		Timeout:     timeout,
	}
}

// NewTestConfigFromEnv creates a test configuration from environment variables
// This function centralizes environment variable access
func NewTestConfigFromEnv() *TestConfig {
	controller := os.Getenv(wnc.EnvVarController)
	accessToken := os.Getenv(wnc.EnvVarAccessToken)

	if controller == "" || accessToken == "" {
		return nil
	}

	return NewTestConfig(controller, accessToken, DefaultTestTimeout)
}

// IsValid checks if the test configuration is valid
func (tc *TestConfig) IsValid() bool {
	return tc.Controller != "" && tc.AccessToken != ""
}

// TestDataCollector represents a generic test data collector
type TestDataCollector struct {
	Data map[string]interface{} `json:"test_data"`
}

// NewTestDataCollector creates a new test data collector
func NewTestDataCollector() *TestDataCollector {
	return &TestDataCollector{
		Data: make(map[string]interface{}),
	}
}

// isTestDataDirExists checks if test data directory exists
func isTestDataDirExists() bool {
	_, err := os.Stat(TestDataDir)
	return !os.IsNotExist(err)
}

// ensureTestDataDir creates the test data directory if it doesn't exist
func ensureTestDataDir() error {
	if isTestDataDirExists() {
		return nil
	}
	return os.MkdirAll(TestDataDir, FilePermissionMode)
}

// SaveTestDataToFile saves test data to a JSON file in the test data directory
func SaveTestDataToFile(filename string, data interface{}) error {
	// Early return if directory creation fails
	if err := ensureTestDataDir(); err != nil {
		return err
	}

	targetFilePath := filepath.Join(TestDataDir, filename)
	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := targetFile.Close(); closeErr != nil {
			// Log error but don't return it since we're in a defer
			fmt.Printf("Warning: Failed to close file %s: %v\n", targetFilePath, closeErr)
		}
	}()

	jsonEncoder := json.NewEncoder(targetFile)
	jsonEncoder.SetIndent(JSONIndentPrefix, JSONIndentString)
	return jsonEncoder.Encode(data)
}

// CreateTestContext creates a context with the specified timeout
func CreateTestContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

// CreateDefaultTestContext creates a context with the default test timeout
func CreateDefaultTestContext() (context.Context, context.CancelFunc) {
	return CreateTestContext(DefaultTestTimeout)
}

// CreateExtendedTestContext creates a context with extended timeout
func CreateExtendedTestContext() (context.Context, context.CancelFunc) {
	return CreateTestContext(ExtendedTestTimeout)
}

// CreateComprehensiveTestContext creates a context with comprehensive timeout
func CreateComprehensiveTestContext() (context.Context, context.CancelFunc) {
	return CreateTestContext(ComprehensiveTestTimeout)
}

// validateEndpoint validates that an endpoint meets basic requirements
func validateEndpoint(t *testing.T, endpointName, endpointValue string) {
	t.Helper()

	// Early return for empty endpoint
	if endpointValue == "" {
		t.Errorf(wnc.EmptyEndpointErrorTemplate, endpointName)
		return
	}

	// Early return for short endpoint
	if len(endpointValue) < MinEndpointLength {
		t.Errorf(wnc.ShortEndpointErrorTemplate, endpointName, endpointValue)
		return
	}
}

// ValidateEndpoints validates a map of endpoints
func ValidateEndpoints(t *testing.T, endpointsToValidate map[string]string) {
	t.Helper()

	for endpointName, endpointValue := range endpointsToValidate {
		t.Run("Validate_"+endpointName, func(t *testing.T) {
			validateEndpoint(t, endpointName, endpointValue)
		})
	}
}

// ValidateEndpointsWithConstants validates endpoints against expected constants using reflection
func ValidateEndpointsWithConstants(t *testing.T, expectedEndpoints map[string]string, constantsStruct interface{}) {
	t.Helper()

	structValue := reflect.ValueOf(constantsStruct)
	if structValue.Kind() == reflect.Ptr {
		structValue = structValue.Elem()
	}

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			// Try to find the field by name
			fieldValue := structValue.FieldByName(name)
			if !fieldValue.IsValid() {
				t.Errorf("Constant %s not found in provided struct", name)
				return
			}

			if fieldValue.Kind() != reflect.String {
				t.Errorf("Constant %s is not a string", name)
				return
			}

			actual := fieldValue.String()
			if actual != expected {
				t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, actual)
			}
		})
	}
}

// ValidateEndpointConstants is a simplified version that validates endpoint constants
func ValidateEndpointConstants(t *testing.T, endpoints map[string]string) {
	t.Helper()

	for name, expected := range endpoints {
		t.Run(name, func(t *testing.T) {
			validateEndpoint(t, name, expected)
		})
	}
}

// CollectTestResult collects test results in a standard format
func CollectTestResult(collector *TestDataCollector, methodName, endpoint string, result interface{}, err error) {
	testData := map[string]interface{}{
		"method":    methodName,
		"endpoint":  endpoint,
		"timestamp": time.Now().Format(time.RFC3339),
	}

	if err != nil {
		testData["error"] = err.Error()
		testData["success"] = false
	} else {
		testData["success"] = true
		testData["response"] = result
	}

	collector.Data[methodName] = testData
}

// SaveCollectedTestData saves collected test data to a file with standard logging
func SaveCollectedTestData(t *testing.T, collector *TestDataCollector, filename string) {
	t.Helper()

	if len(collector.Data) > 0 {
		if err := SaveTestDataToFile(filename, collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Test data saved to %s/%s", TestDataDir, filename)
		}
	}
}

// GenerateEndpointValidationTest creates a standard endpoint validation test
func GenerateEndpointValidationTest(t *testing.T, expectedEndpoints map[string]string, actualEndpoints map[string]string) {
	t.Helper()

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			actual, exists := actualEndpoints[name]
			if !exists {
				t.Errorf("Endpoint constant %s not found", name)
				return
			}
			if actual != expected {
				t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, actual)
			}
		})
	}
}

// CreateStandardTestContext creates a context with a specified timeout using constants
func CreateStandardTestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultTestTimeout)
}

// CreateQuickTestContext creates a context with quick timeout
func CreateQuickTestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), QuickTestTimeout)
}

// CreateMicroTestContext creates a context with micro timeout for cancellation tests
func CreateMicroTestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), MicroTestTimeout)
}

// CreateTestClient creates a test client using provided configuration
func CreateTestClient(t *testing.T, config *TestConfig) *wnc.Client {
	t.Helper()

	if config == nil || !config.IsValid() {
		t.Skip("Invalid test configuration - skipping test")
	}

	wncConfig := wnc.Config{
		Controller:         config.Controller,
		AccessToken:        config.AccessToken,
		Timeout:            config.Timeout,
		InsecureSkipVerify: true,
	}
	client, err := wnc.NewClient(wncConfig)
	if err != nil {
		t.Fatalf("Failed to create test client: %v", err)
	}

	return client
}

// CreateTestClientFromEnv creates a test client using environment variables
func CreateTestClientFromEnv(t *testing.T) *wnc.Client {
	t.Helper()

	config := NewTestConfigFromEnv()
	if config == nil {
		t.Skip("Required environment variables not set - skipping test")
	}

	return CreateTestClient(t, config)
}

// GetTestClient creates a test client using provided configuration
func GetTestClient(config *TestConfig) *wnc.Client {
	if config == nil || !config.IsValid() {
		return nil
	}

	client, err := CreateRealWNCClient(config.Controller, config.AccessToken, int(config.Timeout.Seconds()))
	if err != nil {
		return nil
	}

	return client
}

// GetTestClientFromEnv creates a test client using environment variables
func GetTestClientFromEnv() *wnc.Client {
	config := NewTestConfigFromEnv()
	if config == nil {
		return nil
	}

	return GetTestClient(config)
}

// CheckTestConfig validates a test configuration
func CheckTestConfig(t *testing.T, config *TestConfig) {
	t.Helper()

	if config == nil {
		t.Skip("Test configuration not provided - skipping test")
	}

	if config.Controller == "" {
		t.Skip("Controller not configured - skipping test")
	}

	if config.AccessToken == "" {
		t.Skip("Access token not configured - skipping test")
	}
}

// CreateRealWNCClient creates a WNC client using the provided credentials for testing against real hardware
func CreateRealWNCClient(controller, accessToken string, timeoutSeconds int) (*wnc.Client, error) {
	timeoutDuration := time.Duration(timeoutSeconds) * time.Second

	config := wnc.Config{
		Controller:         controller,
		AccessToken:        accessToken,
		Timeout:            timeoutDuration,
		InsecureSkipVerify: true,
	}
	client, err := wnc.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create WNC client: %w", err)
	}

	return client, nil
}

// DebugJSONResponse logs and saves raw JSON response for debugging data structure issues
func DebugJSONResponse(t *testing.T, endpointName, rawJSON string) {
	t.Helper()

	// Log the raw response
	t.Logf("Raw JSON response from %s:\n%s", endpointName, rawJSON)

	// Save to file for detailed analysis
	filename := fmt.Sprintf("debug_%s_response.json", endpointName)
	debugData := map[string]interface{}{
		"endpoint":  endpointName,
		"timestamp": time.Now().Format(time.RFC3339),
		"raw_json":  rawJSON,
	}

	if err := SaveTestDataToFile(filename, debugData); err != nil {
		t.Logf("Warning: Could not save debug data: %v", err)
	} else {
		t.Logf("Debug data saved to %s/%s", TestDataDir, filename)
	}
}

// createHTTPClient creates an HTTP client with the specified timeout
func createHTTPClient(timeoutSeconds int) *http.Client {
	timeoutDuration := time.Duration(timeoutSeconds) * time.Second

	return &http.Client{
		Timeout: timeoutDuration,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // For testing against controllers with self-signed certs
			},
		},
	}
}

// buildAPIRequest creates an HTTP request with proper headers
func buildAPIRequest(controller, accessToken, endpoint string) (*http.Request, error) {
	requestURL := fmt.Sprintf("%s://%s%s", wnc.HTTPSScheme, controller, endpoint)

	req, err := http.NewRequest(wnc.HTTPMethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add required headers
	req.Header.Set(wnc.HTTPHeaderKeyAuthorization, wnc.HTTPHeaderValueBasicPrefix+accessToken)
	req.Header.Set(wnc.HTTPHeaderKeyAccept, wnc.HTTPHeaderAccept)
	req.Header.Set(wnc.HTTPHeaderKeyContentType, wnc.HTTPHeaderContentType)

	return req, nil
}

// executeAPIRequest executes the HTTP request and returns the response body
func executeAPIRequest(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			// Log error but don't return it since we're in a defer
			fmt.Printf("Warning: Failed to close response body: %v\n", closeErr)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// MakeRawAPIRequest makes a raw API request and returns the response body for debugging
func MakeRawAPIRequest(controller, accessToken string, endpoint string, timeoutSeconds int) ([]byte, error) {
	client := createHTTPClient(timeoutSeconds)

	req, err := buildAPIRequest(controller, accessToken, endpoint)
	if err != nil {
		return nil, err
	}

	return executeAPIRequest(client, req)
}

// SaveTestDataWithLogging saves test data to a file with standard logging
// This replaces all duplicate saveTestDataWithLogging functions across test files
func SaveTestDataWithLogging(filename string, data interface{}) {
	if err := SaveTestDataToFile(filename, data); err != nil {
		fmt.Printf("Warning: Failed to save data to %s: %v\n", filename, err)
	} else {
		fmt.Printf("Test data saved to %s/%s\n", TestDataDir, filename)
	}
}
