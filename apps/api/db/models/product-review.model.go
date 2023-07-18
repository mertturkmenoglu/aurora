package models

import "github.com/google/uuid"

type ProductReview struct {
	BaseModel
	Comment   string    `json:"comment"`
	Rating    int       `json:"rating"`
	ProductId uuid.UUID `json:"productId"`
	UserId    uuid.UUID `json:"userId"`
	Product   Product   `json:"product"`
	User      User      `json:"user"`
}
