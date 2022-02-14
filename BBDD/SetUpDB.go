package bbdd

import (
	"database/sql"
	"fmt"

	utils "github.com/jperarm1971/convenios-golang-api/Utils"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "admin"
	DB_NAME     = "josue"
)

func SetUpDB() *sql.DB {

	host := "192.168.1.39"
	port := "5432"
	user := "postgres"
	password := "admin"
	dbname := "josue"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", psqlInfo)

	utils.CheckErr(err)

	return db
}
