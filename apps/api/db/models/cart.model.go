package models

import "github.com/google/uuid"

type Cart struct {
	BaseModel
	UserID   uuid.UUID     `json:"userId"`
	Products []CartProduct `json:"products"`
}

type CartProduct struct {
	BaseModel
	UserID    uuid.UUID `json:"userId"`
	CartID    uuid.UUID `json:"cartId"`
	ProductID uuid.UUID `json:"productId"`
	ItemCount int       `json:"itemCount"`
}
