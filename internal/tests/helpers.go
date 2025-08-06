package tests

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// TestClient creates a test client using environment variables
func TestClient(t *testing.T) *wnc.Client {
	t.Helper()

	controller := os.Getenv("WNC_CONTROLLER")
	token := os.Getenv("WNC_ACCESS_TOKEN")

	if controller == "" || token == "" {
		t.Skip("WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set for integration tests")
	}

	client, err := wnc.New(controller, token,
		wnc.WithTimeout(30*time.Second),
		wnc.WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create test client: %v", err)
	}

	return client
}

// CreateTestClientFromEnv creates a test client from environment variables
// This is an alias for TestClient to match expected API
func CreateTestClientFromEnv(t *testing.T) *wnc.Client {
	return TestClient(t)
}

// DefaultTestTimeout is the default timeout for tests
const DefaultTestTimeout = 30 * time.Second

// ExtendedTestTimeout is an extended timeout for longer tests
const ExtendedTestTimeout = 60 * time.Second

// TestDataDir is the directory for test data files
const TestDataDir = "./test_data"

// TestContext creates a test context with timeout
func TestContext(t *testing.T) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(cancel)
	return ctx
}

// SkipIfNoConnection skips the test if no network connection to WNC
func SkipIfNoConnection(t *testing.T, client *wnc.Client) {
	t.Helper()

	// Try a simple health check - this assumes there's some basic endpoint available
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// We'll use a very basic endpoint that should exist on most controllers
	var result interface{}
	err := client.Do(ctx, "GET", "/yang-library-version", &result)
	if err != nil {
		t.Skipf("No connection to WNC controller: %v", err)
	}
}

// SaveTestDataToFile saves test data to a JSON file
func SaveTestDataToFile(filename string, data interface{}) error {
	// Create test_data directory if it doesn't exist
	if err := os.MkdirAll(TestDataDir, 0755); err != nil {
		return err
	}

	// Create the full file path
	fullPath := filepath.Join(TestDataDir, filename)

	// Marshal data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(fullPath, jsonData, 0644)
}
