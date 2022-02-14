package routers

import (
	"encoding/json"
	"fmt"

	"net/http"

	_ "github.com/lib/pq"

	bbdd "github.com/jperarm1971/convenios-golang-api/BBDD"
	Model "github.com/jperarm1971/convenios-golang-api/Models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db := bbdd.SetUpDB()

	printMessage("obteniendo libros...")

	// Get all books from books table that don't have bookID = "1"
	rows, err := db.Query("SELECT * FROM books where bookID <> $1", "1")

	checkErr(err)
	var books []Model.Book
	// var response []JsonResponse
	// Foreach book
	for rows.Next() {
		var id int
		var bookID string
		var bookName string

		err = rows.Scan(&id, &bookID, &bookName)

		checkErr(err)

		books = append(books, Model.Book{BookID: bookID, BookName: bookName})
	}

	var response = Model.JsonResponse{Type: "success", Data: books}

	json.NewEncoder(w).Encode(response)
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
