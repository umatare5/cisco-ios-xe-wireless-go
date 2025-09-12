//go:build scenario

package tag

import (
	"testing"
	"time"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rf"
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
	tsc, skip := client.InitializeTagScenario(t, generateTestRFTagName)
	if skip {
		return
	}

	// Initialize service
	service := rf.NewRFTagService(tsc.Client)

	// Cleanup on exit
	defer tsc.CreateCleanupFunction(service.DeleteRFTag)()

	// Execute standard 6-step workflow
	executeStandardRFTagWorkflow(t, tsc, service)
}

// executeStandardRFTagWorkflow performs the standard RF tag scenario
func executeStandardRFTagWorkflow(t *testing.T, tsc *client.TagScenarioContext, service *rf.RFTagService) {
	// Step 1: List initial tags
	initialTags := client.ExecuteStepWithResult(tsc, client.StepInitialList,
		"Getting initial RF tags list", func() ([]model.RfTag, error) {
			return service.ListRFTags(tsc.Ctx)
		})
	initialCount := len(initialTags)
	tsc.SetInitialCount(initialCount)

	// Step 2: Create test tag
	tsc.ExecuteStep(client.StepCreate, "Creating test RF tag", func() error {
		return service.CreateRFTag(tsc.Ctx, &model.RfTag{
			TagName:     tsc.TestTagName,
			Description: expectedRFTagDescription,
		})
	})
	tsc.LogSuccess("created")

	// Step 3: Configure RF profiles
	tsc.ExecuteStep(client.StepConfigure, "Configuring RF profiles for test tag", func() error {
		return configureRFTagProfiles(t, tsc, service)
	})
	tsc.LogSuccess("configured")

	// Step 4: Get and validate configuration
	retrievedConfig := client.ExecuteStepWithResult(tsc, client.StepGet,
		"Getting configuration for test tag", func() (*model.RfTag, error) {
			return service.GetRFTag(tsc.Ctx, tsc.TestTagName)
		})
	validateRFTagConfig(t, tsc, retrievedConfig)
	tsc.LogSuccess("validated")

	// Step 5: Delete test tag
	tsc.ExecuteDeleteStep("Deleting test RF tag", func() error {
		return service.DeleteRFTag(tsc.Ctx, tsc.TestTagName)
	})
	tsc.LogSuccess("deleted")

	// Step 6: Verify deletion
	finalTags := client.ExecuteStepWithResult(tsc, client.StepFinalList,
		"Getting final RF tags list", func() ([]model.RfTag, error) {
			return service.ListRFTags(tsc.Ctx)
		})
	finalCount := len(finalTags)
	validateRFTagDeletion(t, tsc.TestTagName, finalTags)
	tsc.ValidateTagCount(finalCount)
}

// configureRFTagProfiles sets up RF profiles for the test tag
func configureRFTagProfiles(t *testing.T, tsc *client.TagScenarioContext, service *rf.RFTagService) error {
	// Configure 2.4GHz profile
	if err := service.SetDot11BRfProfile(tsc.Ctx, tsc.TestTagName, testRF24GHzProfile); err != nil {
		return err
	}

	// Configure 5GHz profile
	if err := service.SetDot11ARfProfile(tsc.Ctx, tsc.TestTagName, testRF5GHzProfile); err != nil {
		return err
	}

	// Configure 6GHz profile (may fail in some environments)
	if err := service.SetDot116GhzRfProfile(tsc.Ctx, tsc.TestTagName, testRF6GHzProfile); err != nil {
		t.Logf("Warning: Failed to set 6GHz profile (this might be expected): %v", err)
	}

	return nil
}

// validateRFTagConfig validates the retrieved RF tag configuration
func validateRFTagConfig(t *testing.T, tsc *client.TagScenarioContext, config *model.RfTag) {
	client.ValidateConfigNotNil(t, config, tsc.TestTagName)
	client.ValidateCommonTagFields(t, tsc.TestTagName, expectedRFTagDescription,
		config.TagName, config.Description)

	// Validate 2.4GHz profile using standardized assertion
	helper.AssertStringEquals(t, config.Dot11BRfProfileName, testRF24GHzProfile,
		"2.4GHz RF profile validation")

	// Validate 5GHz profile using standardized assertion
	helper.AssertStringEquals(t, config.Dot11ARfProfileName, testRF5GHzProfile,
		"5GHz RF profile validation")

	// Validate 6GHz profile (may be empty in some environments)
	if config.Dot116GhzRfProfName != testRF6GHzProfile {
		t.Logf("Warning: Expected 6GHz profile '%s', got '%s' (this might be expected in some environments due to the default profile doesn't show in the response)",
			testRF6GHzProfile, config.Dot116GhzRfProfName)
	}
}

// validateRFTagDeletion ensures the test tag was properly deleted
func validateRFTagDeletion(t *testing.T, testTagName string, finalTags []model.RfTag) {
	for _, tag := range finalTags {
		if tag.TagName == testTagName {
			t.Errorf("Test tag %s should have been deleted but still exists", testTagName)
		}
	}
}
