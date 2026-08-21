//go:build integration || scenario

package integration

import (
	"strings"
	"testing"
)

// TestRedactCapture pins what may not reach a capture file. The function is pure, so this
// runs without a controller — but it is behind the integration build tag, because the
// package it guards is.
func TestRedactCapture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		body     string
		mustGo   string
		mustHold string
	}{
		{
			name:     "twelve bare hex digits are not a MAC and are left alone",
			body:     `{"wtp-mac":"a1b2c3d4e5f6"}`,
			mustGo:   "",
			mustHold: "a1b2c3d4e5f6",
		},
		{
			name:     "colon MAC becomes a locally administered dummy",
			body:     `{"wtp-mac":"1a:2b:3c:4d:5e:6f"}`,
			mustGo:   "1a:2b:3c:4d:5e:6f",
			mustHold: "02:11:22:33:44:55",
		},
		{
			name:     "dotted MAC",
			body:     `{"ap-mac":"1a2b.3c4d.5e6f"}`,
			mustGo:   "1a2b.3c4d.5e6f",
			mustHold: "0211.2233.4455",
		},
		{
			name:     "psk value",
			body:     `{"psk":"aVerySecretPassphrase"}`,
			mustGo:   "aVerySecretPassphrase",
			mustHold: `"psk":"REDACTED"`,
		},
		{
			name:     "proxy password value",
			body:     `{"password":"hunter2hunter2"}`,
			mustGo:   "hunter2hunter2",
			mustHold: `"password":"REDACTED"`,
		},
		{
			name:     "chassis serial",
			body:     `{"serial-number":"ABC1234WXYZ"}`,
			mustGo:   "ABC1234WXYZ",
			mustHold: "REDACTEDSERIAL",
		},
		{
			name:     "ordinary content survives",
			body:     `{"ap-name":"TEST-AP01","channel":36}`,
			mustGo:   "",
			mustHold: `"ap-name":"TEST-AP01"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := string(redactCapture([]byte(tt.body)))

			if tt.mustGo != "" && strings.Contains(got, tt.mustGo) {
				t.Errorf("redactCapture() left %q in %q", tt.mustGo, got)
			}
			if !strings.Contains(got, tt.mustHold) {
				t.Errorf("redactCapture() = %q, want it to contain %q", got, tt.mustHold)
			}
		})
	}
}
