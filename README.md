![drello-brand-icon-192x192](https://user-images.githubusercontent.com/12164726/150669986-66933d5f-a6bc-4420-b9d7-dd48f8ad17f2.png)

# Drello Overview

## What is Drello?
<img width="1280" alt="drello-homepage" src="https://user-images.githubusercontent.com/12164726/152212089-e151b44f-265b-4b57-b049-b4759cd8d43d.gif">

Drello is a symple kanban web service for your todos. We created Drello as our portfolio. <br>
Take a look 😁 (Please read "Notes" below before checking it out 🙏) <br>
https://drello.netlify.app/

#### ⚠️ Notes
- Our server is probably [sleeping](https://devcenter.heroku.com/articles/free-dyno-hours#dyno-sleeping) when you access it. So it would have a delay (usually around 30 seconds) when you log in at first.
  - This occurs because we use a free plan server of [Dyno](https://devcenter.heroku.com/categories/dynos). See more details [here](https://devcenter.heroku.com/articles/free-dyno-hours#dyno-sleeping).
- You might have to **turn off your ad blocker** to log in since we use a popup window to log in with Firebase Authentication.

## Repositories for Drello
We have two GitHub repositories. You can see all our codes here.
- [`drello-api`](https://github.com/setunas/drello-api): The backend app exposes RESTful APIs connecting MySQL using Go.
- [`drello-web`](https://github.com/setunas/drello-web): The frontend app for web mainly using React, Next.js and TypeScript.


## System Architecture
![drello-architecture](https://user-images.githubusercontent.com/12164726/150669120-874976a2-d8e2-43c3-aae4-0236bf900187.png)


## Authentication & Authorization
We authenticate users by Firebase Authentication. Firebase Authentication provides an ID token. We use this ID token also to authorize a user to use proper API requests by passing the ID token together. Our API server (`drello-api`) verifies the provided ID token with Firebase Admin SDK.


## Logics
### How to Manage the Order of Cards and Columns
![dnd](https://user-images.githubusercontent.com/12164726/151677475-bc65ef3c-6a5e-4d7d-af6a-3f879012331c.gif)<br>
To manage the order of cards and columns, we set a number for each card and column. We named that number `position`. The bigger number of `position` is the more backward in the order. <br>
<br>
Let us break it down. Let's say the default number of `position` is 1024. When the first card in a column is created, it is given a `position` of 1024. Let's see the other examples.
- When a card is added to the end of a column, it is given <`position` of last card> + 1024. 
- When a card is added to the beginning, it is given <`position` of first card> / 2. 
- When a card is added between two cards, it is the average of the two neighbors.
- When the `position` of a card can't be calculated anymore (such as two cards' values get too close or the value gets too big), we re-number the positions of all cards in the column.

We adopted this way to make it easier to keep the consistency of positions among other cards and avoid the calculation getting expensive.<br>
<br>
You can see the position-related code mainly [here](https://github.com/setunas/drello-web/blob/develop/src/features/position/position.ts).






# About `drello-api`
In this section, we are going to explain `drello-api`. If you want to see about `drello-web`, see [here](https://github.com/setunas/drello-web/blob/develop/README.md#about-drello-web).

## Tech Stacks
- [Go](https://go.dev/)
- [MariaDB](https://mariadb.org/) / [MySQL](https://www.mysql.com/)
- [Github Actions](https://github.com/features/actions)
- [Docker](https://www.docker.com/)

## Software Architecture 
### Onion Architecture
![onion-architecture](https://user-images.githubusercontent.com/12164726/150693911-15137f0e-a54e-4d93-88e4-48d4b8a5d323.png)<br>
*Image by [Daniel Rusnok](https://dev.to/danielrusnok)*

We adopted Onion Architecture for our software architecture.<br>
<br>
Onion Architecture is an architectural pattern that keeps maintainability with good separation of concern by splitting application codes into the layers in the image above. To get to know more details about Onion Architecture, see [here](https://marcoatschaefer.medium.com/onion-architecture-explained-building-maintainable-software-54996ff8e464).<br>
<br>
To be honest, Onion Architecture is not fit for our application since our codebase size is small and we don't have business rules so much. But our app is made as a portfolio, so we adopted Onion Architecture so we can show how we code usually.


## CI/CD
We use [Github Actions](https://github.com/features/actions) for CI/CD. It automatically tests, builds, or deploys when you push to a pull request or merge to a certain branch.

# Contribution
This section is for developers who joined us to work on developing `drello-api` together ✌️

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
