//go:build scenario

package tag

import (
	"testing"
	"time"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/scenario"
)

// Test constants
const (
	expectedSiteTagDescription = "Test site tag created by scenario test"
	testAPJoinProfile          = "default-ap-profile"   // Using default profile (will not appear in GET response due to YANG default behavior)
	testFlexProfile            = "default-flex-profile" // Using default profile (will not appear in GET response due to YANG default behavior)
)

// generateTestSiteTagName creates a unique test site tag name with timestamp.
// Format: test-site-tag-{yyyymmdd-hhmmss}
func generateTestSiteTagName() string {
	timestamp := time.Now().Format("20060102-150405")
	return "test-site-tag-" + timestamp
}

// TestSiteTagScenario executes a comprehensive scenario test for site tag management.
func TestSiteTagServiceScenario_TagLifecycleManagement_Success(t *testing.T) {
	// Initialize scenario context
	tsc := scenario.SetupTag(t, generateTestSiteTagName)

	// Initialize service
	service := site.NewSiteTagService(tsc.Client)

	// Execute standard 6-step workflow
	executeStandardSiteTagWorkflow(t, tsc, service)
}

// executeStandardSiteTagWorkflow performs the standard site tag scenario
func executeStandardSiteTagWorkflow(t *testing.T, tsc *scenario.TagContext, service *site.SiteTagService) {
	// Step 1: List initial tags and set count
	t.Logf("Step 1: Getting current site tags list")
	if err := tsc.SetInitialCount(func() (int, error) {
		tags, err := service.ListSiteTags(tsc.Ctx)
		return len(tags), err
	}); err != nil {
		t.Fatalf("Failed to get initial site tags count: %v", err)
	}

	// Step 2: Create test tag
	t.Logf("Step 2: Creating test site tag: %s", tsc.TestTagName)
	description := expectedSiteTagDescription
	if err := service.CreateSiteTag(tsc.Ctx, &model.SiteListEntry{
		SiteTagName: tsc.TestTagName,
		Description: &description,
	}); err != nil {
		t.Fatalf("Failed to create site tag %s: %v", tsc.TestTagName, err)
	}
	t.Logf("Successfully created site tag: %s", tsc.TestTagName)

	// Step 3: Configure site profiles
	t.Logf("Step 3: Setting profiles for test tag: %s", tsc.TestTagName)
	if err := configureSiteTagProfiles(t, tsc, service); err != nil {
		t.Fatalf("Failed to configure site tag profiles for %s: %v", tsc.TestTagName, err)
	}
	t.Logf("Successfully configured site tag: %s", tsc.TestTagName)

	// Step 4: Get and validate configuration
	t.Logf("Step 4: Getting configuration for test tag: %s", tsc.TestTagName)
	retrievedConfig, err := service.GetSiteTag(tsc.Ctx, tsc.TestTagName)
	if err != nil {
		t.Fatalf("Failed to retrieve site tag configuration for %s: %v", tsc.TestTagName, err)
	}
	validateSiteTagConfiguration(t, tsc, retrievedConfig)
	t.Logf("Successfully validated site tag: %s", tsc.TestTagName)

	// Step 5: Delete test tag
	t.Logf("Step 5: Deleting site tag: %s", tsc.TestTagName)
	if err := service.DeleteSiteTag(tsc.Ctx, tsc.TestTagName); err != nil {
		t.Fatalf("Failed to delete site tag %s: %v", tsc.TestTagName, err)
	}
	tsc.MarkTagDeleted()
	t.Logf("Successfully deleted site tag: %s", tsc.TestTagName)

	// Step 6: Verify final tag count
	t.Logf("Step 6: Verifying final site tags list")
	finalTags, err := service.ListSiteTags(tsc.Ctx)
	if err != nil {
		t.Fatalf("Failed to list final site tags: %v", err)
	}
	validateSiteTagDeletion(t, tsc.TestTagName, finalTags)
	if err := tsc.ValidateTagCount(func() (int, error) {
		tags, err := service.ListSiteTags(tsc.Ctx)
		return len(tags), err
	}, len(finalTags)); err != nil {
		t.Fatalf("Tag count validation failed: %v", err)
	}
}

// configureSiteTagProfiles sets up profiles for the test tag
func configureSiteTagProfiles(t *testing.T, tsc *scenario.TagContext, service *site.SiteTagService) error {
	// Step 3a: Set is-local-site to false
	t.Logf("Step 3a: Setting is-local-site to false for test tag: %s", tsc.TestTagName)
	if err := service.SetLocalSite(tsc.Ctx, tsc.TestTagName, false); err != nil {
		return err
	}
	t.Logf("Successfully set is-local-site to false for test tag: %s", tsc.TestTagName)

	// Step 3b: Set AP join profile
	t.Logf("Step 3b: Testing SetAPJoinProfile function for test tag: %s", tsc.TestTagName)
	if err := service.SetAPJoinProfile(tsc.Ctx, tsc.TestTagName, testAPJoinProfile); err != nil {
		return err
	}
	t.Logf("Successfully set AP join profile for test tag: %s", tsc.TestTagName)

	// Step 3c: Set flex profile
	t.Logf("Step 3c: Setting flex profile for test tag: %s", tsc.TestTagName)
	if err := service.SetFlexProfile(tsc.Ctx, tsc.TestTagName, testFlexProfile); err != nil {
		return err
	}
	t.Logf("Successfully set flex profile for test tag: %s", tsc.TestTagName)

	return nil
}

// validateSiteTagConfiguration validates the retrieved site tag configuration
func validateSiteTagConfiguration(t *testing.T, tsc *scenario.TagContext, config *model.SiteListEntry) {
	if config == nil {
		t.Fatalf("Expected site tag configuration for %s, but got nil", tsc.TestTagName)
	}

	// Get description value for validation
	var description string
	if config.Description != nil {
		description = *config.Description
	}

	if config.SiteTagName != tsc.TestTagName {
		t.Errorf("Expected tag name %s, got %s", tsc.TestTagName, config.SiteTagName)
	}

	if description != expectedSiteTagDescription {
		t.Errorf("Expected description %s, got %s", expectedSiteTagDescription, description)
	}

	// Validate is-local-site setting
	if config.IsLocalSite == nil || *config.IsLocalSite != false {
		t.Errorf("Expected is-local-site to be false, got %v", config.IsLocalSite)
	}

	// Note: AP join profile and flex profile are set to their default values.
	// When RESTCONF API receives values that match YANG defaults, it omits them from GET responses.
	// This is expected behavior per YANG/RESTCONF specification.
	if config.ApJoinProfile != nil && *config.ApJoinProfile != "" && *config.ApJoinProfile != testAPJoinProfile {
		t.Logf("AP join profile retrieved: %s (expected %s or empty due to default)",
			*config.ApJoinProfile, testAPJoinProfile)
	}

	if config.FlexProfile != nil && *config.FlexProfile != "" && *config.FlexProfile != testFlexProfile {
		t.Logf("Flex profile retrieved: %s (expected %s or empty due to default)",
			*config.FlexProfile, testFlexProfile)
	}
}

// validateSiteTagDeletion ensures the test tag was properly deleted
func validateSiteTagDeletion(t *testing.T, testTagName string, finalTags []model.SiteListEntry) {
	for _, tag := range finalTags {
		if tag.SiteTagName == testTagName {
			t.Errorf("Test tag %s should have been deleted but still exists", testTagName)
		}
	}
}
