package models

import "time"

type InterviewBody struct {
	Interview
	UserDetail
}

// Interview represents the interview table.
type Interview struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar" json:"name" binding:"required"`
	Description string `gorm:"type:varchar" json:"description"`
	Status      string `gorm:"type:char" json:"status"`
	DefaultSchema
}

type InterviewCommentBody struct {
	InterviewComment
	UserDetail
}

// InterviewComment represents the interview_comment table.
type InterviewComment struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	InterviewID int    `gorm:"type:int" json:"interview_id"`
	Text        string `gorm:"type:varchar" json:"text" binding:"required"`
	DefaultSchema
}

// InterviewLog represents the interview_log table.
type InterviewLog struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	InterviewID int    `gorm:"type:int"`
	Name        string `gorm:"type:varchar"`
	Description string `gorm:"type:varchar"`
	Status      string `gorm:"type:char"`
	DefaultSchema
	LogDate time.Time `json:"log_date"`
}
