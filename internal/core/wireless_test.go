package core_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// TestGetRadioBandInfo tests the radio band information conversion.
func TestGetRadioBandInfo(t *testing.T) {
	tests := []struct {
		name         string
		radioBand    int
		expectedErr  bool
		expectedBand uint32
		expectedSlot uint8
	}{
		{
			name:         "2.4GHz band",
			radioBand:    0, // RadioBand24GHzEnum
			expectedErr:  false,
			expectedBand: 1, // RadioBand24GHzValue
			expectedSlot: 0, // RadioSlot24GHz
		},
		{
			name:         "5GHz band",
			radioBand:    1, // RadioBand5GHzEnum
			expectedErr:  false,
			expectedBand: 2, // RadioBand5GHzValue
			expectedSlot: 1, // RadioSlot5GHz
		},
		{
			name:        "Invalid band - negative",
			radioBand:   -1,
			expectedErr: true,
		},
		{
			name:        "Invalid band - too high",
			radioBand:   999,
			expectedErr: true,
		},
		{
			name:         "6GHz band",
			radioBand:    2, // RadioBand6GHzEnum
			expectedErr:  false,
			expectedBand: 3, // RadioBand6GHzValue
			expectedSlot: 2, // RadioSlot6GHz
		},
		{
			name:        "Invalid band - 3",
			radioBand:   3,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := core.GetRadioBandInfo(tt.radioBand)

			if tt.expectedErr {
				testutil.AssertError(t, err, "GetRadioBandInfo expected error")
				return
			}

			testutil.AssertNoError(t, err, "GetRadioBandInfo unexpected error")

			testutil.AssertIntEquals(t, int(info.Band), int(tt.expectedBand), "GetRadioBandInfo Band")

			testutil.AssertIntEquals(t, int(info.SlotID), int(tt.expectedSlot), "GetRadioBandInfo SlotID")
		})
	}
}

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

// TestRadioBandConstants tests that radio band constants have expected values.
func TestRadioBandConstants(t *testing.T) {
	// Test constant values to ensure they don't change unexpectedly

	t.Run("RadioBand24GHzEnum", func(t *testing.T) {
		testutil.AssertIntEquals(t, core.RadioBand24GHzEnum, 0, "RadioBand24GHzEnum")
	})

	t.Run("RadioBand5GHzEnum", func(t *testing.T) {
		testutil.AssertIntEquals(t, core.RadioBand5GHzEnum, 1, "RadioBand5GHzEnum")
	})

	t.Run("RadioBand24GHzValue", func(t *testing.T) {
		expected := uint32(1)
		testutil.AssertIntEquals(t, int(core.RadioBand24GHzValue), int(expected), "RadioBand24GHzValue")
	})

	t.Run("RadioBand5GHzValue", func(t *testing.T) {
		expected := uint32(2)
		testutil.AssertIntEquals(t, int(core.RadioBand5GHzValue), int(expected), "RadioBand5GHzValue")
	})

	t.Run("RadioSlot24GHz", func(t *testing.T) {
		expected := uint8(0)
		testutil.AssertIntEquals(t, int(core.RadioSlot24GHz), int(expected), "RadioSlot24GHz")
	})

	t.Run("RadioSlot5GHz", func(t *testing.T) {
		expected := uint8(1)
		testutil.AssertIntEquals(t, int(core.RadioSlot5GHz), int(expected), "RadioSlot5GHz")
	})

	t.Run("AdminStateEnabled", func(t *testing.T) {
		expected := "admin-state-enabled"
		testutil.AssertStringEquals(t, core.AdminStateEnabled, expected, "AdminStateEnabled")
	})

	t.Run("AdminStateDisabled", func(t *testing.T) {
		expected := "admin-state-disabled"
		testutil.AssertStringEquals(t, core.AdminStateDisabled, expected, "AdminStateDisabled")
	})

	t.Run("OperationEnable", func(t *testing.T) {
		expected := "enable"
		testutil.AssertStringEquals(t, core.OperationEnable, expected, "OperationEnable")
	})

	t.Run("OperationDisable", func(t *testing.T) {
		expected := "disable"
		testutil.AssertStringEquals(t, core.OperationDisable, expected, "OperationDisable")
	})
}
