package routes

import (
	"net/http"

	"http-server/src/handlers"
	"http-server/src/middleware"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.HandleRoot)
	mux.HandleFunc("GET /json", handlers.HandleJSON)
	mux.HandleFunc("GET /json-struct", handlers.HandleJSONWithStruct)
	mux.HandleFunc("POST /user", handlers.HandleUserPost)
	mux.HandleFunc("GET /user", handlers.HandleUsersGet)
	mux.HandleFunc("DELETE /user/{id}", handlers.HandleUserDelete)

	mux.Handle("GET /log-test", middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Logging test endpoint"))
	})))

	return mux
}
