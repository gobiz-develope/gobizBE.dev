package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	CategoryName string `bson:"category_name" json:"category_name,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"` 
}