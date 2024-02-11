package handlers

import (
	"log"
	"main/internal/api/models"
	"main/internal/api/repository"
	"main/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	start := time.Now()

	log.Printf("1 %s", time.Since(start))

	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.Set("validate", err.Error())
		return
	}
	user, err := repository.GetUserByUsername(creds.Username)
	if err != nil {
		c.Set("error", err.Error())
		return
	}
	log.Printf("2 %s", time.Since(start))

	err = utils.ComparePasswordHash(creds.Password, user.Password, user.Salt)
	if err != nil {
		c.Set("error", err.Error())
		return
	}

	token, err := utils.GenerateToken(creds.Username)
	if err != nil {
		c.Set("error", err.Error())
		return
	}

	c.Set("data", token)

	log.Printf("3 %s", time.Since(start))

}
