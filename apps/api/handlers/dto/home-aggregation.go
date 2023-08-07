package dto

import "aurora/db/models"

type HomeAggregation struct {
	FeaturedProducts []*models.Product `json:"featured"`
	NewProducts      []*models.Product `json:"new"`
	SaleProducts     []*models.Product `json:"sale"`
	PopularProducts  []*models.Product `json:"popular"`
}
