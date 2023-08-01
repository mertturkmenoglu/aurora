package dto

type CreateProductDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsFeatured  bool   `json:"isFeatured" validate:"exists"`
	IsNew       bool   `json:"isNew" validate:"exists"`
	IsOnSale    bool   `json:"isOnSale" validate:"exists"`
	IsPopular   bool   `json:"isPopular" validate:"exists"`

	BrandId    string `json:"brandId" binding:"required"`
	CategoryId string `json:"categoryId" binding:"required"`

	ProductVariants []ProductVariantDto `json:"variants" binding:"required"`
}

type ProductVariantDto struct {
	IsDefault      bool            `json:"isDefault" binding:"exists"`
	CurrentPrice   float64         `json:"currentPrice" binding:"required"`
	OldPrice       float64         `json:"oldPrice" binding:"required"`
	Inventory      int             `json:"inventory" binding:"required"`
	ShippingPrice  float64         `json:"shippingPrice" binding:"required"`
	ShippingTime   string          `json:"shippingTime" binding:"required"`
	ShippingType   string          `json:"shippingType" binding:"required"`
	ProductStyleId string          `json:"styleId" binding:"required"`
	ProductSizeId  string          `json:"sizeId" binding:"required"`
	Image          ProductImageDto `json:"image" binding:"required"`
}

type ProductImageDto struct {
	Url string `json:"url" binding:"required"`
}
