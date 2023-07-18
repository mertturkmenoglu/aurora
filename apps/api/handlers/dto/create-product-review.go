package dto

type CreateProductReviewDto struct {
	Comment   string `json:"comment" binding:"required"`
	Rating    int    `json:"rating" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
}
