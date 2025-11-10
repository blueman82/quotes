# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based CLI tool that generates random inspiring quotes with zero configuration. It's built with Cobra and designed to work seamlessly with the Conductor orchestration system.

## Development Commands

### Build and Install
```bash
# Build binary locally to ./quotes
make quotes-build

# Install to $GOPATH/bin (makes it system-wide)
make quotes-install

# The install target automatically checks if $GOPATH/bin is in PATH
# and offers to add it to ~/.zshrc if needed
```

### Testing
```bash
# Run all tests with coverage report
make quotes-test

# Run tests with verbose output
make test-verbose

# Generate HTML coverage report
make coverage

# Run specific test
go test ./cmd/quotes -run TestFormatText -v

# Run tests for a specific file
go test ./cmd/quotes -run TestFormatter -v
```

### Code Quality
```bash
# Format code
make fmt

# Run go vet
make vet

# Tidy dependencies
make tidy

# Run all checks (fmt + vet + test)
make check

# Clean build artifacts
make clean
```

## Architecture

### Core Design Pattern: Graceful Fallback

The CLI follows a "never fail" philosophy - it always provides usable output by falling back to defaults when user customization fails.

**Quote Loading Flow:**
1. Try to load `~/.quotes.json` (user's custom quotes)
2. If file doesn't exist → use `defaultQuotes` from `types.go`
3. If file is invalid JSON → use `defaultQuotes`
4. If file is empty array → use `defaultQuotes`

This is implemented in `loader.go:LoadQuotes()` which **always returns a valid slice**, never nil or empty.

### Random Selection with Seed

The quote selection uses Go's `math/rand` with explicit seeding:
- **Without `--seed`**: Uses `time.Now().UnixNano()` for randomness
- **With `--seed`**: Uses provided seed for reproducible output
- **Multiple quotes**: Each quote gets `seed + index` to avoid duplicates while maintaining reproducibility

See `types.go:SelectRandom()` and `main.go:runQuotes()` lines 59-70.

### Three-Tier Formatting System

All formatters in `formatter.go` accept `[]Quote` and return `string`:

1. **FormatText**: Numbered list for count > 1, plain text for single quote
2. **FormatJSON**: JSON array with proper indentation, handles nil gracefully
3. **FormatMarkdown**: Blockquote format with em-dash attribution

The formatter selection happens in `main.go:runQuotes()` lines 72-81 using a switch statement.

### Cobra CLI Structure

Command setup follows a factory pattern:
- `newRootCommand()` creates the command structure
- Flags are defined with short and long forms (-f/--format, -n/--count)
- `runQuotes()` is the main execution function attached to RunE
- Validation happens at the start of `runQuotes()` for format and count ranges

## Key Implementation Details

### Why tests might fail
- If modifying `defaultQuotes` in `types.go`, update corresponding test expectations
- Formatters have specific output formats - check `formatter_test.go` for expected strings
- The JSON formatter uses 2-space indentation - maintain this for test consistency

### Adding new output formats
1. Add format validation to `isValidFormat()` in `main.go`
2. Create formatter function in `formatter.go` with signature: `func FormatX([]Quote) string`
3. Add case to switch statement in `runQuotes()`
4. Add table-driven tests in `formatter_test.go`

### Custom quote file behavior
The CLI checks `~/.quotes.json` **on every run** - no caching. This allows users to edit their quotes without restarting anything. The file must be valid JSON array of objects with "text" and "author" fields (case-sensitive).

### Build versioning
The Makefile injects version info at build time using ldflags:
- `VERSION`: From environment or defaults to 1.0.0
- `BUILD_TIME`: UTC timestamp
- `GIT_COMMIT`: Short git hash

These are injected into `main` package variables.

## Testing Strategy

Tests follow table-driven design for all major functions:
- `types_test.go`: Tests `SelectRandom()` with various inputs including edge cases
- `formatter_test.go`: Tests all three formatters with single/multiple quotes
- `loader_test.go`: Tests quote loading with valid/invalid/missing files
- `main_test.go`: Integration tests for the full command flow

Target coverage: 70%+ overall, 90%+ for critical paths (formatters, selection logic).

## Integration Context

This CLI is part of the Conductor project and designed for:
- Embedding in scripts with JSON output
- Daily quote generation with date-based seeds
- README/documentation generation with markdown format
- Zero-config usage in automated systems

The quotes collection (70+ programming/tech quotes) is intentionally hardcoded to ensure the CLI always works even in restricted environments.
