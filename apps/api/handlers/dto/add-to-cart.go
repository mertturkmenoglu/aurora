package dto

type AddToCartDto struct {
	ProductId        string `json:"productId" binding:"required"`
	ProductVariantId string `json:"productVariantId" binding:"required"`
	Quantity         uint8  `json:"quantity" binding:"required"`
}
