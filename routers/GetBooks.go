package routers

import (
	"encoding/json"

	"net/http"

	_ "github.com/lib/pq"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db := SetUpDB()

	printMessage("obteniendo libros...")

	// Get all books from books table that don't have bookID = "1"
	rows, err := db.Query("SELECT * FROM books where bookID <> $1", "1")

	checkErr(err)
	var books []Book
	// var response []JsonResponse
	// Foreach book
	for rows.Next() {
		var id int
		var bookID string
		var bookName string

		err = rows.Scan(&id, &bookID, &bookName)

		checkErr(err)

		books = append(books, Book{BookID: bookID, BookName: bookName})
	}

	var response = JsonResponse{Type: "success", Data: books}

	json.NewEncoder(w).Encode(response)
}
