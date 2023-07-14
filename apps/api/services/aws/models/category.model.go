package models

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Category struct {
	Id       string          `dynamodbav:"id" json:"id"`
	Name     string          `dynamodbav:"name" json:"name"`
	Slug     string          `dynamodbav:"slug" json:"slug"`
	Children []ChildCategory `dynamodbav:"children" json:"children"`
	Parent   *Parent         `dynamodbav:"parent" json:"parent"`
}

type ChildCategory struct {
	Id       string          `dynamodbav:"id" json:"id"`
	Name     string          `dynamodbav:"name" json:"name"`
	Slug     string          `dynamodbav:"slug" json:"slug"`
	Children []ChildCategory `dynamodbav:"children" json:"children"`
}

type Parent struct {
	Id     string  `dynamodbav:"id" json:"id"`
	Name   string  `dynamodbav:"name" json:"name"`
	Slug   string  `dynamodbav:"slug" json:"slug"`
	Parent *Parent `dynamodbav:"parent" json:"parent"`
}

func (category *Category) GetCategoryById(id string) (*Category, error) {
	key, err := attributevalue.Marshal(id)

	if err != nil {
		return nil, err
	}

	categoryResult, err := GetByKey[Category](CategoriesTable, map[string]types.AttributeValue{
		"id": key,
	})

	if err != nil {
		return nil, err
	}

	if categoryResult.Id == "" {
		return categoryResult, errors.New("category not found")
	}

	return categoryResult, nil
}
