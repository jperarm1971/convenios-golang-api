package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	bbdd "github.com/jperarm1971/convenios-golang-api/BBDD"
)

func GetConvenio(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	id, err := strconv.Atoi(ID)

	if err == nil {
		result, status := bbdd.ObtieneConvenios(id)
		if status == false {
			http.Error(w, "Error al leer los usuarios en la llamada a bbd", http.StatusBadRequest)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(result)
		}
	} else {
		http.Error(w, "Error al leer el convenio, el id no es numérico", http.StatusBadRequest)
	}

}
