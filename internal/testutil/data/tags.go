package data

// =============================================================================
// Service-Specific Tag Test Constants
// =============================================================================

// WLANTagTestConstants represents WLAN/policy tag-specific test constants.
type WLANTagTestConstants struct {
	TestTagName           string
	TestDescription       string
	TestWLANProfileName   string
	TestPolicyProfileName string
}

// StandardWLANTagTestConstants returns default test constants for WLAN tag operations.
func StandardWLANTagTestConstants() WLANTagTestConstants {
	return WLANTagTestConstants{
		TestTagName:           "test-pol-tag",
		TestDescription:       "Test policy tag description",
		TestWLANProfileName:   "test-wlan",
		TestPolicyProfileName: "default-policy-profile",
	}
}

// RFTagTestConstants represents RF tag-specific test constants.
type RFTagTestConstants struct {
	TestTagName             string
	TestDescription         string
	TestDot11ARfProfileName string
	TestDot11BRfProfileName string
	TestDot116GhzRfProfName string
}

// StandardRFTagTestConstants returns default test constants for RF tag operations.
func StandardRFTagTestConstants() RFTagTestConstants {
	return RFTagTestConstants{
		TestTagName:             "test-rf-tag",
		TestDescription:         "Test RF tag description",
		TestDot11ARfProfileName: "Typical_Client_Density_rf_24gh",
		TestDot11BRfProfileName: "Typical_Client_Density_rf_5gh",
		TestDot116GhzRfProfName: "default-rf-profile-6ghz",
	}
}

// SiteTagTestConstants represents site tag-specific test constants.
type SiteTagTestConstants struct {
	TestTagName     string
	TestDescription string
	TestFloorRFMap  string
}

// StandardSiteTagTestConstants returns default test constants for site tag operations.
func StandardSiteTagTestConstants() SiteTagTestConstants {
	return SiteTagTestConstants{
		TestTagName:     "test-site-tag",
		TestDescription: "Test site tag description",
		TestFloorRFMap:  "test-floor-map",
	}
}
