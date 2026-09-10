package db

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ejecuta las pruebas para cada una de las operaciones generadas
func TestQueries_CRUD(t *testing.T) {
	// nos conectamos a la base de datos
	connStr := "user=user password=password dbname=tp2_test host=localhost port=5432 sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Fallo al conectar a la BD: %v", err)
	}
	defer conn.Close()

	schema, err := os.ReadFile("../schema/schema.sql")
	if err != nil {
		t.Fatalf("Error al leer schema.sql: %v", err)
	}
	if _, err := conn.Exec(string(schema)); err != nil {
		t.Fatalf("Error al crear la tabla en la BD: %v", err)
	}

	queries := New(conn)
	
	ctx := context.Background() 
	
	var createdTP Tp

	// 1. CrearTP
	t.Run("CrearTP", func(t *testing.T) {
		params := CrearTPParams{
			Titulo:       "TP2 - Persistencia con sqlc",
			Materia:      "Programación Web",
			FechaEntrega: time.Date(2026, time.September, 14, 0, 0, 0, 0, time.UTC),
			Estado:       "Pendiente",
		}
		
		tp, err := queries.CrearTP(ctx, params)
		if err != nil {
			t.Fatalf("Error al crear TP: %v", err)
		}
		if tp.ID == 0 {
			t.Errorf("Se esperaba un ID válido generado por la BD")
		}
		createdTP = tp
	})

	// 2. ObtenerTP
	t.Run("ObtenerTP", func(t *testing.T) {
		tp, err := queries.ObtenerTP(ctx, createdTP.ID)
		if err != nil {
			t.Fatalf("Error al obtener TP: %v", err)
		}
		if tp.Titulo != "TP2 - Persistencia con sqlc" {
			t.Errorf("Los datos obtenidos no coinciden")
		}
	})

	// 3. ActualizarTP
	t.Run("ActualizarTP", func(t *testing.T) {
		params := ActualizarTPParams{
			ID:           createdTP.ID,
			Titulo:       "TP2 - Persistencia Finalizado",
			Materia:      "Programación Web",
			FechaEntrega: createdTP.FechaEntrega,
			Estado:       "Entregado",
		}
		
		err := queries.ActualizarTP(ctx, params)
		if err != nil {
			t.Fatalf("Error al actualizar TP: %v", err)
		}

		tpActualizado, err := queries.ObtenerTP(ctx, createdTP.ID)
		if err != nil {
			t.Fatalf("Error al verificar actualización: %v", err)
		}
		if tpActualizado.Estado != "Entregado" {
			t.Errorf("El estado no se actualizó correctamente")
		}
	})

	// 4. ListarTP
	t.Run("ListarTP", func(t *testing.T) {
		tps, err := queries.ListarTP(ctx)
		if err != nil {
			t.Fatalf("Error al listar TPs: %v", err)
		}
		if len(tps) == 0 {
			t.Errorf("La lista de TPs está vacía")
		}
	})

	// 5. EliminarTP
	t.Run("EliminarTP", func(t *testing.T) {
		err := queries.EliminarTP(ctx, createdTP.ID)
		if err != nil {
			t.Fatalf("Error al eliminar TP: %v", err)
		}

		_, err = queries.ObtenerTP(ctx, createdTP.ID)
		if err == nil {
			t.Errorf("Se esperaba error al buscar un TP borrado")
		}
		if err != sql.ErrNoRows {
			t.Errorf("Se esperaba sql.ErrNoRows, se obtuvo: %v", err)
		}
	})
}
