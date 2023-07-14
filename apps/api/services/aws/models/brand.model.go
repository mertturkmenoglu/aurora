package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Brand struct {
	Id          string `dynamodbav:"id" json:"id"`
	Name        string `dynamodbav:"name" json:"name"`
	Description string `dynamodbav:"description" json:"description"`
}

func (brand *Brand) GetBrandById(id string) (*Brand, error) {
	key, err := attributevalue.Marshal(id)

	if err != nil {
		return nil, err
	}

	return GetByKey[Brand](BrandsTable, map[string]types.AttributeValue{
		"id": key,
	})
}
