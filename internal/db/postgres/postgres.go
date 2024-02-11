package postgres

import (
	"fmt"
	"main/internal/api/models"
	"main/internal/config"
	"main/internal/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	// Load configuration
	cfg := config.Config

	// Create DSN string using configuration values
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabaseName, cfg.DatabasePassword, cfg.DatabasePort)

	// Establish database connection
	db, err := establishConnection(dsn)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func establishConnection(dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Retry connection with a backoff mechanism
	for i := 0; i < 15; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		fmt.Printf("Failed to connect to database (attempt %d). Retrying...\n", i+1)
		time.Sleep(time.Second * time.Duration(i+1))
	}

	return nil, fmt.Errorf("failed to connect to database after multiple attempts: %v", err)
}

func InitializeDatabase() error {
	// Create the database if it does not exist

	// Auto-migrate the schema
	if err := DB.AutoMigrate(
		&models.UserInfo{},
		&models.UserRole{},
		&models.RoleDescription{}); err != nil {
		return fmt.Errorf("failed to auto-migrate database schema: %v", err)
	}

	// Seed initial data if the table is empty
	if err := seedData(DB); err != nil {
		return fmt.Errorf("failed to seed data: %v", err)
	}

	return nil
}

func seedData(db *gorm.DB) error {
	var count int64
	// Check if the table is empty
	if err := db.Model(&models.RoleDescription{}).Count(&count).Error; err != nil {
		return err
	}

	// Insert initial data if the table is empty
	if count == 0 {
		err := db.Create(&models.RoleDescription{
			RoleName:     "admin",
			RecordStatus: "A",
			CreateUser:   99,
			CreateDate:   time.Now(),
		}).Error
		if err != nil {
			return err
		}
	}

	// Check if the table is empty
	if err := db.Model(&models.UserInfo{}).Count(&count).Error; err != nil {
		return err
	}

	// Insert initial data if the table is empty
	if count == 0 {
		salt, hashedPassword, err := utils.HashPassword("password")
		if err != nil {
			return err
		}
		// Insert initial data into the table
		err = db.Create(&models.UserInfo{
			UserName:     "admin",
			Password:     string(hashedPassword),
			Salt:         string(salt),
			RecordStatus: "A",
			CreateUser:   99,
			CreateDate:   time.Now(),
		}).Error
		if err != nil {
			return err
		}
	}

	// Check if the table is empty
	if err := db.Model(&models.UserRole{}).Count(&count).Error; err != nil {
		return err
	}

	// Insert initial data if the table is empty
	if count == 0 {
		// Insert initial data into the table
		// Retrieve the auto-generated ID after creating the record
		var user models.UserInfo
		var role models.RoleDescription
		if err := db.Where("user_name = ?", "admin").First(&user).Error; err != nil {
			return err
		}
		if err := db.Where("role_name = ?", "admin").First(&role).Error; err != nil {
			return err
		}
		if err := db.Create(&models.UserRole{
			UserID: user.ID,
			RoleID: role.ID,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}
