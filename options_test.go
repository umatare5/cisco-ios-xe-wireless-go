package wnc

import (
	"io"
	"log/slog"
	"testing"
	"time"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestOptionNilClient tests option behavior with nil client
func TestOptionNilClient(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Options should not panic with nil client, but got panic: %v", r)
		}
	}()

	// Create a simple option
	opt := WithTimeout(10 * time.Second)

	// Test that the option function can be created without panicking
	if opt == nil {
		t.Error("Expected option function to be created")
	}

	// Note: We cannot safely apply options to nil client as it will cause panic
	// This test validates that option creation itself is safe
}

// =============================================================================

// TestClientOptionType tests the ClientOption type
func TestClientOptionType(t *testing.T) {
	// Test that ClientOption is a function type
	var option ClientOption

	// Test creating a simple option
	option = func(c *Client) {
		c.timeout = 30 * time.Second
	}

	if option == nil {
		t.Error("Expected created ClientOption to not be nil")
	}
}

// =============================================================================
// 2. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// TestWithTimeout tests the WithTimeout option
func TestWithTimeout(t *testing.T) {
	tests := []struct {
		name    string
		timeout time.Duration
	}{
		{"StandardTimeout", 30 * time.Second},
		{"QuickTimeout", 5 * time.Second},
		{"ExtendedTimeout", 90 * time.Second},
		{"ZeroTimeout", 0},
		{"NegativeTimeout", -10 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{}
			option := WithTimeout(tt.timeout)

			if option == nil {
				t.Fatal("Expected WithTimeout to return a non-nil option")
			}

			// Apply the option
			option(client)

			if client.timeout != tt.timeout {
				t.Errorf("Expected timeout to be %v, got %v", tt.timeout, client.timeout)
			}
		})
	}
}

// TestWithInsecureSkipVerify tests the WithInsecureSkipVerify option
func TestWithInsecureSkipVerify(t *testing.T) {
	tests := []struct {
		name string
		skip bool
	}{
		{"SkipVerification", true},
		{"DoNotSkipVerification", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{}
			option := WithInsecureSkipVerify(tt.skip)

			if option == nil {
				t.Fatal("Expected WithInsecureSkipVerify to return a non-nil option")
			}

			// Apply the option
			option(client)

			if client.insecureSkipVerify != tt.skip {
				t.Errorf("Expected insecureSkipVerify to be %v, got %v", tt.skip, client.insecureSkipVerify)
			}
		})
	}
}

// TestWithLogger tests the WithLogger option
func TestWithLogger(t *testing.T) {
	tests := []struct {
		name   string
		logger *slog.Logger
	}{
		{"TextLogger", slog.New(slog.NewTextHandler(io.Discard, nil))},
		{"JSONLogger", slog.New(slog.NewJSONHandler(io.Discard, nil))},
		{"NilLogger", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{}
			option := WithLogger(tt.logger)

			if option == nil {
				t.Fatal("Expected WithLogger to return a non-nil option")
			}

			// Apply the option
			option(client)

			if client.logger != tt.logger {
				t.Errorf("Expected logger to be %v, got %v", tt.logger, client.logger)
			}
		})
	}
}

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// TestMultipleOptions tests applying multiple options to a client
func TestMultipleOptions(t *testing.T) {
	client := &Client{}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	timeout := 45 * time.Second
	skipVerify := true

	// Apply multiple options
	options := []ClientOption{
		WithTimeout(timeout),
		WithInsecureSkipVerify(skipVerify),
		WithLogger(logger),
	}

	for _, option := range options {
		option(client)
	}

	// Verify all options were applied
	if client.timeout != timeout {
		t.Errorf("Expected timeout to be %v, got %v", timeout, client.timeout)
	}

	if client.insecureSkipVerify != skipVerify {
		t.Errorf("Expected insecureSkipVerify to be %v, got %v", skipVerify, client.insecureSkipVerify)
	}

	if client.logger != logger {
		t.Errorf("Expected logger to be %v, got %v", logger, client.logger)
	}
}

// TestOptionOverride tests that later options override earlier ones
func TestOptionOverride(t *testing.T) {
	client := &Client{}

	// Apply timeout options in sequence
	timeout1 := 30 * time.Second
	timeout2 := 60 * time.Second

	option1 := WithTimeout(timeout1)
	option2 := WithTimeout(timeout2)

	option1(client)
	if client.timeout != timeout1 {
		t.Errorf("Expected timeout to be %v after first option, got %v", timeout1, client.timeout)
	}

	option2(client)
	if client.timeout != timeout2 {
		t.Errorf("Expected timeout to be %v after second option, got %v", timeout2, client.timeout)
	}
}

// TestFunctionalOptionsPattern tests the functional options pattern
func TestFunctionalOptionsPattern(t *testing.T) {
	// Create a function that accepts options (similar to NewClient)
	createTestClient := func(options ...ClientOption) *Client {
		client := &Client{
			timeout:            DefaultTimeout,
			insecureSkipVerify: false,
			logger:             slog.New(slog.NewTextHandler(io.Discard, nil)),
		}

		for _, option := range options {
			option(client)
		}

		return client
	}

	// Test with no options
	t.Run("NoOptions", func(t *testing.T) {
		client := createTestClient()

		if client.timeout != DefaultTimeout {
			t.Errorf("Expected default timeout %v, got %v", DefaultTimeout, client.timeout)
		}

		if client.insecureSkipVerify {
			t.Error("Expected insecureSkipVerify to be false by default")
		}

		if client.logger == nil {
			t.Error("Expected logger to be set by default")
		}
	})

	// Test with multiple options
	t.Run("MultipleOptions", func(t *testing.T) {
		customTimeout := 45 * time.Second
		customLogger := slog.New(slog.NewJSONHandler(io.Discard, nil))

		client := createTestClient(
			WithTimeout(customTimeout),
			WithInsecureSkipVerify(true),
			WithLogger(customLogger),
		)

		if client.timeout != customTimeout {
			t.Errorf("Expected timeout %v, got %v", customTimeout, client.timeout)
		}

		if !client.insecureSkipVerify {
			t.Error("Expected insecureSkipVerify to be true")
		}

		if client.logger != customLogger {
			t.Error("Expected custom logger to be set")
		}
	})
}
