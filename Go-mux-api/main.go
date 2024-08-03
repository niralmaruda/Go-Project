package main

import (
	"Go-mux-api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var items = make(map[string]models.Item)

func init() {
	items["1"] = models.Item{ID: "1", Name: "Item 1", Description: "This is item 1"}
	items["2"] = models.Item{ID: "2", Name: "Item 2", Description: "This is item 2"}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/api/items", GetItemsHandler).Methods("GET")
	router.HandleFunc("/api/item", CreateItemHandler).Methods("POST")
	router.HandleFunc("/api/item/{id}", GetItemByIDHandler).Methods("Get")
	router.HandleFunc("/api/item/{id}", UpdateItemHandler).Methods("PUT")
	router.HandleFunc("/api/item/{id}", DeleteItemHandler).Methods("DELETE")

	fmt.Println("Server is running on Port :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API."))
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = uuid.New().String()
	items[item.ID] = item
	json.NewEncoder(w).Encode(item)
}

func GetItemByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var item, exists = items[id]

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Item not found")
		return
	}

	json.NewEncoder(w).Encode(item)
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedItem models.Item
	_ = json.NewDecoder(r.Body).Decode(&updatedItem)
	item, exists := items[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Item not found")
		return
	}
	updatedItem.ID = item.ID
	items[id] = updatedItem
	json.NewEncoder(w).Encode(updatedItem)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_, exists := items[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Item not found")
		return
	}
	delete(items, id)
	w.WriteHeader(http.StatusNoContent)
}
