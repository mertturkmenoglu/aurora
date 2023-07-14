package models

type Product struct {
	BaseModel
	Name         string
	Description  string
	CurrentPrice float64
	OldPrice     float64
	Inventory    int
	// Images        []string
	IsFeatured    bool
	IsNew         bool
	IsOnSale      bool
	IsPopular     bool
	ShippingPrice float64
	ShippingTime  string
	ShippingType  string
	Slug          string
	BrandId       string
	Brand         Brand
}
