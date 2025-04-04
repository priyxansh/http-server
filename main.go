package main

import (
	"fmt"
	"net/http"
)

const (
	PORT = ":5000"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func createServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleRoot)

	fmt.Println("Starting server on http://localhost" + PORT)

	return http.ListenAndServe(PORT, mux)
}

func main() {
	if err := createServer(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
