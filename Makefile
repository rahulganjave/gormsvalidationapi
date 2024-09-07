postgres:
	docker run --name postgres14.5 --network mo-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretkey -d postgres:14.5-alpine

createdb:
	docker exec -it postgres14.5 createdb --username=root --owner=root MiddleOffice

dropdb:
	docker exec -it postgres14.5 dropdb --username=root --owner=root MiddleOffice

migrateup:
	migrate -path db/migration -database "postgresql://root:secretkey@localhost:5432/MiddleOffice?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secretkey@localhost:5432/MiddleOffice?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
