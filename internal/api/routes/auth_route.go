package routes

import (
	"main/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	// api_user := r.Group("/users")
	r.POST("/login", handlers.LoginHandler)
}
