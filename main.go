package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jperarm1971/convenios-golang-api/routers"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()

	// Get all books
	router.HandleFunc("/books/", routers.GetBooks).Methods("GET")

	router.HandleFunc("/convenios", routers.GetConvenios).Methods("GET")

	router.HandleFunc("/altaconvenio", routers.CreaConvenio).Methods("POST")

	router.HandleFunc("/modificaconvenio", routers.ActualizaConvenio).Methods("POST")

	router.HandleFunc("/convenio", routers.GetConvenio).Methods("GET")

	// Create a book
	router.HandleFunc("/books/", routers.CreateBook).Methods("POST")

	// Delete a specific book by the bookID
	router.HandleFunc("/books/{bookid}", routers.DeleteBook).Methods("DELETE")

	// Delete all books
	router.HandleFunc("/books/", routers.DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
