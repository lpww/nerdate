# nerdate API

An API for interacting with nerdate. Written in Go.

## API Routes

| Implemented        | Method | URL Pattern        | Handler              | Action                |
| ------------------ | ------ | ------------------ | -------------------- | --------------------- |
| :white_check_mark: | GET    | /v1/healthcheck    | healthcheckHandler   | Show application info |
| :white_check_mark: | POST   | /v1/users          | registerUserHandler  | Register a new user   |
| :white_check_mark: | GET    | /v1/users/discover | discoverUsersHandler | Show available users  |
| :white_check_mark: | POST   | /v1/swipes         | createSwipeHandler   | Create a new swipe    |
| :x:                | GET    | /v1/users/matches  | showMatchesHandler   | Show matches          |

## Prerequisites

1. [go](https://go.dev)
2. [make](https://www.gnu.org/software/make)
3. [golang-migrate](https://github.com/golang-migrate/migrate)
4. [docker](https://www.docker.com)
5. [docker-compose](https://docs.docker.com/compose/)

## Env Vars

The `.example.envrc` file lists the variables that that application requires. Move them to `.envrc` and populate the values, they will be automatically loaded by `make` before running commands.

```
cp .example.envrc .envrc
```

## Quick Start

1. `make db/up`: run the db
2. `make db/migrations/up`: run up migrations
3. `make audit`: run code quality checks
4. `make run/api`: run the server

## Important Commands

* `make help`: list the available make commands
* `make db/up`: run the db in a docker container
* `make db/down`: stop the db 
* `make db/migrations/up`: run up migrations
* `make run/api`: run the server

## Folder structure

* `bin`: Contains compiled application binaries, ready for deployment.
* `cmd/api`: Contains application specific code for the nerdate API. Any code for running the server, handling http requests, and handling authentication.
* `internal`: Contains various ancillary packages used by the API. Any code which isn't application specific and can potentially be reused.
* `migrations`: Contains SQL migration files for the DB.
* `remote`: Contains config files and setup scripts for the production server.
* `go.mod`: Will declare project dependencies, versions, and module path.
* `Makefile`: Contains recipes for automating common admin tasks.
* `database`: Contains an init file for the database. Used for installing extensions.

## Router

This project uses httprouter. It is a good option for public APIs because it allows custom JSON responses when automatically hanlding 404 and 405 responses. It also automatically handles OPTIONS requests.

## Example requests

### GET /v1/healthcheck

```
curl -i localhost:4000/v1/healthcheck
```

### POST /v1/users

```
export BODY='{"name": "lpww", "gender": "male", "dob": "2006-01-02T15:04:05Z", "ascii_art": "image here", "description": "test description", "email": "lpww@github.com", "password": "abcd1234"}'
curl -i -d $BODY localhost:4000/v1/users
```

### GET /v1/users/discover

```
curl -i localhost:4000/v1/users/discover
```

### POST /v1/swipes

```
export BODY='{"swiped_user_id":"eab9be97-fb18-449b-b568-5c37b89bff18","liked": true}'
curl -i -d $BODY localhost:4000/v1/swipes
```
