package models

import "github.com/google/uuid"

type Admin struct {
	BaseModel
	UserId uuid.UUID `gorm:"type:uuid;not null;unique" json:"userId"`
	User   User      `json:"user"`
}
