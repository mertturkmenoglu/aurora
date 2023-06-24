package routes

import "github.com/gin-gonic/gin"

func Bootstrap(router *gin.Engine) {
	app := router.Group("/v1")

	app.GET("/test", GetTestData)
}
