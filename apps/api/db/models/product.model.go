package models

import "github.com/google/uuid"

type Product struct {
	BaseModel
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	IsFeatured       bool             `json:"isFeatured"`
	IsNew            bool             `json:"isNew"`
	IsOnSale         bool             `json:"isOnSale"`
	IsPopular        bool             `json:"isPopular"`
	ProductVariants  []ProductVariant `json:"variants"`
	BrandId          uuid.UUID        `json:"brandId"`
	Brand            Brand            `json:"brand"`
	CategoryId       uuid.UUID        `json:"categoryId"`
	Category         Category         `json:"category"`
	DefaultVariantId uuid.UUID        `json:"defaultVariantId"`
	DefaultVariant   ProductVariant   `json:"defaultVariant"`
}

type ProductVariant struct {
	BaseModel
	ProductId      uuid.UUID    `json:"productId"`
	CurrentPrice   float64      `json:"currentPrice"`
	OldPrice       float64      `json:"oldPrice"`
	Inventory      int          `json:"inventory"`
	ImageId        uuid.UUID    `json:"imageId"`
	Image          ProductImage `json:"image"`
	ShippingPrice  float64      `json:"shippingPrice"`
	ShippingTime   string       `json:"shippingTime"`
	ShippingType   string       `json:"shippingType"`
	ProductStyleId uuid.UUID    `json:"styleId"`
	ProductStyle   ProductStyle `json:"style"`
	ProductSizeId  uuid.UUID    `json:"sizeId"`
	ProductSize    ProductSize  `json:"size"`
}

type ProductImage struct {
	BaseModel
	ProductId        uuid.UUID `json:"productId"`
	ProductVariantId uuid.UUID `json:"productVariantId"`
	Url              string    `json:"url"`
}

type ProductStyle struct {
	BaseModel
	Name string `json:"name"`
}

type ProductSize struct {
	BaseModel
	Name string `json:"name"`
}
