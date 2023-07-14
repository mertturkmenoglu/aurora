package models

type Auth struct {
	BaseModel
	FullName string
	Email    string
	Password string
}
