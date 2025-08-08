package core

import (
	"testing"
	"time"
)

// TestStandardTimeoutConsistency tests that StandardTimeout equals DefaultTimeout
func TestStandardTimeoutConsistency(t *testing.T) {
	// Test that StandardTimeout equals DefaultTimeout
	if StandardTimeout != DefaultTimeout {
		t.Errorf("Expected StandardTimeout (%v) to equal DefaultTimeout (%v)", StandardTimeout, DefaultTimeout)
	}

	// Test that the default timeout is 60 seconds
	expectedTimeout := 60 * time.Second
	if DefaultTimeout != expectedTimeout {
		t.Errorf("Expected DefaultTimeout to be %v, got %v", expectedTimeout, DefaultTimeout)
	}
}
