# Getting Started

## Preparation

Set the following environment variables to connect your local database.

```
DB_USER
DB_PASS
DB_TCP_HOST
DB_PORT
DB_NAME
```

Run the container of MariaDB

```bash
// Move to the project root
cd drello-api

docker compose up
```

## Run Server

```bash
// Move to the project root
cd drello-api

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
