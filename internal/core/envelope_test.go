package core

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

func TestCoreEnvelopeUnit_NodeName_Success(t *testing.T) {
	const base = "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data"

	tests := []struct {
		name     string
		endpoint string
		expected string
	}{
		{
			name:     "container",
			endpoint: base,
			expected: "common-oper-data",
		},
		{
			name:     "root container carries the module prefix",
			endpoint: "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data",
			expected: "client-oper-data",
		},
		{
			name:     "list key holds colons",
			endpoint: base + "=aa:bb:cc:dd:ee:ff",
			expected: "common-oper-data",
		},
		{
			name:     "list key holds a slash",
			endpoint: base + "=aa/bb",
			expected: "common-oper-data",
		},
		{
			name:     "composite list key",
			endpoint: base + "=aa:bb:cc:dd:ee:ff,1",
			expected: "common-oper-data",
		},
		{
			name:     "query follows the list key",
			endpoint: base + "=aa:bb:cc:dd:ee:ff?with-defaults=report-all",
			expected: "common-oper-data",
		},
		{
			name:     "query without a list key",
			endpoint: base + "?with-defaults=report-all",
			expected: "common-oper-data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AssertStringEquals(t, nodeName(tt.endpoint), tt.expected, "nodeName()")
		})
	}
}

func TestCoreEnvelopeUnit_DecodeSoleKey_Success(t *testing.T) {
	const endpoint = "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data"

	type response struct {
		CommonOperData []struct {
			ClientMAC string `json:"client-mac"`
		} `json:"Cisco-IOS-XE-wireless-client-oper:common-oper-data"`
	}

	body := []byte(`{"Cisco-IOS-XE-wireless-client-oper:common-oper-data":[{"client-mac":"aa:bb:cc:dd:ee:ff"}]}`)

	out, err := decodeSoleKey[response](body, endpoint+"=aa:bb:cc:dd:ee:ff")
	testutil.AssertNoError(t, err, "decodeSoleKey() on a keyed read")
	testutil.AssertIntEquals(t, len(out.CommonOperData), 1, "decoded record count")
}

func TestCoreEnvelopeUnit_DecodeSoleKey_Error(t *testing.T) {
	const endpoint = "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data"

	type response struct {
		CommonOperData []struct{} `json:"Cisco-IOS-XE-wireless-client-oper:common-oper-data"`
	}

	tests := []struct {
		name     string
		body     string
		contains string
	}{
		{
			name:     "no top-level key",
			body:     `{}`,
			contains: "response carries 0 top-level keys",
		},
		{
			name:     "two top-level keys",
			body:     `{"a":1,"b":2}`,
			contains: "response carries 2 top-level keys",
		},
		{
			name:     "key is not module-qualified",
			body:     `{"common-oper-data":[]}`,
			contains: `want a module-qualified "common-oper-data"`,
		},
		{
			name:     "key names another node",
			body:     `{"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data":[]}`,
			contains: `want a module-qualified "common-oper-data"`,
		},
		{
			name:     "body is null",
			body:     `null`,
			contains: "response carries 0 top-level keys",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeSoleKey[response]([]byte(tt.body), endpoint)
			testutil.AssertError(t, err, "decodeSoleKey()")
			testutil.AssertStringContains(t, err.Error(), tt.contains, "error message")
		})
	}
}

func TestCoreEnvelopeUnit_DecodeSoleKey_UndeclaredField(t *testing.T) {
	const endpoint = "/restconf/data/Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-stats"

	type bareResponse struct {
		RestartCount int `json:"restart-count"`
	}

	body := []byte(`{"Cisco-IOS-XE-wireless-rogue-oper:rogue-stats":{"restart-count":7}}`)

	_, err := decodeSoleKey[bareResponse](body, endpoint)
	testutil.AssertError(t, err, "decodeSoleKey() with a type missing the envelope")
	testutil.AssertStringContains(t, err.Error(), "declares no field for response key", "error message")
}
