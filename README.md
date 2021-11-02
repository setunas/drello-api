# Start the Server

Run this command at the root directory of this project to start the server and the database with docker containers.

```
docker compose up
```

# Enter a Container

When you want to enter a running container of the server (e.g. to run database migration or tests), you can use this command.

```
docker compose exec app bash 
```


# Database Migration

Use the following `make` commands inside the container.

```
// Apply one up migration file.
make db-migrate

// Apply all up migration files.
make db-migrate-all

// Apply one down migration file.
make db-rollback

// Apply all down migration files.
make db-rollback-all

// Apply all up migration files for the DB in test environment.
make db-migrate-all-test
```

These `make` commands use [golang-migrate/migrate](https://github.com/golang-migrate/migrate).

# Testing

Run this command inside the container.

```
// Run all tests
go test ./...
```

# Deploy to Staging
You can deploy to the staging Heroku environment by pushing any branches whose name starts with `stg-`. <br>
e.g. `stg-new-feature-1`
