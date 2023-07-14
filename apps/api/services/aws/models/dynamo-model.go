package models

import (
	awsService "aurora/services/aws"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoModel interface {
	Auth | Product | Brand | User | Address | AdPreference | Category
}

func GetByKey[Table DynamoModel](tableName string, key map[string]types.AttributeValue) (*Table, error) {
	client := awsService.GetDynamoClient()
	output, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})

	if err != nil {
		return nil, err
	}

	var table Table
	err = attributevalue.UnmarshalMap(output.Item, &table)

	if err != nil {
		return nil, err
	}

	return &table, nil
}

func Save[T DynamoModel](data *T, tableName string) (*dynamodb.PutItemOutput, error) {
	item, err := attributevalue.MarshalMap(data)

	if err != nil {
		return nil, err
	}

	client := awsService.GetDynamoClient()
	return client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
}
