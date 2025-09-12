//go:build scenario

package tag

import (
	"testing"
	"time"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/site"
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
	tsc, skip := client.InitializeTagScenario(t, generateTestSiteTagName)
	if skip {
		return
	}

	// Initialize service
	service := site.NewSiteTagService(tsc.Client)

	// Cleanup on exit
	defer tsc.CreateCleanupFunction(service.DeleteSiteTag)()

	// Execute standard 6-step workflow
	executeStandardSiteTagWorkflow(t, tsc, service)
}

// executeStandardSiteTagWorkflow performs the standard site tag scenario
func executeStandardSiteTagWorkflow(t *testing.T, tsc *client.TagScenarioContext, service *site.SiteTagService) {
	// Step 1: List initial tags
	initialTags := client.ExecuteStepWithResult(tsc, client.StepInitialList,
		"Getting current site tags list", func() ([]model.SiteListEntry, error) {
			return service.ListSiteTags(tsc.Ctx)
		})
	tsc.SetInitialCount(len(initialTags))

	// Step 2: Create test tag
	tsc.ExecuteStep(client.StepCreate, "Creating test site tag", func() error {
		description := expectedSiteTagDescription
		return service.CreateSiteTag(tsc.Ctx, &model.SiteListEntry{
			SiteTagName: tsc.TestTagName,
			Description: &description,
		})
	})
	tsc.LogSuccess("created")

	// Step 3: Configure site profiles
	tsc.ExecuteStep(client.StepConfigure, "Setting profiles for test tag", func() error {
		return configureSiteTagProfiles(t, tsc, service)
	})
	tsc.LogSuccess("configured")

	// Step 4: Get and validate configuration
	retrievedConfig := client.ExecuteStepWithResult(tsc, client.StepGet,
		"Getting configuration for test tag", func() (*model.SiteListEntry, error) {
			return service.GetSiteTag(tsc.Ctx, tsc.TestTagName)
		})
	validateSiteTagConfiguration(t, tsc, retrievedConfig)
	tsc.LogSuccess("validated")

	// Step 5: Delete test tag
	tsc.ExecuteDeleteStep("Deleting test site tag", func() error {
		return service.DeleteSiteTag(tsc.Ctx, tsc.TestTagName)
	})
	tsc.LogSuccess("deleted")

	// Step 6: List final tags and verify deletion
	finalTags := client.ExecuteStepWithResult(tsc, client.StepFinalList,
		"Getting final site tags list", func() ([]model.SiteListEntry, error) {
			return service.ListSiteTags(tsc.Ctx)
		})
	validateSiteTagDeletion(t, tsc.TestTagName, finalTags)
	tsc.ValidateTagCount(len(finalTags))
}

// configureSiteTagProfiles sets up profiles for the test tag
func configureSiteTagProfiles(t *testing.T, tsc *client.TagScenarioContext, service *site.SiteTagService) error {
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
func validateSiteTagConfiguration(t *testing.T, tsc *client.TagScenarioContext, config *model.SiteListEntry) {
	client.ValidateConfigNotNil(t, config, tsc.TestTagName)

	// Get description value for validation
	var description string
	if config.Description != nil {
		description = *config.Description
	}
	client.ValidateCommonTagFields(t, tsc.TestTagName, expectedSiteTagDescription,
		config.SiteTagName, description)

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
