package models

import "github.com/google/uuid"

type Product struct {
	BaseModel
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	CurrentPrice  float64        `json:"currentPrice"`
	OldPrice      float64        `json:"oldPrice"`
	Inventory     int            `json:"inventory"`
	Images        []ProductImage `json:"images"`
	IsFeatured    bool           `json:"isFeatured"`
	IsNew         bool           `json:"isNew"`
	IsOnSale      bool           `json:"isOnSale"`
	IsPopular     bool           `json:"isPopular"`
	ShippingPrice float64        `json:"shippingPrice"`
	ShippingTime  string         `json:"shippingTime"`
	ShippingType  string         `json:"shippingType"`
	Slug          string         `json:"slug"`
	BrandId       uuid.UUID      `json:"brandId"`
	Brand         Brand          `json:"brand"`
	CategoryId    uuid.UUID      `json:"categoryId"`
	Category      Category       `json:"category"`
}

type ProductImage struct {
	BaseModel
	ProductId uuid.UUID `json:"productId"`
	Url       string    `json:"url"`
}
