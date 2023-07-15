package models

import "github.com/google/uuid"

type User struct {
	BaseModel
	FullName     string       `json:"fullName"`
	Email        string       `json:"email"`
	Phone        string       `json:"phone"`
	AdPreference AdPreference `json:"adPreference"`
	Addresses    []Address    `json:"addresses"`
}

type AdPreference struct {
	BaseModel
	UserID uuid.UUID `json:"userId"`
	Email  bool      `json:"email"`
	Sms    bool      `json:"sms"`
	Phone  bool      `json:"phone"`
}

type Address struct {
	BaseModel
	UserID      uuid.UUID `json:"userId"`
	City        string    `json:"city"`
	Description string    `json:"description"`
	IsDefault   bool      `json:"isDefault"`
	Line1       string    `json:"line1"`
	Line2       string    `json:"line2"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	State       string    `json:"state"`
	Type        string    `json:"type"`
	ZipCode     string    `json:"zipCode"`
}
