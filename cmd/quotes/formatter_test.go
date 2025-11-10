package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestFormatText(t *testing.T) {
	tests := []struct {
		name   string
		quotes []Quote
		want   string
	}{
		{
			name:   "single quote",
			quotes: []Quote{{Text: "Be", Author: "Gandhi"}},
			want:   "Be\n   - Gandhi\n",
		},
		{
			name: "multiple quotes",
			quotes: []Quote{
				{Text: "Be", Author: "Gandhi"},
				{Text: "Code", Author: "Unknown"},
			},
			want: "1. Be\n   - Gandhi\n2. Code\n   - Unknown\n",
		},
		{
			name: "three quotes",
			quotes: []Quote{
				{Text: "First", Author: "Author1"},
				{Text: "Second", Author: "Author2"},
				{Text: "Third", Author: "Author3"},
			},
			want: "1. First\n   - Author1\n2. Second\n   - Author2\n3. Third\n   - Author3\n",
		},
		{
			name:   "empty quotes slice",
			quotes: []Quote{},
			want:   "",
		},
		{
			name:   "quote with empty text",
			quotes: []Quote{{Text: "", Author: "Someone"}},
			want:   "\n   - Someone\n",
		},
		{
			name:   "quote with empty author",
			quotes: []Quote{{Text: "Anonymous quote", Author: ""}},
			want:   "Anonymous quote\n   - \n",
		},
		{
			name:   "quote with special characters",
			quotes: []Quote{{Text: "Quote with \"quotes\" and \nnewline", Author: "Test"}},
			want:   "Quote with \"quotes\" and \nnewline\n   - Test\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatText(tt.quotes)
			if got != tt.want {
				t.Errorf("FormatText() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatJSON(t *testing.T) {
	tests := []struct {
		name   string
		quotes []Quote
		want   []Quote
	}{
		{
			name:   "single quote",
			quotes: []Quote{{Text: "Be", Author: "Gandhi"}},
			want:   []Quote{{Text: "Be", Author: "Gandhi"}},
		},
		{
			name: "multiple quotes",
			quotes: []Quote{
				{Text: "Be", Author: "Gandhi"},
				{Text: "Code", Author: "Unknown"},
			},
			want: []Quote{
				{Text: "Be", Author: "Gandhi"},
				{Text: "Code", Author: "Unknown"},
			},
		},
		{
			name:   "empty quotes slice",
			quotes: []Quote{},
			want:   []Quote{},
		},
		{
			name:   "quote with special characters",
			quotes: []Quote{{Text: "Quote with \"quotes\" and \nnewline", Author: "Test"}},
			want:   []Quote{{Text: "Quote with \"quotes\" and \nnewline", Author: "Test"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatJSON(tt.quotes)

			// Verify it's valid JSON
			var parsed []Quote
			if err := json.Unmarshal([]byte(got), &parsed); err != nil {
				t.Errorf("FormatJSON() produced invalid JSON: %v", err)
				return
			}

			// Verify it's indented (should contain newlines and spaces)
			if len(tt.quotes) > 0 && !strings.Contains(got, "\n") {
				t.Errorf("FormatJSON() should produce indented output, got: %q", got)
			}

			// Verify the parsed content matches expected
			if len(parsed) != len(tt.want) {
				t.Errorf("FormatJSON() parsed length = %d, want %d", len(parsed), len(tt.want))
				return
			}

			for i := range parsed {
				if parsed[i].Text != tt.want[i].Text || parsed[i].Author != tt.want[i].Author {
					t.Errorf("FormatJSON() parsed[%d] = %+v, want %+v", i, parsed[i], tt.want[i])
				}
			}
		})
	}
}

func TestFormatMarkdown(t *testing.T) {
	tests := []struct {
		name   string
		quotes []Quote
		want   string
	}{
		{
			name:   "single quote",
			quotes: []Quote{{Text: "Be", Author: "Gandhi"}},
			want:   "> Be\n\n— Gandhi\n\n",
		},
		{
			name: "multiple quotes",
			quotes: []Quote{
				{Text: "Be", Author: "Gandhi"},
				{Text: "Code", Author: "Unknown"},
			},
			want: "> Be\n\n— Gandhi\n\n> Code\n\n— Unknown\n\n",
		},
		{
			name: "three quotes",
			quotes: []Quote{
				{Text: "First", Author: "Author1"},
				{Text: "Second", Author: "Author2"},
				{Text: "Third", Author: "Author3"},
			},
			want: "> First\n\n— Author1\n\n> Second\n\n— Author2\n\n> Third\n\n— Author3\n\n",
		},
		{
			name:   "empty quotes slice",
			quotes: []Quote{},
			want:   "",
		},
		{
			name:   "quote with empty text",
			quotes: []Quote{{Text: "", Author: "Someone"}},
			want:   "> \n\n— Someone\n\n",
		},
		{
			name:   "quote with empty author",
			quotes: []Quote{{Text: "Anonymous quote", Author: ""}},
			want:   "> Anonymous quote\n\n— \n\n",
		},
		{
			name:   "quote with special characters",
			quotes: []Quote{{Text: "Quote with \"quotes\"", Author: "Test"}},
			want:   "> Quote with \"quotes\"\n\n— Test\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatMarkdown(tt.quotes)
			if got != tt.want {
				t.Errorf("FormatMarkdown() = %q, want %q", got, tt.want)
			}

			// Verify markdown quote block syntax
			if len(tt.quotes) > 0 {
				if !strings.Contains(got, ">") {
					t.Errorf("FormatMarkdown() should contain '>' for quote blocks")
				}
				if !strings.Contains(got, "—") {
					t.Errorf("FormatMarkdown() should contain '—' for attribution")
				}
			}
		})
	}
}

func TestFormatters_NilInput(t *testing.T) {
	// Test that formatters handle nil slices gracefully
	t.Run("FormatText with nil", func(t *testing.T) {
		got := FormatText(nil)
		if got != "" {
			t.Errorf("FormatText(nil) = %q, want empty string", got)
		}
	})

	t.Run("FormatJSON with nil", func(t *testing.T) {
		got := FormatJSON(nil)
		// Should produce valid JSON for empty array
		var parsed []Quote
		if err := json.Unmarshal([]byte(got), &parsed); err != nil {
			t.Errorf("FormatJSON(nil) produced invalid JSON: %v", err)
		}
	})

	t.Run("FormatMarkdown with nil", func(t *testing.T) {
		got := FormatMarkdown(nil)
		if got != "" {
			t.Errorf("FormatMarkdown(nil) = %q, want empty string", got)
		}
	})
}
