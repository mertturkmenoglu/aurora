package models

import "github.com/google/uuid"

type BrandReview struct {
	BaseModel
	Comment string    `json:"comment"`
	Rating  int       `json:"rating"`
	BrandId uuid.UUID `json:"brandId"`
	UserId  uuid.UUID `json:"userId"`
	Brand   Brand     `json:"brand"`
	User    User      `json:"user"`
}
