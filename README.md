# Start Server

Run this command to start the server and the database.

```bash
docker compose up
```

When you want to enter the running container (e.g. to run database migration or tests), you can use this command.

```bash
docker compose exec go bash 
```


# Database Migration

Use the following `make` commands inside the container.

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

These `make` commands use [golang-migrate/migrate](https://github.com/golang-migrate/migrate).

# Testing

Run this command inside the container.

```bash
// Run all tests
go test ./...
```

# Deploy to Staging
You can deploy to the staging Heroku environment by pushing any branches whose name starts with `stg-`. <br>
e.g. `stg-new-feature-1`
