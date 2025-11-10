package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// Helper function to capture command output
func executeCommand(cmd *cobra.Command, args ...string) (string, error) {
	// Save original stdout
	oldStdout := os.Stdout

	// Create pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Set args and execute
	cmd.SetArgs(args)
	err := cmd.Execute()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Reset command for next test
	cmd.SetArgs([]string{})
	format = "text"
	count = 1
	seed = 0

	return buf.String(), err
}

func TestQuotesCommand_DefaultBehavior(t *testing.T) {
	cmd := newRootCommand()
	output, err := executeCommand(cmd)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if output == "" {
		t.Error("expected output, got empty string")
	}

	// Should be text format (no JSON brackets, no markdown blockquotes)
	if strings.Contains(output, "{") || strings.Contains(output, "[") {
		t.Error("default output should be text format, not JSON")
	}
	if strings.HasPrefix(output, ">") {
		t.Error("default output should be text format, not markdown")
	}
}

func TestQuotesCommand_FormatFlag(t *testing.T) {
	tests := []struct {
		name           string
		format         string
		wantErr        bool
		validateOutput func(string) bool
	}{
		{
			name:    "text format",
			format:  "text",
			wantErr: false,
			validateOutput: func(output string) bool {
				return !strings.Contains(output, "{") && !strings.Contains(output, "##")
			},
		},
		{
			name:    "json format",
			format:  "json",
			wantErr: false,
			validateOutput: func(output string) bool {
				// Should be valid JSON array
				var data []interface{}
				return json.Unmarshal([]byte(output), &data) == nil
			},
		},
		{
			name:    "markdown format",
			format:  "markdown",
			wantErr: false,
			validateOutput: func(output string) bool {
				// Should contain markdown blockquote syntax
				return strings.Contains(output, ">") && strings.Contains(output, "—")
			},
		},
		{
			name:    "invalid format",
			format:  "xml",
			wantErr: true,
			validateOutput: func(output string) bool {
				return true // Don't validate output on error
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newRootCommand()
			output, err := executeCommand(cmd, "--format", tt.format)

			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr = %v, got error: %v", tt.wantErr, err)
				return
			}

			if !tt.wantErr && !tt.validateOutput(output) {
				t.Errorf("output validation failed for format %s: %s", tt.format, output)
			}
		})
	}
}

func TestQuotesCommand_CountFlag(t *testing.T) {
	tests := []struct {
		name         string
		count        string
		wantErr      bool
		validateFunc func(string) bool
	}{
		{
			name:    "count 1",
			count:   "1",
			wantErr: false,
			validateFunc: func(output string) bool {
				return output != ""
			},
		},
		{
			name:    "count 5",
			count:   "5",
			wantErr: false,
			validateFunc: func(output string) bool {
				// In text format, should have multiple quotes separated by blank lines
				lines := strings.Split(strings.TrimSpace(output), "\n")
				return len(lines) > 5 // Each quote has text + author + separator
			},
		},
		{
			name:    "count 0",
			count:   "0",
			wantErr: true,
			validateFunc: func(output string) bool {
				return true
			},
		},
		{
			name:    "count 101",
			count:   "101",
			wantErr: true,
			validateFunc: func(output string) bool {
				return true
			},
		},
		{
			name:    "negative count",
			count:   "-1",
			wantErr: true,
			validateFunc: func(output string) bool {
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newRootCommand()
			output, err := executeCommand(cmd, "--count", tt.count)

			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr = %v, got error: %v", tt.wantErr, err)
				return
			}

			if !tt.wantErr && !tt.validateFunc(output) {
				t.Errorf("output validation failed for count %s", tt.count)
			}
		})
	}
}

func TestQuotesCommand_SeedFlag(t *testing.T) {
	cmd1 := newRootCommand()
	output1, err1 := executeCommand(cmd1, "--seed", "42", "--count", "3")
	if err1 != nil {
		t.Fatalf("first execution failed: %v", err1)
	}

	cmd2 := newRootCommand()
	output2, err2 := executeCommand(cmd2, "--seed", "42", "--count", "3")
	if err2 != nil {
		t.Fatalf("second execution failed: %v", err2)
	}

	if output1 != output2 {
		t.Error("same seed should produce identical output")
		t.Logf("Output 1:\n%s", output1)
		t.Logf("Output 2:\n%s", output2)
	}
}

func TestQuotesCommand_FlagCombinations(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "format json + count 3",
			args:    []string{"--format", "json", "--count", "3"},
			wantErr: false,
		},
		{
			name:    "all flags",
			args:    []string{"--format", "markdown", "--count", "5", "--seed", "99"},
			wantErr: false,
		},
		{
			name:    "short flags",
			args:    []string{"-f", "text", "-n", "2"},
			wantErr: false,
		},
		{
			name:    "invalid format + valid count",
			args:    []string{"--format", "yaml", "--count", "3"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newRootCommand()
			_, err := executeCommand(cmd, tt.args...)

			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr = %v, got error: %v", tt.wantErr, err)
			}
		})
	}
}

func TestQuotesCommand_InvalidFormat(t *testing.T) {
	invalidFormats := []string{"xml", "yaml", "csv", "html", ""}

	for _, format := range invalidFormats {
		t.Run("format_"+format, func(t *testing.T) {
			cmd := newRootCommand()
			_, err := executeCommand(cmd, "--format", format)

			if err == nil {
				t.Errorf("expected error for invalid format %q, got nil", format)
			}

			if err != nil && !strings.Contains(err.Error(), "invalid format") {
				t.Errorf("error should mention 'invalid format', got: %v", err)
			}
		})
	}
}

func TestQuotesCommand_CountWithJSON(t *testing.T) {
	cmd := newRootCommand()
	output, err := executeCommand(cmd, "--format", "json", "--count", "3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var quotes []map[string]string
	if err := json.Unmarshal([]byte(output), &quotes); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}

	if len(quotes) != 3 {
		t.Errorf("expected 3 quotes, got %d", len(quotes))
	}
}
