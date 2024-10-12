package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	No_Telp   string             `bson:"no_telp,omitempty" json:"no_telp,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Alamat    string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
