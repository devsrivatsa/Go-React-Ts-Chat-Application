postgres-init:
	docker run --name postgres-17-alpine -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=go-chat -p 5432:5432 -d postgres:17-alpine

postgres:
	docker exec -it postgres-17-alpine psql -U postgres -d go-chat
createdb:
	docker exec -it postgres-17-alpine createdb --username=postgres --owner=postgres go-chat
dropdb:
	docker exec -it postgres-17-alpine dropdb go-chat

create-migrations:
	migrate create -ext sql -dir db/migrations add_users_table

migrate-up:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/go-chat?sslmode=disable" --verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/go-chat?sslmode=disable" down

.PHONY: postgres-init postgres createdb dropdb
