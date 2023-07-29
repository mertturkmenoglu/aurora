package dto

type UpdateCartItemDto struct {
	Quantity uint8 `json:"quantity" binding:"required"`
}
