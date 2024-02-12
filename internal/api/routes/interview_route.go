package routes

import (
	"main/internal/api/handlers"
	"main/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func InterviewRoutes(r *gin.RouterGroup) {

	r.Use(middleware.AuthMiddleware())

	r.GET("/", handlers.GetAllInterviews)
	r.GET("/:id", handlers.GetInterviewByID)
	r.POST("/", handlers.CreateInterview)
	r.PUT("/:id", handlers.UpdateInterview)
	r.DELETE("/:id", handlers.DeleteInterview)

	r.PUT("/:id/archives", handlers.ArchiveInterview)

	r_comment := r.Group("/:id/comments")
	r_comment.GET("/", handlers.GetInterviewCommentAll)
	r_comment.POST("/", handlers.CreateInterviewComment)
	r_comment.PUT("/:cid", handlers.UpdateInterviewComment)
	r_comment.DELETE("/:cid", handlers.DeleteInterviewComment)

	r_log := r.Group("/:id/logs")
	r_log.GET("/", handlers.GetInterviewLogAll)
}
