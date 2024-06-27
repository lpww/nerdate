confirm:
		@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

run/api:
		go run ./cmd/api

db/up:
		docker-compose up -d

db/logs:
		docker-compose logs -f -t

db/down:
		docker-compose down

db/psql:
		psql ${NERDATE_DB_DSN}

db/migrations/new:
		@echo 'Creating migration files for ${name}...'
		migrate create -seq -ext=.sql -dir=./migrations ${name}

db/migrations/up: confirm
		@echo 'Running up migrations...'
		migrate -path ./migrations -database ${NERDATE_DB_DSN} up
