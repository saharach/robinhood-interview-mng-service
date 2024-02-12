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

	routes.AuthRoutes(rg.Group("/auth"))
	routes.UserRoutes(rg.Group("/users"))
	routes.InterviewRoutes(rg.Group("/interviews"))
	routes.MasterDataRoutes(rg.Group("/master"))

	// Set the 404 handler function
	r.NoRoute(middleware.NotFound)

	r.Run(":8080")
}
