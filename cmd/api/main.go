package main

import (
	"fmt"
	"main/internal/api/middleware"
	"main/internal/api/routes"
	"main/internal/config"
	"main/internal/db/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to the database

	if err := postgres.ConnectDB(); err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	if err := postgres.InitializeDatabase(); err != nil {
		fmt.Println("Failed to initialize the database:", err)
		return
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	// r.Use(middleware.AuthMiddleware())

	r.Use(middleware.JSONResponse())
	r.Use(middleware.SetPaginationData())

	rg := r.Group("/api/v1")
	routes.UserRoutes(rg.Group("/users"))

	routes.TestRoutes(rg.Group("/tests"))

	routes.AuthRoutes(rg.Group("/auth"))

	// Set the 404 handler function
	r.NoRoute(middleware.NotFound)

	r.Run(":8080")
}

// func main() {
// 	// Example usage of HashPassword
// 	password := "mysecretpassword"
// 	saltBase64, hashedPasswordBase64, err := utils.HashPassword(password)
// 	if err != nil {
// 		log.Fatal("Error hashing password:", err)
// 	}

// 	// Print the base64-encoded salt and hashed password
// 	fmt.Println("Base64-encoded salt:", saltBase64)
// 	fmt.Println("Base64-encoded hashed password:", hashedPasswordBase64)

// 	// Example usage of ComparePasswordHash
// 	providedPassword := "mysecretpasswords"

// 	// Retrieve the salt and hashed password from the database
// 	// In this example, we assume the saltBase64 and hashedPasswordBase64 values are retrieved from the database

// 	// Verify the provided password against the stored hashed password
// 	err = utils.ComparePasswordHash(providedPassword, hashedPasswordBase64, saltBase64)
// 	if err != nil {
// 		log.Fatal("Password verification failed:", err)
// 	}

// 	fmt.Println("Password verification successful!")
// }
