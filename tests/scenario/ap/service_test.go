//go:build scenario

package ap_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ap"
)

// TestAPServiceScenario_AdminStateManagement_Success validates the complete
// admin state management workflow for Access Points through the WNC controller.
//
// This scenario test executes a 4-step workflow that covers the complete
// enable/disable cycle for AP admin state management:
//  1. Retrieve initial AP CAPWAP status
//  2. Disable AP admin state
//  3. Re-enable AP admin state
//  4. Validate final state shows 'adminstate-enabled'
//
// Test Environment:
//   - Requires live WNC controller connectivity
//   - Uses environment variables for authentication
//   - Tests against configured test AP (WNC_AP_MAC_ADDR)
//
// Note: Not marked as parallel due to AP state modifications
// that could interfere with concurrent AP operations.
func TestAPServiceScenario_AdminStateManagement_Success(t *testing.T) {
	// Initialize scenario context
	asc, skip := initializeAPScenario(t, client.TestAPMac(), testAPSlotID)
	if skip {
		return
	}

	// Initialize service
	service := ap.NewService(asc.Client)

	// Execute standard 4-step workflow
	executeAPAdminStateWorkflow(asc, &service)
}

// TestAPServiceScenario_RadioStateManagement_Success validates the complete
// radio state management workflow for Access Point radio interfaces.
//
// This scenario test executes a 4-step workflow that covers the complete
// enable/disable cycle for AP radio admin state management:
//  1. Retrieve initial radio operational status
//  2. Disable radio admin state
//  3. Re-enable radio admin state
//  4. Validate final state shows 'enabled' and 'radio-up'
//
// Test Environment:
//   - Requires live WNC controller connectivity
//   - Uses environment variables for authentication
//   - Tests against configured test AP (WNC_AP_MAC_ADDR) Slot 0
//
// Note: Not marked as parallel due to radio state modifications
// that could interfere with concurrent radio operations.
func TestAPServiceScenario_RadioStateManagement_Success(t *testing.T) {
	// Initialize scenario context
	asc, skip := initializeAPScenario(t, client.TestAPMac(), testAPSlotID)
	if skip {
		return
	}

	// Initialize service
	service := ap.NewService(asc.Client)

	// Execute standard 4-step workflow
	executeAPRadioStateWorkflow(asc, &service)
}
