## help: print this help message
.PHONY: help
help:
		@echo 'Usage:'
		@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
		@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
		go run ./cmd/api -db-dsn=${NERDATE_DB_DSN}

## db/up: start the database docker container
.PHONY: db/up
db/up:
		docker-compose up -d

## db/logs: tail the database logs
.PHONY: db/logs
db/logs:
		docker-compose logs -f -t

## db/down: stop the database docker container
.PHONY: db/down
db/down:
		docker-compose down

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
		psql ${NERDATE_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
		@echo 'Creating migration files for ${name}...'
		migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all database up migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
		@echo 'Running up migrations...'
		migrate -path ./migrations -database ${NERDATE_DB_DSN} up
