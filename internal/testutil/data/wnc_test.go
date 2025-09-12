package data

import (
	"os"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
)

func TestTestUtilDataUnit_StandardConstants_Success(t *testing.T) {
	constants := StandardTestConstants()

	helper.AssertStringNotEmpty(t, constants.TestAPMac, "TestAPMac")
	helper.AssertStringNotEmpty(t, constants.TestRequestID, "TestRequestID")
	if constants.TestSlotID < 0 {
		helper.AssertTrue(t, constants.TestSlotID >= 0, "TestSlotID should not be negative")
	}
	if constants.TestWlanID <= 0 {
		helper.AssertTrue(t, constants.TestWlanID > 0, "TestWlanID should be positive")
	}
	helper.AssertStringNotEmpty(t, constants.TestLocation, "TestLocation")

	// Verify expected default values (environment variable or fallback)
	expectedAPMac := os.Getenv("WNC_AP_MAC_ADDR")
	if expectedAPMac == "" {
		expectedAPMac = DefaultTestAPMac // Fallback value
	}
	expectedRequestID := "test-request-id"
	expectedSlotID := 0
	expectedWlanID := 1
	expectedLocation := "building-1"

	helper.AssertStringEquals(t, constants.TestAPMac, expectedAPMac, "TestAPMac")
	helper.AssertStringEquals(t, constants.TestRequestID, expectedRequestID, "TestRequestID")
	helper.AssertIntEquals(t, constants.TestSlotID, expectedSlotID, "TestSlotID")
	helper.AssertIntEquals(t, constants.TestWlanID, expectedWlanID, "TestWlanID")
	helper.AssertStringEquals(t, constants.TestLocation, expectedLocation, "TestLocation")
}

func TestTestUtilDataUnit_ConstantsStructure_Success(t *testing.T) {
	constants := StandardTestConstants()

	// Test that the struct has all expected fields
	var structType TestConstants
	structType.TestAPMac = constants.TestAPMac
	structType.TestRequestID = constants.TestRequestID
	structType.TestSlotID = constants.TestSlotID
	structType.TestWlanID = constants.TestWlanID
	structType.TestLocation = constants.TestLocation

	helper.AssertStringEquals(t, structType.TestAPMac, constants.TestAPMac, "TestConstants struct field TestAPMac")
	helper.AssertStringEquals(t, structType.TestRequestID, constants.TestRequestID,
		"TestConstants struct field TestRequestID")
	helper.AssertIntEquals(t, structType.TestSlotID, constants.TestSlotID, "TestConstants struct field TestSlotID")
	helper.AssertIntEquals(t, structType.TestWlanID, constants.TestWlanID, "TestConstants struct field TestWlanID")
	helper.AssertStringEquals(t, structType.TestLocation, constants.TestLocation,
		"TestConstants struct field TestLocation")
}

func TestTestUtilDataUnit_ConstantsValidation_Success(t *testing.T) {
	constants := StandardTestConstants()

	// Validate MAC address format (basic check)
	if len(constants.TestAPMac) != 17 { // "a4:8c:db:ac:ba:20" format
		helper.AssertIntEquals(t, len(constants.TestAPMac), 17, "TestAPMac should be in MAC address format")
	}

	// Validate slot ID range (typical slot IDs are 0-3)
	if constants.TestSlotID < 0 || constants.TestSlotID > 3 {
		helper.AssertTrue(t, constants.TestSlotID >= 0 && constants.TestSlotID <= 3,
			"TestSlotID should be in valid range (0-3)")
	}

	// Validate WLAN ID range (typical WLAN IDs are 1-512)
	if constants.TestWlanID < 1 || constants.TestWlanID > 512 {
		helper.AssertTrue(t, constants.TestWlanID >= 1 && constants.TestWlanID <= 512,
			"TestWlanID should be in valid range (1-512)")
	}
}

func TestTestUtilDataUnit_AllConstantsCompletion_Success(t *testing.T) {
	// Ensure all constant types can be instantiated
	testConstants := StandardTestConstants()
	wlanTagConstants := StandardWLANTagTestConstants()
	rfTagConstants := StandardRFTagTestConstants()
	siteTagConstants := StandardSiteTagTestConstants()

	// Basic non-nil checks
	helper.AssertStringNotEmpty(t, testConstants.TestAPMac, "StandardTestConstants")
	helper.AssertStringNotEmpty(t, wlanTagConstants.TestTagName, "StandardWLANTagTestConstants")
	helper.AssertStringNotEmpty(t, rfTagConstants.TestTagName, "StandardRFTagTestConstants")
	helper.AssertStringNotEmpty(t, siteTagConstants.TestTagName, "StandardSiteTagTestConstants")
}
