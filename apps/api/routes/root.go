package routes

import (
	"aurora/routes/auth"
	"aurora/routes/aws"
	"github.com/gin-gonic/gin"
)

func Bootstrap(router *gin.Engine) {
	app := router.Group("/v1")

	app.GET("/test", GetTestData)
	app.GET("/s3", aws.GetAwsDummyS3Data)
	app.GET("/dynamo", aws.GetAwsDummyDynamoData)
	app.POST("/auth/register", auth.RegisterUser)
	app.POST("/auth/login", auth.LoginUser)
}
