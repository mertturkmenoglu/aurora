package dto

type AddToCartDto struct {
	ProductId string `json:"productId" binding:"required"`
	Quantity  uint8  `json:"quantity" binding:"required"`
}
