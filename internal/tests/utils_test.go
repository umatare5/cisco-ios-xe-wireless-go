package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// TestBasicUtilityFunctions tests basic utility functions to improve coverage
func TestBasicUtilityFunctions(t *testing.T) {
	t.Run("Constants", func(t *testing.T) {
		if DefaultTestTimeoutSeconds != 40 {
			t.Errorf("Expected DefaultTestTimeoutSeconds to be 40, got %d", DefaultTestTimeoutSeconds)
		}
		if ExtendedTestTimeoutSeconds != 60 {
			t.Errorf("Expected ExtendedTestTimeoutSeconds to be 60, got %d", ExtendedTestTimeoutSeconds)
		}
		if TestDataDirName != "test_data" {
			t.Errorf("Expected TestDataDirName to be 'test_data', got %s", TestDataDirName)
		}
		if HTTPSScheme != "https" {
			t.Errorf("Expected HTTPSScheme to be 'https', got %s", HTTPSScheme)
		}
		if HTTPMethodGet != "GET" {
			t.Errorf("Expected HTTPMethodGet to be 'GET', got %s", HTTPMethodGet)
		}
	})

	t.Run("NewTestConfig", func(t *testing.T) {
		config := NewTestConfig("test-controller", "test-token", 30*time.Second)
		if config == nil {
			t.Error("Expected non-nil config")
		}
		if config.Controller != "test-controller" {
			t.Errorf("Expected controller 'test-controller', got %s", config.Controller)
		}
		if config.AccessToken != "test-token" {
			t.Errorf("Expected token 'test-token', got %s", config.AccessToken)
		}
	})

	t.Run("NewTestDataCollector", func(t *testing.T) {
		collector := NewTestDataCollector()
		if collector == nil {
			t.Error("Expected non-nil collector")
		}
		if collector.Data == nil {
			t.Error("Expected initialized Data map")
		}
	})

	t.Run("CreateTestContext", func(t *testing.T) {
		timeout := 5 * time.Second
		ctx, cancel := CreateTestContext(timeout)
		defer cancel()
		if ctx == nil {
			t.Error("Expected non-nil context")
		}
		if deadline, ok := ctx.Deadline(); !ok || time.Until(deadline) > timeout {
			t.Error("Context deadline not set correctly")
		}
	})

	t.Run("CreateDefaultTestContext", func(t *testing.T) {
		ctx, cancel := CreateDefaultTestContext()
		defer cancel()
		if ctx == nil {
			t.Error("Expected non-nil context")
		}
	})

	t.Run("CreateCancelledTestContext", func(t *testing.T) {
		ctx, cancel := CreateCancelledTestContext()
		defer cancel()
		if ctx == nil {
			t.Error("Expected non-nil context")
		}
		select {
		case <-ctx.Done():
			// Expected - context should be cancelled
		default:
			t.Error("Expected cancelled context")
		}
	})
}

// TestMockHTTPServer tests the MockHTTPServer functionality
func TestMockHTTPServer(t *testing.T) {
	t.Run("NewMockHTTPServer", func(t *testing.T) {
		server := NewMockHTTPServer()
		if server == nil {
			t.Error("Expected non-nil server")
		}
		defer server.Close()

		if server.Server == nil {
			t.Error("Expected server.Server to be initialized")
		}
		if server.Handlers == nil {
			t.Error("Expected server.Handlers to be initialized")
		}
	})

	t.Run("AddHandler", func(t *testing.T) {
		server := NewMockHTTPServer()
		defer server.Close()

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("test response"))
		}

		server.AddHandler("/test", handler)
		if len(server.Handlers) != 1 {
			t.Errorf("Expected 1 handler, got %d", len(server.Handlers))
		}
	})

	t.Run("CreateJSONResponse", func(t *testing.T) {
		response := TestHTTPResponse{
			StatusCode: http.StatusOK,
			Body:       `{"test": "data"}`,
			Headers:    map[string]string{"X-Test": "value"},
		}

		handler := CreateJSONResponse(response)
		if handler == nil {
			t.Error("Expected non-nil handler")
		}
	})

	t.Run("CreateTestClientForMockServer", func(t *testing.T) {
		server := NewMockHTTPServer()
		defer server.Close()

		client := CreateTestClientForMockServer(t, server)
		if client == nil {
			t.Error("Expected non-nil client")
		}
	})
}

