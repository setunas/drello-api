# Getting Started

Run all containers you need with this command.

```bash
docker compose up
```

# Database Migration

Install [golang-migrate/migrate](https://github.com/golang-migrate/migrate).

Then use the following commands.

```bash
// Apply one up migration file.
make db-migrate

// Apply all up migration files.
make db-migrate-all

// Apply one down migration file.
make db-rollback

// Apply all down migration files.
make db-rollback-all
```

# Testing

```bash
// Run all tests
go test ./...
```
