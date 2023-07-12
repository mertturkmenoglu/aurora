package routes

import (
	"aurora/middlewares"
	"aurora/routes/auth"
	"aurora/routes/products"
	"aurora/routes/users"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"time"
)

func Bootstrap(router *gin.Engine) {
	app := router.Group("/v1")

	middlewares.Limit = ratelimit.New(1000, ratelimit.Per(time.Minute))

	app.Use(middlewares.LeakBucket())

	app.POST("/auth/register", auth.RegisterUser)
	app.POST("/auth/login", auth.LoginUser)
	app.GET("/users/:email", middlewares.IsAuth(), users.GetUserById)
	app.GET("/products/:id", products.GetProductById)
}
