package config

import (
	"fmt"
	"net/http"

	"http-server/src/routes"
)

const PORT = ":5000"

func CreateServer() error {

	fmt.Println("Starting server on http://localhost" + PORT)

	return http.ListenAndServe(PORT, routes.SetupRoutes())
}
