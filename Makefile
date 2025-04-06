include .env

APIDOC_BASE = cmd/api
APIDOC_INFO = internal/http/handler
PGSQL_URI = postgres://$(PGSQL_USERNAME):$(PGSQL_PASSWORD)@$(PGSQL_HOST):$(PGSQL_PORT)/$(PGSQL_DATABASE_NAME)?sslmode=disable

## ----- MIGRATION -----
auto_migrate:
	go run tools/auto_migrate.go

migrate:
ifndef create
	$(error Please provide name: make migrate create=create_users_table)
endif
	migrate create -ext sql -dir database/migration/ -seq $(create)

migrate_up:
	migrate -path database/migration -database '$(PGSQL_URI)' -verbose up

migrate_down:
	migrate -path database/migration -database '$(PGSQL_URI)' -verbose down

migrate_rollback:
ifndef step
	$(error Please provide step: make migrate_rollback step=1)
endif
	migrate -path database/migration -database '$(PGSQL_URI)' -verbose down $(step)

migrate_fix:
ifndef version
	$(error Please provide version: make migrate_fix version=3)
endif
	migrate -path database/migration -database '$(PGSQL_URI)' force $(version)

## ----- DEVELOPMENT -----
dev:
	air

run:
	go run cmd/api/main.go

build:
	go build -o bin/app cmd/api/main.go

## ----- TESTING -----
test:
	go test -v -cover ./...

coverage:
	go test ./... -coverprofile=cover.out
	go tool cover -func=cover.out

mock:
ifndef d
	$(error Please provide interface name: make mock d=UserService)
endif
	mockery --name=$(d) --recursive=true --output=tests/mocks

## ----- API DOC -----
apidoc:
	swag init -d $(APIDOC_BASE) --parseInternal --output docs

## ----- PROTOBUF -----
protob:
	protoc --go_out=proto/pb --go_opt=paths=source_relative \
	       --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative proto/*.proto