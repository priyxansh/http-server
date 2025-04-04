package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	PORT = ":5000"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, JSON!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type Response struct {
	Message string `json:"message"`
}

func handleJSONWithStruct(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, JSON with Struct!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("GET /json", handleJSON)
	mux.HandleFunc("GET /json-struct", handleJSONWithStruct)

	fmt.Println("Starting server on http://localhost" + PORT)

	return http.ListenAndServe(PORT, mux)
}

func main() {
	if err := createServer(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
