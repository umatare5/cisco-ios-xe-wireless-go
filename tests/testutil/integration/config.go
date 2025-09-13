//go:build integration

package integration

import (
	"os"
	"testing"
)

// Config provides access to integration test environment configuration.
type Config struct {
	Controller          string
	AccessToken         string
	TestAPMac           string
	TestClientMac       string
	TestAPWlanBSSID     string
	TestAPNeighborBSSID string
}

// APMac returns the test AP MAC address.
func (c *Config) APMac() string {
	return c.TestAPMac
}

// ClientMac returns the test client MAC address.
func (c *Config) ClientMac() string {
	return c.TestClientMac
}

// APWlanBSSID returns the test AP WLAN BSSID.
func (c *Config) APWlanBSSID() string {
	return c.TestAPWlanBSSID
}

// APNeighborBSSID returns the test AP neighbor BSSID.
func (c *Config) APNeighborBSSID() string {
	return c.TestAPNeighborBSSID
}

// LoadConfig loads integration test environment configuration.
func LoadConfig() Config {
	return Config{
		Controller:          os.Getenv("WNC_CONTROLLER"),
		AccessToken:         os.Getenv("WNC_ACCESS_TOKEN"),
		TestAPMac:           os.Getenv("WNC_AP_MAC_ADDR"),
		TestClientMac:       os.Getenv("WNC_CLIENT_MAC_ADDR"),
		TestAPWlanBSSID:     os.Getenv("WNC_AP_WLAN_BSSID"),
		TestAPNeighborBSSID: os.Getenv("WNC_AP_NEIGHBOR_BSSID"),
	}
}

// RequireConfig validates that all required integration test configuration is available.
// Fails the test if any required environment variable is missing.
func RequireConfig(t *testing.T) {
	t.Helper()
	var missing []string

	if os.Getenv("WNC_CONTROLLER") == "" {
		missing = append(missing, "WNC_CONTROLLER")
	}
	if os.Getenv("WNC_ACCESS_TOKEN") == "" {
		missing = append(missing, "WNC_ACCESS_TOKEN")
	}
	if os.Getenv("WNC_AP_MAC_ADDR") == "" {
		missing = append(missing, "WNC_AP_MAC_ADDR")
	}
	if os.Getenv("WNC_CLIENT_MAC_ADDR") == "" {
		missing = append(missing, "WNC_CLIENT_MAC_ADDR")
	}
	if os.Getenv("WNC_AP_WLAN_BSSID") == "" {
		missing = append(missing, "WNC_AP_WLAN_BSSID")
	}
	if os.Getenv("WNC_AP_NEIGHBOR_BSSID") == "" {
		missing = append(missing, "WNC_AP_NEIGHBOR_BSSID")
	}

	if len(missing) > 0 {
		t.Fatalf("Required integration test configuration is not available: %v", missing)
	}
}

// TestAPMac returns the test AP MAC address from environment variable.
// Should only be called after RequireConfig() validation.
func TestAPMac() string {
	return os.Getenv("WNC_AP_MAC_ADDR")
}

// TestClientMac returns the test client MAC address from environment variable.
// Should only be called after RequireConfig() validation.
func TestClientMac() string {
	return os.Getenv("WNC_CLIENT_MAC_ADDR")
}

// TestAPWlanBSSID returns the test AP WLAN BSSID from environment variable.
// Should only be called after RequireConfig() validation.
func TestAPWlanBSSID() string {
	return os.Getenv("WNC_AP_WLAN_BSSID")
}

// TestAPNeighborBSSID returns the test AP neighbor BSSID from environment variable.
// Should only be called after RequireConfig() validation.
func TestAPNeighborBSSID() string {
	return os.Getenv("WNC_AP_NEIGHBOR_BSSID")
}
