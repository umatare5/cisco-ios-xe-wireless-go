//go:build scenario

package tag

import (
	"testing"
	"time"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/wlan"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/scenario"
)

// Test constants
const (
	expectedPolicyTagDescription = "Test policy tag created by scenario test"
	testWLANProfileName          = "test-wlan"
	testPolicyProfileName        = "default-policy-profile"
)

// generateTestPolicyTagName creates a unique test policy tag name with timestamp.
// Format: test-policy-tag-{yyyymmdd-hhmmss}
func generateTestPolicyTagName() string {
	timestamp := time.Now().Format("20060102-150405")
	return "test-policy-tag-" + timestamp
}

// TestPolicyTagScenario executes a comprehensive scenario test for policy tag management.
func TestPolicyTagServiceScenario_TagLifecycleManagement_Success(t *testing.T) {
	// Initialize scenario context
	tsc := scenario.SetupTag(t, generateTestPolicyTagName)

	// Initialize service
	service := wlan.NewPolicyTagService(tsc.Client)

	// Execute standard 6-step workflow
	executeStandardPolicyTagWorkflow(t, tsc, service)
}

// executeStandardPolicyTagWorkflow performs the standard policy tag scenario
func executeStandardPolicyTagWorkflow(t *testing.T, tsc *scenario.TagContext, service *wlan.PolicyTagService) {
	// Step 1: List initial tags and set count
	t.Logf("Step 1: Getting current policy tags list")
	if err := tsc.SetInitialCount(func() (int, error) {
		tags, err := service.ListPolicyTags(tsc.Ctx)
		return len(tags), err
	}); err != nil {
		t.Fatalf("Failed to get initial policy tags count: %v", err)
	}

	// Step 2: Create test tag
	t.Logf("Step 2: Creating test policy tag: %s", tsc.TestTagName)
	if err := service.CreatePolicyTag(tsc.Ctx, &model.PolicyListEntry{
		TagName:     tsc.TestTagName,
		Description: expectedPolicyTagDescription,
	}); err != nil {
		t.Fatalf("Failed to create policy tag %s: %v", tsc.TestTagName, err)
	}
	t.Logf("Successfully created policy tag: %s", tsc.TestTagName)

	// Step 3: Configure WLAN policies
	t.Logf("Step 3: Setting WLAN policy for test tag: %s", tsc.TestTagName)
	if err := service.SetPolicyTag(tsc.Ctx, &model.PolicyListEntry{
		TagName: tsc.TestTagName,
		WLANPolicies: &model.WLANPolicies{
			WLANPolicy: []model.WLANPolicyMap{
				{
					WLANProfileName:   testWLANProfileName,
					PolicyProfileName: testPolicyProfileName,
				},
			},
		},
	}); err != nil {
		t.Fatalf("Failed to configure WLAN policy for tag %s: %v", tsc.TestTagName, err)
	}
	t.Logf("Successfully configured policy tag: %s", tsc.TestTagName)

	// Step 4: Get and validate configuration
	t.Logf("Step 4: Getting configuration for test tag: %s", tsc.TestTagName)
	retrievedConfig, err := service.GetPolicyTag(tsc.Ctx, tsc.TestTagName)
	if err != nil {
		t.Fatalf("Failed to retrieve policy tag configuration for %s: %v", tsc.TestTagName, err)
	}
	validatePolicyTagConfiguration(t, tsc, retrievedConfig)
	t.Logf("Successfully validated policy tag: %s", tsc.TestTagName)

	// Step 5: Delete test tag
	t.Logf("Step 5: Deleting policy tag: %s", tsc.TestTagName)
	if err := service.DeletePolicyTag(tsc.Ctx, tsc.TestTagName); err != nil {
		t.Fatalf("Failed to delete policy tag %s: %v", tsc.TestTagName, err)
	}
	tsc.MarkTagDeleted()
	t.Logf("Successfully deleted policy tag: %s", tsc.TestTagName)

	// Step 6: Verify final tag count
	t.Logf("Step 6: Verifying final policy tags list")
	finalTags, err := service.ListPolicyTags(tsc.Ctx)
	if err != nil {
		t.Fatalf("Failed to list final policy tags: %v", err)
	}
	if err := tsc.ValidateTagCount(func() (int, error) {
		tags, err := service.ListPolicyTags(tsc.Ctx)
		return len(tags), err
	}, len(finalTags)); err != nil {
		t.Fatalf("Tag count validation failed: %v", err)
	}
}

// validatePolicyTagConfiguration validates the retrieved policy tag configuration
func validatePolicyTagConfiguration(t *testing.T, tsc *scenario.TagContext, config *model.PolicyListEntry) {
	if config == nil {
		t.Errorf("Tag '%s' configuration should not be nil", tsc.TestTagName)
		return
	}

	// Validate common tag fields
	if config.TagName != tsc.TestTagName {
		t.Errorf("Tag name mismatch: expected %s, got %s", tsc.TestTagName, config.TagName)
	}
	if config.Description != expectedPolicyTagDescription {
		t.Errorf("Description mismatch: expected %s, got %s", expectedPolicyTagDescription, config.Description)
	}

	// Validate WLAN policies
	if config.WLANPolicies == nil || len(config.WLANPolicies.WLANPolicy) == 0 {
		t.Error("WLAN policies should not be empty")
		return
	}

	found := false
	for _, policy := range config.WLANPolicies.WLANPolicy {
		if policy.WLANProfileName == testWLANProfileName &&
			policy.PolicyProfileName == testPolicyProfileName {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected WLAN policy mapping not found: %s -> %s",
			testWLANProfileName, testPolicyProfileName)
	}
}
