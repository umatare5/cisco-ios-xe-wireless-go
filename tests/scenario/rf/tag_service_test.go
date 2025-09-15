//go:build scenario

package tag

import (
	"testing"
	"time"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/scenario"
)

// Test constants
const (
	expectedRFTagDescription = "Test RF tag created by scenario test"
	testRF24GHzProfile       = "default-rf-profile-24ghz"
	testRF5GHzProfile        = "default-rf-profile-5ghz"
	testRF6GHzProfile        = "default-rf-profile-6ghz"
)

// generateTestRFTagName creates a unique test RF tag name with timestamp.
// Format: test-rf-tag-{yyyymmdd-hhmmss}
func generateTestRFTagName() string {
	timestamp := time.Now().Format("20060102-150405")
	return "test-rf-tag-" + timestamp
}

// TestRFTagScenario executes a comprehensive scenario test for RF tag management.
func TestRFTagServiceScenario_TagLifecycleManagement_Success(t *testing.T) {
	// Initialize scenario context
	tsc := scenario.SetupTag(t, generateTestRFTagName)

	// Initialize service
	service := rf.NewRFTagService(tsc.Client)

	// Execute standard 6-step workflow
	executeStandardRFTagWorkflow(t, tsc, service)
}

// executeStandardRFTagWorkflow performs the standard RF tag scenario
func executeStandardRFTagWorkflow(t *testing.T, tsc *scenario.TagContext, service *rf.RFTagService) {
	// Step 1: List initial tags and set count
	t.Logf("Step 1: Getting initial RF tags list")
	if err := tsc.SetInitialCount(func() (int, error) {
		tags, err := service.ListRFTags(tsc.Ctx)
		return len(tags), err
	}); err != nil {
		t.Fatalf("Failed to get initial RF tags count: %v", err)
	}

	// Step 2: Create test tag
	t.Logf("Step 2: Creating test RF tag: %s", tsc.TestTagName)
	if err := service.CreateRFTag(tsc.Ctx, &model.RFTag{
		TagName:     tsc.TestTagName,
		Description: expectedRFTagDescription,
	}); err != nil {
		t.Fatalf("Failed to create RF tag %s: %v", tsc.TestTagName, err)
	}
	t.Logf("Successfully created RF tag: %s", tsc.TestTagName)

	// Step 3: Configure RF profiles
	t.Logf("Step 3: Configuring RF profiles for test tag: %s", tsc.TestTagName)
	if err := configureRFTagProfiles(t, tsc, service); err != nil {
		t.Fatalf("Failed to configure RF tag profiles for %s: %v", tsc.TestTagName, err)
	}
	t.Logf("Successfully configured RF tag: %s", tsc.TestTagName)

	// Step 4: Get and validate configuration
	t.Logf("Step 4: Getting configuration for test tag: %s", tsc.TestTagName)
	retrievedConfig, err := service.GetRFTag(tsc.Ctx, tsc.TestTagName)
	if err != nil {
		t.Fatalf("Failed to retrieve RF tag configuration for %s: %v", tsc.TestTagName, err)
	}
	validateRFTagConfig(t, tsc, retrievedConfig)
	t.Logf("Successfully validated RF tag: %s", tsc.TestTagName)

	// Step 5: Delete test tag
	t.Logf("Step 5: Deleting RF tag: %s", tsc.TestTagName)
	if err := service.DeleteRFTag(tsc.Ctx, tsc.TestTagName); err != nil {
		t.Fatalf("Failed to delete RF tag %s: %v", tsc.TestTagName, err)
	}
	tsc.MarkTagDeleted()
	t.Logf("Successfully deleted RF tag: %s", tsc.TestTagName)

	// Step 6: Verify final tag count
	t.Logf("Step 6: Verifying final RF tags list")
	finalTags, err := service.ListRFTags(tsc.Ctx)
	if err != nil {
		t.Fatalf("Failed to list final RF tags: %v", err)
	}
	validateRFTagDeletion(t, tsc.TestTagName, finalTags)
	tsc.ValidateTagCount(func() (int, error) {
		tags, err := service.ListRFTags(tsc.Ctx)
		return len(tags), err
	}, len(finalTags))
}

// configureRFTagProfiles sets up RF profiles for the test tag
func configureRFTagProfiles(t *testing.T, tsc *scenario.TagContext, service *rf.RFTagService) error {
	// Configure 2.4GHz profile
	if err := service.SetDot11BRfProfile(tsc.Ctx, tsc.TestTagName, testRF24GHzProfile); err != nil {
		return err
	}

	// Configure 5GHz profile
	if err := service.SetDot11ARfProfile(tsc.Ctx, tsc.TestTagName, testRF5GHzProfile); err != nil {
		return err
	}

	// Configure 6GHz profile (may fail in some environments)
	if err := service.SetDot116GhzRFProfile(tsc.Ctx, tsc.TestTagName, testRF6GHzProfile); err != nil {
		t.Logf("Warning: Failed to set 6GHz profile (this might be expected): %v", err)
	}

	return nil
}

// validateRFTagConfig validates the retrieved RF tag configuration
func validateRFTagConfig(t *testing.T, tsc *scenario.TagContext, config *model.RFTag) {
	if config == nil {
		t.Fatalf("Expected RF tag configuration for %s, but got nil", tsc.TestTagName)
	}

	if config.TagName != tsc.TestTagName {
		t.Errorf("Expected tag name %s, got %s", tsc.TestTagName, config.TagName)
	}

	if config.Description != expectedRFTagDescription {
		t.Errorf("Expected description %s, got %s", expectedRFTagDescription, config.Description)
	}

	// Validate 2.4GHz profile
	if config.Dot11BRfProfileName != testRF24GHzProfile {
		t.Errorf("2.4GHz RF profile validation: expected '%s', got '%s'", testRF24GHzProfile, config.Dot11BRfProfileName)
	}

	// Validate 5GHz profile
	if config.Dot11ARfProfileName != testRF5GHzProfile {
		t.Errorf("5GHz RF profile validation: expected '%s', got '%s'", testRF5GHzProfile, config.Dot11ARfProfileName)
	}

	// Validate 6GHz profile (may be empty in some environments)
	if config.Dot116GhzRFProfName != testRF6GHzProfile {
		t.Logf("Warning: Expected 6GHz profile '%s', got '%s' (this might be expected in some environments due to the default profile doesn't show in the response)",
			testRF6GHzProfile, config.Dot116GhzRFProfName)
	}
}

// validateRFTagDeletion ensures the test tag was properly deleted
func validateRFTagDeletion(t *testing.T, testTagName string, finalTags []model.RFTag) {
	for _, tag := range finalTags {
		if tag.TagName == testTagName {
			t.Errorf("Test tag %s should have been deleted but still exists", testTagName)
		}
	}
}