// TestAPIHelpers tests the API helper functions
func TestAPIHelpers(t *testing.T) {
	t.Run("TestAPIFunction", func(t *testing.T) {
		TestAPIFunction(t, "SimpleTest", func(mock *MockHTTPServer) {
			mock.AddHandler("/test", CreateJSONResponse(TestHTTPResponse{
				StatusCode: http.StatusOK,
				Body:       `{"status": "ok"}`,
			}))
		}, func(c *wnc.Client) error {
			return nil // Simple success case
		})
	})

	t.Run("TestAPIFunctionWithContext", func(t *testing.T) {
		TestAPIFunctionWithContext(t, "ContextTest", func(mock *MockHTTPServer) {
			mock.AddHandler("/test", CreateJSONResponse(TestHTTPResponse{
				StatusCode: http.StatusOK,
				Body:       `{"status": "ok"}`,
			}))
		}, func(ctx context.Context, c *wnc.Client) error {
			return nil // Simple success case
		})
	})

	t.Run("TestTimeoutAPI", func(t *testing.T) {
		TestTimeoutAPI(t, "TimeoutTest", func(ctx context.Context, c *wnc.Client) error {
			// Simulate a long-running operation that should timeout
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(200 * time.Millisecond):
				return nil // This should not be reached if timeout works properly
			}
		})
	})
}

// TestErrorHelpers tests the error helper functions
func TestErrorHelpers(t *testing.T) {
	t.Run("GetNilClientErrorTests", func(t *testing.T) {
		tests := GetNilClientErrorTests()
		if len(tests) == 0 {
			t.Error("Expected non-empty test cases")
		}

		for _, testCase := range tests {
			err := testCase.TestFunc(nil)
			if err == nil {
				t.Errorf("Test case %s: expected error with nil client", testCase.Name)
			}
		}
	})

	t.Run("TestWithCancelledContext", func(t *testing.T) {
		TestWithCancelledContext(t, func(ctx context.Context, c *wnc.Client) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return fmt.Errorf("context not cancelled")
			}
		})
	})

	t.Run("TestWithTimeout", func(t *testing.T) {
		TestWithTimeout(t, func(ctx context.Context, c *wnc.Client) error {
			return nil // Simple success case
		}, 100*time.Millisecond)
	})

	t.Run("TestContextHandling", func(t *testing.T) {
		TestContextHandling(t, func(ctx context.Context, c *wnc.Client) error {
			// Check if context is cancelled first
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				// If context is not cancelled, simulate a long operation that will be interrupted
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(200 * time.Millisecond):
					return fmt.Errorf("operation completed without context cancellation")
				}
			}
		})
	})
}

// TestValidationHelpers tests the validation helper functions
func TestValidationHelpers(t *testing.T) {
	t.Run("ValidateStringNotEmpty", func(t *testing.T) {
		// Test with valid string
		ValidateStringNotEmpty(t, "test", "TestField")
	})

	t.Run("ValidateErrorContains", func(t *testing.T) {
		err := fmt.Errorf("test error message")
		ValidateErrorContains(t, err, "test error")
	})

	t.Run("ValidateNoError", func(t *testing.T) {
		ValidateNoError(t, nil, "test operation")
	})

	t.Run("ValidateEndpoints", func(t *testing.T) {
		endpoints := map[string]string{
			"TestEndpoint": "/api/v1/test-endpoint",
		}
		ValidateEndpoints(t, endpoints)
	})
}

