package framework

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

// =============================================================================
// Generic Tag Test Framework
// =============================================================================

// TagConfig represents a generic tag configuration interface
type TagConfig interface {
	GetTagName() string
	SetTagName(string)
	Validate() error
}

// TagService represents a generic tag service interface
type TagService[T TagConfig] interface {
	CreateTag(ctx context.Context, config T) error
	GetTag(ctx context.Context, tagName string) (T, error)
	SetTag(ctx context.Context, config T) error
	DeleteTag(ctx context.Context, tagName string) error
	ListTags(ctx context.Context) ([]T, error)
}

// TagTestCase represents a generic tag test case
type TagTestCase[T TagConfig] struct {
	Name        string
	Client      *core.Client
	Config      T
	ExpectError bool
	ErrorMsg    string
}

// TagTestSuite represents a complete test suite for tag operations
type TagTestSuite[T TagConfig] struct {
	ServiceName   string
	TagType       string // "site", "policy", "rf"
	BaseEndpoint  string
	CreateConfig  func(tagName string) T
	CreateService func(*core.Client) TagService[T]
	TestConstants TagTestConstants
}

// TagTestConstants represents generic tag test constants interface
type TagTestConstants interface {
	GetTestTagName() string
	GetTestDescription() string
}

// =============================================================================
// Unified Unit Test Patterns
// =============================================================================

// TagUnitTestPattern provides unified unit test execution for tag operations
type TagUnitTestPattern struct {
	ServiceName string
	Operation   string // "Create", "Get", "Set", "Delete", "List"

	// Test functions that should return error with nil client
	NilClientTest func() error

	// Test functions that should return error with empty/invalid parameters
	EmptyParamTest func() error

	// Optional additional validation tests
	AdditionalTests []TagUnitValidationTest
}

// TagUnitValidationTest represents a single unit validation test
type TagUnitValidationTest struct {
	Name string
	Test func() error
}

// RunUnifiedTagUnitTests executes standardized unit tests for tag operations
func RunUnifiedTagUnitTests(t *testing.T, pattern TagUnitTestPattern) {
	t.Run("NilClient", func(t *testing.T) {
		err := pattern.NilClientTest()
		if err == nil {
			t.Errorf("%s with nil client should return error", pattern.Operation)
		}
		expectedErrMsg := "client cannot be nil"
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Errorf("Expected error containing '%s', got: %v", expectedErrMsg, err)
		}
	})

	t.Run("EmptyParameters", func(t *testing.T) {
		err := pattern.EmptyParamTest()
		if err == nil {
			t.Errorf("%s with empty parameters should return error", pattern.Operation)
		}
	})

	// Run additional tests if provided
	for _, test := range pattern.AdditionalTests {
		t.Run(test.Name, func(t *testing.T) {
			err := test.Test()
			if err != nil {
				t.Errorf("Additional test '%s' failed: %v", test.Name, err)
			}
		})
	}
}

// =============================================================================
// Unified Table-Driven Test Patterns
// =============================================================================

// TagTableDrivenTestPattern provides unified table-driven test execution
type TagTableDrivenTestPattern struct {
	ServiceName  string
	Operation    string
	TestCases    []TagTableDrivenTestCase
	SetupServer  func(testCase TagTableDrivenTestCase) interface{}
	ExecuteTest  func(client *core.Client, testCase TagTableDrivenTestCase) error
	ValidateTest func(t *testing.T, testCase TagTableDrivenTestCase, err error)
}

// TagTableDrivenTestCase represents a single table-driven test case
type TagTableDrivenTestCase struct {
	Name            string
	TagName         string
	ExpectError     bool
	ExpectedErrType string
	Description     string
	ServerResponse  string
	HTTPStatus      int
}

// RunUnifiedTagTableDrivenTests executes standardized table-driven tests
func RunUnifiedTagTableDrivenTests(t *testing.T, pattern TagTableDrivenTestPattern) {
	for _, tc := range pattern.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Setup test server
			server := pattern.SetupServer(tc)
			defer func() {
				if closer, ok := server.(interface{ Close() }); ok {
					closer.Close()
				}
			}()

			// Create client for this test case
			client, err := createTestClient(server)
			if err != nil {
				t.Fatalf("Failed to create test client: %v", err)
			}

			// Execute the test
			err = pattern.ExecuteTest(client, tc)

			// Validate the result
			pattern.ValidateTest(t, tc, err)
		})
	}
}

// =============================================================================
// Unified Integration Test Patterns
// =============================================================================

// TagIntegrationTestPattern provides unified integration test execution
type TagIntegrationTestPattern struct {
	ServiceName     string
	Operation       string
	TestCases       []TagIntegrationTestCase
	SetupClient     func() (*core.Client, error)
	ExecuteTest     func(client *core.Client, testCase TagIntegrationTestCase) error
	CleanupTest     func(client *core.Client, testCase TagIntegrationTestCase) error
	ValidateTest    func(t *testing.T, testCase TagIntegrationTestCase, err error)
	SkipCondition   func() bool
	TimeoutDuration time.Duration
}

// TagIntegrationTestCase represents a single integration test case
type TagIntegrationTestCase struct {
	Name         string
	TagName      string
	Description  string
	ExpectError  bool
	ExpectResult bool
	CleanupAfter bool
}

// RunUnifiedTagIntegrationTests executes standardized integration tests
func RunUnifiedTagIntegrationTests(t *testing.T, pattern TagIntegrationTestPattern) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	if pattern.SkipCondition != nil && pattern.SkipCondition() {
		t.Skip("Skipping integration tests - condition not met")
	}

	client, err := pattern.SetupClient()
	if err != nil {
		t.Fatalf("Failed to setup integration test client: %v", err)
	}

	for _, tc := range pattern.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Set timeout for integration test
			timeout := pattern.TimeoutDuration
			if timeout == 0 {
				timeout = 30 * time.Second
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			// Execute the test with timeout context
			done := make(chan error, 1)
			go func() {
				done <- pattern.ExecuteTest(client, tc)
			}()

			var testErr error
			select {
			case testErr = <-done:
				// Test completed
			case <-ctx.Done():
				testErr = fmt.Errorf("test timed out after %v", timeout)
			}

			// Validate the result
			pattern.ValidateTest(t, tc, testErr)

			// Cleanup if requested and test was successful
			if tc.CleanupAfter && (testErr == nil || !tc.ExpectError) {
				if cleanupErr := pattern.CleanupTest(client, tc); cleanupErr != nil {
					t.Logf("Cleanup failed for %s: %v", tc.Name, cleanupErr)
				}
			}
		})
	}
}

// createTestClient creates a test client from a server interface
func createTestClient(server interface{}) (*core.Client, error) {
	// This is a placeholder - actual implementation would depend on server type
	// In real usage, this would extract URL from server and create core.Client
	return nil, fmt.Errorf("createTestClient not implemented for server type %T", server)
}
