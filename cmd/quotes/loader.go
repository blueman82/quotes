package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// LoadQuotes returns quotes from ~/.quotes.json if it exists and is valid,
// otherwise returns the default hardcoded quotes.
// Never returns nil or an empty slice - always provides usable quotes.
func LoadQuotes() []Quote {
	// Try to get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return defaultQuotes
	}

	// Check for override file
	overridePath := filepath.Join(home, ".quotes.json")
	data, err := os.ReadFile(overridePath)
	if err != nil {
		// File doesn't exist or is unreadable - use defaults
		return defaultQuotes
	}

	// Try to parse the override file
	var quotes []Quote
	if err := json.Unmarshal(data, &quotes); err != nil || len(quotes) == 0 {
		// Invalid JSON or empty array - fall back to defaults
		return defaultQuotes
	}

	// Successfully loaded override quotes
	return quotes
}
