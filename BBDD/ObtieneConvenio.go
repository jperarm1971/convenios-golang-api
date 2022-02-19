package bbdd

import (
	"fmt"

	Model "github.com/jperarm1971/convenios-golang-api/Models"
	utils "github.com/jperarm1971/convenios-golang-api/Utils"
	_ "github.com/lib/pq"
)

func ObtieneConvenios(idconvenio int) (*Model.Convenio, bool) {

	db := SetUpDB()
	var s Model.Convenio
	utils.PrintMessage("obteniendo Convenios...")

	query := "SELECT * FROM convenios WHERE id=" + fmt.Sprint(idconvenio)
	err := db.QueryRow(query).Scan(&s.ConvenioID, &s.ConvenioName, &s.ConvenioFechaFirma, &s.ConvenioGestor, &s.ConvenioOrganismo)
	if err != nil {
		fmt.Println("error1:" + err.Error())
	}

	return &s, true

}
