package routes

import (
	"main/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func TestRoutes(r *gin.RouterGroup) {
	// api_user := r.Group("/users")
	r.GET("/test", handlers.TestDB)
}
