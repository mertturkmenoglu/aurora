package models

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Product struct {
	Id            string   `dynamodbav:"id" json:"id"`
	Name          string   `dynamodbav:"name" json:"name"`
	Description   string   `dynamodbav:"description" json:"description"`
	CurrentPrice  float64  `dynamodbav:"currentPrice" json:"currentPrice"`
	OldPrice      float64  `dynamodbav:"oldPrice" json:"oldPrice"`
	Inventory     int      `dynamodbav:"inventory" json:"inventory"`
	Images        []string `dynamodbav:"images" json:"images"`
	IsFeatured    bool     `dynamodbav:"isFeatured" json:"isFeatured"`
	IsNew         bool     `dynamodbav:"isNew" json:"isNew"`
	IsOnSale      bool     `dynamodbav:"isOnSale" json:"isOnSale"`
	IsPopular     bool     `dynamodbav:"isPopular" json:"isPopular"`
	ShippingPrice float64  `dynamodbav:"shippingPrice" json:"shippingPrice"`
	ShippingTime  string   `dynamodbav:"shippingTime" json:"shippingTime"`
	ShippingType  string   `dynamodbav:"shippingType" json:"shippingType"`
	Slug          string   `dynamodbav:"slug" json:"slug"`
	BrandId       string   `dynamodbav:"brandId" json:"brandId"`
	Brand         Brand    `json:"brand"`
	CategoryId    string   `dynamodbav:"categoryId" json:"categoryId"`
	Category      Category `json:"category"`
}

func (product *Product) GetProductById(id string) (*Product, error) {
	key, err := attributevalue.Marshal(id)

	if err != nil {
		return nil, err
	}

	productResult, err := GetByKey[Product](ProductsTable, map[string]types.AttributeValue{
		"id": key,
	})

	if err != nil {
		return nil, err
	}

	if productResult.Id == "" {
		return productResult, errors.New("product not found")
	}

	brand, err := productResult.Brand.GetBrandById(productResult.BrandId)

	if err != nil {
		return nil, err
	}

	productResult.Brand = *brand

	category, err := productResult.Category.GetCategoryById(productResult.CategoryId)

	if err != nil {
		return nil, err
	}

	productResult.Category = *category

	return productResult, nil

}
