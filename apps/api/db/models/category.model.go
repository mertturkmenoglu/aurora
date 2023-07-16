package models

import "github.com/google/uuid"

type Category struct {
	BaseModel
	Name     string     `json:"name"`
	ParentId *uuid.UUID `json:"parentId"`
	Parent   *Category  `json:"parent"`
}
