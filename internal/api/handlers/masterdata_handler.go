package handlers

import (
	"main/internal/api/repository"

	"github.com/gin-gonic/gin"
)

func GetMasterStatus(c *gin.Context) {

	status, err := repository.GetStatusByParam("status")
	if err != nil {
		c.Set("error", err.Error())
		return
	}
	c.Set("data", status)

}

func GetMasterRecordStatus(c *gin.Context) {

	status, err := repository.GetStatusByParam("record_status")
	if err != nil {
		c.Set("error", err.Error())
		return
	}
	c.Set("data", status)
}
