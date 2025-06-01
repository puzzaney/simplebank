postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
	
sqlc:
	sqlc generate

test:
	go test -v -cover -coverprofile coverage.out ./...

testapi:
	go test -v -cover -coverprofile coverage.out ./api

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/puzzaney/simplebank/db/sqlc Store

.PHONY:postgres createdb dropdb migrateup migratedown sqlc test server mock testapi migrateup1 migratedown1
