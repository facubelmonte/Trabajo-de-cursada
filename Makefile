.PHONY: env sqlc db-down db-up test test-api clean

#Crear el .env desde .env.example si no existe
env:
	@cp -n .env.example .env || true

#Generar codigo de sqlc
sqlc:
	@sqlc generate

#Detiene y borra los contenedores y volumenes de docker
db-down: env
	@docker compose down -v
	
#Levanta el contenedor de postgresql y espera
db-up: env
	@docker compose up -d
	@echo "esperando a que postrgresql este listo"
	@until docker compose exec -T database pg_isready -U user -d tp2_test; do sleep 1; done
	
#Ejecuta los tests unitarios de persistencia y limpia al terminar
test: sqlc db-down db-up
	@go test -v ./...; status=$$?; $(MAKE) db-down; exit $$status
	
#Ejecuta el flujo completo de pruebas de la API REST
test-api: sqlc db-down db-up
	@go build
	@docker compose exec -T database psql -U user -d tp2_test < db/schema/schema.sql
	@./trabajo-cursada &
	@sleep 1
	@./request.bash
	@pkill trabajo-cursada
	@rm -f trabajo-cursada
	@$(MAKE) db-down

#Limpia contenedores y archivos
clean: db-down
	@rm -rf ./db/sqlc trabajo-cursada
