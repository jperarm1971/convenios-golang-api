package routers

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"

	bbdd "github.com/jperarm1971/convenios-golang-api/BBDD"
)

func GetConvenios(w http.ResponseWriter, r *http.Request) {
	serviciogestor := r.URL.Query().Get("servicio")
	organismo := r.URL.Query().Get("organismo")
	search := r.URL.Query().Get("search")
	result, status := bbdd.LeoConvenios(serviciogestor, organismo, search)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
