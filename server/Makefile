postgres-init:
	docker run --name postgres-init -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -p 5432:5432 -d postgres:16.1-alpine

postgres:
	docker exec -it postgres-17-alpine psql
createdb:
	docker exec -it postgres-17-alpine createdb --username=postgres --owner=postgres go-chat
dropdb:
	docker exec -it postgres-17-alpine dropdb go-chat

.PHONY: postgres-init postgres createdb dropdb
