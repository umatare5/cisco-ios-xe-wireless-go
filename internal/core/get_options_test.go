package core

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

func TestCoreGetOptionsUnit_ApplyGetOptions_Success(t *testing.T) {
	const endpoint = "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries"

	tests := []struct {
		name     string
		opts     []GetOption
		expected string
	}{
		{
			name:     "no option leaves the endpoint untouched",
			opts:     nil,
			expected: endpoint,
		},
		{
			name:     "nil option is ignored",
			opts:     []GetOption{nil},
			expected: endpoint,
		},
		{
			name:     "report-all mode",
			opts:     []GetOption{WithDefaults(DefaultsReportAll)},
			expected: endpoint + "?with-defaults=report-all",
		},
		{
			name:     "explicit mode",
			opts:     []GetOption{WithDefaults(DefaultsExplicit)},
			expected: endpoint + "?with-defaults=explicit",
		},
		{
			name:     "last option wins",
			opts:     []GetOption{WithDefaults(DefaultsReportAll), WithDefaults(DefaultsExplicit)},
			expected: endpoint + "?with-defaults=explicit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := applyGetOptions(endpoint, tt.opts)
			testutil.AssertStringEquals(t, result, tt.expected, "applyGetOptions()")
		})
	}
}

func TestCoreGetOptionsUnit_Constants_Success(t *testing.T) {
	testutil.AssertStringEquals(t, string(DefaultsReportAll), "report-all", "DefaultsReportAll")
	testutil.AssertStringEquals(t, string(DefaultsExplicit), "explicit", "DefaultsExplicit")
}
