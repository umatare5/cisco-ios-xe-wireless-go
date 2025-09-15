package wnc

import (
	"log/slog"
	"testing"
	"time"
)

// TestNewClient tests the creation of a new unified client.
func TestNewClient(t *testing.T) {
	testCases := []struct {
		name        string
		host        string
		token       string
		opts        []Option
		expectError bool
	}{
		{
			name:        "ValidClient",
			host:        "192.168.1.100",
			token:       "YWRtaW46cGFzc3dvcmQ=", // base64 encoded "admin:password"
			opts:        nil,
			expectError: false,
		},
		{
			name:        "ValidClientWithOptions",
			host:        "wnc.example.internal",
			token:       "YWRtaW46cGFzc3dvcmQ=",
			opts:        []Option{WithTimeout(30 * time.Second), WithInsecureSkipVerify(true)},
			expectError: false,
		},
		{
			name:        "ValidClientWithLoggerAndUserAgent",
			host:        "controller.example.internal",
			token:       "YWRtaW46cGFzc3dvcmQ=",
			opts:        []Option{WithLogger(slog.New(slog.DiscardHandler)), WithUserAgent("custom-agent/1.0")},
			expectError: false,
		},
		{
			name:        "InvalidHost",
			host:        "",
			token:       "YWRtaW46cGFzc3dvcmQ=",
			opts:        nil,
			expectError: true,
		},
		{
			name:        "InvalidToken",
			host:        "controller.example.com",
			token:       "",
			opts:        nil,
			expectError: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.host, tt.token, tt.opts...)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error, but got none")
				}
				if client != nil {
					t.Error("Expected nil client on error")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if client == nil {
					t.Error("Expected client, but got nil")
				}
			}
		})
	}
}

// TestClientServiceAccessors tests that all service accessors return non-nil services.
func TestClientServiceAccessors(t *testing.T) {
	client, err := NewClient("controller.example.com", "dGVzdDp0ZXN0")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test Core() method
	coreClient := client.Core()
	if coreClient == nil {
		t.Error("Core() returned nil")
	}

	// Test all service accessors - verify they don't panic and return valid structs
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Service accessor panicked: %v", r)
		}
	}()

	// Test service accessors return valid services
	_ = client.AFC()           // Should not panic
	_ = client.AP()            // Should not panic
	_ = client.APF()           // Should not panic
	_ = client.AWIPS()         // Should not panic
	_ = client.BLE()           // Should not panic
	_ = client.Client()        // Should not panic
	_ = client.Controller()    // Should not panic
	_ = client.CTS()           // Should not panic
	_ = client.Dot11()         // Should not panic
	_ = client.Dot15()         // Should not panic
	_ = client.Fabric()        // Should not panic
	_ = client.Flex()          // Should not panic
	_ = client.General()       // Should not panic
	_ = client.Geolocation()   // Should not panic
	_ = client.Hyperlocation() // Should not panic
	_ = client.LISP()          // Should not panic
	_ = client.Location()      // Should not panic
	_ = client.Mcast()         // Should not panic
	_ = client.MDNS()          // Should not panic
	_ = client.Mesh()          // Should not panic
	_ = client.Mobility()      // Should not panic
	_ = client.NMSP()          // Should not panic
	_ = client.Radio()         // Should not panic
	_ = client.RF()            // Should not panic
	_ = client.RFID()          // Should not panic
	_ = client.Rogue()         // Should not panic
	_ = client.RRM()           // Should not panic
	_ = client.Site()          // Should not panic
	_ = client.Spaces()        // Should not panic
	_ = client.URWB()          // Should not panic
	_ = client.WAT()           // Should not panic
	_ = client.WLAN()          // Should not panic

	// Test tag service accessors
	_ = client.PolicyTag() // Should not panic
	_ = client.RFTag()     // Should not panic
	_ = client.SiteTag()   // Should not panic
}
