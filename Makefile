-include .env
export

.PHONY: up down test clean run test-api

.env:
	@if [ ! -f .env ]; then cp .env.example .env; fi

up: .env
	@docker compose up -d
	@echo "Esperando a que PostgreSQL este listo..."
	@until docker compose exec -T database pg_isready -U $(POSTGRES_USER) -d $(POSTGRES_DB); do sleep 1; done
	@docker compose exec -T database psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) < db/schema/schema.sql
	@sqlc generate
	@echo "Base de datos configurada y sqlc listo"

test: up
	@go test -v ./db/sqlc
	@$(MAKE) down

run: up
	@go run main.go

test-api: up
	@go build -o trabajo-cursada main.go
	@./trabajo-cursada &
	@until curl -s http://localhost:8080/tps > /dev/null; do sleep 1; done
	@bash request.bash
	@pkill -x trabajo-cursada || true
	@$(MAKE) down

down:
	@docker compose down -v

clean: down
	@rm -f db/sqlc/db.go db/sqlc/models.go db/sqlc/queries.sql.go
	@rm -f trabajo-cursada
