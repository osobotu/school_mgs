postgres:
	docker run --name postgres-school-mgs -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:16-alpine

createdb:
	docker exec -it postgres-school-mgs createdb --username=postgres --owner=postgres school_mgs

dropdb:
	docker exec -it postgres-school-mgs dropdb -U postgress school_mgs

migrate_up:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable" -verbose down


.PHONY: postgres createdb dropdb migrate_down migrate_up