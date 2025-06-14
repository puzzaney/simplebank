DB_URL = postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name postgres17 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1
	
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

dbdocs:
	dbdocs build doc/db.dbml

dbschema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	rm -rf pb/*go
	rm -rf doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
		proto/*.proto
	statik -src ./doc/swagger -dest ./doc

evans:
	evans --host localhost --port 9090 -r repl

.PHONY:postgres createdb dropdb migrateup migratedown sqlc test server mock testapi migrateup1 migratedown1 dbdocs dbschema proto evans
