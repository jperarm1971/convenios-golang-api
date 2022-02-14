package routers

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	bbdd "github.com/jperarm1971/convenios-golang-api/BBDD"
	Model "github.com/jperarm1971/convenios-golang-api/Models"
	utils "github.com/jperarm1971/convenios-golang-api/Utils"
)

// Delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID := params["bookid"]

	var response = Model.JsonResponse{}

	if bookID == "" {
		response = Model.JsonResponse{Type: "error", Message: "You are missing bookID parameter."}
	} else {
		db := bbdd.SetUpDB()

		utils.PrintMessage("Deleting book from DB")

		_, err := db.Exec("DELETE FROM books where bookID = $1", bookID)
		utils.CheckErr(err)

		response = Model.JsonResponse{Type: "success", Message: "The book has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
