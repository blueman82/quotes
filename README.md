# Quotes CLI

A simple, zero-configuration CLI tool that generates random inspiring quotes.

## Quick Start

```bash
# Install
make install

# Get a random quote
quotes

# Get multiple quotes in JSON format
quotes --format json --count 3

# Reproducible quotes with seed
quotes --seed 42
```

## Features

- **Zero Configuration**: Works immediately with 50+ built-in quotes
- **Multiple Formats**: Text, JSON, and Markdown output
- **Reproducible**: Optional seed for deterministic output
- **Customizable**: Override quotes with `~/.quotes.json`
- **Fast**: Single binary, instant execution

## Documentation

For complete documentation, see [docs/quotes-cli.md](docs/quotes-cli.md)

Topics covered:
- Installation and setup
- Usage examples
- Output formats (text, JSON, markdown)
- Custom quote collections
- Integration with Conductor
- Scripting examples
- Testing and build information

## Examples

**Basic usage:**
```bash
quotes
```

**JSON output:**
```bash
quotes --format json
```

**Multiple quotes:**
```bash
quotes --count 5
```

**Markdown for README:**
```bash
quotes --format markdown >> README.md
```

**Daily quote script:**
```bash
quotes --seed $(date +%Y%m%d)
```

## Customization

Create `~/.quotes.json` to use your own quotes:

```json
[
  {
    "text": "Your inspiring quote here",
    "author": "Your Name"
  }
]
```

## Building

```bash
# Build binary
make quotes-build

# Run tests
make quotes-test

# Install globally
make quotes-install
```

## Command Reference

```
quotes [flags]

Flags:
  -f, --format string   Output format: text|json|markdown (default "text")
  -n, --count int       Number of quotes (1-100) (default 1)
      --seed int        Random seed for reproducibility
  -h, --help            Help for quotes
```

## Integration with Conductor

The Quotes CLI is designed to work seamlessly with Conductor for orchestrated task execution. See [docs/quotes-cli.md](docs/quotes-cli.md#integration-with-conductor) for integration examples.

## Technical Details

- **Language**: Go 1.21+
- **Framework**: Cobra CLI
- **Testing**: 70%+ coverage with table-driven tests
- **Build**: Makefile with standard targets

## License

Part of the Conductor project.
