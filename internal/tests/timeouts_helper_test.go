package tests

import (
	"path/filepath"
	"testing"
	"time"
)

func TestConstants(t *testing.T) {
	if DefaultTestTimeout != 30*time.Second {
		t.Errorf("DefaultTestTimeout expected 30s, got %v", DefaultTestTimeout)
	}
	if ExtendedTestTimeout != 60*time.Second {
		t.Errorf("ExtendedTestTimeout expected 60s, got %v", ExtendedTestTimeout)
	}
	if TestDataDir != "./test_data" {
		t.Errorf("TestDataDir expected './test_data', got %v", TestDataDir)
	}
	_ = filepath.Separator
}
