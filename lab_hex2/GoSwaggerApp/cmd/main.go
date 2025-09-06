package main

import (
	"GoSwaggerApp/internal/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items", handlers.GetItems).Methods(http.MethodGet, http.MethodPost)
	// r.HandleFunc("/items/{id:[0-9]+}", handlers.GetItem).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	fmt.Println("Starting server on :8080")
	http.Handle("/", http.FileServer(http.Dir("./docs")))
	http.ListenAndServe(":8080", r)
}
