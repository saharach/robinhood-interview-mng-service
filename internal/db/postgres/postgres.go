package postgres

import (
	"fmt"
	"main/internal/api/models"
	"main/internal/config"
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

	// Ensure database is initialized
	if err := initializeDatabase(db); err != nil {
		return err
	}

	DB = db
	return nil
}

func establishConnection(dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Retry connection with a backoff mechanism
	for i := 0; i < 50; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		fmt.Printf("Failed to connect to database (attempt %d). Retrying...\n", i+1)
		time.Sleep(time.Second * time.Duration(i+1))
	}

	return nil, fmt.Errorf("failed to connect to database after multiple attempts: %v", err)
}

func initializeDatabase(db *gorm.DB) error {
	// Create the database if it does not exist

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.UserInfo{}); err != nil {
		return fmt.Errorf("failed to auto-migrate database schema: %v", err)
	}

	// Seed initial data if the table is empty
	if err := seedData(db); err != nil {
		return fmt.Errorf("failed to seed data: %v", err)
	}

	return nil
}

func seedData(db *gorm.DB) error {
	// var count int64

	// // Check if the table is empty
	// if err := db.Model(&YourModel{}).Count(&count).Error; err != nil {
	// 	return err
	// }

	// // Insert initial data if the table is empty
	// if count == 0 {
	// 	// Insert initial data into the table
	// 	// Example:
	// 	// err := db.Create(&YourModel{Field1: "value1", Field2: "value2"}).Error
	// 	// if err != nil {
	// 	//     return err
	// 	// }
	// }

	return nil
}
