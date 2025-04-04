package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"http-server/src/models"

	"github.com/google/uuid"
)

var (
	userMap = make(map[string]models.User)
	mu      sync.RWMutex
)

func HandleUserPost(w http.ResponseWriter, r *http.Request) {
	var user models.User
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

func HandleUsersGet(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(userMap); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleUserDelete(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")

	mu.Lock()
	defer mu.Unlock()

	if _, exists := userMap[userId]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(userMap, userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"message": "User deleted successfully"}

	json.NewEncoder(w).Encode(response)
}
