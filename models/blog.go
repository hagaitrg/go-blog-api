package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Blog struct {
	ID 			primitive.ObjectID
	Title 		string		
	Content 	string		
	Slug		string 		
	CreatedAt	time.Time 
	UpdatedAt	time.Time 	
}