// TestUtilityFunctions tests various utility functions
func TestUtilityFunctions(t *testing.T) {
	t.Run("GetTestCredentials", func(t *testing.T) {
		// Save original environment
		origController := os.Getenv("WNC_CONTROLLER")
		origToken := os.Getenv("WNC_ACCESS_TOKEN")
		defer func() {
			if origController != "" {
				os.Setenv("WNC_CONTROLLER", origController)
			} else {
				os.Unsetenv("WNC_CONTROLLER")
			}
			if origToken != "" {
				os.Setenv("WNC_ACCESS_TOKEN", origToken)
			} else {
				os.Unsetenv("WNC_ACCESS_TOKEN")
			}
		}()

		// Test with environment variables
		os.Setenv("WNC_CONTROLLER", "test-controller")
		os.Setenv("WNC_ACCESS_TOKEN", "test-token")

		controller, token, ok := GetTestCredentials()
		if !ok {
			t.Error("Expected credentials to be available")
		}
		if controller != "test-controller" {
			t.Errorf("Expected controller 'test-controller', got %s", controller)
		}
		if token != "test-token" {
			t.Errorf("Expected token 'test-token', got %s", token)
		}
	})

	t.Run("TestServiceMethod", func(t *testing.T) {
		TestServiceMethod(t, func() error {
			return nil // Simple success case
		})
	})

	t.Run("RunTableTests", func(t *testing.T) {
		executed := false
		tests := []TableTest{
			{
				Name: "test case",
				Test: func(t *testing.T) {
					executed = true
				},
			},
		}

		RunTableTests(t, tests)
		if !executed {
			t.Error("Table test was not executed")
		}
	})

	t.Run("CollectTestResult", func(t *testing.T) {
		collector := NewTestDataCollector()
		result := map[string]interface{}{"test": "data"}

		CollectTestResult(collector, "TestMethod", "/test", result, nil)
		if len(collector.Data) != 1 {
			t.Errorf("Expected 1 result, got %d", len(collector.Data))
		}
	})

	t.Run("isTestDataDirExists", func(t *testing.T) {
		exists := isTestDataDirExists()
		// Just test that the function returns without error
		_ = exists
	})

	t.Run("ensureTestDataDir", func(t *testing.T) {
		err := ensureTestDataDir()
		if err != nil {
			t.Errorf("Failed to ensure test data dir: %v", err)
		}
	})

	t.Run("SaveTestDataToFile", func(t *testing.T) {
		testData := map[string]interface{}{
			"test": "data",
			"time": time.Now(),
		}

		err := SaveTestDataToFile("test_coverage_data.json", testData)
		if err != nil {
			t.Errorf("Failed to save test data: %v", err)
		}
	})

	t.Run("SaveCollectedTestData", func(t *testing.T) {
		collector := NewTestDataCollector()
		collector.Data["test"] = map[string]interface{}{"test": "data"}

		SaveCollectedTestData(t, collector, "test_collected_data.json")
		// Function logs on success/failure, so just ensure it doesn't panic
	})

	t.Run("SaveTestDataWithLogging", func(t *testing.T) {
		testData := map[string]interface{}{"test": "logging"}
		SaveTestDataWithLogging("test_logging_data.json", testData)
		// Function logs on success/failure, so just ensure it doesn't panic
	})

	t.Run("DebugJSONResponse", func(t *testing.T) {
		rawJSON := `{"debug": "test", "data": [1, 2, 3]}`
		DebugJSONResponse(t, "test_endpoint", rawJSON)
		// Function logs and saves data, so just ensure it doesn't panic
	})
}

