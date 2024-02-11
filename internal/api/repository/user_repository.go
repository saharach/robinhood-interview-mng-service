package repository

import (
	"main/internal/api/models"
	"main/internal/db/postgres"

	"github.com/jinzhu/gorm"
)

// GetUserByUsername retrieves a user from the database based on the provided username
func GetUserByUsername(username string) (*models.UserInfo, error) {
	var user models.UserInfo
	if err := postgres.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetUserRole retrieves the role of a user from the database based on their user ID
func GetUserRole(userID int) (string, error) {
	// Implement logic to retrieve user's role from the database
	return "", nil
}
