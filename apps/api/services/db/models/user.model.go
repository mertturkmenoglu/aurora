package models

import "github.com/google/uuid"

type User struct {
	BaseModel
	FullName     string
	Email        string
	Phone        string
	AdPreference AdPreference
	Addresses    []Address
}

type AdPreference struct {
	BaseModel
	UserID uuid.UUID
	Email  bool
	Sms    bool
	Phone  bool
}

type Address struct {
	BaseModel
	UserID      uuid.UUID
	City        string
	Description string
	IsDefault   bool
	Line1       string
	Line2       string
	Name        string
	Phone       string
	State       string
	Type        string
	ZipCode     string
}
