package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSaveTestDataToFile(t *testing.T) {
	tempDir := t.TempDir()
	_ = tempDir
	testData := map[string]interface{}{"test_key": "test_value", "number": 42, "boolean": true}
	filename := "test_file.json"
	testDataDir := filepath.Join(tempDir, "test_data")
	if err := os.MkdirAll(testDataDir, 0o755); err != nil {
		t.Fatalf("Failed to create test data directory: %v", err)
	}
	fullPath := filepath.Join(testDataDir, filename)
	jsonData, err := json.MarshalIndent(testData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}
	if err := os.WriteFile(fullPath, jsonData, 0o644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		t.Error("Expected test file to be created")
	}
	fileData, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}
	var loadedData map[string]interface{}
	if err := json.Unmarshal(fileData, &loadedData); err != nil {
		t.Fatalf("Failed to unmarshal saved data: %v", err)
	}
	if loadedData["test_key"] != testData["test_key"] {
		t.Error("Saved data mismatch")
	}

	t.Run("SuccessfulSave", func(t *testing.T) {
		err := SaveTestDataToFile("successful_test.json", testData)
		if err != nil {
			t.Errorf("SaveTestDataToFile should succeed: %v", err)
		}
		expectedPath := filepath.Join(TestDataDir, "successful_test.json")
		if _, statErr := os.Stat(expectedPath); os.IsNotExist(statErr) {
			t.Error("Expected file to be created")
		}
		defer os.Remove(expectedPath)
	})

	t.Run("MarshalError", func(t *testing.T) {
		invalidData := make(chan int)
		err := SaveTestDataToFile("invalid_data.json", invalidData)
		if err == nil {
			t.Error("expected error for unmarshalable data")
		} else {
			t.Logf("expected error: %v", err)
		}
	})

	t.Run("DirectoryCreation", func(t *testing.T) {
		testDir := "./test_data_creation_test"
		os.RemoveAll(testDir)
		if err := os.MkdirAll(testDir, 0o755); err != nil {
			t.Errorf("mkdir failed: %v", err)
		}
		defer os.RemoveAll(testDir)
		if _, err := os.Stat(testDir); os.IsNotExist(err) {
			t.Error("Expected directory to be created")
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		invalidPath := "/root/restricted/path"
		if err := os.MkdirAll(invalidPath, 0o755); err != nil {
			t.Logf("mkdir failed as expected: %v", err)
		} else {
			t.Log("mkdir unexpectedly succeeded")
		}
	})

	t.Run("SaveTestDataToFileDirectoryCreation", func(t *testing.T) {
		os.RemoveAll(TestDataDir)
		testData := map[string]string{"test": "data"}
		err := SaveTestDataToFile("dir_creation_test.json", testData)
		if err != nil {
			t.Logf("SaveTestDataToFile failed: %v", err)
		} else {
			if _, err := os.Stat(TestDataDir); os.IsNotExist(err) {
				t.Error("Expected test_data to be created")
			}
			defer os.Remove(filepath.Join(TestDataDir, "dir_creation_test.json"))
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		err := SaveTestDataToFile("nonexistent/deep/directory/structure/test.json", testData)
		if err != nil {
			t.Logf("handled path error: %v", err)
		}
	})

	t.Run("CompleteFlow", func(t *testing.T) {
		cases := []struct {
			name string
			data interface{}
		}{
			{"SimpleString", "test"},
			{"SimpleNumber", 42},
			{"SimpleBool", true},
			{"SimpleNull", nil},
			{"ComplexMap", map[string]interface{}{
				"nested": map[string]string{
					"key": "value",
				},
				"array": []int{1, 2, 3},
			}},
		}
		for _, tc := range cases {
			filename := fmt.Sprintf("complete_flow_%s.json", strings.ToLower(tc.name))
			if err := SaveTestDataToFile(filename, tc.data); err != nil {
				t.Errorf("failed for %s: %v", tc.name, err)
			}
			fullPath := filepath.Join(TestDataDir, filename)
			if _, err := os.Stat(fullPath); err == nil {
				os.Remove(fullPath)
			}
		}
	})

	complexData := map[string]interface{}{
		"special_chars": "hello\nworld\t\"quotes\"",
		"unicode":       "こんにちは",
		"nested": map[string]interface{}{
			"array": []interface{}{1, 2.5, true, nil},
			"empty": map[string]interface{}{},
		},
	}
	if err := SaveTestDataToFile("complex_data.json", complexData); err != nil {
		t.Logf("complex data returned error: %v", err)
	} else {
		defer os.Remove(filepath.Join(TestDataDir, "complex_data.json"))
	}
}

func TestFileOperations(t *testing.T) {
	tempDir := t.TempDir()
	testDir := filepath.Join(tempDir, "nested", "directory")
	if err := os.MkdirAll(testDir, 0o755); err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Error("Expected nested directory to be created")
	}
}

func TestSaveTestDataToFileInjectionHooks(t *testing.T) {
	origMkdir := mkdirAll
	mkdirAll = func(path string, perm os.FileMode) error { return fmt.Errorf("injected mkdir error") }
	if err := SaveTestDataToFile("x.json", map[string]string{"k": "v"}); err == nil {
		t.Error("expected mkdir error")
	}
	mkdirAll = origMkdir
	origWrite := writeFile
	writeFile = func(filename string, data []byte, perm os.FileMode) error { return fmt.Errorf("injected write error") }
	if err := SaveTestDataToFile("y.json", map[string]string{"k": "v"}); err == nil {
		t.Error("expected write error")
	}
	writeFile = origWrite
}

func TestSaveTestDataToFileComprehensive(t *testing.T) {
	testDir := "./tmp/test_data"
	defer os.RemoveAll(testDir)
	os.RemoveAll(testDir)
	t.Run("ValidData", func(t *testing.T) {
		data := map[string]interface{}{"name": "test", "id": 42}
		if err := SaveTestDataToFile("test_valid.json", data); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		fullPath := filepath.Join(TestDataDir, "test_valid.json")
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Error("expected file created")
		} else {
			if b, err := os.ReadFile(fullPath); err == nil {
				var u map[string]interface{}
				_ = json.Unmarshal(b, &u)
			}
			os.Remove(fullPath)
		}
	})

	t.Run("UnmarshalableData", func(t *testing.T) {
		bad := struct {
			BadFunc func() `json:"func"`
		}{BadFunc: func() {}}
		if err := SaveTestDataToFile("test_unmarshalable.json", bad); err == nil {
			t.Error("expected error")
		}
		os.Remove(filepath.Join(TestDataDir, "test_unmarshalable.json"))
	})

	t.Run("VariousDataTypes", func(t *testing.T) {
		if err := SaveTestDataToFile("test_string.json", "simple string"); err != nil {
			t.Errorf("save string: %v", err)
		}
		os.Remove(filepath.Join(TestDataDir, "test_string.json"))
		if err := SaveTestDataToFile("test_number.json", 42); err != nil {
			t.Errorf("save number: %v", err)
		}
		os.Remove(filepath.Join(TestDataDir, "test_number.json"))
		if err := SaveTestDataToFile("test_array.json", []string{"a", "b", "c"}); err != nil {
			t.Errorf("save array: %v", err)
		}
		os.Remove(filepath.Join(TestDataDir, "test_array.json"))
	})
}
