package helpers

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// TestDataDir is the directory for test data files
const TestDataDir = "./testdata"

// Dependency injection hooks for filesystem operations (overridden in tests only).
var (
	mkdirAll  = os.MkdirAll
	writeFile = os.WriteFile
)

// SaveTestDataToFile saves test data to a JSON file.
func SaveTestDataToFile(filename string, data any) error { //nolint:revive // helper clarity
	// Create testdata directory if it doesn't exist
	if err := mkdirAll(TestDataDir, 0o755); err != nil { //nolint:gosec // Test directory permissions
		return err
	}
	fullPath := filepath.Join(TestDataDir, filename)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(fullPath, jsonData, 0o644) //nolint:gosec // Test file permissions
}
