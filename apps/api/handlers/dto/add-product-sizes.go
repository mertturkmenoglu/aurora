package dto

type AddProductSizesDto struct {
	Sizes []AddProductSizesItemDto `json:"sizes" binding:"required"`
}

type AddProductSizesItemDto struct {
	Name string `json:"name" binding:"required"`
}
