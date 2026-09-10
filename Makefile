.PHONY: sqlc db-up db-down test clean

#Generar codigo de acceso a dato
sqlc:
	@sqlc generate

#Limpiar contenedores y volumenes previos
db-down:
	@echo "Deteniendo Postgres y eliminando volumenes"
	@docker compose down -v

#Levantar contenedor de base de datos
db-up:
	@echo "Levantando Postgres en Docker"
	@docker compose up -d
	@echo "Esperando que Postgres este listo"
	@sleep 4

#Ejecucion de tests
test: sqlc db-down db-up
	@echo "Ejecutando pruebas"
	@go test -v ./...; status=$$?; $(MAKE) db-down; exit $$status

#Finalizacion y limpieza general
clean: db-down
	@echo "Limpiando codigo generado"
	@rm -rf ./db/sqlc


