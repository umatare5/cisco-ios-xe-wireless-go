// Package testhelpers provides common test helper functions to avoid circular imports
package testhelpers

import (
	"os"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// CreateTestClientFromEnv creates a test client using environment variables
func CreateTestClientFromEnv(t *testing.T) *wnc.Client {
	t.Helper()

	config := testutil.NewTestConfigFromEnv()
	if config == nil {
		t.Skip("Required environment variables not set - skipping test")
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

// CreateTestClientWithTimeout creates a test client with custom timeout
func CreateTestClientWithTimeout(t *testing.T, timeout time.Duration) *wnc.Client {
	t.Helper()

	config := testutil.NewTestConfigFromEnv()
	if config == nil {
		t.Skip("Required environment variables not set - skipping test")
	}

	config.Timeout = timeout

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

// GetTestCredentials returns test credentials from environment variables
func GetTestCredentials(t *testing.T) (controller, accessToken string) {
	t.Helper()

	controller = os.Getenv("WNC_CONTROLLER")
	accessToken = os.Getenv("WNC_ACCESS_TOKEN")

	if controller == "" || accessToken == "" {
		t.Skip("Required environment variables not set - skipping test")
	}

	return controller, accessToken
}
