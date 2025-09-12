package client

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

// Test helper variables for flexible testing behavior.
var (
	createCoreClient   = core.New            // Allows test to inject custom client creation behavior
	failOnClientError  = true                // Controls whether TestClient fatals (true) or skips (false) on error
	simulateFatalAsLog = false               // If true, logs instead of calling t.Fatalf for coverage testing
	testFatalf         = (*testing.T).Fatalf // Hook for testing fatal path without aborting test suite
)

// ServiceSetup holds the common test setup components.
type ServiceSetup struct {
	Client  *core.Client
	Context context.Context
}

// IntegrationTestEnv provides access to integration test environment variables.
type IntegrationTestEnv struct {
	Controller  string
	AccessToken string
	TestAPMac   string
}

// TestClient creates a test client using environment variables (original behavior retained).
func TestClient(t *testing.T) *core.Client {
	t.Helper()

	controller := os.Getenv("WNC_CONTROLLER")
	token := os.Getenv("WNC_ACCESS_TOKEN")

	if controller == "" || token == "" {
		t.Skip("WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set for integration tests")
	}

	client, err := createCoreClient(controller, token,
		core.WithTimeout(5*time.Second),
		core.WithInsecureSkipVerify(true))
	if err == nil {
		return client
	}

	if !failOnClientError {
		t.Skipf("Failed to create test client (downgraded to skip for coverage): %v", err)
		return nil
	}

	if simulateFatalAsLog {
		t.Logf("(simulated fatal) Failed to create test client: %v", err)
		return nil
	}

	testFatalf(t, "Failed to create test client: %v", err)
	return nil
}

// OptionalTestClient returns a *core.Client if required env vars are set, otherwise nil without skipping.
// This enables tests to obtain a client opportunistically without triggering a Skip or Fatal path,
// reducing duplicated env checks across service tests while preserving existing coverage behavior.
func OptionalTestClient(t *testing.T) *core.Client {
	t.Helper()
	controller := os.Getenv("WNC_CONTROLLER")
	token := os.Getenv("WNC_ACCESS_TOKEN")

	if controller == "" || token == "" {
		return nil
	}

	client, err := createCoreClient(
		controller,
		token,
		core.WithTimeout(5*time.Second),
		core.WithInsecureSkipVerify(true))
	if err != nil {
		t.Logf("OptionalTestClient: failed to create client: %v", err)
		return nil
	}

	return client
}

// SetupOptionalClient creates a ServiceSetup with optional client and appropriate context.
// This function does not skip tests if no client is available, making it suitable for both unit and integration tests.
func SetupOptionalClient(t *testing.T) ServiceSetup {
	t.Helper()

	client := OptionalTestClient(t)
	if client == nil {
		return ServiceSetup{
			Client:  nil,
			Context: context.Background(),
		}
	}

	return ServiceSetup{
		Client:  client,
		Context: testutil.TestContext(t),
	}
}

// SetupRequiredClient creates a ServiceSetup for integration tests that require a client.
// This function skips the test if no client is available, making it suitable for integration tests.
func SetupRequiredClient(t *testing.T) ServiceSetup {
	t.Helper()

	setup := SetupOptionalClient(t)
	if setup.Client == nil {
		t.Skip("Skipping integration tests: no client available")
	}
	return setup
}

// SkipIfNoConnection skips in absence of connectivity; no-ops for nil client.
func SkipIfNoConnection(t *testing.T, client *core.Client) {
	t.Helper()

	if client == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Do(ctx, http.MethodGet, "Cisco-IOS-XE-wireless-general-oper:general-oper-data"); err != nil {
		t.Skipf("No connection to WNC controller: %v", err)
	}
}

// LoadIntegrationEnv loads integration test environment configuration.
func LoadIntegrationEnv() IntegrationTestEnv {
	return IntegrationTestEnv{
		Controller:  os.Getenv("WNC_CONTROLLER"),
		AccessToken: os.Getenv("WNC_ACCESS_TOKEN"),
		TestAPMac:   os.Getenv("WNC_AP_MAC_ADDR"),
	}
}

// TestAPMac returns the test AP MAC address from environment variable with fallback.
// This function provides integration with unit tests by falling back to a default value
// when the environment variable is not set.
func TestAPMac() string {
	if testAPMac := os.Getenv("WNC_AP_MAC_ADDR"); testAPMac != "" {
		return testAPMac
	}
	return "28:ac:9e:bb:3c:80" // Default fallback for unit tests
}
