package models

import (
	"time"
)

type UserInfo struct {
	ID           int       `gorm:"primary_key" json:"id"`
	UserName     string    `json:"user_name"`
	Password     string    `json:"password"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	RecordStatus string    `json:"record_status"`
	CreateDate   time.Time `json:"create_date"`
	CreateUser   int       `json:"create_user"`
	UpdateDate   time.Time `json:"update_date"`
	UpdateUser   int       `json:"update_user"`
	LastLogin    time.Time `json:"last_login"`
}
