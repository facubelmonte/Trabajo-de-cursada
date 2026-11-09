package handlers

import (
	"net/http"
	"time"

	"trabajo-cursada/db/sqlc"
	"trabajo-cursada/views"
)

var Queries *db.Queries

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lista, err := Queries.ListarTP(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	views.IndexPage("Gestor de Trabajos Practicos", lista).Render(r.Context(), w)
}

func CreateTPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	titulo := r.FormValue("titulo")
	materia := r.FormValue("materia")
	fechaStr := r.FormValue("fecha_entrega")
	estado := r.FormValue("estado")

	if titulo == "" || materia == "" || fechaStr == "" {
		http.Error(w, "Campos obligatorios faltantes", http.StatusBadRequest)
		return
	}

	fecha, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		http.Error(w, "Fecha invalida", http.StatusBadRequest)
		return
	}

	if estado == "" {
		estado = "Pendiente"
	}

	_, err = Queries.CrearTP(r.Context(), db.CrearTPParams{
		Titulo:       titulo,
		Materia:      materia,
		FechaEntrega: fecha,
		Estado:       estado,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
