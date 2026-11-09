-include .env
export

.PHONY: up down test clean run run-app

.env:
	@if [ ! -f .env ]; then cp .env.example .env; fi

up: .env
	@docker compose down -v
	@docker compose up -d
	@echo "Esperando a que PostgreSQL este listo..."
	@until docker compose exec -T database pg_isready -U $(POSTGRES_USER) -d $(POSTGRES_DB); do sleep 1; done
	@docker compose exec -T database psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) < db/schema/schema.sql
	@sqlc generate
	@go run github.com/a-h/templ/cmd/templ@latest generate
	@echo "Base de datos, sqlc y templ listos"

test: up
	@go test -v ./db/sqlc
	@$(MAKE) down

run-app: up
	@go build -o trabajo-cursada main.go
	@./trabajo-cursada; $(MAKE) down

run: run-app

down:
	@docker compose down -v

clean: down
	@rm -f db/sqlc/db.go db/sqlc/models.go db/sqlc/queries.sql.go
	@rm -f views/_templ.go
	@rm -f trabajo-cursada
