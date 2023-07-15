package dto

type CreateProductDto struct {
	Name          string            `json:"name" binding:"required"`
	Description   string            `json:"description" binding:"required"`
	CurrentPrice  float64           `json:"currentPrice" binding:"required"`
	OldPrice      float64           `json:"oldPrice" binding:"required"`
	Inventory     int               `json:"inventory" binding:"required"`
	Images        []ProductImageDto `json:"images" binding:"required"`
	IsFeatured    bool              `json:"isFeatured" validate:"exists"`
	IsNew         bool              `json:"isNew" validate:"exists"`
	IsOnSale      bool              `json:"isOnSale" validate:"exists"`
	IsPopular     bool              `json:"isPopular" validate:"exists"`
	ShippingPrice float64           `json:"shippingPrice" binding:"required"`
	ShippingTime  string            `json:"shippingTime" binding:"required"`
	ShippingType  string            `json:"shippingType" binding:"required"`
	Slug          string            `json:"slug" binding:"required"`
	BrandId       string            `json:"brandId" binding:"required"`
}

type ProductImageDto struct {
	Url string `json:"url" binding:"required"`
}
