package dto

import "github.com/google/uuid"

type CreateCategoryDto struct {
	Name     string     `json:"name"`
	ParentId *uuid.UUID `json:"parentId"`
}
