# Quotes CLI

A simple, zero-configuration CLI tool that generates random inspiring quotes with multiple output formats.

## Overview

The Quotes CLI is a lightweight command-line tool designed to provide random inspiring quotes for your terminal, scripts, or daily motivation. Built with Go and the Cobra CLI framework, it features:

- **Zero Configuration**: Works immediately after installation with 50+ hardcoded quotes
- **Flexible Output**: Three formats (text, JSON, markdown) for different use cases
- **Reproducible Randomness**: Optional seed flag for deterministic output
- **Customizable**: Override default quotes with your own via `~/.quotes.json`
- **Fast**: Single binary with no external dependencies

## Installation

### Using Make (Recommended)

```bash
make install
```

This will build and install the `quotes` binary to your `$GOPATH/bin`, making it available globally.

### Manual Installation

```bash
go install ./cmd/quotes
```

### Building from Source

```bash
# Build local binary
make quotes-build

# Run locally
./quotes
```

## Usage

### Basic Usage

Get a random quote in plain text format:

```bash
quotes
```

**Output:**
```
Be the change you wish to see in the world
   - Mahatma Gandhi
```

### Output Formats

#### Text Format (Default)

Simple, human-readable format with author attribution:

```bash
quotes --format text
```

**Output:**
```
Code is poetry
   - Unknown
```

#### JSON Format

Perfect for scripting and programmatic use:

```bash
quotes --format json
```

**Output:**
```json
[
  {
    "text": "The only way to do great work is to love what you do",
    "author": "Steve Jobs"
  }
]
```

#### Markdown Format

Great for documentation and README files:

```bash
quotes --format markdown
```

**Output:**
```markdown
> The only way to do great work is to love what you do

— Steve Jobs
```

### Multiple Quotes

Get multiple random quotes at once (1-100):

```bash
quotes --count 3
```

**Output:**
```
1. Be the change you wish to see in the world
   - Mahatma Gandhi
2. Code is poetry
   - Unknown
3. The only way to do great work is to love what you do
   - Steve Jobs
```

Combine with JSON format for data processing:

```bash
quotes --format json --count 5
```

### Reproducible Randomness

Use the `--seed` flag for deterministic output (useful for testing or consistent daily quotes):

```bash
quotes --seed 42
```

Run multiple times with the same seed to get the same quote:

```bash
quotes --seed 42 --count 3
quotes --seed 42 --count 3  # Same output
```

### Combining Flags

All flags can be combined:

```bash
# Get 5 quotes in markdown format with reproducible randomness
quotes --format markdown --count 5 --seed 12345
```

## Command Reference

```
quotes [flags]

Flags:
  -f, --format string   Output format: text|json|markdown (default "text")
  -n, --count int       Number of quotes (1-100) (default 1)
      --seed int        Random seed for reproducibility (default 0, random)
  -h, --help            Help for quotes
```

### Flag Details

- **--format, -f**: Controls output format
  - `text`: Plain text with author attribution (default)
  - `json`: JSON array of quote objects
  - `markdown`: Markdown blockquote format

- **--count, -n**: Number of quotes to generate
  - Range: 1-100
  - Default: 1
  - With count > 1, text format uses numbered list

- **--seed**: Random number generator seed
  - Default: 0 (uses current time, random)
  - Any integer value produces deterministic output
  - Useful for testing, daily quotes, or reproducible scripts

## Customization

### Custom Quote Collection

Override the default quotes by creating `~/.quotes.json`:

```json
[
  {
    "text": "Your custom inspiring quote here",
    "author": "Your Name"
  },
  {
    "text": "Another custom quote",
    "author": "Another Author"
  }
]
```

**Features:**
- Graceful fallback: If the file is missing, invalid, or empty, the CLI falls back to hardcoded defaults
- No configuration required: The CLI works immediately without any setup
- Standard JSON format: Easy to create and edit

**Example:**

```bash
# Create custom quotes
cat > ~/.quotes.json << 'EOF'
[
  {
    "text": "Keep calm and code on",
    "author": "Developer Wisdom"
  },
  {
    "text": "Simplicity is the ultimate sophistication",
    "author": "Leonardo da Vinci"
  }
]
EOF

# Use custom quotes
quotes
```

### Quote Format

Each quote in `~/.quotes.json` must have:
- `text`: The quote text (string, required)
- `author`: The quote author (string, required)

## Integration with Conductor

The Quotes CLI is designed to work seamlessly with the Conductor orchestration system.

