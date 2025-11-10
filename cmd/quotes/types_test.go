package main

import (
	"testing"
)

// Test fixtures
var sampleQuotes = []Quote{
	{Text: "Be the change you wish to see in the world", Author: "Gandhi"},
	{Text: "Code is poetry", Author: "Unknown"},
	{Text: "The only way to do great work is to love what you do", Author: "Steve Jobs"},
	{Text: "Innovation distinguishes between a leader and a follower", Author: "Steve Jobs"},
	{Text: "Stay hungry, stay foolish", Author: "Steve Jobs"},
}

func TestQuoteStruct(t *testing.T) {
	quote := Quote{
		Text:   "Test quote",
		Author: "Test author",
	}

	if quote.Text != "Test quote" {
		t.Errorf("Quote.Text = %v, want %v", quote.Text, "Test quote")
	}

	if quote.Author != "Test author" {
		t.Errorf("Quote.Author = %v, want %v", quote.Author, "Test author")
	}
}

func TestRandomSelection_DeterministicSeed(t *testing.T) {
	seed := int64(42)

	// Select with same seed multiple times
	result1, err1 := SelectRandom(sampleQuotes, seed)
	if err1 != nil {
		t.Fatalf("SelectRandom failed: %v", err1)
	}

	result2, err2 := SelectRandom(sampleQuotes, seed)
	if err2 != nil {
		t.Fatalf("SelectRandom failed: %v", err2)
	}

	// Should get the same quote
	if result1.Text != result2.Text || result1.Author != result2.Author {
		t.Errorf("Same seed produced different results: got %v and %v", result1, result2)
	}
}

func TestRandomSelection_DifferentSeeds(t *testing.T) {
	tests := []struct {
		name string
		seed int64
		want int // expected index
	}{
		{"seed 42", 42, -1}, // We'll verify it's deterministic, not the exact index
		{"seed 99", 99, -1},
		{"seed 1000", 1000, -1},
	}

	results := make(map[int64]Quote)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectRandom(sampleQuotes, tt.seed)
			if err != nil {
				t.Fatalf("SelectRandom failed: %v", err)
			}

			// Verify result is from our sample quotes
			found := false
			for _, q := range sampleQuotes {
				if q.Text == got.Text && q.Author == got.Author {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("SelectRandom returned quote not in sample list: %v", got)
			}

			// Store result for comparison
			results[tt.seed] = got
		})
	}

	// Verify at least some different seeds produce different results
	// (with 5 quotes and 3 different seeds, we expect at least one difference)
	allSame := true
	firstQuote := results[42]
	for _, q := range results {
		if q.Text != firstQuote.Text {
			allSame = false
			break
		}
	}
	if allSame {
		t.Errorf("All different seeds produced the same quote, expected some variation")
	}
}

func TestRandomSelection_EdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		quotes    []Quote
		seed      int64
		wantError bool
	}{
		{
			name:      "empty list",
			quotes:    []Quote{},
			seed:      42,
			wantError: true,
		},
		{
			name: "single item",
			quotes: []Quote{
				{Text: "Only quote", Author: "Solo"},
			},
			seed:      42,
			wantError: false,
		},
		{
			name:      "negative seed",
			quotes:    sampleQuotes,
			seed:      -42,
			wantError: false,
		},
		{
			name:      "zero seed",
			quotes:    sampleQuotes,
			seed:      0,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectRandom(tt.quotes, tt.seed)

			if tt.wantError {
				if err == nil {
					t.Errorf("SelectRandom() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("SelectRandom() unexpected error: %v", err)
				return
			}

			// For single item, verify it's the correct one
			if len(tt.quotes) == 1 {
				if got.Text != tt.quotes[0].Text || got.Author != tt.quotes[0].Author {
					t.Errorf("SelectRandom() = %v, want %v", got, tt.quotes[0])
				}
			}

			// For non-empty lists, verify result is from the list
			found := false
			for _, q := range tt.quotes {
				if q.Text == got.Text && q.Author == got.Author {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("SelectRandom() returned quote not in input list: %v", got)
			}
		})
	}
}
