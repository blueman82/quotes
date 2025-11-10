# AGENTS.md

A simple, zero-configuration CLI tool that generates random inspiring quotes. Built with Go and Cobra, designed for scripting, automation, and integration with the Conductor orchestration system.

## Build & Test

```bash
# Build the binary locally
make quotes-build

# Run all tests with coverage
make quotes-test

# Run specific test
go test ./cmd/quotes -run TestFormatText -v

# Check code quality (fmt + vet + test)
make check
```

## Installation

```bash
# Install to $GOPATH/bin (system-wide)
make quotes-install

# The install target checks if $GOPATH/bin is in PATH
# and offers to add it to ~/.zshrc automatically
```

## Quick Usage

```bash
# Get a random quote
quotes

# JSON format for scripting
quotes --format json

# Multiple quotes with reproducible seed
quotes --count 3 --seed 42
```

## Architecture

**Core Pattern: Graceful Fallback**
- `LoadQuotes()` tries `~/.quotes.json`, falls back to 70+ hardcoded quotes
- Never fails - always returns valid quotes even with missing/invalid custom file
- No caching - checks `~/.quotes.json` on every run for live editing

**Package Structure:**
- `main.go` - Cobra CLI setup and command execution
- `types.go` - Quote struct, defaultQuotes (70+ quotes), SelectRandom()
- `loader.go` - LoadQuotes() with fallback logic
- `formatter.go` - FormatText(), FormatJSON(), FormatMarkdown()

**Random Selection:**
- Without `--seed`: Uses `time.Now().UnixNano()`
- With `--seed`: Reproducible output using provided seed
- Multiple quotes: Each gets `seed + index` to avoid duplicates

## Code Style

**Testing:**
- Table-driven tests for all major functions
- Target coverage: 70%+ overall, 90%+ for formatters/selection
- Test files: `*_test.go` mirror implementation files

**Adding New Formatters:**
1. Add format to `isValidFormat()` in `main.go`
2. Create `FormatX([]Quote) string` in `formatter.go`
3. Add case to switch in `runQuotes()`
4. Add table tests in `formatter_test.go`

**Quote File Format:**
- `~/.quotes.json` must be JSON array
- Each quote: `{"text": "...", "author": "..."}` (case-sensitive)
- Empty or invalid = automatic fallback to defaults

## Development Environment

**Requirements:**
- Go 1.25.4+
- Dependencies: `github.com/spf13/cobra` (only external dependency)

**Testing Changes:**
```bash
# Local test
./quotes --format json --count 3

# After modifying defaultQuotes, update test expectations
go test ./cmd/quotes -v

# Generate coverage HTML
make coverage
```

## Security Notes

- No external network calls
- Reads only `~/.quotes.json` (user's home directory)
- All input validation happens in `runQuotes()`:
  - Format must be: text|json|markdown
  - Count must be: 1-100
  - Invalid input returns clear error messages

## PR Guidelines

**Before Merging:**
- Run `make check` (fmt + vet + test)
- Ensure test coverage remains 70%+
- Update test expectations if modifying `defaultQuotes`
- Keep quote collection focused on programming/tech wisdom

**Commit Convention:**
- Use conventional commits (feat:, fix:, chore:, docs:)
- See recent commits for style examples

## Integration Context

Part of the Conductor project. Design decisions:
- **Zero-config**: Works immediately with no setup
- **Hardcoded quotes**: 70+ quotes embedded for reliability in restricted environments
- **Three formats**: text (human), JSON (scripts), markdown (docs)
- **Never fails**: Always provides output via fallback system

## Troubleshooting

```bash
# Command not found after install
echo $PATH | grep $(go env GOPATH)/bin
export PATH=$PATH:$(go env GOPATH)/bin

# Custom quotes not loading
cat ~/.quotes.json | jq .  # Validate JSON

# Test failures
go test ./cmd/quotes -v    # See detailed output
```
