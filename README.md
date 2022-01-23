![drello-brand-icon-192x192](https://user-images.githubusercontent.com/12164726/150669986-66933d5f-a6bc-4420-b9d7-dd48f8ad17f2.png)



# What is Drello?
Drello is a symple kanban web service for your todos. We created Drello as our portfolio. <br>
Take a look 😁 https://drello.netlify.app/

# Repositories for Drello
- [`drello-api`](https://github.com/setunas/drello-api): The backend app exposes RESTful APIs connecting MySQL using Go.
- [`drello-web`](https://github.com/setunas/drello-web): The frontend app for web mainly using React, Next.js and TypeScript.

# System Architecture
![drello-architecture](https://user-images.githubusercontent.com/12164726/150669120-874976a2-d8e2-43c3-aae4-0236bf900187.png)

# Authentication & Authorization
We authenticate users by Firebase Authentication. Firebase Authentication provide an ID token. We use this ID token to authorize a user to use proper API request by passing the ID token together. Our API server (`drello-api`) verifies the provided ID token.


# File Structure

# Contribution

## Get Started

### Set Environment Variables

We use `.envrc` file to set environment variables with [`direnv`](https://direnv.net/).<br>
Install `direnv` and ask another developer to share `.envrc` file, to make it easy to set environment variables for this app.


### Place GOOGLE_APPLICATION_CREDENTIALS file

Ask another developer to share GOOGLE_APPLICATION_CREDENTIALS file. <br>
You need to place this file in the right place specified by `GOOGLE_APPLICATION_CREDENTIALS` environment variable.<br><br>
Without this file being placed properly, some functions with google cloud platform wouldn't work.


### Create Database and Tables
See [this section](#database-migration) to get to know how to run commands for database migration.


### Start the Server

Run this command at the root directory of this project to start the server and the database with docker containers.

```
docker compose up
```

## Enter a Running Container

When you want to enter a running container of the server (e.g. to run database migration or tests), you can use this command.

```
docker compose exec app bash 
```


## Database Migration

Run the following `make` commands in the running `app` container.<br>
You can see how to enter the app container [here](#enter-a-container).

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

## Testing

Run this command inside the container.

```
// Run all tests
go test ./...
```

## Release to Staging

#### 1. Deploy to Heroku

You can deploy to the staging Heroku environment by just pushing any branches whose name starts with `stg-`. <br>
e.g. `stg-new-feature-1`

#### 2. Database Migration (if needed)

You can ask a developer who has admin permission to migrate database.
