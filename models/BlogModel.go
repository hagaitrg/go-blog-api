package models

import (
	"time"
)

type Blog struct {
	Title 		string				`json:"title"`		
	Content 	string				`json:"content"`		
	Slug		string				`json:"slug"` 		
	CreatedAt	time.Time			`json:"created_at"` 
	UpdatedAt	time.Time			`json:"updated_at"`  	
}