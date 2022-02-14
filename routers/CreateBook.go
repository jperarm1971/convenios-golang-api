package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	bbdd "github.com/jperarm1971/convenios-golang-api/BBDD"
	Model "github.com/jperarm1971/convenios-golang-api/Models"
	utils "github.com/jperarm1971/convenios-golang-api/Utils"
)

// Create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookid")
	bookName := r.FormValue("bookname")

	var response = Model.JsonResponse{}

	if bookID == "" || bookName == "" {
		response = Model.JsonResponse{Type: "error", Message: "You are missing bookID or bookName parameter."}
	} else {
		db := bbdd.SetUpDB()

		utils.PrintMessage("Inserting book into DB")

		fmt.Println("Inserting new book with ID: " + bookID + " and name: " + bookName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO books(bookID, bookName) VALUES($1, $2) returning id;", bookID, bookName).Scan(&lastInsertID)

		utils.CheckErr(err)

		response = Model.JsonResponse{Type: "success", Message: "The book has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
