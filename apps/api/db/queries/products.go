package queries

import (
	"aurora/db"
	"aurora/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

const (
	conditionFeatured = "is_featured = ?"
	conditionNew      = "is_new = ?"
	conditionSale     = "is_on_sale = ?"
	conditionPopular  = "is_popular = ?"
)

func getProductsByCondition(condition string) ([]*models.Product, error) {
	var products []*models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
		Preload("ProductVariants.Image").
		Preload("ProductVariants.ProductStyle").
		Preload("ProductVariants.ProductSize").
		Order("created_at desc").
		Limit(25).
		Find(&products, condition, true)

	return products, res.Error
}

func GetProductById(id uuid.UUID) (*models.Product, error) {
	var product *models.Product

	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
		Preload("ProductVariants.Image").
		Preload("ProductVariants.ProductStyle").
		Preload("ProductVariants.ProductSize").
		First(&product, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return product, nil
}

func GetFeaturedProducts() ([]*models.Product, error) {
	return getProductsByCondition(conditionFeatured)
}

func GetNewProducts() ([]*models.Product, error) {
	return getProductsByCondition(conditionNew)
}

func GetSaleProducts() ([]*models.Product, error) {
	return getProductsByCondition(conditionSale)
}

func GetPopularProducts() ([]*models.Product, error) {
	return getProductsByCondition(conditionPopular)
}
