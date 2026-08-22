package restconf

import (
	"strconv"
	"strings"
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
		{
			name:         "module-qualified node keeps its colon",
			endpointPath: "/p",
			param:        FieldsParam,
			value:        "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries",
			expected:     "/p?fields=Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries",
		},
		{
			name:         "value cannot open a second parameter",
			endpointPath: "/p",
			param:        FieldsParam,
			value:        "wlan-id&depth=1",
			expected:     "/p?fields=wlan-id%26depth%3D1",
		},
		{
			name:         "value cannot open a fragment",
			endpointPath: "/p",
			param:        FieldsParam,
			value:        "wlan-id#frag",
			expected:     "/p?fields=wlan-id%23frag",
		},
		{
			name:         "space and plus are both encoded",
			endpointPath: "/p",
			param:        FieldsParam,
			value:        "a b+c",
			expected:     "/p?fields=a%20b%2Bc",
		},
		{
			name:         "percent is encoded once, so a pre-encoded value is encoded twice",
			endpointPath: "/p",
			param:        FieldsParam,
			value:        "already%20encoded",
			expected:     "/p?fields=already%2520encoded",
		},
		{
			name:         "multibyte is encoded as uppercase UTF-8 triplets",
			endpointPath: "/p",
			param:        FieldsParam,
			value:        "社",
			expected:     "/p?fields=%E7%A4%BE",
		},
		{
			name:         "depth bounds are digits and survive",
			endpointPath: "/p",
			param:        DepthParam,
			value:        "65535",
			expected:     "/p?depth=65535",
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
	testutil.AssertStringEquals(t, FieldsParam, "fields", "FieldsParam")
	testutil.AssertStringEquals(t, DepthParam, "depth", "DepthParam")
}

// TestRESTCONFQueryUnit_EscapeQueryValue_ByteSet pins the escaper on every one of the
// 256 byte values, so neither widening nor narrowing the literal set goes unnoticed.
//
// A byte counts as kept only when the output is that same byte: classifying by output
// length alone would record the input for any one-byte output, and a substitution that
// answers a different single byte would pass.
func TestRESTCONFQueryUnit_EscapeQueryValue_ByteSet(t *testing.T) {
	const literalBytes = "()-./0123456789:;ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

	var kept strings.Builder
	for c := range 256 {
		got := escapeQueryValue(string([]byte{byte(c)}))
		if got == string([]byte{byte(c)}) {
			kept.WriteByte(byte(c))
			continue
		}
		testutil.AssertIntEquals(t, len(got), 3, "escaped length of byte "+strconv.Itoa(c))
	}

	testutil.AssertStringEquals(t, kept.String(), literalBytes, "bytes kept literally")
	testutil.AssertIntEquals(t, len(kept.String()), 71, "count of bytes kept literally")
}

// TestRESTCONFQueryUnit_EscapeQueryValue_UpperHex pins the hex alphabet: RFC 3986 6.2.2.1
// prefers upper case, and a controller comparing a key byte for byte sees the difference.
func TestRESTCONFQueryUnit_EscapeQueryValue_UpperHex(t *testing.T) {
	testutil.AssertStringEquals(t, escapeQueryValue("?"), "%3F", "escapeQueryValue(\"?\")")
	testutil.AssertStringEquals(t, escapeQueryValue("\xff"), "%FF", "escapeQueryValue(0xFF)")
	testutil.AssertStringEquals(t, escapeQueryValue("\x00"), "%00", "escapeQueryValue(0x00)")
}
