package main

import (
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
	postgres.ConnectDB()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	// r.Use(middleware.AuthMiddleware())

	r.Use(middleware.JSONResponse())
	r.Use(middleware.SetPaginationData())

	rg := r.Group("/api/v1")
	routes.UserRoutes(rg.Group("/users"))

	routes.TestRoutes(rg.Group("/tests"))

	// Set the 404 handler function
	r.NoRoute(middleware.NotFound)

	r.Run(":8080")
}
