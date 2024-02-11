package handlers

import (
	"main/internal/api/repository"

	"github.com/gin-gonic/gin"
)

func TestDB(c *gin.Context) {
	repository.GetUserByUsername("tess")
}
