package handlers

import (
	"main/internal/api/models"
	"main/internal/api/repository"
	"main/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func LoginHandler(c *gin.Context) {

	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.Set("validate", err.Error())
		return
	}
	user, err := repository.GetUserByUsername(creds.Username)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.Set("unauthorized", "The username/password not correct.")
			return
		} else {
			c.Set("error", err.Error())
			return
		}
	}

	err = utils.ComparePasswordHash(creds.Password, user.Password, user.Salt)
	if err != nil {
		if err.Error() == "Not match" {
			c.Set("unauthorized", "The username/password not correct.")
			return
		} else {
			c.Set("error", err.Error())
			return
		}
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		c.Set("error", err.Error())
		return
	}

	apiResponse := &models.LoginResponse{
		Success:    true,
		Token:      token,
		StatusCode: http.StatusOK,
	}
	c.JSON(http.StatusOK, apiResponse)
}
