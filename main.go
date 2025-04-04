package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
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

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var (
	userMap = make(map[string]User)
	mu      sync.Mutex
)

func handleUserPost(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId := uuid.NewString()

	mu.Lock()
	defer mu.Unlock()
	userMap[userId] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]any{
		"id":   userId,
		"user": user,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("GET /json", handleJSON)
	mux.HandleFunc("GET /json-struct", handleJSONWithStruct)
	mux.HandleFunc("POST /user", handleUserPost)

	fmt.Println("Starting server on http://localhost" + PORT)

	return http.ListenAndServe(PORT, mux)
}

func main() {
	if err := createServer(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
