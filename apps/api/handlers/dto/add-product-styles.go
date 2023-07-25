package dto

type AddProductStylesDto struct {
	Styles []AddProductStylesItemDto `json:"styles" binding:"required"`
}

type AddProductStylesItemDto struct {
	Name string `json:"name" binding:"required"`
}
