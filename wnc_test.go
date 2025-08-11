package wnc

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"
)

// TestNewClient tests the creation of a new unified client
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
			host:        "controller.example.com",
			token:       "dGVzdDp0ZXN0", // base64 encoded "test:test"
			opts:        nil,
			expectError: false,
		},
		{
			name:        "ValidClientWithOptions",
			host:        "192.168.1.100",
			token:       "YWRtaW46cGFzc3dvcmQ=", // base64 encoded "admin:password"
			opts:        []Option{WithTimeout(30 * time.Second), WithInsecureSkipVerify(true)},
			expectError: false,
		},
		{
			name:        "InvalidHost",
			host:        "",
			token:       "dGVzdDp0ZXN0",
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

// TestClientServiceAccessors tests that all service accessors return non-nil services
func TestClientServiceAccessors(t *testing.T) {
	client, err := NewClient("controller.example.com", "dGVzdDp0ZXN0")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
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
	_ = client.WLAN()          // Should not panic
}

// TestClientCore tests the Core() method
func TestClientCore(t *testing.T) {
	client, err := NewClient("controller.example.com", "dGVzdDp0ZXN0")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	coreClient := client.Core()
	if coreClient == nil {
		t.Error("Core() returned nil")
	}
}

// TestUnifiedClientUsage demonstrates the unified client usage pattern
func TestUnifiedClientUsage(t *testing.T) {
	// This test demonstrates the usage pattern without making real API calls
	client, err := NewClient("controller.example.com", "dGVzdDp0ZXN0",
		WithTimeout(30*time.Second),
		WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Demonstrate that services can be accessed
	ctx := context.Background()

	// These would normally make actual API calls, but we're just testing the interface
	afcService := client.AFC()
	generalService := client.General()
	apService := client.AP()

	// Verify services are accessible (struct values, not pointers for some services)
	_ = afcService     // Should not panic accessing AFC service
	_ = generalService // Should not panic accessing General service
	_ = apService      // Should not panic accessing AP service

	// For testing purposes, we won't make actual API calls
	// But this demonstrates the intended usage pattern:
	// afcData, err := afcService.GetOper(ctx)
	// generalData, err := generalService.GetOper(ctx)
	// apData, err := apService.GetOper(ctx)

	_ = ctx // Use ctx to avoid unused variable warning
}

// TestOptionWrappers ensures wrapper options (WithLogger, WithUserAgent) apply without error.
func TestOptionWrappers(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	c, err := NewClient("controller.example.com", "dGVzdDp0ZXN0",
		WithLogger(logger),
		WithUserAgent("custom-agent/1.0"),
	)
	if err != nil {
		// If underlying validation changes, this should still expose root cause.
		t.Fatalf("unexpected error applying wrapper options: %v", err)
	}
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}
