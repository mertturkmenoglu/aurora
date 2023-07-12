package models

import (
	awsService "aurora/services/aws"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Auth struct {
	Id       string `dynamodbav:"id"`
	FullName string `dynamodbav:"fullName"`
	Email    string `dynamodbav:"email"`
	Password string `dynamodbav:"password"`
}

func (auth *Auth) GetByEmail(email string) (*Auth, error) {
	key, err := attributevalue.Marshal(email)

	if err != nil {
		return nil, err
	}

	authResult, err := GetByKey[Auth](AuthTable, map[string]types.AttributeValue{
		"email": key,
	})

	return authResult, err
}

func (auth *Auth) Save() (*dynamodb.PutItemOutput, error) {
	item, err := attributevalue.MarshalMap(auth)

	if err != nil {
		return nil, err
	}

	client := awsService.GetDynamoClient()
	output, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(AuthTable),
		Item:      item,
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}
