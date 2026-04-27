# Quick Start

## 1. Set up credentials

Create `~/.bbscope.yaml` with your platform credentials:

```yaml
hackerone:
  username: "your_username"
  token: "your_api_token"
```

See [Configuration](./configuration.md) for all platforms.

## 2. Print scopes (no database)

```bash
# All configured platforms
bbscope poll

# Just HackerOne, only bug bounty programs
bbscope poll h1 -b

# Custom output: target + description + program URL
bbscope poll -o tdu
```

## 3. Track changes with a database

The database is automatically created at `~/.bbscope/bbscope.db`. You can override the path in your config:

```yaml
db_path: "/custom/path/to/bbscope.db"
```

Then poll with `--db`:

```bash
# First run populates the database silently
bbscope poll --db

# Subsequent runs print only changes (new/removed/updated targets)
bbscope poll --db
```

## 4. Query the database

```bash
# View stats
bbscope db stats

# Search for a target
bbscope db find "example.com"

# Extract wildcards for recon
bbscope db get wildcards

# Recent changes
bbscope db changes
```

## 5. Enable AI normalization (optional)

Add an OpenAI API key to your config:

```yaml
ai:
  api_key: "sk-..."
```

```bash
bbscope poll --db --ai
```

This cleans up messy scope entries (e.g., `"*.example.com (main site)"` becomes `*.example.com`) and caches results in the database to avoid redundant API calls.
