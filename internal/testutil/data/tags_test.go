package data

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
)

func TestTestUtilDataUnit_WLANTagConstants_Success(t *testing.T) {
	constants := StandardWLANTagTestConstants()

	helper.AssertStringNotEmpty(t, constants.TestTagName, "TestTagName should not be empty")
	helper.AssertStringNotEmpty(t, constants.TestDescription, "TestDescription should not be empty")
	helper.AssertStringNotEmpty(t, constants.TestWLANProfileName, "TestWLANProfileName should not be empty")
	if constants.TestPolicyProfileName == "" {
		helper.AssertStringNotEmpty(t, constants.TestPolicyProfileName, "TestPolicyProfileName should not be empty")
	}

	// Verify expected default values
	expectedTagName := "test-pol-tag"
	expectedDescription := "Test policy tag description"
	expectedWLANProfile := "test-wlan"
	expectedPolicyProfile := "default-policy-profile"

	if constants.TestTagName != expectedTagName {
		helper.AssertStringEquals(t, constants.TestTagName, expectedTagName, "TestTagName mismatch")
	}
	if constants.TestDescription != expectedDescription {
		helper.AssertStringEquals(t, constants.TestDescription, expectedDescription, "TestDescription mismatch")
	}
	if constants.TestWLANProfileName != expectedWLANProfile {
		helper.AssertStringEquals(t, constants.TestWLANProfileName, expectedWLANProfile,
			"TestWLANProfileName mismatch")
	}
	if constants.TestPolicyProfileName != expectedPolicyProfile {
		helper.AssertStringEquals(t, constants.TestPolicyProfileName, expectedPolicyProfile,
			"TestPolicyProfileName mismatch")
	}
}

func TestTestUtilDataUnit_RFTagConstants_Success(t *testing.T) {
	constants := StandardRFTagTestConstants()

	helper.AssertStringNotEmpty(t, constants.TestTagName, "TestTagName should not be empty")
	helper.AssertStringNotEmpty(t, constants.TestDescription, "TestDescription should not be empty")
	helper.AssertStringNotEmpty(t, constants.TestDot11ARfProfileName, "TestDot11ARfProfileName should not be empty")
	helper.AssertStringNotEmpty(t, constants.TestDot11BRfProfileName, "TestDot11BRfProfileName should not be empty")
	helper.AssertStringNotEmpty(t, constants.TestDot116GhzRfProfName, "TestDot116GhzRfProfName should not be empty")

	// Verify expected default values
	expectedTagName := "test-rf-tag"
	expectedDescription := "Test RF tag description"
	expected24GhzProfile := "Typical_Client_Density_rf_24gh"
	expected5GhzProfile := "Typical_Client_Density_rf_5gh"
	expected6GhzProfile := "default-rf-profile-6ghz"

	if constants.TestTagName != expectedTagName {
		helper.AssertStringEquals(t, constants.TestTagName, expectedTagName, "TestTagName mismatch")
	}
	if constants.TestDescription != expectedDescription {
		helper.AssertStringEquals(t, constants.TestDescription, expectedDescription, "TestDescription mismatch")
	}
	if constants.TestDot11ARfProfileName != expected24GhzProfile {
		helper.AssertStringEquals(t, constants.TestDot11ARfProfileName, expected24GhzProfile,
			"TestDot11ARfProfileName mismatch")
	}
	if constants.TestDot11BRfProfileName != expected5GhzProfile {
		helper.AssertStringEquals(t, constants.TestDot11BRfProfileName, expected5GhzProfile,
			"TestDot11BRfProfileName mismatch")
	}
	if constants.TestDot116GhzRfProfName != expected6GhzProfile {
		helper.AssertStringEquals(t, constants.TestDot116GhzRfProfName, expected6GhzProfile,
			"TestDot116GhzRfProfName mismatch")
	}
}

func TestTestUtilDataUnit_SiteTagConstants_Success(t *testing.T) {
	constants := StandardSiteTagTestConstants()

	helper.AssertStringNotEmpty(t, constants.TestTagName, "TestTagName should not be empty")
	helper.AssertStringNotEmpty(t, constants.TestDescription, "TestDescription should not be empty")
	helper.AssertStringNotEmpty(t, constants.TestFloorRFMap, "TestFloorRFMap should not be empty")

	// Verify expected default values
	expectedTagName := "test-site-tag"
	expectedDescription := "Test site tag description"
	expectedFloorRFMap := "test-floor-map"

	if constants.TestTagName != expectedTagName {
		helper.AssertStringEquals(t, constants.TestTagName, expectedTagName, "TestTagName mismatch")
	}
	if constants.TestDescription != expectedDescription {
		helper.AssertStringEquals(t, constants.TestDescription, expectedDescription, "TestDescription mismatch")
	}
	if constants.TestFloorRFMap != expectedFloorRFMap {
		helper.AssertStringEquals(t, constants.TestFloorRFMap, expectedFloorRFMap, "TestFloorRFMap mismatch")
	}
}

func TestTestUtilDataUnit_WLANTagStructure_Success(t *testing.T) {
	constants := StandardWLANTagTestConstants()

	// Test that the struct has all expected fields
	var structType WLANTagTestConstants
	structType.TestTagName = constants.TestTagName
	structType.TestDescription = constants.TestDescription
	structType.TestWLANProfileName = constants.TestWLANProfileName
	structType.TestPolicyProfileName = constants.TestPolicyProfileName

	helper.AssertStringEquals(t, structType.TestTagName, constants.TestTagName,
		"WLANTagTestConstants struct field TestTagName")
}

func TestTestUtilDataUnit_RFTagStructure_Success(t *testing.T) {
	constants := StandardRFTagTestConstants()

	// Test that the struct has all expected fields
	var structType RFTagTestConstants
	structType.TestTagName = constants.TestTagName
	structType.TestDescription = constants.TestDescription
	structType.TestDot11ARfProfileName = constants.TestDot11ARfProfileName
	structType.TestDot11BRfProfileName = constants.TestDot11BRfProfileName
	structType.TestDot116GhzRfProfName = constants.TestDot116GhzRfProfName

	helper.AssertStringEquals(t, structType.TestTagName, constants.TestTagName,
		"RFTagTestConstants struct field TestTagName")
}

func TestTestUtilDataUnit_SiteTagStructure_Success(t *testing.T) {
	constants := StandardSiteTagTestConstants()

	// Test that the struct has all expected fields
	var structType SiteTagTestConstants
	structType.TestTagName = constants.TestTagName
	structType.TestDescription = constants.TestDescription
	structType.TestFloorRFMap = constants.TestFloorRFMap

	helper.AssertStringEquals(t, structType.TestTagName, constants.TestTagName,
		"SiteTagTestConstants struct field TestTagName")
}
