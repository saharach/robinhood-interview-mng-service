package models

import (
	"time"
)

type UserDetail struct {
	CreateUserFullname string `json:"create_user_fullname"`
	CreateUserEmail    string `json:"create_user_email"`
}

// UserInfo represents the user_info table
type UserInfo struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email"`
	Salt      string `json:"salt,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DefaultSchema
	LastLogin time.Time `json:"last_login"`
}

// UserRole represents the user_role table
type UserRole struct {
	ID     int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
	DefaultSchema
}

// RoleDescription represents the role_description table
type RoleDescription struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
	DefaultSchema
}
