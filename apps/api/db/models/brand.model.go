package models

type Brand struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
