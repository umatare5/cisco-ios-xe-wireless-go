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
		{
			name:     "fields alone opens the query",
			opts:     []GetOption{WithFields("wlan-cfg-entry(profile-name;wlan-id)")},
			expected: endpoint + "?fields=wlan-cfg-entry(profile-name;wlan-id)",
		},
		{
			name:     "empty fields expression is ignored",
			opts:     []GetOption{WithFields("")},
			expected: endpoint,
		},
		{
			name:     "last fields expression wins",
			opts:     []GetOption{WithFields("wlan-id"), WithFields("profile-name")},
			expected: endpoint + "?fields=profile-name",
		},
		{
			name:     "a fields expression cannot open a second parameter",
			opts:     []GetOption{WithFields("wlan-id&with-defaults=report-all")},
			expected: endpoint + "?fields=wlan-id%26with-defaults%3Dreport-all",
		},
		{
			name:     "depth alone opens the query",
			opts:     []GetOption{WithDepth(3)},
			expected: endpoint + "?depth=3",
		},
		{
			name:     "depth 1 is the lowest accepted value",
			opts:     []GetOption{WithDepth(1)},
			expected: endpoint + "?depth=1",
		},
		{
			name:     "depth 65535 is the highest documented value",
			opts:     []GetOption{WithDepth(65535)},
			expected: endpoint + "?depth=65535",
		},
		{
			name:     "depth zero is ignored",
			opts:     []GetOption{WithDepth(0)},
			expected: endpoint,
		},
		{
			name:     "negative depth is ignored",
			opts:     []GetOption{WithDepth(-1)},
			expected: endpoint,
		},
		{
			name: "all three fold in a fixed order",
			opts: []GetOption{
				WithDepth(2),
				WithFields("wlan-cfg-entries"),
				WithDefaults(DefaultsReportAll),
			},
			expected: endpoint + "?with-defaults=report-all&fields=wlan-cfg-entries&depth=2",
		},
		{
			name:     "fields and depth without defaults still use one question mark",
			opts:     []GetOption{WithFields("wlan-cfg-entries"), WithDepth(2)},
			expected: endpoint + "?fields=wlan-cfg-entries&depth=2",
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
