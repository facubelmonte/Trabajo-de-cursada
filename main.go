package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	
	_ "github.com/lib/pq"
	"trabajo-cursada/db/sqlc"

)

var queries *db.Queries

type Tarea struct {
	Titulo string `json:"titulo"`
	Materia string `json:"materia"`
	FechaEntrega string `json:"fecha_entrega"`
	Estado string `json:"estado"`
}


func main() {
    	dsn := "user=user password=password dbname=tp2_test sslmode=disable"
    	conn, err := sql.Open("postgres", dsn)
    	if err != nil {
    	    log.Fatalf("failed to connect to DB: %v", err)
    	}
    	defer conn.Close()
    	
    	queries = db.New(conn)
    	fmt.Println("Base de datos configurada y sqlc listo")
    	
    	http.HandleFunc("/tps", tpsHandler)
    	http.HandleFunc("/tps/", tpHandler)

    	fmt.Println("Servidor escuchando en localhost 8080")
    	log.Fatal(http.ListenAndServe(":8080", nil))
}

func tpsHandler(w http.ResponseWriter, r *http.Request) {
    	switch r.Method {
    	case http.MethodGet:
    	    getTPs(w, r)
    	case http.MethodPost:
    	    createTP(w, r)
    	default:
    	    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    	}
}

func tpHandler(w http.ResponseWriter, r *http.Request) {
    	parts := strings.Split(r.URL.Path, "/")
    	if len(parts) != 3 {
    	    http.Error(w, "Invalid ID", http.StatusBadRequest)
    	    return
    	}
    	id64, err := strconv.ParseInt(parts[2], 10, 64)
    	if err != nil {
    	    http.Error(w, "invalid ID", http.StatusBadRequest)
    	    return
    	}
	id := int32(id64)

    	switch r.Method {
    	case http.MethodGet:
    	    getTP(w, r, id)
    	case http.MethodPut:
    	    updateTP(w, r, id)
    	case http.MethodDelete:
        	deleteTP(w, r, id)
    	default:
        	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    	}
}

func getTPs(w http.ResponseWriter, r *http.Request) {
	lista, err := queries.ListarTP(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}

func createTP(w http.ResponseWriter, r *http.Request) {
	var nuevaTarea Tarea
	err := json.NewDecoder(r.Body).Decode(&nuevaTarea)
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusBadRequest)
        	return
    	}
    	if nuevaTarea.Titulo == "" || nuevaTarea.Materia == "" || nuevaTarea.FechaEntrega == "" {
        	http.Error(w, "titulo, materia y fecha_entrega son obligatorios", http.StatusBadRequest)
        	return
    	}
    	fecha, err := time.Parse("2006-01-02", nuevaTarea.FechaEntrega)
    	if err != nil {
        	http.Error(w, "fecha_entrega debe tener formato YYYY-MM-DD", http.StatusBadRequest)
        	return
    	}
    	estado := nuevaTarea.Estado
    	if estado == "" {
    	    	estado = "Pendiente"
    	}
    	tpGuardado, err := queries.CrearTP(r.Context(), db.CrearTPParams{
        	Titulo:       nuevaTarea.Titulo,
        	Materia:      nuevaTarea.Materia,
        	FechaEntrega: fecha,
        	Estado:       estado,
    	})
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        	return
    	}
        w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusCreated)
    	json.NewEncoder(w).Encode(tpGuardado)
}

func getTP(w http.ResponseWriter, r *http.Request, id int32) {
    	tp, err := queries.ObtenerTP(r.Context(), id)
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusNotFound)
        	return
    	}
    	w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(tp)
}

func updateTP(w http.ResponseWriter, r *http.Request, id int32) {
    	var tareaActualizada Tarea
    	err := json.NewDecoder(r.Body).Decode(&tareaActualizada)
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusBadRequest)
        	return
    	}

    	fecha, err := time.Parse("2006-01-02", tareaActualizada.FechaEntrega)
    	if err != nil {
        	http.Error(w, "fechaentrega invalida (formato YYYY-MM-DD)", http.StatusBadRequest)
        	return
    	}

    	_, err = queries.ObtenerTP(r.Context(), id)
    	if err != nil {
        	http.Error(w, "TP not found", http.StatusNotFound)
        	return
    	}

    	err = queries.ActualizarTP(r.Context(), db.ActualizarTPParams{
        	ID:           id,
        	Titulo:       tareaActualizada.Titulo,
        	Materia:      tareaActualizada.Materia,
        	FechaEntrega: fecha,
        	Estado:       tareaActualizada.Estado,
    	})
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        	return
    	}

        tp, _ := queries.ObtenerTP(r.Context(), id)
    	w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(tp)
}

func deleteTP(w http.ResponseWriter, r *http.Request, id int32) {
    	_, err := queries.ObtenerTP(r.Context(), id)
    	if err != nil {
        	http.Error(w, "TP not found", http.StatusNotFound)
        	return
    	}

    	err = queries.EliminarTP(r.Context(), id)
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        	return
    	}

    	w.WriteHeader(http.StatusNoContent)
}






