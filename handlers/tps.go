package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"trabajo-cursada/db/sqlc"
)

var Queries *db.Queries

type Tarea struct {
	Titulo       string `json:"titulo"`
	Materia      string `json:"materia"`
	FechaEntrega string `json:"fecha_entrega"`
	Estado       string `json:"estado"`
}

func TpsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTPs(w, r)
	case http.MethodPost:
		createTP(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func TpHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	id64, err := strconv.ParseInt(parts[2], 10, 32)
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
	lista, err := Queries.ListarTP(r.Context())
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
	tpGuardado, err := Queries.CrearTP(r.Context(), db.CrearTPParams{
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
	tp, err := Queries.ObtenerTP(r.Context(), id)
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

	_, err = Queries.ObtenerTP(r.Context(), id)
	if err != nil {
		http.Error(w, "TP not found", http.StatusNotFound)
		return
	}

	err = Queries.ActualizarTP(r.Context(), db.ActualizarTPParams{
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

	tp, _ := Queries.ObtenerTP(r.Context(), id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tp)
}

func deleteTP(w http.ResponseWriter, r *http.Request, id int32) {
	_, err := Queries.ObtenerTP(r.Context(), id)
	if err != nil {
		http.Error(w, "TP not found", http.StatusNotFound)
		return
	}

	err = Queries.EliminarTP(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

