postgres:
	docker run --name postgres-school-mgs --network school_mgs_network -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:16-alpine

start_container:
	docker start postgres-school-mgs

stop_container:
	docker stop postgres-school-mgs

createdb:
	docker exec -it postgres-school-mgs createdb --username=postgres --owner=postgres school_mgs

dropdb:
	docker exec -it postgres-school-mgs dropdb -U postgres school_mgs

migrate_up:
	igrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable" -verbose down

migrate_1up:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable" -verbose up 1

migrate_1down:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable" -verbose down 1

migrate_fix:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable" force $(version)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/osobotu/school_mgs/db/sqlc Store


.PHONY: postgres createdb dropdb migrate_down migrate_up start_container stop_container sqlc test server mockdb migrate_1up migrate_1down