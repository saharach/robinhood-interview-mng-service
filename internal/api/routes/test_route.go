package routes

import (
	"main/internal/api/handlers"
	"main/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func TestRoutes(r *gin.RouterGroup) {
	// api_user := r.Group("/users")
	r.GET("/test", middleware.AuthMiddleware(), handlers.TestHandler)
}
