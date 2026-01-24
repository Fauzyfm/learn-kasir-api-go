package main

import (
	"fmt"
	"net/http"

	"kasir-api/internal/router"
)

func main() {
	// Setup semua routes
	router.SetupRoutes()

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
