package bbdd

import (
	"fmt"

	Model "github.com/jperarm1971/convenios-golang-api/Models"
	utils "github.com/jperarm1971/convenios-golang-api/Utils"
	_ "github.com/lib/pq"
)

func LeoConvenios(serviciogestor string, organismo string, search string) ([]*Model.Convenio, bool) {

	db := SetUpDB()
	var convenios []*Model.Convenio
	utils.PrintMessage("obteniendo Convenios...")
	str_where := ""
	query := "SELECT * FROM convenios"
	if serviciogestor != "" {
		str_where = str_where + "serviciogestor like '%" + serviciogestor + "%'"
	}
	if organismo != "" {
		if str_where != "" {
			str_where = str_where + " AND "
		}
		str_where = str_where + "organismo like '%" + organismo + "%'"
	}
	if search != "" {
		if str_where != "" {
			str_where = str_where + " AND "
		}
		str_where = str_where + "name like '%" + search + "%'"
	}
	if str_where != "" {
		query = query + " WHERE " + str_where
	}
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("error:" + err.Error())
		return convenios, false
	}

	for rows.Next() {
		var s Model.Convenio
		err = rows.Scan(&s.ConvenioID, &s.ConvenioName, &s.ConvenioFechaFirma, &s.ConvenioGestor, &s.ConvenioOrganismo)
		if err != nil {
			fmt.Println("error2:" + err.Error())
			return convenios, false
		}
		convenios = append(convenios, &s)
	}
	return convenios, true

}
