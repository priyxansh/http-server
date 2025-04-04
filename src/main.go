package main

import (
	"fmt"

	"http-server/src/config"
)

func main() {
	if err := config.CreateServer(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
