package handlers

import (
	"main/internal/api/repository"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	c.Set("data", "1234")
}

func TestDB(c *gin.Context) {
	repository.GetUserByUsername("tess")
}
