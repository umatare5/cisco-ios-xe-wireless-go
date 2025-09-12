//go:build scenario

package tag

import (
	"testing"
	"time"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/wlan"
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
	tsc, skip := client.InitializeTagScenario(t, generateTestPolicyTagName)
	if skip {
		return
	}

	// Initialize service
	service := wlan.NewPolicyTagService(tsc.Client)

	// Cleanup on exit
	defer tsc.CreateCleanupFunction(service.DeletePolicyTag)()

	// Execute standard 6-step workflow
	executeStandardPolicyTagWorkflow(t, tsc, service)
}

// executeStandardPolicyTagWorkflow performs the standard policy tag scenario
func executeStandardPolicyTagWorkflow(t *testing.T, tsc *client.TagScenarioContext, service *wlan.PolicyTagService) {
	// Step 1: List initial tags
	initialTags := client.ExecuteStepWithResult(tsc, client.StepInitialList,
		"Getting current policy tags list", func() ([]model.PolicyListEntry, error) {
			return service.ListPolicyTags(tsc.Ctx)
		})
	tsc.SetInitialCount(len(initialTags))

	// Step 2: Create test tag
	tsc.ExecuteStep(client.StepCreate, "Creating test policy tag", func() error {
		return service.CreatePolicyTag(tsc.Ctx, &model.PolicyListEntry{
			TagName:     tsc.TestTagName,
			Description: expectedPolicyTagDescription,
		})
	})
	tsc.LogSuccess("created")

	// Step 3: Configure WLAN policies
	tsc.ExecuteStep(client.StepConfigure, "Setting WLAN policy for test tag", func() error {
		return service.SetPolicyTag(tsc.Ctx, &model.PolicyListEntry{
			TagName: tsc.TestTagName,
			WLANPolicies: &model.WLANPolicies{
				WLANPolicy: []model.WLANPolicyMap{
					{
						WLANProfileName:   testWLANProfileName,
						PolicyProfileName: testPolicyProfileName,
					},
				},
			},
		})
	})
	tsc.LogSuccess("configured")

	// Step 4: Get and validate configuration
	retrievedConfig := client.ExecuteStepWithResult(tsc, client.StepGet,
		"Getting configuration for test tag", func() (*model.PolicyListEntry, error) {
			return service.GetPolicyTag(tsc.Ctx, tsc.TestTagName)
		})
	validatePolicyTagConfiguration(t, tsc, retrievedConfig)
	tsc.LogSuccess("validated")

	// Step 5: Delete test tag
	tsc.ExecuteDeleteStep("Deleting test policy tag", func() error {
		return service.DeletePolicyTag(tsc.Ctx, tsc.TestTagName)
	})
	tsc.LogSuccess("deleted")

	// Step 6: Verify deletion
	finalTags := client.ExecuteStepWithResult(tsc, client.StepFinalList,
		"Getting final policy tags list", func() ([]model.PolicyListEntry, error) {
			return service.ListPolicyTags(tsc.Ctx)
		})
	validatePolicyTagDeletion(t, tsc.TestTagName, finalTags)
	tsc.ValidateTagCount(len(finalTags))
}

// validatePolicyTagConfiguration validates the retrieved policy tag configuration
func validatePolicyTagConfiguration(t *testing.T, tsc *client.TagScenarioContext, config *model.PolicyListEntry) {
	if config == nil {
		t.Fatalf("Retrieved config is nil for test tag: %s", tsc.TestTagName)
		return
	}

	client.ValidateConfigNotNil(t, config, tsc.TestTagName)
	client.ValidateCommonTagFields(t, tsc.TestTagName, expectedPolicyTagDescription,
		config.TagName, config.Description)

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

// validatePolicyTagDeletion ensures the test tag was properly deleted
func validatePolicyTagDeletion(t *testing.T, testTagName string, finalTags []model.PolicyListEntry) {
	for _, tag := range finalTags {
		if tag.TagName == testTagName {
			t.Errorf("Test tag %s should have been deleted but still exists", testTagName)
		}
	}
}
