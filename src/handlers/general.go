package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, JSON!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type Response struct {
	Message string `json:"message"`
}

func HandleJSONWithStruct(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, JSON with Struct!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
