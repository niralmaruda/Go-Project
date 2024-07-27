package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func customHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	name := vars.Get("name")
	if name == "" {
		name = "Stranger"
	}

	response := Response{Message: fmt.Sprintf("Hello, %s!", name)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func htmxHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Stranger"
	}
	message := fmt.Sprintf("Hello, %s!", name)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>%s</h1>", message)
}

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/api/custom", customHandler)
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/htmx", htmxHandler)

	fmt.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
