package aws

import (
	awsService "aurora/services/aws"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAwsDummyS3Data(c *gin.Context) {
	client := awsService.GetS3Client()

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("aurora-dev-eu-test"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	data := make([]string, 0, 10)

	for _, object := range output.Contents {
		s := fmt.Sprintf("key=%s size=%d", aws.ToString(object.Key), object.Size)
		data = append(data, s)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func GetAwsDummyDynamoData(c *gin.Context) {
	client := awsService.GetDynamoClient()
	output, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("testtable"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": output.Items,
	})
}
