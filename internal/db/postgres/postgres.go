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
		&models.RoleDescription{},
		&models.Interview{},
		&models.InterviewComment{},
		&models.InterviewLog{},
		&models.StatusDescription{}); err != nil {
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

	defaultSchema := models.DefaultSchema{
		RecordStatus: "A",
		CreateUser:   99,
		CreateDate:   time.Now(),
	}
	// Check if the table is empty
	if err := db.Model(&models.StatusDescription{}).Count(&count).Error; err != nil {
		return err
	}

	// Insert initial data if the table is empty
	if count == 0 {
		datas := []*models.StatusDescription{
			{Param: "record_status", Key: "A", Description: "Active", DefaultSchema: defaultSchema},
			{Param: "record_status", Key: "I", Description: "Inctive", DefaultSchema: defaultSchema},
			{Param: "record_status", Key: "P", Description: "Archive", DefaultSchema: defaultSchema},
			{Param: "status", Key: "T", Description: "To do", DefaultSchema: defaultSchema},
			{Param: "status", Key: "I", Description: "In progres", DefaultSchema: defaultSchema},
			{Param: "status", Key: "D", Description: "Done", DefaultSchema: defaultSchema},
		}
		for _, data := range datas {
			err := db.Create(data).Error
			if err != nil {
				return err
			}
		}
	}

	// Check if the table is empty
	if err := db.Model(&models.RoleDescription{}).Count(&count).Error; err != nil {
		return err
	}

	// Insert initial data if the table is empty
	if count == 0 {
		datas := []*models.RoleDescription{
			{RoleName: "admin", DefaultSchema: defaultSchema},
			{RoleName: "user", DefaultSchema: defaultSchema},
		}
		for _, data := range datas {
			err := db.Create(data).Error
			if err != nil {
				return err
			}
		}
	}

	// Check if the table is empty
	if err := db.Model(&models.UserInfo{}).Count(&count).Error; err != nil {
		return err
	}

	userDatas := []*models.UserInfo{
		{UserName: "admin", Email: "admin@robihood.co.th", FirstName: "Robinhood", LastName: "Admin", DefaultSchema: defaultSchema},
		{UserName: "user1", Email: "user1@robihood.co.th", FirstName: "User1", LastName: "", DefaultSchema: defaultSchema},
		{UserName: "user2", Email: "user2@robihood.co.th", FirstName: "User2", LastName: "", DefaultSchema: defaultSchema},
	}
	// Insert initial data if the table is empty
	if count == 0 {

		for _, data := range userDatas {
			// Fix all user password at initial
			salt, hashedPassword, err := utils.HashPassword("password")
			if err != nil {
				return err
			}
			data.Password = string(hashedPassword)
			data.Salt = string(salt)
			err = db.Create(data).Error
			if err != nil {
				return err
			}
		}
	}

	// Check if the table is empty
	if err := db.Model(&models.UserRole{}).Count(&count).Error; err != nil {
		return err
	}
	// Insert initial data if the table is empty
	if count == 0 {
		for _, data := range userDatas {
			var user models.UserInfo
			var role models.RoleDescription
			if err := db.Where("user_name = ?", data.UserName).First(&user).Error; err != nil {
				return err
			}
			role_name := "user"
			if data.UserName == "admin" {
				role_name = "admin"
			}
			if err := db.Where("role_name = ?", role_name).First(&role).Error; err != nil {
				return err
			}
			if err := db.Create(&models.UserRole{
				UserID:        user.ID,
				RoleID:        role.ID,
				DefaultSchema: defaultSchema,
			}).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
