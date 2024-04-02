# export $(grep -v '^#' .env | xargs)
# Database connection details
DATABASE_URL=scylla://127.0.0.1:9042/mykeyspace

# Path to your migration directory
MIGRATION_PATH=./migrations

start:
	go run cmd/server/main.go

build:
	go build cmd/server/main.go

test:
	go test	-run=Test -v ./...

# Run migrations
migrate-up:
	migrate -database $(DATABASE_URL) -path $(MIGRATION_PATH) up

# Rollback migrations
migrate-down:
	migrate -database $(DATABASE_URL) -path $(MIGRATION_PATH) down

# Ensure the database is up and running
check-db:
	docker exec -it scylla-node1 cqlsh 172.21.0.2 -e "DESCRIBE KEYSPACES"

# Target to apply migrations and then check the database
run-migrations: migrate-up check-db

# Target to rollback migrations and then check the database
rollback-migrations: migrate-down check-db
