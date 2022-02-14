package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jperarm1971/convenios-golang-api/routers"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "admin"
	DB_NAME     = "josue"
)

type Book struct {
	BookID   string `json:"bookid"`
	BookName string `json:"bookname"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Book `json:"data"`
	Message string `json:"message"`
}

func main() {
	router := mux.NewRouter()

	// Get all books
	router.HandleFunc("/books/", routers.GetBooks).Methods("GET")

	// Create a book
	router.HandleFunc("/books/", CreateBook).Methods("POST")

	// Delete a specific book by the bookID
	router.HandleFunc("/books/{bookid}", DeleteBook).Methods("DELETE")

	// Delete all books
	router.HandleFunc("/books/", DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Get all books

// Create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookid")
	bookName := r.FormValue("bookname")

	var response = JsonResponse{}

	if bookID == "" || bookName == "" {
		response = JsonResponse{Type: "error", Message: "You are missing bookID or bookName parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting book into DB")

		fmt.Println("Inserting new book with ID: " + bookID + " and name: " + bookName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO books(bookID, bookName) VALUES($1, $2) returning id;", bookID, bookName).Scan(&lastInsertID)

		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The book has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID := params["bookid"]

	var response = JsonResponse{}

	if bookID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing bookID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting book from DB")

		_, err := db.Exec("DELETE FROM books where bookID = $1", bookID)
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The book has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete all books
func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all books...")

	_, err := db.Exec("DELETE FROM books")
	checkErr(err)

	printMessage("All books have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All books have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

// DB set up
func setupDB() *sql.DB {

	host := "192.168.1.39"
	port := "5432"
	user := "postgres"
	password := "admin"
	dbname := "josue"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", psqlInfo)

	checkErr(err)

	return db
}

// Function for handling messages
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
