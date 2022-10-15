package models

import (
	"time"
)

type Blog struct {
	ID 			uint 		`json:"id" gorm:"primary_key"`
	title 		string		`json:"title gorm: "type:varchar(120)""`
	content 	string		`json:"content" gorm: "type:varchar(500)"`
	slug		string 		`json:"slug" gorm: "type:varchar(500)"`
	CreatedAt	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at`
	UpdatedAt	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at`
}