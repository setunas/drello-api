# Getting Started

Set the following environment variables to connect to your local database.

```
DB_USER         // e.g. 'root'
DB_PASS         // e.g. 'password'
DB_TCP_HOST     // e.g. '127.0.0.1'
DB_PORT         // e.g. '4306'
DB_NAME         // e.g. 'drello-dev'
```

Run the container of MariaDB.

```bash
docker compose up
```

Run the server.

```bash
go run .
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
