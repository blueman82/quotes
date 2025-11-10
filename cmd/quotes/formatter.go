package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FormatText formats quotes as plain text with author attribution.
// For a single quote, outputs: "Text\n   - Author\n"
// For multiple quotes, outputs numbered list: "1. Text\n   - Author\n"
func FormatText(quotes []Quote) string {
	var result strings.Builder

	for i, q := range quotes {
		if len(quotes) > 1 {
			fmt.Fprintf(&result, "%d. ", i+1)
		}
		fmt.Fprintf(&result, "%s\n   - %s\n", q.Text, q.Author)
	}

	return result.String()
}

// FormatJSON formats quotes as indented JSON array.
// Returns a JSON array of objects with "Text" and "Author" fields.
func FormatJSON(quotes []Quote) string {
	// Handle nil slice by treating it as empty array
	if quotes == nil {
		quotes = []Quote{}
	}

	b, err := json.MarshalIndent(quotes, "", "  ")
	if err != nil {
		// This should rarely happen with simple Quote structs
		// but handle it gracefully by returning empty array
		return "[]"
	}

	return string(b)
}

// FormatMarkdown formats quotes as markdown quote blocks.
// Each quote is formatted as: "> Text\n\n— Author\n\n"
func FormatMarkdown(quotes []Quote) string {
	var result strings.Builder

	for _, q := range quotes {
		fmt.Fprintf(&result, "> %s\n\n— %s\n\n", q.Text, q.Author)
	}

	return result.String()
}
