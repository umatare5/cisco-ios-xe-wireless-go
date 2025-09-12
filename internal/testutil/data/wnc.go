package data

import "os"

const DefaultTestAPMac = "28:ac:9e:bb:3c:80"

// TestConstants represents common test constants used across all WNC service tests.
type TestConstants struct {
	TestAPMac     string
	TestWtpMAC    string
	TestEthMAC    string
	TestRequestID string
	TestSlotID    int
	TestWlanID    int
	TestLocation  string
}

// StandardTestConstants returns default test constants for all WNC operations.
func StandardTestConstants() TestConstants {
	// Get TestAPMac from environment variable or use hardcoded fallback for unit tests
	testAPMac := os.Getenv("WNC_AP_MAC_ADDR")
	if testAPMac == "" {
		testAPMac = DefaultTestAPMac // Fallback for unit tests
	}

	return TestConstants{
		TestAPMac:     testAPMac,
		TestWtpMAC:    testAPMac,           // Use same MAC for WTP
		TestEthMAC:    "28:ac:9e:11:48:10", // TEST-AP01 ethernet MAC
		TestRequestID: "test-request-id",
		TestSlotID:    0,
		TestWlanID:    1,
		TestLocation:  "building-1",
	}
}
