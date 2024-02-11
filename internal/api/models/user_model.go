package models

import (
	"time"
)

// UserInfo represents the user_info table
type UserInfo struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName     string    `json:"user_name"`
	Password     string    `json:"password,omitempty"`
	Salt         string    `json:"salt,omitempty"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	RecordStatus string    `json:"record_status"`
	CreateDate   time.Time `json:"create_date"`
	CreateUser   uint      `json:"create_user"`
	UpdateDate   time.Time `json:"update_date"`
	UpdateUser   uint      `json:"update_user"`
	LastLogin    time.Time `json:"last_login"`
}

// UserRole represents the user_role table
type UserRole struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint      `json:"user_id"`
	RoleID       uint      `json:"role_id"`
	RecordStatus string    `json:"record_status"`
	CreateDate   time.Time `json:"create_date"`
	CreateUser   uint      `json:"create_user"`
	UpdateDate   time.Time `json:"update_date"`
	UpdateUser   uint      `json:"update_user"`
}

// RoleDescription represents the role_description table
type RoleDescription struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName     string    `json:"role_name"`
	Description  string    `json:"description"`
	RecordStatus string    `json:"record_status"`
	CreateDate   time.Time `json:"create_date"`
	CreateUser   uint      `json:"create_user"`
	UpdateDate   time.Time `json:"update_date"`
	UpdateUser   uint      `json:"update_user"`
}