// TestAdditionalUtilityFunctions tests more utility functions for coverage
func TestAdditionalUtilityFunctions(t *testing.T) {
	t.Run("CreateExtendedTestContext", func(t *testing.T) {
		ctx, cancel := CreateExtendedTestContext()
		defer cancel()
		if ctx == nil {
			t.Error("Expected non-nil context")
		}
		deadline, ok := ctx.Deadline()
		if !ok {
			t.Error("Expected context with deadline")
		}
		if time.Until(deadline) > ExtendedTestTimeout+time.Second {
			t.Error("Context deadline not set correctly for extended timeout")
		}
	})

	t.Run("CreateComprehensiveTestContext", func(t *testing.T) {
		ctx, cancel := CreateComprehensiveTestContext()
		defer cancel()
		if ctx == nil {
			t.Error("Expected non-nil context")
		}
	})

	t.Run("CreateQuickTestContext", func(t *testing.T) {
		ctx, cancel := CreateQuickTestContext()
		defer cancel()
		if ctx == nil {
			t.Error("Expected non-nil context")
		}
	})

	t.Run("CreateMicroTestContext", func(t *testing.T) {
		ctx, cancel := CreateMicroTestContext()
		defer cancel()
		if ctx == nil {
			t.Error("Expected non-nil context")
		}
	})

	t.Run("CreateStandardTestContext", func(t *testing.T) {
		ctx, cancel := CreateStandardTestContext()
		defer cancel()
		if ctx == nil {
			t.Error("Expected non-nil context")
		}
	})

	t.Run("GetTestClientWithTimeout", func(t *testing.T) {
		// Skip if no environment variables
		if os.Getenv("WNC_CONTROLLER") == "" || os.Getenv("WNC_ACCESS_TOKEN") == "" {
			t.Skip("Required environment variables not set - skipping test")
		}

		client := GetTestClientWithTimeout(t, 10*time.Second)
		if client == nil {
			t.Error("Expected non-nil client")
		}
	})

	t.Run("GetTestClientWithContext", func(t *testing.T) {
		// Skip if no environment variables
		if os.Getenv("WNC_CONTROLLER") == "" || os.Getenv("WNC_ACCESS_TOKEN") == "" {
			t.Skip("Required environment variables not set - skipping test")
		}

		ctx := context.Background()
		client := GetTestClientWithContext(t, ctx)
		if client == nil {
			t.Error("Expected non-nil client")
		}
	})

	t.Run("CreateTestClientWithTimeout", func(t *testing.T) {
		// Skip if no environment variables
		if os.Getenv("WNC_CONTROLLER") == "" || os.Getenv("WNC_ACCESS_TOKEN") == "" {
			t.Skip("Required environment variables not set - skipping test")
		}

		client := CreateTestClientWithTimeout(t, 15*time.Second)
		if client == nil {
			t.Error("Expected non-nil client")
		}
	})

	t.Run("ValidateClient", func(t *testing.T) {
		// Skip if no environment variables
		if os.Getenv("WNC_CONTROLLER") == "" || os.Getenv("WNC_ACCESS_TOKEN") == "" {
			t.Skip("Required environment variables not set - skipping test")
		}

		client := GetTestClient(t)
		ValidateClient(t, client)
	})

	t.Run("GenerateEndpointValidationTest", func(t *testing.T) {
		expected := map[string]string{
			"TestEndpoint": "/api/v1/test",
		}
		actual := map[string]string{
			"TestEndpoint": "/api/v1/test",
		}

		GenerateEndpointValidationTest(t, expected, actual)
	})
}

// TestStringHelperFunctions tests string helper functions for coverage
func TestStringHelperFunctions(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		tests := []struct {
			s      string
			substr string
			want   bool
		}{
			{"hello world", "hello", true},
			{"hello world", "world", true},
			{"hello world", "lo wo", true},
			{"hello world", "xyz", false},
			{"test", "test", true},
			{"", "", true},
			{"test", "", true},
			{"", "test", false},
		}

		for _, tt := range tests {
			got := contains(tt.s, tt.substr)
			if got != tt.want {
				t.Errorf("contains(%q, %q) = %v, want %v", tt.s, tt.substr, got, tt.want)
			}
		}
	})

	t.Run("indexOf", func(t *testing.T) {
		tests := []struct {
			s      string
			substr string
			want   int
		}{
			{"hello world", "hello", 0},
			{"hello world", "world", 6},
			{"hello world", "lo wo", 3},
			{"hello world", "xyz", -1},
			{"test", "test", 0},
			{"", "", 0},
			{"test", "", 0},
			{"", "test", -1},
		}

		for _, tt := range tests {
			got := indexOf(tt.s, tt.substr)
			if got != tt.want {
				t.Errorf("indexOf(%q, %q) = %v, want %v", tt.s, tt.substr, got, tt.want)
			}
		}
	})

	t.Run("isAcceptableServiceError", func(t *testing.T) {
		tests := []struct {
			err  error
			want bool
		}{
			{nil, true},
			{fmt.Errorf("connection refused"), true},
			{fmt.Errorf("timeout occurred"), true},
			{fmt.Errorf("no such host"), true},
			{fmt.Errorf("network unreachable"), true},
			{fmt.Errorf("404 Not Found"), true},
			{fmt.Errorf("500 Internal Server Error"), true},
			{fmt.Errorf("some other error"), false},
		}

		for _, tt := range tests {
			got := isAcceptableServiceError(tt.err)
			if got != tt.want {
				t.Errorf("isAcceptableServiceError(%v) = %v, want %v", tt.err, got, tt.want)
			}
		}
	})
}

