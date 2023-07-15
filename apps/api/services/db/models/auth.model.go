package models

type Auth struct {
	BaseModel
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
