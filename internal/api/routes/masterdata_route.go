package routes

import (
	"main/internal/api/handlers"
	"main/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func MasterDataRoutes(r *gin.RouterGroup) {

	r.Use(middleware.AuthMiddleware())

	r.GET("/recordstatus", handlers.GetMasterRecordStatus)
	r.GET("/status", handlers.GetMasterStatus)

}
