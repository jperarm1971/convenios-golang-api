package routers

import (
	"encoding/json"

	"net/http"

	_ "github.com/lib/pq"

	bbdd "github.com/jperarm1971/convenios-golang-api/BBDD"
	Model "github.com/jperarm1971/convenios-golang-api/Models"
	utils "github.com/jperarm1971/convenios-golang-api/Utils"
)

// Delete all books
func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	db := bbdd.SetUpDB()

	utils.PrintMessage("Deleting all books...")

	_, err := db.Exec("DELETE FROM books")
	utils.CheckErr(err)

	utils.PrintMessage("All books have been deleted successfully!")

	var response = Model.JsonResponse{Type: "success", Message: "All books have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}
