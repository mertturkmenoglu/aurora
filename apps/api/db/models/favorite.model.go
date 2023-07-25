package models

import "github.com/google/uuid"

type Favorite struct {
	BaseModel
	UserId    uuid.UUID `json:"userId"`
	ProductId uuid.UUID `json:"productId"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
}
