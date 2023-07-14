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

func (auth *Auth) Update() (*dynamodb.UpdateItemOutput, error) {
	key, err := attributevalue.Marshal(auth.Email)

	if err != nil {
		return nil, err
	}

	client := awsService.GetDynamoClient()
	return client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(AuthTable),
		Key: map[string]types.AttributeValue{
			"email": key,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":password": &types.AttributeValueMemberS{
				Value: auth.Password,
			},
		},
		UpdateExpression: aws.String("set password = :password"),
	})
}
