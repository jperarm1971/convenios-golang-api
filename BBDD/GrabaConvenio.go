package bbdd

import (
	_ "github.com/lib/pq"

	Model "github.com/jperarm1971/convenios-golang-api/Models"
	utils "github.com/jperarm1971/convenios-golang-api/Utils"
)

// Create a convenio

func GrabaConvenio(u Model.Convenio) (string, bool, error) {

	if u.ConvenioName == "" || u.ConvenioOrganismo == "" {
		return "Error con el nombre / organismo del convenio", false, nil
	} else {
		db := SetUpDB()

		utils.PrintMessage("Inserting convenio into DB")

		var lastInsertID int
		//query := fmt.Sprintf("INSERT INTO convenios(name, fechafirma,serviciogestor,organismo) VALUES($1, $2,$3,$4) returning id;", u.ConvenioName, u.ConvenioFechaFirma, u.ConvenioGestor, u.ConvenioOrganismo)
		//fmt.Println("la consulta sería: " + query)

		err := db.QueryRow("INSERT INTO convenios(name, fechafirma,serviciogestor,organismo) VALUES($1, $2,$3,$4) returning id;", u.ConvenioName, u.ConvenioFechaFirma, u.ConvenioGestor, u.ConvenioOrganismo).Scan(&lastInsertID)
		if err != nil {
			return "Error insertando registro en bbdd", false, err
		}
	}
	return "", true, nil
}
