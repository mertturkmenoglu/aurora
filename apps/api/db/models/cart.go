package models

import "github.com/google/uuid"

type Cart struct {
	BaseModel
	UserId uuid.UUID  `json:"userId"`
	Items  []CartItem `json:"items"`
}

type CartItem struct {
	BaseModel
	CartId    uuid.UUID `json:"cartId"`
	ProductId uuid.UUID `json:"productId"`
	Product   Product   `json:"product"`
	Quantity  uint8     `json:"quantity"`
}
