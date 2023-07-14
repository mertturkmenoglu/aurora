package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User struct {
	Id           string       `dynamodbav:"id" json:"id"`
	FullName     string       `dynamodbav:"fullName" json:"fullName"`
	Email        string       `dynamodbav:"email" json:"email"`
	AdPreference AdPreference `dynamodbav:"adPreference" json:"adPreference"`
	Addresses    []Address    `dynamodbav:"addresses" json:"addresses"`
	Phone        string       `dynamodbav:"phone" json:"phone"`
}

type AdPreference struct {
	Email bool `dynamodbav:"email" json:"email"`
	Sms   bool `dynamodbav:"sms" json:"sms"`
	Phone bool `dynamodbav:"phone" json:"phone"`
}

type Address struct {
	City        string `dynamodbav:"city" json:"city"`
	Description string `dynamodbav:"description" json:"description"`
	IsDefault   bool   `dynamodbav:"isDefault" json:"isDefault"`
	Line1       string `dynamodbav:"line1" json:"line1"`
	Line2       string `dynamodbav:"line2" json:"line2"`
	Name        string `dynamodbav:"name" json:"name"`
	Phone       string `dynamodbav:"phone" json:"phone"`
	State       string `dynamodbav:"state" json:"state"`
	Type        string `dynamodbav:"type" json:"type"`
	ZipCode     string `dynamodbav:"zipCode" json:"zipCode"`
}

func (user *User) GetUserByEmail(email string) (*User, error) {
	key, err := attributevalue.Marshal(email)

	if err != nil {
		return nil, err
	}

	return GetByKey[User](UserTable, map[string]types.AttributeValue{
		"email": key,
	})
}
