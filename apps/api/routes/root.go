package routes

import (
	"aurora/middlewares"
	"aurora/routes/auth"
	"aurora/routes/aws"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"net/http"
	"time"
)

func Bootstrap(router *gin.Engine) {
	app := router.Group("/v1")

	middlewares.Limit = ratelimit.New(1000, ratelimit.Per(time.Minute))

	app.Use(middlewares.LeakBucket())

	app.GET("/test", GetTestData)
	app.GET("/s3", aws.GetAwsDummyS3Data)
	app.GET("/dynamo", aws.GetAwsDummyDynamoData)
	app.POST("/auth/register", auth.RegisterUser)
	app.POST("/auth/login", auth.LoginUser)
	app.GET("/secret", middlewares.IsAuth(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Secret data",
		})
	})
}
