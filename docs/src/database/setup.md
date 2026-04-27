# Database Setup

bbscope uses SQLite to store program scopes and track changes over time.

## Requirements

- None — SQLite is embedded.

## Database path

The database file is automatically created at `~/.bbscope/bbscope.db` by default.

You can override it in `~/.bbscope/config.yaml`:

```yaml
db_path: "/custom/path/to/bbscope.db"
```

Or via environment variable (used by the web server):

```bash
export DB_PATH="/custom/path/to/bbscope.db"
```

## Schema auto-migration

bbscope automatically creates all tables and indexes on first connection. There's no manual migration step. The schema is idempotent — safe to run against an existing database.

## Docker

When running via Docker Compose, the database is persisted in a named volume mounted at `/data/bbscope.db`.
