package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/api/item", GetItemHandler).Methods("GET")
	router.HandleFunc("/api/item", CreateItemHandler).Methods("POST")
	router.HandleFunc("/api/item/{id}", GetItemByIDHandler).Methods("Get")

	fmt.Println("Server is running on Port :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API."))
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get items"))
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create an item"))
}

func GetItemByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create item by ID."))
}