### Using Quotes in Conductor Tasks

```yaml
tasks:
  - task_number: 1
    name: "Generate motivational quote"
    agent: "general-purpose"
    commands:
      - "quotes --format text"
```

### Scripting Examples

**Daily Quote Script:**

```bash
#!/bin/bash
# daily-quote.sh - Get quote based on current date

SEED=$(date +%Y%m%d)
quotes --seed $SEED --format text
```

**Quote of the Day in JSON:**

```bash
quotes --seed $(date +%Y%m%d) --format json > quote-of-the-day.json
```

**Multiple Quotes for README:**

```bash
echo "# Daily Inspiration" > QUOTES.md
quotes --format markdown --count 3 >> QUOTES.md
```

### Pipeline Integration

```bash
# Use in a pipeline
quotes --format json | jq '.[] | .text'

# Store in variable
QUOTE=$(quotes --format text)
echo "Today's wisdom: $QUOTE"

# Error handling
if quotes --format json > quotes.json; then
    echo "Quotes saved successfully"
else
    echo "Failed to generate quotes"
fi
```

## Examples

### Quick Motivation

```bash
quotes
```

### Generate README Content

```bash
quotes --format markdown --count 3 >> README.md
```

### Daily Quote with Seed

```bash
# Same quote all day, changes daily
quotes --seed $(date +%Y%m%d)
```

### Export to JSON File

```bash
quotes --format json --count 10 > quotes.json
```

### Multiple Formats

```bash
# Text
quotes --format text > quote.txt

# JSON
quotes --format json > quote.json

# Markdown
quotes --format markdown > quote.md
```

## Error Handling

The Quotes CLI is designed to never fail:

- **Missing ~/.quotes.json**: Falls back to default quotes
- **Invalid JSON**: Falls back to default quotes
- **Empty quote file**: Falls back to default quotes
- **Invalid flags**: Returns clear error message with help text
- **Count out of range**: Error message with valid range

**Example Error:**

```bash
quotes --format xml
# Error: invalid format: xml
# Use: text, json, or markdown
```

## Testing

Run the test suite:

```bash
# Run all tests
make quotes-test

# Run with coverage
go test ./cmd/quotes -cover

# Run specific tests
go test ./cmd/quotes -run TestFormatText
```

Expected test coverage:
- Overall: 70%+
- Critical paths (formatters, selection): 90%+

## Build Information

```bash
# Build the binary
make quotes-build

# Run tests
make quotes-test

# Install globally
make quotes-install

# Clean build artifacts
make clean
```

## Use Cases

- **Daily Motivation**: Add to your shell profile for daily quotes
- **README Files**: Generate inspirational quotes for documentation
- **Scripts**: Add motivational messages to build scripts
- **Testing**: Use seeded quotes for reproducible test data
- **Learning**: Example of clean CLI design with Cobra
- **Integration**: Embed in larger systems via JSON output

## Technical Details

- **Language**: Go 1.21+
- **Framework**: Cobra CLI
- **Dependencies**: Only Cobra (same as Conductor)
- **Build System**: Makefile
- **Testing**: Go standard testing with table-driven tests
- **Binary Size**: ~5MB (includes 50+ default quotes)
- **Performance**: Instant execution (<10ms)

## Troubleshooting

### Command Not Found

If `quotes` is not found after installation:

```bash
# Check GOPATH
echo $GOPATH

# Add to PATH if needed
export PATH=$PATH:$GOPATH/bin

# Or reinstall
make quotes-install
```

### Custom Quotes Not Loading

```bash
# Check file exists
ls -la ~/.quotes.json

# Validate JSON
cat ~/.quotes.json | jq .

# Test with defaults
mv ~/.quotes.json ~/.quotes.json.backup
quotes
```

### Invalid Format Error

Ensure format is one of: `text`, `json`, or `markdown`

```bash
# Correct
quotes --format json

# Incorrect
quotes --format xml  # Error
```

## Contributing

The Quotes CLI follows these principles:

- **Simplicity**: Single purpose, no feature creep
- **Zero Config**: Works out of the box
- **Test-Driven**: TDD approach with 70%+ coverage
- **UNIX Philosophy**: Do one thing well
- **Graceful Degradation**: Never fails, always has fallback

## License

Part of the Conductor project.

## See Also

- [Conductor Documentation](../README.md)
- [Go Cobra Framework](https://github.com/spf13/cobra)
- [Go Standard Testing](https://golang.org/pkg/testing/)
