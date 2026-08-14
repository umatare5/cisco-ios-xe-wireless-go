package restconf

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

func TestRESTCONFQueryUnit_AppendQueryParam_Success(t *testing.T) {
	tests := []struct {
		name         string
		endpointPath string
		param        string
		value        string
		expected     string
	}{
		{
			name:         "path without query",
			endpointPath: "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries",
			param:        WithDefaultsParam,
			value:        "report-all",
			expected:     "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries?with-defaults=report-all",
		},
		{
			name:         "path already carrying a query",
			endpointPath: "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries?depth=1",
			param:        WithDefaultsParam,
			value:        "explicit",
			expected: "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries?depth=1" +
				"&with-defaults=explicit",
		},
		{
			name:         "value kept verbatim",
			endpointPath: "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data",
			param:        "fields",
			value:        "wlan-cfg-entries/wlan-cfg-entry(profile-name;wlan-id)",
			expected: "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data" +
				"?fields=wlan-cfg-entries/wlan-cfg-entry(profile-name;wlan-id)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AppendQueryParam(tt.endpointPath, tt.param, tt.value)
			testutil.AssertStringEquals(t, result, tt.expected, "AppendQueryParam()")
		})
	}
}

func TestRESTCONFQueryUnit_Constants_Success(t *testing.T) {
	testutil.AssertStringEquals(t, WithDefaultsParam, "with-defaults", "WithDefaultsParam")
}
