package models

import "time"

const (
	StatusToDo                   = 'T'
	StatusInProgress             = 'I'
	StatusDone                   = 'D'
	RecordStatusActive           = 'A'
	RecordStatusInactiveOrDelete = 'I'
)

type DefaultSchema struct {
	RecordStatus string    `gorm:"type:char" json:"record_status"`
	CreateDate   time.Time `gorm:"type:timestamp" json:"create_date"`
	CreateUser   int       `gorm:"type:int" json:"create_user"`
	UpdateDate   time.Time `gorm:"type:timestamp" json:"update_date"`
	UpdateUser   int       `gorm:"type:int" json:"update_user"`
}

type StatusDescription struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Param       string `gorm:"type:varchar" json:"param"`
	Key         string `gorm:"type:char" json:"key"`
	Description string `gorm:"type:varchar" json:"description"`
	DefaultSchema
}
