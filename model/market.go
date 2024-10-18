package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ProductName string             `bson:"product_name" json:"product_name,omitempty"`
	Price       float32            `bson:"price" json:"price,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Category    string             `bson:"category" json:"category,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type Market struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MarketName string             `bson:"nama_toko" json:"nama_toko"`
	Slug       string             `bson:"slug" json:"slug"`
	Alamat     Address             `bson:"alamat" json:"alamat"`
	Menu       []Menu             `bson:"menu" json:"menu"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type Address struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id,omitempty"`
	Street      string             `bson:"street" json:"street,omitempty"`
	Province    string             `bson:"province" json:"province,omitempty"`
	City        string             `bson:"city" json:"city,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	PostalCode  string             `bson:"postal_code" json:"postal_code,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
