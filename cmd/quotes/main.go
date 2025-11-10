package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var (
	format string
	count  int
	seed   int64
)

// newRootCommand creates and returns the root command
func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quotes",
		Short: "Generate random inspiring quotes",
		Long:  "A CLI tool to generate random inspiring quotes with various output formats",
		RunE:  runQuotes,
	}

	cmd.Flags().StringVarP(&format, "format", "f", "text", "Output format: text|json|markdown")
	cmd.Flags().IntVarP(&count, "count", "n", 1, "Number of quotes (1-100)")
	cmd.Flags().Int64Var(&seed, "seed", 0, "Random seed for reproducibility")

	return cmd
}

// isValidFormat checks if the provided format is valid
func isValidFormat(format string) bool {
	switch format {
	case "text", "json", "markdown":
		return true
	default:
		return false
	}
}

// runQuotes is the main command execution function
func runQuotes(cmd *cobra.Command, args []string) error {
	// Validate format
	if !isValidFormat(format) {
		return fmt.Errorf("invalid format: %s (must be one of: text, json, markdown)", format)
	}

	// Validate count
	if count < 1 || count > 100 {
		return fmt.Errorf("count must be 1-100, got %d", count)
	}

	// Load quotes
	quotes := LoadQuotes()

	// Use current time as seed if not specified
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	// Select random quotes
	selected := make([]Quote, 0, count)
	for i := 0; i < count; i++ {
		// Use different seed for each selection to avoid duplicates
		q, _ := SelectRandom(quotes, seed+int64(i))
		selected = append(selected, q)
	}

	// Format and print
	var output string
	switch format {
	case "json":
		output = FormatJSON(selected)
	case "markdown":
		output = FormatMarkdown(selected)
	default:
		output = FormatText(selected)
	}

	fmt.Print(output)
	return nil
}

func main() {
	// Initialize random seed for reproducibility
	rand.Seed(time.Now().UnixNano())

	// Create and execute root command
	rootCmd := newRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
