package routes

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAwsDummyS3Data(c *gin.Context) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	client := s3.NewFromConfig(cfg)

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
