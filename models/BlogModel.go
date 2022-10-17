package models

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Blog struct {
	// Id			primitive.ObjectID	`json:"id,omitempty"`
	Title 		string				`json:"title"`		
	Content 	string				`json:"content"`		
	Slug		string				`json:"slug"` 		
	CreatedAt	time.Time			`json:"created_at"` 
	UpdatedAt	time.Time			`json:"updated_at"`  	
}