package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadQuotes_DefaultsWhenNoOverride(t *testing.T) {
	// Remove any existing override file
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get user home directory")
	}

	overridePath := filepath.Join(home, ".quotes.json")
	os.Remove(overridePath)

	quotes := LoadQuotes()

	if quotes == nil {
		t.Fatal("LoadQuotes returned nil")
	}

	if len(quotes) == 0 {
		t.Fatal("LoadQuotes returned empty slice")
	}

	// Should have default quotes (at least 50)
	if len(quotes) < 50 {
		t.Errorf("Expected at least 50 default quotes, got %d", len(quotes))
	}
}

func TestLoadQuotes_OverrideWhenFileExists(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get user home directory")
	}

	overridePath := filepath.Join(home, ".quotes.json")

	// Create a temporary override file
	testQuotes := `[
		{"text": "Test quote 1", "author": "Test Author 1"},
		{"text": "Test quote 2", "author": "Test Author 2"},
		{"text": "Test quote 3", "author": "Test Author 3"}
	]`

	err = os.WriteFile(overridePath, []byte(testQuotes), 0644)
	if err != nil {
		t.Fatalf("Failed to create test override file: %v", err)
	}
	defer os.Remove(overridePath)

	quotes := LoadQuotes()

	if quotes == nil {
		t.Fatal("LoadQuotes returned nil")
	}

	if len(quotes) != 3 {
		t.Errorf("Expected 3 override quotes, got %d", len(quotes))
	}

	if len(quotes) > 0 && quotes[0].Text != "Test quote 1" {
		t.Errorf("Expected first quote text to be 'Test quote 1', got '%s'", quotes[0].Text)
	}
}

func TestLoadQuotes_InvalidJSON_FallsBackToDefaults(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get user home directory")
	}

	overridePath := filepath.Join(home, ".quotes.json")

	// Create an invalid JSON file
	invalidJSON := `{this is not valid json`

	err = os.WriteFile(overridePath, []byte(invalidJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create test override file: %v", err)
	}
	defer os.Remove(overridePath)

	quotes := LoadQuotes()

	if quotes == nil {
		t.Fatal("LoadQuotes returned nil")
	}

	if len(quotes) == 0 {
		t.Fatal("LoadQuotes returned empty slice")
	}

	// Should fall back to defaults
	if len(quotes) < 50 {
		t.Errorf("Expected at least 50 default quotes on invalid JSON, got %d", len(quotes))
	}
}

func TestLoadQuotes_MissingFile_UsesDefaults(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get user home directory")
	}

	overridePath := filepath.Join(home, ".quotes.json")

	// Ensure file doesn't exist
	os.Remove(overridePath)

	quotes := LoadQuotes()

	if quotes == nil {
		t.Fatal("LoadQuotes returned nil")
	}

	if len(quotes) == 0 {
		t.Fatal("LoadQuotes returned empty slice")
	}

	// Should use defaults
	if len(quotes) < 50 {
		t.Errorf("Expected at least 50 default quotes, got %d", len(quotes))
	}
}

func TestLoadQuotes_EmptyJSONArray_FallsBackToDefaults(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get user home directory")
	}

	overridePath := filepath.Join(home, ".quotes.json")

	// Create an empty JSON array
	emptyJSON := `[]`

	err = os.WriteFile(overridePath, []byte(emptyJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create test override file: %v", err)
	}
	defer os.Remove(overridePath)

	quotes := LoadQuotes()

	if quotes == nil {
		t.Fatal("LoadQuotes returned nil")
	}

	if len(quotes) == 0 {
		t.Fatal("LoadQuotes returned empty slice")
	}

	// Should fall back to defaults when array is empty
	if len(quotes) < 50 {
		t.Errorf("Expected at least 50 default quotes on empty array, got %d", len(quotes))
	}
}

func TestLoadQuotes_NeverReturnsNilOrEmpty(t *testing.T) {
	// Test multiple times to ensure consistency
	for i := 0; i < 5; i++ {
		quotes := LoadQuotes()

		if quotes == nil {
			t.Fatal("LoadQuotes returned nil")
		}

		if len(quotes) == 0 {
			t.Fatal("LoadQuotes returned empty slice")
		}
	}
}
