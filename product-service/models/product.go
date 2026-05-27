package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Price     float64            `json:"price" bson:"price"`
	Stock     int                `json:"stock" bson:"stock"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ProductResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Price     float64            `json:"price"`
	Stock     int                `json:"stock"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at,omitempty"`
}
