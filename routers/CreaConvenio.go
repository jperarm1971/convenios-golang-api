package routers

import (
	"encoding/json"
	"net/http"

	bbdd "github.com/jperarm1971/convenios-golang-api/BBDD"
	Model "github.com/jperarm1971/convenios-golang-api/Models"
)

/* funcion para grabar en bd el usuario*/

func CreaConvenio(w http.ResponseWriter, r *http.Request) {

	var t Model.Convenio
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.ConvenioName) == 0 {
		http.Error(w, "El nombre del convenio es requerido", 400)
		return
	}
	if len(t.ConvenioOrganismo) == 0 {
		http.Error(w, "El convenio debe tener un organismo", 400)
		return
	}

	_, status, err := bbdd.GrabaConvenio(t)
	if err != nil {
		http.Error(w, "Se produjo en error en el guardado del registro en la BD:"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se pudo insertar el convenio en la BD:"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
