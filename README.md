# nerdate API

An API for interacting with nerdate. Written in Go.

## API Routes

| Method | URL Pattern        | Handler              | Action                   |
| ------ | ------------------ | -------------------- | ------------------------ |
| GET    | /v1/healthcheck    | healthcheckHandler   | Show application info    |
| POST   | /v1/users          | createUserHandler    | Register a new user      |
| GET    | /v1/users/discover | discoverUsersHandler | Show the available users |
| POST   | /v1/swipe          | createSwipeHandler   | Register a new swipe     |

## Folder structure

* `bin`: Contains compiled application binaries, ready for deployment.
* `cmd/api`: Contains application specific code for the nerdate API. Any code for running the server, handling http requests, and handling authentication.
* `internal`: Contains various ancillary packages used by the API. Any code which isn't application specific and can potentially be reused.
* `migrations`: Contains SQL migration files for the DB.
* `remote`: Contains config files and setup scripts for the production server.
* `go.mod`: Will declare project dependencies, versions, and module path.
* `Makefile`: Contains recipes for automating common admin tasks.
