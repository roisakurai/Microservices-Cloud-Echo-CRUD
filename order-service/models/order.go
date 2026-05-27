package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderRequest struct {
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
}

type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
	Total     float64            `json:"total" bson:"total"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type OrderResponse struct {
	ID           primitive.ObjectID `json:"id"`
	UserID       primitive.ObjectID `json:"user_id"`
	Username     string             `json:"username,omitempty"`
	ProductID    primitive.ObjectID `json:"product_id"`
	ProductName  string             `json:"product_name,omitempty"`
	ProductPrice float64            `json:"product_price,omitempty"`
	Quantity     int                `json:"quantity"`
	Total        float64            `json:"total"`
	Status       string             `json:"status"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at,omitempty"`
}
