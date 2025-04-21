postgresql:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16.8-alpine3.20
createdb:
	docker exec -it postgres createdb --username root --owner root simple_bank;
dropdb:
	docker exec -it postgres psql -U root -d postgres -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'simple_bank' AND pid <> pg_backend_pid();"
	docker exec -it postgres dropdb -U root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: postgresql createdb dropdb migrateup migratedown sqlc