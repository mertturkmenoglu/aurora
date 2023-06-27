package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO())
}

func GetS3Client() *s3.Client {
	cfg, err := getConfig()
	if err != nil {
		panic(err)
	}
	return s3.NewFromConfig(cfg)
}

func GetDynamoClient() *dynamodb.Client {
	cfg, err := getConfig()
	if err != nil {
		panic(err)
	}
	return dynamodb.NewFromConfig(cfg)
}
