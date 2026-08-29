package core_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// TestGetAdminStateMode tests the admin state mode conversion.
func TestGetAdminStateMode(t *testing.T) {
	tests := []struct {
		name     string
		enabled  bool
		expected string
	}{
		{
			name:     "Enabled state",
			enabled:  true,
			expected: "admin-state-enabled",
		},
		{
			name:     "Disabled state",
			enabled:  false,
			expected: "admin-state-disabled",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := core.GetAdminStateMode(tt.enabled)
			testutil.AssertStringEquals(t, result, tt.expected, "GetAdminStateMode result")
		})
	}
}

// TestAdminStateConstants pins the two spellings the AP admin-state and slot-admin RPCs send, so a
// changed constant fails here rather than as a 400 on a controller.
func TestAdminStateConstants(t *testing.T) {
	t.Run("AdminStateEnabled", func(t *testing.T) {
		expected := "admin-state-enabled"
		testutil.AssertStringEquals(t, core.AdminStateEnabled, expected, "AdminStateEnabled")
	})

	t.Run("AdminStateDisabled", func(t *testing.T) {
		expected := "admin-state-disabled"
		testutil.AssertStringEquals(t, core.AdminStateDisabled, expected, "AdminStateDisabled")
	})
}