// TestConfigurationHelpers tests configuration helper functions
func TestConfigurationHelpers(t *testing.T) {
	t.Run("TestConfig_IsValid", func(t *testing.T) {
		tests := []struct {
			name     string
			config   *TestConfig
			expected bool
		}{
			{
				name:     "valid config",
				config:   &TestConfig{Controller: "test", AccessToken: "token"},
				expected: true,
			},
			{
				name:     "empty controller",
				config:   &TestConfig{Controller: "", AccessToken: "token"},
				expected: false,
			},
			{
				name:     "empty token",
				config:   &TestConfig{Controller: "test", AccessToken: ""},
				expected: false,
			},
			{
				name:     "both empty",
				config:   &TestConfig{Controller: "", AccessToken: ""},
				expected: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.config.IsValid()
				if got != tt.expected {
					t.Errorf("IsValid() = %v, want %v", got, tt.expected)
				}
			})
		}
	})

	t.Run("NewTestConfigFromEnv_EmptyEnv", func(t *testing.T) {
		// Save original environment
		origController := os.Getenv("WNC_CONTROLLER")
		origToken := os.Getenv("WNC_ACCESS_TOKEN")
		defer func() {
			if origController != "" {
				os.Setenv("WNC_CONTROLLER", origController)
			} else {
				os.Unsetenv("WNC_CONTROLLER")
			}
			if origToken != "" {
				os.Setenv("WNC_ACCESS_TOKEN", origToken)
			} else {
				os.Unsetenv("WNC_ACCESS_TOKEN")
			}
		}()

		// Test with empty environment
		os.Unsetenv("WNC_CONTROLLER")
		os.Unsetenv("WNC_ACCESS_TOKEN")

		config := NewTestConfigFromEnv()
		if config != nil {
			t.Error("Expected nil config with empty environment")
		}
	})
}

// TestJSONHelpers tests JSON-related helper functions (updated with enhanced coverage)
func TestJSONHelpers(t *testing.T) {
	t.Run("TestJSONUnmarshal", func(t *testing.T) {
		type TestStruct struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		}

		jsonData := `{"name":"test","id":123}`
		var result TestStruct

		TestJSONUnmarshal(t, jsonData, &result, "TestStruct")
		if result.Name != "test" {
			t.Errorf("Expected name 'test', got %s", result.Name)
		}
		if result.ID != 123 {
			t.Errorf("Expected ID 123, got %d", result.ID)
		}
	})

	t.Run("TestJSONUnmarshalError", func(t *testing.T) {
		type TestStruct struct {
			Name string `json:"name"`
		}

		invalidJSON := `{"name": invalid}`
		var result TestStruct

		TestJSONUnmarshalError(t, invalidJSON, &result, "TestStruct")
	})

	t.Run("RunJSONTests", func(t *testing.T) {
		type TestData struct {
			Name string `json:"name"`
		}

		tests := []JSONTestCase{
			{
				Name:       "valid JSON",
				JSONData:   `{"name":"test"}`,
				Target:     &TestData{},
				TypeName:   "TestData",
				ShouldFail: false,
			},
		}

		RunJSONTests(t, tests)
	})

	t.Run("ValidateStructFields", func(t *testing.T) {
		type TestStruct struct {
			Name string
			ID   int
		}

		testStruct := TestStruct{}
		expectedFields := []string{"Name", "ID"}

		ValidateStructFields(t, testStruct, expectedFields)
	})

	t.Run("ValidateJSONTags", func(t *testing.T) {
		type TestStruct struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		}

		testStruct := TestStruct{}
		ValidateJSONTags(t, testStruct)
	})
}
