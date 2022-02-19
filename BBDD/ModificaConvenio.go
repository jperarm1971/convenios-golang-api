package bbdd

import (
	Model "github.com/jperarm1971/convenios-golang-api/Models"
	utils "github.com/jperarm1971/convenios-golang-api/Utils"
	_ "github.com/lib/pq"
)

// Create a convenio

func ModificaConvenio(u Model.Convenio) (string, bool, error) {

	db := SetUpDB()

	utils.PrintMessage("Modificando convenio en la DB")
	_, err := db.Exec("UPDATE convenios SET name=$1, fechafirma=$2, serviciogestor=$3, organismo=$4 WHERE ID=$5;", u.ConvenioName, u.ConvenioFechaFirma, u.ConvenioGestor, u.ConvenioOrganismo, u.ConvenioID)
	if err != nil {
		return "Error actualizando registro en bbdd", false, err
	}
	return "", true, nil
}
