package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	wireConfig "github.com/pedro-muniz/myPrice/auth/infra/config/wire"
)

func main() {
	// Load environment variables from .env file if it exists
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, using existing environment variables")
	}

	authController := wireConfig.InitializeAuthController()

	// Setup routes
	http.HandleFunc("/authorize", authController.Authorize)
	http.HandleFunc("/validate", authController.Validate)

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
