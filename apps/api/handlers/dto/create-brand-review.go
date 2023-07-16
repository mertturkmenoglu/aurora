package dto

type CreateBrandReviewDto struct {
	Comment string `json:"comment"`
	Rating  int    `json:"rating"`
	BrandId string `json:"brandId"`
}
