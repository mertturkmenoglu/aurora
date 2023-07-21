package dto

type AddFavoriteDto struct {
	ProductId string `json:"productId" binding:"required"`
}
