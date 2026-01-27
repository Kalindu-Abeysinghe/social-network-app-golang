include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=./cmd/migrate/migrations -database=$(DB_MIGRATION_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=./cmd/migrate/migrations -database=$(DB_MIGRATION_ADDR) down $(filter-out $@,$(MAKECMDGOALS))