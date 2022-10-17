package models

import (
	"time"
)

type Blog struct {
	Id 			uint64 		`gorm:"primayKey" json:"id"`
	Title 		string		`gorm:"type:varchar(300)" json:"title"`
	Content 	string		`gorm:"type:longtext" json:"content"`
	Slug		string 		`gorm:"type:varchar(300)" json:"slug"`
	CreatedAt	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}