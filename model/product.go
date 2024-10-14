package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ProductName string             `bson:"product_name" json:"product_name,omitempty"`
	Price       float32            `bson:"price" json:"price,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	CategoryID  primitive.ObjectID `bson:"category_id" json:"category_id"`
}
