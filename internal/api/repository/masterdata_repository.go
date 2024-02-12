package repository

import (
	"main/internal/api/models"
	"main/internal/db/postgres"
)

func GetStatusByParam(param string) ([]*models.StatusDescription, error) {
	var status []*models.StatusDescription
	if err := postgres.DB.Where("param = ? AND record_status = ?", param, "A").Find(&status).Error; err != nil {
		return nil, err
	}
	return status, nil
}
